package apiform

import (
	"bytes"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"io"
	"mime/multipart"
	"strings"
	"testing"
	"time"
)

func P[T any](v T) *T { return &v }

type Primitives struct {
	A bool    `form:"a"`
	B int     `form:"b"`
	C uint    `form:"c"`
	D float64 `form:"d"`
	E float32 `form:"e"`
	F []int   `form:"f"`
}

// These aliases are necessary to bypass the cache.
// This only relevant during testing.
type int_ int
type PrimitivesBrackets struct {
	F []int_ `form:"f"`
}

type PrimitivePointers struct {
	A *bool    `form:"a"`
	B *int     `form:"b"`
	C *uint    `form:"c"`
	D *float64 `form:"d"`
	E *float32 `form:"e"`
	F *[]int   `form:"f"`
}

type Slices struct {
	Slice []Primitives `form:"slices"`
}

type DateTime struct {
	Date     time.Time `form:"date" format:"date"`
	DateTime time.Time `form:"date-time" format:"date-time"`
}

type AdditionalProperties struct {
	A      bool           `form:"a"`
	Extras map[string]any `form:"-,extras"`
}

type TypedAdditionalProperties struct {
	A      bool           `form:"a"`
	Extras map[string]int `form:"-,extras"`
}

type EmbeddedStructs struct {
	AdditionalProperties
	A      *int           `form:"number2"`
	Extras map[string]any `form:"-,extras"`
}

type Recursive struct {
	Name  string     `form:"name"`
	Child *Recursive `form:"child"`
}

type UnknownStruct struct {
	Unknown any `form:"unknown"`
}

type UnionStruct struct {
	Union Union `form:"union" format:"date"`
}

type Union interface {
	union()
}

type UnionInteger int64

func (UnionInteger) union() {}

type UnionStructA struct {
	Type string `form:"type"`
	A    string `form:"a"`
	B    string `form:"b"`
}

func (UnionStructA) union() {}

type UnionStructB struct {
	Type string `form:"type"`
	A    string `form:"a"`
}

func (UnionStructB) union() {}

type UnionTime time.Time

func (UnionTime) union() {}

type ReaderStruct struct {
	File io.Reader `form:"file"`
}

type NamedEnum string

const NamedEnumFoo NamedEnum = "foo"

type StructUnionWrapper struct {
	Union StructUnion `form:"union"`
}

type StructUnion struct {
	OfInt    param.Opt[int64]     `form:",omitzero,inline"`
	OfString param.Opt[string]    `form:",omitzero,inline"`
	OfEnum   param.Opt[NamedEnum] `form:",omitzero,inline"`
	OfA      UnionStructA         `form:",omitzero,inline"`
	OfB      UnionStructB         `form:",omitzero,inline"`
	param.APIUnion
}

var tests = map[string]struct {
	buf string
	val any
}{
	"file": {
		buf: `--xxx
Content-Disposition: form-data; name="file"; filename="anonymous_file"
Content-Type: application/octet-stream

some file contents...
--xxx--
`,
		val: ReaderStruct{
			File: io.Reader(bytes.NewBuffer([]byte("some file contents..."))),
		},
	},
	"map_string": {
		`--xxx
Content-Disposition: form-data; name="foo"

bar
--xxx--
`,
		map[string]string{"foo": "bar"},
	},

	"map_interface": {
		`--xxx
Content-Disposition: form-data; name="a"

1
--xxx
Content-Disposition: form-data; name="b"

str
--xxx
Content-Disposition: form-data; name="c"

false
--xxx--
`,
		map[string]any{"a": float64(1), "b": "str", "c": false},
	},

	"primitive_struct": {
		`--xxx
Content-Disposition: form-data; name="a"

false
--xxx
Content-Disposition: form-data; name="b"

237628372683
--xxx
Content-Disposition: form-data; name="c"

654
--xxx
Content-Disposition: form-data; name="d"

9999.43
--xxx
Content-Disposition: form-data; name="e"

43.76
--xxx
Content-Disposition: form-data; name="f.0"

1
--xxx
Content-Disposition: form-data; name="f.1"

2
--xxx
Content-Disposition: form-data; name="f.2"

3
--xxx
Content-Disposition: form-data; name="f.3"

4
--xxx--
`,
		Primitives{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
	},
	"primitive_struct,brackets": {
		`--xxx
Content-Disposition: form-data; name="f[]"

1
--xxx
Content-Disposition: form-data; name="f[]"

2
--xxx
Content-Disposition: form-data; name="f[]"

3
--xxx
Content-Disposition: form-data; name="f[]"

4
--xxx--
`,
		PrimitivesBrackets{F: []int_{1, 2, 3, 4}},
	},

	"slices": {
		`--xxx
Content-Disposition: form-data; name="slices.0.a"

false
--xxx
Content-Disposition: form-data; name="slices.0.b"

237628372683
--xxx
Content-Disposition: form-data; name="slices.0.c"

654
--xxx
Content-Disposition: form-data; name="slices.0.d"

9999.43
--xxx
Content-Disposition: form-data; name="slices.0.e"

43.76
--xxx
Content-Disposition: form-data; name="slices.0.f.0"

1
--xxx
Content-Disposition: form-data; name="slices.0.f.1"

2
--xxx
Content-Disposition: form-data; name="slices.0.f.2"

3
--xxx
Content-Disposition: form-data; name="slices.0.f.3"

4
--xxx--
`,
		Slices{
			Slice: []Primitives{{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}},
		},
	},
	"primitive_pointer_struct": {
		`--xxx
Content-Disposition: form-data; name="a"

false
--xxx
Content-Disposition: form-data; name="b"

237628372683
--xxx
Content-Disposition: form-data; name="c"

654
--xxx
Content-Disposition: form-data; name="d"

9999.43
--xxx
Content-Disposition: form-data; name="e"

43.76
--xxx
Content-Disposition: form-data; name="f.0"

1
--xxx
Content-Disposition: form-data; name="f.1"

2
--xxx
Content-Disposition: form-data; name="f.2"

3
--xxx
Content-Disposition: form-data; name="f.3"

4
--xxx
Content-Disposition: form-data; name="f.4"

5
--xxx--
`,
		PrimitivePointers{
			A: P(false),
			B: P(237628372683),
			C: P(uint(654)),
			D: P(9999.43),
			E: P(float32(43.76)),
			F: &[]int{1, 2, 3, 4, 5},
		},
	},

	"datetime_struct": {
		`--xxx
Content-Disposition: form-data; name="date"

2006-01-02
--xxx
Content-Disposition: form-data; name="date-time"

2006-01-02T15:04:05Z
--xxx--
`,
		DateTime{
			Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		},
	},

	"additional_properties": {
		`--xxx
Content-Disposition: form-data; name="a"

true
--xxx
Content-Disposition: form-data; name="bar"

value
--xxx
Content-Disposition: form-data; name="foo"

true
--xxx--
`,
		AdditionalProperties{
			A: true,
			Extras: map[string]any{
				"bar": "value",
				"foo": true,
			},
		},
	},

	"recursive_struct": {
		`--xxx
Content-Disposition: form-data; name="child.name"

Alex
--xxx
Content-Disposition: form-data; name="name"

Robert
--xxx--
`,
		Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}},
	},

	"unknown_struct_number": {
		`--xxx
Content-Disposition: form-data; name="unknown"

12
--xxx--
`,
		UnknownStruct{
			Unknown: 12.,
		},
	},

	"unknown_struct_map": {
		`--xxx
Content-Disposition: form-data; name="unknown.foo"

bar
--xxx--
`,
		UnknownStruct{
			Unknown: map[string]any{
				"foo": "bar",
			},
		},
	},

	"struct_union_integer": {
		`--xxx
Content-Disposition: form-data; name="union"

12
--xxx--
`,
		StructUnionWrapper{
			Union: StructUnion{OfInt: param.NewOpt[int64](12)},
		},
	},

	"union_integer": {
		`--xxx
Content-Disposition: form-data; name="union"

12
--xxx--
`,
		UnionStruct{
			Union: UnionInteger(12),
		},
	},

	"struct_union_struct_discriminated_a": {
		`--xxx
Content-Disposition: form-data; name="union.a"

foo
--xxx
Content-Disposition: form-data; name="union.b"

bar
--xxx
Content-Disposition: form-data; name="union.type"

typeA
--xxx--
`,
		StructUnionWrapper{
			Union: StructUnion{OfA: UnionStructA{
				Type: "typeA",
				A:    "foo",
				B:    "bar",
			}},
		},
	},

	"union_struct_discriminated_a": {
		`--xxx
Content-Disposition: form-data; name="union.a"

foo
--xxx
Content-Disposition: form-data; name="union.b"

bar
--xxx
Content-Disposition: form-data; name="union.type"

typeA
--xxx--
`,

		UnionStruct{
			Union: UnionStructA{
				Type: "typeA",
				A:    "foo",
				B:    "bar",
			},
		},
	},

	"struct_union_struct_discriminated_b": {
		`--xxx
Content-Disposition: form-data; name="union.a"

foo
--xxx
Content-Disposition: form-data; name="union.type"

typeB
--xxx--
`,
		StructUnionWrapper{
			Union: StructUnion{OfB: UnionStructB{
				Type: "typeB",
				A:    "foo",
			}},
		},
	},

	"union_struct_discriminated_b": {
		`--xxx
Content-Disposition: form-data; name="union.a"

foo
--xxx
Content-Disposition: form-data; name="union.type"

typeB
--xxx--
`,
		UnionStruct{
			Union: UnionStructB{
				Type: "typeB",
				A:    "foo",
			},
		},
	},

	"union_struct_time": {
		`--xxx
Content-Disposition: form-data; name="union"

2010-05-23
--xxx--
`,
		UnionStruct{
			Union: UnionTime(time.Date(2010, 05, 23, 0, 0, 0, 0, time.UTC)),
		},
	},
}

func TestEncode(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			writer := multipart.NewWriter(buf)
			writer.SetBoundary("xxx")

			var arrayFmt string = "indices:dots"
			if tags := strings.Split(name, ","); len(tags) > 1 {
				arrayFmt = tags[1]
			}

			err := MarshalWithSettings(test.val, writer, arrayFmt)
			if err != nil {
				t.Errorf("serialization of %v failed with error %v", test.val, err)
			}
			err = writer.Close()
			if err != nil {
				t.Errorf("serialization of %v failed with error %v", test.val, err)
			}
			raw := buf.Bytes()
			if string(raw) != strings.ReplaceAll(test.buf, "\n", "\r\n") {
				t.Errorf("expected %+#v to serialize to '%s' but got '%s'", test.val, test.buf, string(raw))
			}
		})
	}
}
