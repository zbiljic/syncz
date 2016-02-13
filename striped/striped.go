// Copyright 2016 Nemanja Zbiljic
//

// Striped is a tool for assigning locks to objects in a configurable, flexible manner.
//
// Striped allows the programmer to select a number of locks, which are
// distributed between keys based on their hash code. This allows the programmer
// to dynamically select a tradeoff between concurrency and memory consumption,
// while retaining the key invariant that if key1 == key2,
// then striped.Get(key1) == striped.Get(key2).
package striped

// Returns `v` if it is a power-of-two, or else the next-highest power-of-two.
func nextPowerOfTwo(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
