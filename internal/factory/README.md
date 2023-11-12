# Factory linter

[![CI](https://github.com/maranqz/go-factory-lint/actions/workflows/ci.yml/badge.svg)](https://github.com/maranqz/go-factory-lint/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/maranqz/go-factory-lint)](https://goreportcard.com/report/github.com/maranqz/go-factory-lint?dummy=unused)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Coverage Status](https://coveralls.io/repos/github/maranqz/go-factory-lint/badge.svg?branch=main)](https://coveralls.io/github/maranqz/go-factory-lint?branch=main)

The linter checks that the Structes are created by the Factory, and not directly.

The checking helps to provide invariants without exclusion and helps avoid creating an invalid object.


## Use

Installation

    go install github.com/maranqz/go-factory-lint/cmd/go-factory-lint@latest

### Options

- `-b`, `--blockedPkgs` - list of packages, where the structures should be created by factories. By default, all structures in all packages should be created by factories, [tests](testdata/src/factory/blockedPkgs).
    - `-ob`, `onlyBlockedPkgs` - only blocked packages should use factory to initiate struct, [tests](testdata/src/factory/onlyBlockedPkgs).

## Example

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
package main

import (
	"fmt"

	"bad"
)

func main() {
	// Use factory for bad.User
	u := &bad.User{
		ID: -1,
	}

	fmt.Println(u.ID) // -1
	fmt.Println(u.CreatedAt) // time.Time{}
}

```

```go
package bad

import "time"

type User struct {
	ID        int64
	CreatedAt time.Time
}

var sequenceID = int64(0)

func NextID() int64 {
	sequenceID++

	return sequenceID
}


```

</td><td>

```go
package main

import (
	"fmt"

	"good"
)

func main() {
	u := good.NewUser()
	
	fmt.Println(u.ID)        // auto increment
	fmt.Println(u.CreatedAt) // time.Now()
}

```

```go
package user

import "time"

type User struct {
	ID        int64
	CreatedAt time.Time
}

func NewUser() *User {
	return &User{
		ID: nextID(),
		CreatedAt: time.Now(),
	}
}

var sequenceID = int64(0)

func nextID() int64 {
	sequenceID++

	return sequenceID
}

```

</td></tr>
</tbody></table>

## TODO

### Feature, heavy implementable but not planned

1. Type assertion, type declaration and type underlying, [tests](testdata/src/factory/default/type_nested.go.skip).