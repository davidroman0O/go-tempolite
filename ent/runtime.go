// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/davidroman0O/go-tempolite/ent/activity"
	"github.com/davidroman0O/go-tempolite/ent/activityexecution"
	"github.com/davidroman0O/go-tempolite/ent/run"
	"github.com/davidroman0O/go-tempolite/ent/saga"
	"github.com/davidroman0O/go-tempolite/ent/sagaexecution"
	"github.com/davidroman0O/go-tempolite/ent/sagastepexecution"
	"github.com/davidroman0O/go-tempolite/ent/schema"
	"github.com/davidroman0O/go-tempolite/ent/sideeffect"
	"github.com/davidroman0O/go-tempolite/ent/sideeffectexecution"
	"github.com/davidroman0O/go-tempolite/ent/signal"
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
	// activityDescHandlerName is the schema descriptor for handler_name field.
	activityDescHandlerName := activityFields[2].Descriptor()
	// activity.HandlerNameValidator is a validator for the "handler_name" field. It is called by the builders before save.
	activity.HandlerNameValidator = activityDescHandlerName.Validators[0].(func(string) error)
	// activityDescCreatedAt is the schema descriptor for created_at field.
	activityDescCreatedAt := activityFields[6].Descriptor()
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
	runFields := schema.Run{}.Fields()
	_ = runFields
	// runDescCreatedAt is the schema descriptor for created_at field.
	runDescCreatedAt := runFields[3].Descriptor()
	// run.DefaultCreatedAt holds the default value on creation for the created_at field.
	run.DefaultCreatedAt = runDescCreatedAt.Default.(func() time.Time)
	sagaFields := schema.Saga{}.Fields()
	_ = sagaFields
	// sagaDescName is the schema descriptor for name field.
	sagaDescName := sagaFields[1].Descriptor()
	// saga.NameValidator is a validator for the "name" field. It is called by the builders before save.
	saga.NameValidator = sagaDescName.Validators[0].(func(string) error)
	// sagaDescCreatedAt is the schema descriptor for created_at field.
	sagaDescCreatedAt := sagaFields[5].Descriptor()
	// saga.DefaultCreatedAt holds the default value on creation for the created_at field.
	saga.DefaultCreatedAt = sagaDescCreatedAt.Default.(func() time.Time)
	sagaexecutionFields := schema.SagaExecution{}.Fields()
	_ = sagaexecutionFields
	// sagaexecutionDescAttempt is the schema descriptor for attempt field.
	sagaexecutionDescAttempt := sagaexecutionFields[3].Descriptor()
	// sagaexecution.DefaultAttempt holds the default value on creation for the attempt field.
	sagaexecution.DefaultAttempt = sagaexecutionDescAttempt.Default.(int)
	// sagaexecutionDescStartedAt is the schema descriptor for started_at field.
	sagaexecutionDescStartedAt := sagaexecutionFields[5].Descriptor()
	// sagaexecution.DefaultStartedAt holds the default value on creation for the started_at field.
	sagaexecution.DefaultStartedAt = sagaexecutionDescStartedAt.Default.(func() time.Time)
	// sagaexecutionDescUpdatedAt is the schema descriptor for updated_at field.
	sagaexecutionDescUpdatedAt := sagaexecutionFields[6].Descriptor()
	// sagaexecution.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sagaexecution.DefaultUpdatedAt = sagaexecutionDescUpdatedAt.Default.(func() time.Time)
	// sagaexecution.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sagaexecution.UpdateDefaultUpdatedAt = sagaexecutionDescUpdatedAt.UpdateDefault.(func() time.Time)
	sagastepexecutionFields := schema.SagaStepExecution{}.Fields()
	_ = sagastepexecutionFields
	// sagastepexecutionDescHandlerName is the schema descriptor for handler_name field.
	sagastepexecutionDescHandlerName := sagastepexecutionFields[1].Descriptor()
	// sagastepexecution.HandlerNameValidator is a validator for the "handler_name" field. It is called by the builders before save.
	sagastepexecution.HandlerNameValidator = sagastepexecutionDescHandlerName.Validators[0].(func(string) error)
	// sagastepexecutionDescSequence is the schema descriptor for sequence field.
	sagastepexecutionDescSequence := sagastepexecutionFields[4].Descriptor()
	// sagastepexecution.SequenceValidator is a validator for the "sequence" field. It is called by the builders before save.
	sagastepexecution.SequenceValidator = sagastepexecutionDescSequence.Validators[0].(func(int) error)
	// sagastepexecutionDescAttempt is the schema descriptor for attempt field.
	sagastepexecutionDescAttempt := sagastepexecutionFields[5].Descriptor()
	// sagastepexecution.DefaultAttempt holds the default value on creation for the attempt field.
	sagastepexecution.DefaultAttempt = sagastepexecutionDescAttempt.Default.(int)
	// sagastepexecutionDescStartedAt is the schema descriptor for started_at field.
	sagastepexecutionDescStartedAt := sagastepexecutionFields[8].Descriptor()
	// sagastepexecution.DefaultStartedAt holds the default value on creation for the started_at field.
	sagastepexecution.DefaultStartedAt = sagastepexecutionDescStartedAt.Default.(func() time.Time)
	// sagastepexecutionDescUpdatedAt is the schema descriptor for updated_at field.
	sagastepexecutionDescUpdatedAt := sagastepexecutionFields[9].Descriptor()
	// sagastepexecution.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sagastepexecution.DefaultUpdatedAt = sagastepexecutionDescUpdatedAt.Default.(func() time.Time)
	// sagastepexecution.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sagastepexecution.UpdateDefaultUpdatedAt = sagastepexecutionDescUpdatedAt.UpdateDefault.(func() time.Time)
	sideeffectFields := schema.SideEffect{}.Fields()
	_ = sideeffectFields
	// sideeffectDescIdentity is the schema descriptor for identity field.
	sideeffectDescIdentity := sideeffectFields[1].Descriptor()
	// sideeffect.IdentityValidator is a validator for the "identity" field. It is called by the builders before save.
	sideeffect.IdentityValidator = sideeffectDescIdentity.Validators[0].(func(string) error)
	// sideeffectDescHandlerName is the schema descriptor for handler_name field.
	sideeffectDescHandlerName := sideeffectFields[2].Descriptor()
	// sideeffect.HandlerNameValidator is a validator for the "handler_name" field. It is called by the builders before save.
	sideeffect.HandlerNameValidator = sideeffectDescHandlerName.Validators[0].(func(string) error)
	// sideeffectDescCreatedAt is the schema descriptor for created_at field.
	sideeffectDescCreatedAt := sideeffectFields[6].Descriptor()
	// sideeffect.DefaultCreatedAt holds the default value on creation for the created_at field.
	sideeffect.DefaultCreatedAt = sideeffectDescCreatedAt.Default.(func() time.Time)
	sideeffectexecutionFields := schema.SideEffectExecution{}.Fields()
	_ = sideeffectexecutionFields
	// sideeffectexecutionDescAttempt is the schema descriptor for attempt field.
	sideeffectexecutionDescAttempt := sideeffectexecutionFields[3].Descriptor()
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
	// signalDescName is the schema descriptor for name field.
	signalDescName := signalFields[1].Descriptor()
	// signal.NameValidator is a validator for the "name" field. It is called by the builders before save.
	signal.NameValidator = signalDescName.Validators[0].(func(string) error)
	// signalDescCreatedAt is the schema descriptor for created_at field.
	signalDescCreatedAt := signalFields[4].Descriptor()
	// signal.DefaultCreatedAt holds the default value on creation for the created_at field.
	signal.DefaultCreatedAt = signalDescCreatedAt.Default.(func() time.Time)
	// signalDescUpdatedAt is the schema descriptor for updated_at field.
	signalDescUpdatedAt := signalFields[5].Descriptor()
	// signal.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	signal.DefaultUpdatedAt = signalDescUpdatedAt.Default.(func() time.Time)
	// signal.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	signal.UpdateDefaultUpdatedAt = signalDescUpdatedAt.UpdateDefault.(func() time.Time)
	workflowFields := schema.Workflow{}.Fields()
	_ = workflowFields
	// workflowDescIdentity is the schema descriptor for identity field.
	workflowDescIdentity := workflowFields[1].Descriptor()
	// workflow.IdentityValidator is a validator for the "identity" field. It is called by the builders before save.
	workflow.IdentityValidator = workflowDescIdentity.Validators[0].(func(string) error)
	// workflowDescHandlerName is the schema descriptor for handler_name field.
	workflowDescHandlerName := workflowFields[2].Descriptor()
	// workflow.HandlerNameValidator is a validator for the "handler_name" field. It is called by the builders before save.
	workflow.HandlerNameValidator = workflowDescHandlerName.Validators[0].(func(string) error)
	// workflowDescCreatedAt is the schema descriptor for created_at field.
	workflowDescCreatedAt := workflowFields[6].Descriptor()
	// workflow.DefaultCreatedAt holds the default value on creation for the created_at field.
	workflow.DefaultCreatedAt = workflowDescCreatedAt.Default.(func() time.Time)
	workflowexecutionFields := schema.WorkflowExecution{}.Fields()
	_ = workflowexecutionFields
	// workflowexecutionDescStartedAt is the schema descriptor for started_at field.
	workflowexecutionDescStartedAt := workflowexecutionFields[5].Descriptor()
	// workflowexecution.DefaultStartedAt holds the default value on creation for the started_at field.
	workflowexecution.DefaultStartedAt = workflowexecutionDescStartedAt.Default.(func() time.Time)
	// workflowexecutionDescUpdatedAt is the schema descriptor for updated_at field.
	workflowexecutionDescUpdatedAt := workflowexecutionFields[6].Descriptor()
	// workflowexecution.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	workflowexecution.DefaultUpdatedAt = workflowexecutionDescUpdatedAt.Default.(func() time.Time)
	// workflowexecution.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	workflowexecution.UpdateDefaultUpdatedAt = workflowexecutionDescUpdatedAt.UpdateDefault.(func() time.Time)
}
