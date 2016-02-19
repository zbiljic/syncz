// Copyright 2016 Nemanja Zbiljic
//

package striped

import (
	"fmt"
	"hash/crc32"
	"sync"
)

type StripedMutex struct {
	/////////////////////////
	// Constant data
	/////////////////////////

	// INVARIANT: stripesCount > 0
	stripesCount uint32

	// Array of mutexes
	//
	// INVARIANT: len(array) >= stripesCount
	// INVARIANT: Each element is of type sync.Mutex
	array []sync.Mutex
}

func NewStripedMutex(stripes int) *StripedMutex {
	// INVARIANT: stripesCount > 0
	if !(stripes > 0) {
		panic(fmt.Sprintf("Invalid stripes count: %v", stripes))
	}

	return &StripedMutex{
		stripesCount: nextPowerOfTwo(uint32(stripes)),
		array:        make([]sync.Mutex, stripes),
	}
}

// Returns the stripe that corresponds to the passed key. It is always
// guaranteed that if key1 == key2, then striped.Get(key1) == striped.Get(key2).
func (s *StripedMutex) Get(key string) *sync.Mutex {

	hash := crc32.ChecksumIEEE([]byte(key))

	// This will be skewed toward smaller stripe indexes if numstripes does not divide 256.
	stripeIndex := hash % s.stripesCount

	return &s.array[stripeIndex]
}

// Returns the stripe at the specified index. Valid indexes are 0, inclusively,
// to Len(), exclusively.
func (s *StripedMutex) GetAt(index int) *sync.Mutex {

	return &s.array[index]
}

// Returns the total number of stripes in this instance.
func (s *StripedMutex) Len() int {

	return int(s.stripesCount)
}
