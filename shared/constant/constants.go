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

type Assistant string      // Always "assistant"
type ChatCompletion string // Always "chat.completion"
type Custom string         // Always "custom"
type Disabled string       // Always "disabled"
type Duration string       // Always "duration"
type Enabled string        // Always "enabled"
type Function string       // Always "function"
type Tokens string         // Always "tokens"
type URLCitation string    // Always "url_citation"

func (c Assistant) Default() Assistant           { return "assistant" }
func (c ChatCompletion) Default() ChatCompletion { return "chat.completion" }
func (c Custom) Default() Custom                 { return "custom" }
func (c Disabled) Default() Disabled             { return "disabled" }
func (c Duration) Default() Duration             { return "duration" }
func (c Enabled) Default() Enabled               { return "enabled" }
func (c Function) Default() Function             { return "function" }
func (c Tokens) Default() Tokens                 { return "tokens" }
func (c URLCitation) Default() URLCitation       { return "url_citation" }

func (c Assistant) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ChatCompletion) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Custom) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Disabled) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Duration) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Enabled) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Function) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Tokens) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c URLCitation) MarshalJSON() ([]byte, error)    { return marshalString(c) }

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
