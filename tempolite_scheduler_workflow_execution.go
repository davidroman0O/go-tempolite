package tempolite

import (
	"runtime"

	"github.com/davidroman0O/retrypool"
	"github.com/davidroman0O/tempolite/ent"
	"github.com/davidroman0O/tempolite/ent/workflow"
	"github.com/davidroman0O/tempolite/ent/workflowexecution"
	"github.com/google/uuid"
)

func (tp *Tempolite) schedulerExecutionWorkflow() {
	defer close(tp.schedulerExecutionWorkflowDone)
	for {
		select {
		case <-tp.ctx.Done():
			return
		default:
			pendingWorkflows, err := tp.client.WorkflowExecution.Query().
				Where(workflowexecution.StatusEQ(workflowexecution.StatusPending)).
				Order(ent.Asc(workflowexecution.FieldStartedAt)).
				WithWorkflow().
				Limit(1).
				All(tp.ctx)
			if err != nil {
				continue
			}

			tp.schedulerWorkflowStarted.Store(true)

			if len(pendingWorkflows) == 0 {
				continue
			}

			var value any
			var ok bool

			for _, pendingWorkflowExecution := range pendingWorkflows {
				tp.logger.Debug(tp.ctx, "Scheduler workflow execution: pending workflow", "workflow_id", pendingWorkflowExecution.Edges.Workflow.ID, "workflow_execution_id", pendingWorkflowExecution.ID)

				var workflowEntity *ent.Workflow
				if workflowEntity, err = tp.client.Workflow.Get(tp.ctx, pendingWorkflowExecution.Edges.Workflow.ID); err != nil {
					tp.logger.Error(tp.ctx, "Scheduler workflow execution: workflow.Get failed", "error", err)
					continue
				}

				if value, ok = tp.workflows.Load(HandlerIdentity(workflowEntity.Identity)); ok {
					var workflowHandlerInfo Workflow
					if workflowHandlerInfo, ok = value.(Workflow); !ok {
						tp.logger.Error(tp.ctx, "Scheduler workflow execution: workflow handler info not found", "workflow_id", pendingWorkflowExecution.Edges.Workflow.ID)
						continue
					}

					inputs, err := tp.convertInputs(HandlerInfo(workflowHandlerInfo), workflowEntity.Input)
					if err != nil {
						tp.logger.Error(tp.ctx, "Scheduler workflow execution: convertInputs failed", "error", err)
						continue
					}

					contextWorkflow := WorkflowContext{
						tp:              tp,
						workflowID:      workflowEntity.ID,
						executionID:     pendingWorkflowExecution.ID,
						runID:           pendingWorkflowExecution.RunID,
						workflowType:    workflowEntity.Identity,
						stepID:          workflowEntity.StepID,
						handlerIdentity: HandlerIdentity(workflowEntity.Identity),
					}

					task := &workflowTask{
						ctx:         contextWorkflow,
						handlerName: workflowHandlerInfo.HandlerLongName,
						handler:     workflowHandlerInfo.Handler,
						params:      inputs,
						maxRetry:    workflowEntity.RetryPolicy.MaximumAttempts,
						retryCount:  0,
					}

					retryIt := func() error {
						tx, err := tp.client.Tx(tp.ctx)
						if err != nil {
							return err
						}
						var workflowExecution *ent.WorkflowExecution
						if workflowExecution, err = tx.WorkflowExecution.
							Create().
							SetID(uuid.NewString()).
							SetRunID(pendingWorkflowExecution.RunID).
							SetWorkflow(workflowEntity).
							Save(tp.ctx); err != nil {
							tp.logger.Error(tp.ctx, "Scheduler workflow execution: retry create new execution failed", "error", err)
							return tx.Rollback()
						}

						tp.logger.Debug(tp.ctx, "Scheduler workflow execution: retrying", "workflow_id", pendingWorkflowExecution.Edges.Workflow.ID, "workflow_execution_id", workflowExecution.ID)

						task.ctx.executionID = workflowExecution.ID
						task.retryCount++

						tp.logger.Debug(tp.ctx, "Scheduler workflow execution: retrying", "workflow_id", pendingWorkflowExecution.Edges.Workflow.ID, "workflow_execution_id", workflowExecution.ID, "retry_count", task.retryCount)

						if _, err = tx.Workflow.UpdateOneID(contextWorkflow.workflowID).SetStatus(workflow.StatusRunning).Save(tp.ctx); err != nil {
							tp.logger.Error(tp.ctx, "Scheduler workflow execution: retry update workflow status failed", "error", err)
							return tx.Rollback()
						}

						return tx.Commit()
					}

					task.retry = retryIt

					total, err := tp.client.WorkflowExecution.Query().Where(workflowexecution.HasWorkflowWith(workflow.IDEQ(workflowEntity.ID))).Count(tp.ctx)
					if err != nil {
						tp.logger.Error(tp.ctx, "Scheduler workflow execution: WorkflowExecution.Query total failed", "error", err)
						continue
					}

					if total > 1 {
						task.retryCount = total
					}

					tp.logger.Debug(tp.ctx, "Scheduler workflow execution: Dispatching", "workflow_id", pendingWorkflowExecution.Edges.Workflow.ID, "workflow_execution_id", pendingWorkflowExecution.ID)

					if err := tp.workflowPool.Dispatch(task, retrypool.WithImmediateRetry[*workflowTask]()); err != nil {
						tp.logger.Error(tp.ctx, "Scheduler workflow execution: Dispatch failed", "error", err)
						tx, txErr := tp.client.Tx(tp.ctx)
						if txErr != nil {
							tp.logger.Error(tp.ctx, "Scheduler workflow execution: Failed to start transaction", "error", txErr)
							continue
						}
						if _, err = tx.WorkflowExecution.UpdateOneID(pendingWorkflowExecution.ID).SetStatus(workflowexecution.StatusFailed).SetError(err.Error()).Save(tp.ctx); err != nil {
							tp.logger.Error(tp.ctx, "Scheduler workflow execution: WorkflowExecution.UpdateOneID failed when failed to dispatch", "error", err)
							if rollbackErr := tx.Rollback(); rollbackErr != nil {
								tp.logger.Error(tp.ctx, "Scheduler workflow execution: Failed to rollback transaction", "error", rollbackErr)
							}
							continue
						}
						if _, err = tx.Workflow.UpdateOneID(workflowEntity.ID).SetStatus(workflow.StatusFailed).Save(tp.ctx); err != nil {
							tp.logger.Error(tp.ctx, "Scheduler workflow execution: Workflow.UpdateOneID failed when failed to dispatch", "error", err)
							if rollbackErr := tx.Rollback(); rollbackErr != nil {
								tp.logger.Error(tp.ctx, "Scheduler workflow execution: Failed to rollback transaction", "error", rollbackErr)
							}
							continue
						}
						if err = tx.Commit(); err != nil {
							tp.logger.Error(tp.ctx, "Scheduler workflow execution: Failed to commit transaction", "error", err)
						}
						continue
					}

					tx, err := tp.client.Tx(tp.ctx)
					if err != nil {
						tp.logger.Error(tp.ctx, "Scheduler workflow execution: Failed to start transaction", "error", err)
						continue
					}
					if _, err = tx.WorkflowExecution.UpdateOneID(pendingWorkflowExecution.ID).SetStatus(workflowexecution.StatusRunning).Save(tp.ctx); err != nil {
						tp.logger.Error(tp.ctx, "Scheduler workflow execution: WorkflowExecution.UpdateOneID failed", "error", err)
						if rollbackErr := tx.Rollback(); rollbackErr != nil {
							tp.logger.Error(tp.ctx, "Scheduler workflow execution: Failed to rollback transaction", "error", rollbackErr)
						}
						continue
					}
					if _, err = tx.Workflow.UpdateOneID(workflowEntity.ID).SetStatus(workflow.StatusRunning).Save(tp.ctx); err != nil {
						tp.logger.Error(tp.ctx, "Scheduler workflow execution: Workflow.UpdateOneID failed", "error", err)
						if rollbackErr := tx.Rollback(); rollbackErr != nil {
							tp.logger.Error(tp.ctx, "Scheduler workflow execution: Failed to rollback transaction", "error", rollbackErr)
						}
						continue
					}
					if err = tx.Commit(); err != nil {
						tp.logger.Error(tp.ctx, "Scheduler workflow execution: Failed to commit transaction", "error", err)
					}
				} else {
					tp.logger.Error(tp.ctx, "Workflow not found", "workflow_id", pendingWorkflowExecution.Edges.Workflow.ID)
					continue
				}
			}

			runtime.Gosched()
		}
	}
}
