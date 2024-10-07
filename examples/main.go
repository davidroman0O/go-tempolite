package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/davidroman0O/go-tempolite"
	_ "github.com/mattn/go-sqlite3"
)

var saga = tempolite.Saga().
	With(SagaStep1{Message: " - Step 1"}).
	With(SagaStep2{Message: " - Step 2"}).
	Build()

// SimpleTask represents a simple task with a message
type SimpleTask struct {
	Message string
}

// SagaTask represents a task within a saga
type SagaTask struct {
	StepMessage string
}

// SimpleSideEffect implements the SideEffect interface
type SimpleSideEffect struct {
	Message string
}

func (s SimpleSideEffect) Run(ctx tempolite.SideEffectContext) (interface{}, error) {
	return "Side effect result for " + s.Message, nil
}

// SimpleHandler handles SimpleTask
func SimpleHandler(ctx tempolite.HandlerContext, task SimpleTask) (interface{}, error) {
	log.Printf("Executing simple task: %s", task.Message)

	// Use a side effect
	result, err := ctx.SideEffect("simple-task-side-effect", SimpleSideEffect{Message: task.Message})
	if err != nil {
		return nil, fmt.Errorf("side effect failed: %v", err)
	}

	log.Printf("SimpleHandler Side Effect Example: Side effect result: %v", result)

	return result, nil
}

func SagaHandler(ctx tempolite.HandlerContext, task SimpleTask) (interface{}, error) {

	// Create a children saga
	sagaTaskID, err := ctx.EnqueueSaga(saga, SagaTask{StepMessage: "Saga steps"})
	if err != nil {
		log.Fatalf("Failed to enqueue saga: %v", err)
	}
	log.Printf("Enqueued saga with ID: %s", sagaTaskID)

	value, err := ctx.WaitForCompletion(sagaTaskID)
	if err != nil {
		log.Fatalf("Failed to wait for task completion: %v", err)
	}
	log.Printf("Saga Task completed with value: %v", value)

	return value, nil
}

// SagaStep1 implements the SagaStep interface for step 1
type SagaStep1 struct {
	Message string
}

func (s SagaStep1) Transaction(ctx tempolite.TransactionContext) (interface{}, error) {
	log.Printf("Executing saga step 1: %s", s.Message)
	// time.Sleep(1 * time.Second)
	return nil, nil
}

func (s SagaStep1) Compensation(ctx tempolite.CompensationContext) (interface{}, error) {
	log.Printf("Compensating saga step 1: %s", s.Message)
	return nil, nil
}

// SagaStep2 implements the SagaStep interface for step 2
type SagaStep2 struct {
	Message string
}

func (s SagaStep2) Transaction(ctx tempolite.TransactionContext) (interface{}, error) {
	log.Printf("Executing saga step 2: %s", s.Message)
	// time.Sleep(1 * time.Second)
	// Simulate a failure in step 2
	// 70% chance of failure
	// if rand.Float32() < 0.7 {
	// 	return nil, fmt.Errorf("simulated failure in step 2")
	// }
	return nil, nil
}

func (s SagaStep2) Compensation(ctx tempolite.CompensationContext) (interface{}, error) {
	log.Printf("Compensating saga step 2: %s", s.Message)
	return nil, nil
}

func main() {
	// Open SQLite database
	db, err := sql.Open("sqlite3", "tempolite.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create Tempolite instance
	tp, err := tempolite.New(context.Background(), db,
		tempolite.WithHandlerWorkers(2),
		tempolite.WithSagaWorkers(2),
		tempolite.WithCompensationWorkers(2),
		tempolite.WithSideEffectWorkers(2),
	)
	if err != nil {
		log.Fatalf("Failed to create Tempolite instance: %v", err)
	}
	defer tp.Close()

	tp.RegisterHandler(SagaHandler)
	tp.RegisterHandler(SimpleHandler)

	// Create a simple task
	simpleTaskID, err := tp.Enqueue(context.Background(), SimpleHandler, SimpleTask{Message: "Hello, Tempolite!"})
	if err != nil {
		log.Fatalf("Failed to enqueue simple task: %v", err)
	}
	log.Printf("Enqueued simple task with ID: %s", simpleTaskID)

	// Create a saga
	sagaTaskID, err := tp.EnqueueSaga(context.Background(), saga, SagaTask{StepMessage: "Saga steps"})
	if err != nil {
		log.Fatalf("Failed to enqueue saga: %v", err)
	}
	log.Printf("Enqueued saga with ID: %s", sagaTaskID)

	value, err := tp.WaitForCompletion(context.Background(), simpleTaskID)
	if err != nil {
		log.Fatalf("Failed to wait for task completion: %v", err)
	}
	log.Printf("Task completed with value: %v", value)

	// Wait for tasks to complete
	err = tp.Wait(func(info tempolite.TempoliteInfo) bool {
		return info.IsCompleted()
	}, 500*time.Millisecond)
	if err != nil {
		log.Fatalf("Error waiting for tasks to complete: %v", err)
	}

	// Get results
	simpleTaskResult, err := tp.GetInfo(context.Background(), simpleTaskID)
	if err != nil {
		log.Fatalf("Failed to get simple task result: %v", err)
	}
	log.Printf("Simple task result: %+v", simpleTaskResult)

	sagaTaskResult, err := tp.GetInfo(context.Background(), sagaTaskID)
	if err != nil {
		log.Fatalf("Failed to get saga task result: %v", err)
	}
	log.Printf("Saga task result: %+v", sagaTaskResult)

	tp.PrintExecutionTree(context.Background(), simpleTaskID)
	tp.PrintExecutionTree(context.Background(), sagaTaskID)

	// // Print execution tree
	executionTree, err := tp.GetExecutionTree(context.Background(), simpleTaskID)
	if err != nil {
		log.Fatalf("Failed to get execution tree: %v", err)
	}
	log.Printf("Execution tree for simple task:\n%s", executionTree.String())

	executionTree, err = tp.GetExecutionTree(context.Background(), sagaTaskID)
	if err != nil {
		log.Fatalf("Failed to get execution tree: %v", err)
	}
	log.Printf("Execution tree for saga task:\n%s", executionTree.String())
}