package tempolite

import (
	"fmt"
	"reflect"
)

type SagaStep[T Identifier] interface {
	Transaction(ctx TransactionContext[T]) (interface{}, error)
	Compensation(ctx CompensationContext[T]) (interface{}, error)
}

type SagaDefinition[T Identifier] struct {
	Steps       []SagaStep[T]
	HandlerInfo *SagaHandlerInfo
}

type SagaDefinitionBuilder[T Identifier] struct {
	steps []SagaStep[T]
}

type SagaHandlerInfo struct {
	TransactionInfo  []HandlerInfo
	CompensationInfo []HandlerInfo
}

// NewSaga creates a new builder instance.
func NewSaga[T Identifier]() *SagaDefinitionBuilder[T] {
	return &SagaDefinitionBuilder[T]{
		steps: make([]SagaStep[T], 0),
	}
}

// AddStep adds a saga step to the builder.
func (b *SagaDefinitionBuilder[T]) AddStep(step SagaStep[T]) *SagaDefinitionBuilder[T] {
	b.steps = append(b.steps, step)
	return b
}

// analyzeMethod helper function to create HandlerInfo for a method
func analyzeMethod(method reflect.Method, name string) (HandlerInfo, error) {
	methodType := method.Type

	if methodType.NumIn() < 2 {
		return HandlerInfo{}, fmt.Errorf("method must have at least two parameters (receiver and context)")
	}

	paramTypes := make([]reflect.Type, methodType.NumIn()-2)
	paramKinds := make([]reflect.Kind, methodType.NumIn()-2)
	for i := 2; i < methodType.NumIn(); i++ {
		paramTypes[i-2] = methodType.In(i)
		paramKinds[i-2] = methodType.In(i).Kind()
	}

	returnTypes := make([]reflect.Type, methodType.NumOut()-1)
	returnKinds := make([]reflect.Kind, methodType.NumOut()-1)
	for i := 0; i < methodType.NumOut()-1; i++ {
		returnTypes[i] = methodType.Out(i)
		returnKinds[i] = methodType.Out(i).Kind()
	}

	handlerName := fmt.Sprintf("%s.%s", name, method.Name)

	return HandlerInfo{
		HandlerName:     handlerName,
		HandlerLongName: HandlerIdentity(name),
		Handler:         method.Func.Interface(),
		ParamTypes:      paramTypes,
		ParamsKinds:     paramKinds,
		ReturnTypes:     returnTypes,
		ReturnKinds:     returnKinds,
		NumIn:           methodType.NumIn() - 2,  // Exclude receiver and context
		NumOut:          methodType.NumOut() - 1, // Exclude error
	}, nil
}

// Build creates a SagaDefinition with the HandlerInfo included.
func (b *SagaDefinitionBuilder[T]) Build() (*SagaDefinition[T], error) {
	sagaInfo := &SagaHandlerInfo{
		TransactionInfo:  make([]HandlerInfo, len(b.steps)),
		CompensationInfo: make([]HandlerInfo, len(b.steps)),
	}

	for i, step := range b.steps {
		stepType := reflect.TypeOf(step)
		if stepType.Kind() == reflect.Ptr {
			stepType = stepType.Elem()
		}

		transactionMethod, ok := stepType.MethodByName("Transaction")
		if !ok {
			return nil, fmt.Errorf("Transaction method not found for step %d", i)
		}
		compensationMethod, ok := stepType.MethodByName("Compensation")
		if !ok {
			return nil, fmt.Errorf("Compensation method not found for step %d", i)
		}

		transactionInfo, err := analyzeMethod(transactionMethod, stepType.Name())
		if err != nil {
			return nil, fmt.Errorf("error analyzing Transaction method for step %d: %w", i, err)
		}

		compensationInfo, err := analyzeMethod(compensationMethod, stepType.Name())
		if err != nil {
			return nil, fmt.Errorf("error analyzing Compensation method for step %d: %w", i, err)
		}

		sagaInfo.TransactionInfo[i] = transactionInfo
		sagaInfo.CompensationInfo[i] = compensationInfo
	}

	return &SagaDefinition[T]{
		Steps:       b.steps,
		HandlerInfo: sagaInfo,
	}, nil
}

func (w WorkflowContext[T]) Saga(stepID T, saga *SagaDefinition[T]) *SagaInfo[T] {
	if err := w.checkIfPaused(); err != nil {
		w.tp.logger.Debug(w.tp.ctx, "Workflow is paused, skipping Saga execution", "stepID", stepID, "error", err)
		return &SagaInfo[T]{err: err}
	}
	// Enqueue the saga for execution
	return w.tp.saga(w, stepID, saga)
}
