// Copyright 2016 Nemanja Zbiljic
//

// Package pool provides an interface for a generic pooled resource manager.
package pool

// A generic pooled resource manager.
//
type Pool interface {

	// Allocate allocates a single resource from the pool.
	Allocate() interface{}

	// AllocateArray bulk-allocates multiple resources from the pool, and
	// inserts them in the provided array respecting the offset and length
	// parameters.
	AllocateArray(array []interface{}, offset, length int)

	// Free returns a resource into its appropriate pool.
	Free(resource interface{})

	// FreeArray returns multiple resources to the poll from the provided array.
	FreeArray(array []interface{}, offset, length int)
}
