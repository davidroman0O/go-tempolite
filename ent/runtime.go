// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/davidroman0O/go-tempolite/ent/activity"
	"github.com/davidroman0O/go-tempolite/ent/activityexecution"
	"github.com/davidroman0O/go-tempolite/ent/executionrelationship"
	"github.com/davidroman0O/go-tempolite/ent/run"
	"github.com/davidroman0O/go-tempolite/ent/saga"
	"github.com/davidroman0O/go-tempolite/ent/sagaexecution"
	"github.com/davidroman0O/go-tempolite/ent/schema"
	"github.com/davidroman0O/go-tempolite/ent/sideeffect"
	"github.com/davidroman0O/go-tempolite/ent/sideeffectexecution"
	"github.com/davidroman0O/go-tempolite/ent/signal"
	"github.com/davidroman0O/go-tempolite/ent/signalexecution"
	"github.com/davidroman0O/go-tempolite/ent/workflow"
	"github.com/davidroman0O/go-tempolite/ent/workflowexecution"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	activityFields := schema.Activity{}.Fields()
	_ = activityFields
	// activityDescIdentity is the schema descriptor for identity field.
	activityDescIdentity := activityFields[1].Descriptor()
	// activity.IdentityValidator is a validator for the "identity" field. It is called by the builders before save.
	activity.IdentityValidator = activityDescIdentity.Validators[0].(func(string) error)
	// activityDescStepID is the schema descriptor for step_id field.
	activityDescStepID := activityFields[2].Descriptor()
	// activity.StepIDValidator is a validator for the "step_id" field. It is called by the builders before save.
	activity.StepIDValidator = activityDescStepID.Validators[0].(func(string) error)
	// activityDescHandlerName is the schema descriptor for handler_name field.
	activityDescHandlerName := activityFields[4].Descriptor()
	// activity.HandlerNameValidator is a validator for the "handler_name" field. It is called by the builders before save.
	activity.HandlerNameValidator = activityDescHandlerName.Validators[0].(func(string) error)
	// activityDescCreatedAt is the schema descriptor for created_at field.
	activityDescCreatedAt := activityFields[8].Descriptor()
	// activity.DefaultCreatedAt holds the default value on creation for the created_at field.
	activity.DefaultCreatedAt = activityDescCreatedAt.Default.(func() time.Time)
	activityexecutionFields := schema.ActivityExecution{}.Fields()
	_ = activityexecutionFields
	// activityexecutionDescAttempt is the schema descriptor for attempt field.
	activityexecutionDescAttempt := activityexecutionFields[3].Descriptor()
	// activityexecution.DefaultAttempt holds the default value on creation for the attempt field.
	activityexecution.DefaultAttempt = activityexecutionDescAttempt.Default.(int)
	// activityexecutionDescStartedAt is the schema descriptor for started_at field.
	activityexecutionDescStartedAt := activityexecutionFields[6].Descriptor()
	// activityexecution.DefaultStartedAt holds the default value on creation for the started_at field.
	activityexecution.DefaultStartedAt = activityexecutionDescStartedAt.Default.(func() time.Time)
	// activityexecutionDescUpdatedAt is the schema descriptor for updated_at field.
	activityexecutionDescUpdatedAt := activityexecutionFields[7].Descriptor()
	// activityexecution.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	activityexecution.DefaultUpdatedAt = activityexecutionDescUpdatedAt.Default.(func() time.Time)
	// activityexecution.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	activityexecution.UpdateDefaultUpdatedAt = activityexecutionDescUpdatedAt.UpdateDefault.(func() time.Time)
	executionrelationshipFields := schema.ExecutionRelationship{}.Fields()
	_ = executionrelationshipFields
	// executionrelationshipDescRunID is the schema descriptor for run_id field.
	executionrelationshipDescRunID := executionrelationshipFields[0].Descriptor()
	// executionrelationship.RunIDValidator is a validator for the "run_id" field. It is called by the builders before save.
	executionrelationship.RunIDValidator = executionrelationshipDescRunID.Validators[0].(func(string) error)
	// executionrelationshipDescParentEntityID is the schema descriptor for parent_entity_id field.
	executionrelationshipDescParentEntityID := executionrelationshipFields[1].Descriptor()
	// executionrelationship.ParentEntityIDValidator is a validator for the "parent_entity_id" field. It is called by the builders before save.
	executionrelationship.ParentEntityIDValidator = executionrelationshipDescParentEntityID.Validators[0].(func(string) error)
	// executionrelationshipDescChildEntityID is the schema descriptor for child_entity_id field.
	executionrelationshipDescChildEntityID := executionrelationshipFields[2].Descriptor()
	// executionrelationship.ChildEntityIDValidator is a validator for the "child_entity_id" field. It is called by the builders before save.
	executionrelationship.ChildEntityIDValidator = executionrelationshipDescChildEntityID.Validators[0].(func(string) error)
	// executionrelationshipDescParentStepID is the schema descriptor for parent_step_id field.
	executionrelationshipDescParentStepID := executionrelationshipFields[7].Descriptor()
	// executionrelationship.ParentStepIDValidator is a validator for the "parent_step_id" field. It is called by the builders before save.
	executionrelationship.ParentStepIDValidator = executionrelationshipDescParentStepID.Validators[0].(func(string) error)
	// executionrelationshipDescChildStepID is the schema descriptor for child_step_id field.
	executionrelationshipDescChildStepID := executionrelationshipFields[8].Descriptor()
	// executionrelationship.ChildStepIDValidator is a validator for the "child_step_id" field. It is called by the builders before save.
	executionrelationship.ChildStepIDValidator = executionrelationshipDescChildStepID.Validators[0].(func(string) error)
	runFields := schema.Run{}.Fields()
	_ = runFields
	// runDescCreatedAt is the schema descriptor for created_at field.
	runDescCreatedAt := runFields[3].Descriptor()
	// run.DefaultCreatedAt holds the default value on creation for the created_at field.
	run.DefaultCreatedAt = runDescCreatedAt.Default.(func() time.Time)
	sagaFields := schema.Saga{}.Fields()
	_ = sagaFields
	// sagaDescStepID is the schema descriptor for step_id field.
	sagaDescStepID := sagaFields[2].Descriptor()
	// saga.StepIDValidator is a validator for the "step_id" field. It is called by the builders before save.
	saga.StepIDValidator = sagaDescStepID.Validators[0].(func(string) error)
	// sagaDescCreatedAt is the schema descriptor for created_at field.
	sagaDescCreatedAt := sagaFields[6].Descriptor()
	// saga.DefaultCreatedAt holds the default value on creation for the created_at field.
	saga.DefaultCreatedAt = sagaDescCreatedAt.Default.(func() time.Time)
	// sagaDescUpdatedAt is the schema descriptor for updated_at field.
	sagaDescUpdatedAt := sagaFields[7].Descriptor()
	// saga.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	saga.DefaultUpdatedAt = sagaDescUpdatedAt.Default.(func() time.Time)
	// saga.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	saga.UpdateDefaultUpdatedAt = sagaDescUpdatedAt.UpdateDefault.(func() time.Time)
	sagaexecutionFields := schema.SagaExecution{}.Fields()
	_ = sagaexecutionFields
	// sagaexecutionDescHandlerName is the schema descriptor for handler_name field.
	sagaexecutionDescHandlerName := sagaexecutionFields[1].Descriptor()
	// sagaexecution.HandlerNameValidator is a validator for the "handler_name" field. It is called by the builders before save.
	sagaexecution.HandlerNameValidator = sagaexecutionDescHandlerName.Validators[0].(func(string) error)
	// sagaexecutionDescSequence is the schema descriptor for sequence field.
	sagaexecutionDescSequence := sagaexecutionFields[4].Descriptor()
	// sagaexecution.SequenceValidator is a validator for the "sequence" field. It is called by the builders before save.
	sagaexecution.SequenceValidator = sagaexecutionDescSequence.Validators[0].(func(int) error)
	// sagaexecutionDescStartedAt is the schema descriptor for started_at field.
	sagaexecutionDescStartedAt := sagaexecutionFields[6].Descriptor()
	// sagaexecution.DefaultStartedAt holds the default value on creation for the started_at field.
	sagaexecution.DefaultStartedAt = sagaexecutionDescStartedAt.Default.(func() time.Time)
	sideeffectFields := schema.SideEffect{}.Fields()
	_ = sideeffectFields
	// sideeffectDescIdentity is the schema descriptor for identity field.
	sideeffectDescIdentity := sideeffectFields[1].Descriptor()
	// sideeffect.IdentityValidator is a validator for the "identity" field. It is called by the builders before save.
	sideeffect.IdentityValidator = sideeffectDescIdentity.Validators[0].(func(string) error)
	// sideeffectDescStepID is the schema descriptor for step_id field.
	sideeffectDescStepID := sideeffectFields[2].Descriptor()
	// sideeffect.StepIDValidator is a validator for the "step_id" field. It is called by the builders before save.
	sideeffect.StepIDValidator = sideeffectDescStepID.Validators[0].(func(string) error)
	// sideeffectDescHandlerName is the schema descriptor for handler_name field.
	sideeffectDescHandlerName := sideeffectFields[3].Descriptor()
	// sideeffect.HandlerNameValidator is a validator for the "handler_name" field. It is called by the builders before save.
	sideeffect.HandlerNameValidator = sideeffectDescHandlerName.Validators[0].(func(string) error)
	// sideeffectDescCreatedAt is the schema descriptor for created_at field.
	sideeffectDescCreatedAt := sideeffectFields[7].Descriptor()
	// sideeffect.DefaultCreatedAt holds the default value on creation for the created_at field.
	sideeffect.DefaultCreatedAt = sideeffectDescCreatedAt.Default.(func() time.Time)
	sideeffectexecutionFields := schema.SideEffectExecution{}.Fields()
	_ = sideeffectexecutionFields
	// sideeffectexecutionDescAttempt is the schema descriptor for attempt field.
	sideeffectexecutionDescAttempt := sideeffectexecutionFields[2].Descriptor()
	// sideeffectexecution.DefaultAttempt holds the default value on creation for the attempt field.
	sideeffectexecution.DefaultAttempt = sideeffectexecutionDescAttempt.Default.(int)
	// sideeffectexecutionDescStartedAt is the schema descriptor for started_at field.
	sideeffectexecutionDescStartedAt := sideeffectexecutionFields[5].Descriptor()
	// sideeffectexecution.DefaultStartedAt holds the default value on creation for the started_at field.
	sideeffectexecution.DefaultStartedAt = sideeffectexecutionDescStartedAt.Default.(func() time.Time)
	// sideeffectexecutionDescUpdatedAt is the schema descriptor for updated_at field.
	sideeffectexecutionDescUpdatedAt := sideeffectexecutionFields[6].Descriptor()
	// sideeffectexecution.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sideeffectexecution.DefaultUpdatedAt = sideeffectexecutionDescUpdatedAt.Default.(func() time.Time)
	// sideeffectexecution.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sideeffectexecution.UpdateDefaultUpdatedAt = sideeffectexecutionDescUpdatedAt.UpdateDefault.(func() time.Time)
	signalFields := schema.Signal{}.Fields()
	_ = signalFields
	// signalDescStepID is the schema descriptor for step_id field.
	signalDescStepID := signalFields[1].Descriptor()
	// signal.StepIDValidator is a validator for the "step_id" field. It is called by the builders before save.
	signal.StepIDValidator = signalDescStepID.Validators[0].(func(string) error)
	// signalDescCreatedAt is the schema descriptor for created_at field.
	signalDescCreatedAt := signalFields[3].Descriptor()
	// signal.DefaultCreatedAt holds the default value on creation for the created_at field.
	signal.DefaultCreatedAt = signalDescCreatedAt.Default.(func() time.Time)
	// signalDescConsumed is the schema descriptor for consumed field.
	signalDescConsumed := signalFields[4].Descriptor()
	// signal.DefaultConsumed holds the default value on creation for the consumed field.
	signal.DefaultConsumed = signalDescConsumed.Default.(bool)
	signalexecutionFields := schema.SignalExecution{}.Fields()
	_ = signalexecutionFields
	// signalexecutionDescStartedAt is the schema descriptor for started_at field.
	signalexecutionDescStartedAt := signalexecutionFields[5].Descriptor()
	// signalexecution.DefaultStartedAt holds the default value on creation for the started_at field.
	signalexecution.DefaultStartedAt = signalexecutionDescStartedAt.Default.(func() time.Time)
	// signalexecutionDescUpdatedAt is the schema descriptor for updated_at field.
	signalexecutionDescUpdatedAt := signalexecutionFields[6].Descriptor()
	// signalexecution.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	signalexecution.DefaultUpdatedAt = signalexecutionDescUpdatedAt.Default.(func() time.Time)
	// signalexecution.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	signalexecution.UpdateDefaultUpdatedAt = signalexecutionDescUpdatedAt.UpdateDefault.(func() time.Time)
	workflowFields := schema.Workflow{}.Fields()
	_ = workflowFields
	// workflowDescStepID is the schema descriptor for step_id field.
	workflowDescStepID := workflowFields[1].Descriptor()
	// workflow.StepIDValidator is a validator for the "step_id" field. It is called by the builders before save.
	workflow.StepIDValidator = workflowDescStepID.Validators[0].(func(string) error)
	// workflowDescIdentity is the schema descriptor for identity field.
	workflowDescIdentity := workflowFields[3].Descriptor()
	// workflow.IdentityValidator is a validator for the "identity" field. It is called by the builders before save.
	workflow.IdentityValidator = workflowDescIdentity.Validators[0].(func(string) error)
	// workflowDescHandlerName is the schema descriptor for handler_name field.
	workflowDescHandlerName := workflowFields[4].Descriptor()
	// workflow.HandlerNameValidator is a validator for the "handler_name" field. It is called by the builders before save.
	workflow.HandlerNameValidator = workflowDescHandlerName.Validators[0].(func(string) error)
	// workflowDescIsPaused is the schema descriptor for is_paused field.
	workflowDescIsPaused := workflowFields[7].Descriptor()
	// workflow.DefaultIsPaused holds the default value on creation for the is_paused field.
	workflow.DefaultIsPaused = workflowDescIsPaused.Default.(bool)
	// workflowDescIsReady is the schema descriptor for is_ready field.
	workflowDescIsReady := workflowFields[8].Descriptor()
	// workflow.DefaultIsReady holds the default value on creation for the is_ready field.
	workflow.DefaultIsReady = workflowDescIsReady.Default.(bool)
	// workflowDescCreatedAt is the schema descriptor for created_at field.
	workflowDescCreatedAt := workflowFields[10].Descriptor()
	// workflow.DefaultCreatedAt holds the default value on creation for the created_at field.
	workflow.DefaultCreatedAt = workflowDescCreatedAt.Default.(func() time.Time)
	workflowexecutionFields := schema.WorkflowExecution{}.Fields()
	_ = workflowexecutionFields
	// workflowexecutionDescIsReplay is the schema descriptor for is_replay field.
	workflowexecutionDescIsReplay := workflowexecutionFields[5].Descriptor()
	// workflowexecution.DefaultIsReplay holds the default value on creation for the is_replay field.
	workflowexecution.DefaultIsReplay = workflowexecutionDescIsReplay.Default.(bool)
	// workflowexecutionDescStartedAt is the schema descriptor for started_at field.
	workflowexecutionDescStartedAt := workflowexecutionFields[6].Descriptor()
	// workflowexecution.DefaultStartedAt holds the default value on creation for the started_at field.
	workflowexecution.DefaultStartedAt = workflowexecutionDescStartedAt.Default.(func() time.Time)
	// workflowexecutionDescUpdatedAt is the schema descriptor for updated_at field.
	workflowexecutionDescUpdatedAt := workflowexecutionFields[7].Descriptor()
	// workflowexecution.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	workflowexecution.DefaultUpdatedAt = workflowexecutionDescUpdatedAt.Default.(func() time.Time)
	// workflowexecution.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	workflowexecution.UpdateDefaultUpdatedAt = workflowexecutionDescUpdatedAt.UpdateDefault.(func() time.Time)
}
