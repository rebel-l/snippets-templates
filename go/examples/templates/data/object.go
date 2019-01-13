package data

import (
	"math/rand"
	"time"
)

// Objects ...
type Objects []string

// String ...
func (o Objects) String() string {
	if len(o) == 0 {
		return ""
	}

	if len(o) == 1 {
		return o[0]
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(o))
	return o[n]
}

// ListObjects contains a list of objects
var ListObjects = Objects{
	"Village",
	"Town",
	"City",
	"Country",
	"World",
	"Galaxy",
}
