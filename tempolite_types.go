package tempolite

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

type HandlerInfo struct {
	Handler    interface{}
	ParamType  reflect.Type
	ReturnType reflect.Type
	NumIn      int
	NumOut     int
}

func (hi HandlerInfo) GetFn() reflect.Value {
	return reflect.ValueOf(hi.Handler)
}

func (hi HandlerInfo) ToInterface(data []byte) (interface{}, error) {
	handlerValue := reflect.ValueOf(hi.Handler)
	paramType := handlerValue.Type().In(1)
	param := reflect.New(paramType).Interface()
	err := json.Unmarshal(data, &param)
	if err != nil {
		log.Printf("Failed to unmarshal task payload: %v", err)
		return nil, fmt.Errorf("failed to unmarshal task payload: %v", err)
	}
	return param, nil
}

type SageInfo struct{}

type EnqueueOption func(*enqueueOptions)

type enqueueOptions struct {
	maxDuration        time.Duration
	timeLimit          time.Duration
	immediate          bool
	panicOnTimeout     bool
	parentID           *string
	nodeID             *string
	executionContextID *string
	index              *int
}

func WithIndex(index int) EnqueueOption {
	return func(o *enqueueOptions) {
		log.Printf("Setting index for enqueue option: %d", index)
		o.index = &index
	}
}

func WithExecutionContextID(id string) EnqueueOption {
	return func(o *enqueueOptions) {
		log.Printf("Setting execution context ID for enqueue option: %s", id)
		o.executionContextID = &id
	}
}

func WithParentID(id string) EnqueueOption {
	return func(o *enqueueOptions) {
		log.Printf("Setting parent ID for enqueue option: %s", id)
		o.parentID = &id
	}
}

func WithNodeID(id string) EnqueueOption {
	return func(o *enqueueOptions) {
		log.Printf("Setting node ID for enqueue option: %s", id)
		o.nodeID = &id
	}
}

func WithMaxDuration(duration time.Duration) EnqueueOption {
	return func(o *enqueueOptions) {
		log.Printf("Setting max duration for enqueue option: %v", duration)
		o.maxDuration = duration
	}
}

func WithTimeLimit(limit time.Duration) EnqueueOption {
	return func(o *enqueueOptions) {
		log.Printf("Setting time limit for enqueue option: %v", limit)
		o.timeLimit = limit
	}
}

func WithImmediateRetry() EnqueueOption {
	return func(o *enqueueOptions) {
		log.Printf("Enabling immediate retry for enqueue option")
		o.immediate = true
	}
}

func WithPanicOnTimeout() EnqueueOption {
	return func(o *enqueueOptions) {
		log.Printf("Enabling panic on timeout for enqueue option")
		o.panicOnTimeout = true
	}
}
