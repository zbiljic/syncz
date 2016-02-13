// Copyright 2016 Nemanja Zbiljic
//

// Package syncz contains a collection of libraries supplementing Go's sync package
//
// For more information about the syncz package, see the README at
//
//	http://github.com/zbiljic/syncz
//
package syncz

import (
	_ "github.com/zbiljic/syncz/pool"
	_ "github.com/zbiljic/syncz/striped"
)
