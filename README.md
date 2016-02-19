# syncz

[![GoDoc](https://godoc.org/github.com/zbiljic/syncz?status.svg)](https://godoc.org/github.com/zbiljic/syncz)

A library of concurrent data structures and synchronization mechanisms for Go.

It contains code intended to supplement the [sync][] package from the Go standard library.

### NOTE: only tested with Go 1.5+.

#### Pool

Package contains interface for a generic Poll of resources.

#### Striped

Striped allows the programmer to select a number of locks, which are distributed
between keys based on their hash code. This allows the programmer to dynamically
select a tradeoff between concurrency and memory consumption, while retaining
the key invariant that if key1 == key2, then
striped.Get(key1) == striped.Get(key2).

Striped contains implementation for Mutex, and RWMutex.

The values used to retreive a lock cannot be nil.

The number should be a factor of 256 for even striped distribution.


### Installation

 1. Install Go 1.5 or higher.
 2. Run `go get github.com/zbiljic/syncz/...`

### Updating

When new code is merged to master, you can use

	go get -u github.com/zbiljic/syncz/...

To retrieve the latest version of syncz.

### Testing

To run all the unit tests use these commands:

	cd $GOPATH/src/github.com/zbiljic/syncz
	go get -t -u ./...
	go test ./...

Once you've done this once, you can simply use this command to run all unit tests:

	go test ./...


See the [reference][] for more info.

[sync]: http://godoc.org/sync
[reference]: http://godoc.org/github.com/zbiljic/syncz

---

Copyright © 2016 Nemanja Zbiljić
