// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/davidroman0O/go-tempolite/ent/executionunit"
	"github.com/davidroman0O/go-tempolite/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	executionunitFields := schema.ExecutionUnit{}.Fields()
	_ = executionunitFields
	// executionunitDescRetryCount is the schema descriptor for retry_count field.
	executionunitDescRetryCount := executionunitFields[5].Descriptor()
	// executionunit.DefaultRetryCount holds the default value on creation for the retry_count field.
	executionunit.DefaultRetryCount = executionunitDescRetryCount.Default.(int)
	// executionunitDescMaxRetries is the schema descriptor for max_retries field.
	executionunitDescMaxRetries := executionunitFields[6].Descriptor()
	// executionunit.DefaultMaxRetries holds the default value on creation for the max_retries field.
	executionunit.DefaultMaxRetries = executionunitDescMaxRetries.Default.(int)
}
