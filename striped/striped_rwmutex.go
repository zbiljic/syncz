// Copyright 2016 Nemanja Zbiljic
//

package striped

import (
	"fmt"
	"hash/crc32"
	"sync"
)

type StripedRWMutex struct {
	/////////////////////////
	// Constant data
	/////////////////////////

	// INVARIANT: stripesCount > 0
	stripesCount uint32

	// Array of mutexes
	//
	// INVARIANT: len(array) >= stripesCount
	// INVARIANT: Each element is of type sync.RWMutex
	array []sync.RWMutex
}

func NewStripedRWMutex(stripes int) *StripedRWMutex {
	// INVARIANT: stripesCount > 0
	if !(stripes > 0) {
		panic(fmt.Sprintf("Invalid stripes count: %v", stripes))
	}

	return &StripedRWMutex{
		stripesCount: nextPowerOfTwo(uint32(stripes)),
		array:        make([]sync.RWMutex, stripes),
	}
}

// Returns the stripe that corresponds to the passed key. It is always
// guaranteed that if key1 == key2, then striped.Get(key1) == striped.Get(key2).
func (s *StripedRWMutex) Get(key string) *sync.RWMutex {

	hash := crc32.ChecksumIEEE([]byte(key))

	// This will be skewed toward smaller stripe indexes if numstripes does not divide 256.
	stripeIndex := hash % uint32(s.stripesCount)

	return &s.array[stripeIndex]
}

// Returns the stripe at the specified index. Valid indexes are 0, inclusively,
// to Len(), exclusively.
func (s *StripedRWMutex) GetAt(index int) *sync.RWMutex {

	return &s.array[index]
}

// Returns the total number of stripes in this instance.
func (s *StripedRWMutex) Len() int {

	return int(s.stripesCount)
}
