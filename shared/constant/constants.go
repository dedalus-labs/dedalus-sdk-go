// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/dedalus-labs/dedalus-sdk-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Assistant string   // Always "assistant"
type Custom string      // Always "custom"
type Function string    // Always "function"
type URLCitation string // Always "url_citation"

func (c Assistant) Default() Assistant     { return "assistant" }
func (c Custom) Default() Custom           { return "custom" }
func (c Function) Default() Function       { return "function" }
func (c URLCitation) Default() URLCitation { return "url_citation" }

func (c Assistant) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Custom) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Function) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c URLCitation) MarshalJSON() ([]byte, error) { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
