// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	shimjson "github.com/dedalus-labs/dedalus-sdk-go/internal/encoding/json"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/respjson"
)

// EmbeddingService contains methods and other services that help with interacting
// with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddingService] method instead.
type EmbeddingService struct {
	Options []option.RequestOption
}

// NewEmbeddingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddingService(opts ...option.RequestOption) (r EmbeddingService) {
	r = EmbeddingService{}
	r.Options = opts
	return
}

// Create embeddings using the configured provider.
func (r *EmbeddingService) New(ctx context.Context, body EmbeddingNewParams, opts ...option.RequestOption) (res *CreateEmbeddingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fields:
//
//   - input (required): str | Annotated[list[str], MinLen(1), MaxLen(2048)] |
//     Annotated[list[int], MinLen(1), MaxLen(2048)] |
//     Annotated[list[Annotated[list[int], MinLen(1)]], MinLen(1), MaxLen(2048)]
//   - model (required): str | Literal['text-embedding-ada-002',
//     'text-embedding-3-small', 'text-embedding-3-large']
//   - encoding_format (optional): Literal['float', 'base64']
//   - dimensions (optional): int
//   - user (optional): str
//
// The properties Input, Model are required.
type CreateEmbeddingRequestParam struct {
	// Input text to embed, encoded as a string or array of tokens. To embed multiple
	// inputs in a single request, pass an array of strings or array of token arrays.
	// The input must not exceed the max input tokens for the model (8192 tokens for
	// all embedding models), cannot be an empty string, and any array must be 2048
	// dimensions or less.
	// [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken)
	// for counting tokens. In addition to the per-input token limit, all embedding
	// models enforce a maximum of 300,000 tokens summed across all inputs in a single
	// request.
	Input CreateEmbeddingRequestInputUnionParam `json:"input,omitzero,required"`
	// ID of the model to use. You can use the
	// [List models](https://platform.openai.com/docs/api-reference/models/list) API to
	// see all of your available models, or see our
	// [Model overview](https://platform.openai.com/docs/models) for descriptions of
	// them.
	Model CreateEmbeddingRequestModel `json:"model,omitzero,required"`
	// The number of dimensions the resulting output embeddings should have. Only
	// supported in `text-embedding-3` and later models.
	Dimensions param.Opt[int64] `json:"dimensions,omitzero"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor
	// and detect abuse.
	// [Learn more](https://platform.openai.com/docs/guides/safety-best-practices#end-user-ids).
	User param.Opt[string] `json:"user,omitzero"`
	// The format to return the embeddings in. Can be either `float` or
	// [`base64`](https://pypi.org/project/pybase64/).
	//
	// Any of "float", "base64".
	EncodingFormat CreateEmbeddingRequestEncodingFormat `json:"encoding_format,omitzero"`
	paramObj
}

func (r CreateEmbeddingRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateEmbeddingRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateEmbeddingRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreateEmbeddingRequestInputUnionParam struct {
	OfString          param.Opt[string] `json:",omitzero,inline"`
	OfStringArray     []string          `json:",omitzero,inline"`
	OfIntArray        []int64           `json:",omitzero,inline"`
	OfArrayOfIntArray [][]int64         `json:",omitzero,inline"`
	paramUnion
}

func (u CreateEmbeddingRequestInputUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfIntArray, u.OfArrayOfIntArray)
}
func (u *CreateEmbeddingRequestInputUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CreateEmbeddingRequestInputUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfIntArray) {
		return &u.OfIntArray
	} else if !param.IsOmitted(u.OfArrayOfIntArray) {
		return &u.OfArrayOfIntArray
	}
	return nil
}

// ID of the model to use. You can use the
// [List models](https://platform.openai.com/docs/api-reference/models/list) API to
// see all of your available models, or see our
// [Model overview](https://platform.openai.com/docs/models) for descriptions of
// them.
type CreateEmbeddingRequestModel string

const (
	CreateEmbeddingRequestModelTextEmbeddingAda002 CreateEmbeddingRequestModel = "text-embedding-ada-002"
	CreateEmbeddingRequestModelTextEmbedding3Small CreateEmbeddingRequestModel = "text-embedding-3-small"
	CreateEmbeddingRequestModelTextEmbedding3Large CreateEmbeddingRequestModel = "text-embedding-3-large"
)

// The format to return the embeddings in. Can be either `float` or
// [`base64`](https://pypi.org/project/pybase64/).
type CreateEmbeddingRequestEncodingFormat string

const (
	CreateEmbeddingRequestEncodingFormatFloat  CreateEmbeddingRequestEncodingFormat = "float"
	CreateEmbeddingRequestEncodingFormatBase64 CreateEmbeddingRequestEncodingFormat = "base64"
)

// Response from embeddings endpoint.
type CreateEmbeddingResponse struct {
	// List of embedding objects
	Data []CreateEmbeddingResponseData `json:"data,required"`
	// The model used for embeddings
	Model string `json:"model,required"`
	// Usage statistics (prompt_tokens, total_tokens)
	Usage map[string]int64 `json:"usage,required"`
	// Object type, always 'list'
	//
	// Any of "list".
	Object CreateEmbeddingResponseObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Model       respjson.Field
		Usage       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateEmbeddingResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateEmbeddingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single embedding object.
type CreateEmbeddingResponseData struct {
	// The embedding vector (float array or base64 string)
	Embedding CreateEmbeddingResponseDataEmbeddingUnion `json:"embedding,required"`
	// Index of the embedding in the list
	Index int64 `json:"index,required"`
	// Object type, always 'embedding'
	//
	// Any of "embedding".
	Object string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embedding   respjson.Field
		Index       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateEmbeddingResponseData) RawJSON() string { return r.JSON.raw }
func (r *CreateEmbeddingResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CreateEmbeddingResponseDataEmbeddingUnion contains all possible properties and
// values from [[]float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloatArray OfString]
type CreateEmbeddingResponseDataEmbeddingUnion struct {
	// This field will be present if the value is a [[]float64] instead of an object.
	OfFloatArray []float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloatArray respjson.Field
		OfString     respjson.Field
		raw          string
	} `json:"-"`
}

func (u CreateEmbeddingResponseDataEmbeddingUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreateEmbeddingResponseDataEmbeddingUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CreateEmbeddingResponseDataEmbeddingUnion) RawJSON() string { return u.JSON.raw }

func (r *CreateEmbeddingResponseDataEmbeddingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object type, always 'list'
type CreateEmbeddingResponseObject string

const (
	CreateEmbeddingResponseObjectList CreateEmbeddingResponseObject = "list"
)

type EmbeddingNewParams struct {
	// Fields:
	//
	//   - input (required): str | Annotated[list[str], MinLen(1), MaxLen(2048)] |
	//     Annotated[list[int], MinLen(1), MaxLen(2048)] |
	//     Annotated[list[Annotated[list[int], MinLen(1)]], MinLen(1), MaxLen(2048)]
	//   - model (required): str | Literal['text-embedding-ada-002',
	//     'text-embedding-3-small', 'text-embedding-3-large']
	//   - encoding_format (optional): Literal['float', 'base64']
	//   - dimensions (optional): int
	//   - user (optional): str
	CreateEmbeddingRequest CreateEmbeddingRequestParam
	paramObj
}

func (r EmbeddingNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateEmbeddingRequest)
}
func (r *EmbeddingNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateEmbeddingRequest)
}
