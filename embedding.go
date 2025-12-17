// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/param"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/shared"
	"github.com/tidwall/gjson"
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
func NewEmbeddingService(opts ...option.RequestOption) (r *EmbeddingService) {
	r = &EmbeddingService{}
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

// Schema for CreateEmbeddingRequest.
//
// Fields:
//
//   - input (required): str | Annotated[list[str], MinLen(1), MaxLen(2048),
//     ArrayTitle("CreateEmbeddingRequestInputArray")] | Annotated[list[int],
//     MinLen(1), MaxLen(2048), ArrayTitle("CreateEmbeddingRequestInputArray")] |
//     Annotated[list[Annotated[list[int], MinLen(1),
//     ArrayTitle("CreateEmbeddingRequestInputItemArray")]], MinLen(1), MaxLen(2048),
//     ArrayTitle("CreateEmbeddingRequestInputArray")]
//   - model (required): str | Literal["text-embedding-ada-002",
//     "text-embedding-3-small", "text-embedding-3-large"]
//   - encoding_format (optional): Literal["float", "base64"]
//   - dimensions (optional): int
//   - user (optional): str
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
	Input param.Field[CreateEmbeddingRequestInputUnionParam] `json:"input,required"`
	// ID of the model to use. You can use the
	// [List models](https://platform.openai.com/docs/api-reference/models/list) API to
	// see all of your available models, or see our
	// [Model overview](https://platform.openai.com/docs/models) for descriptions of
	// them.
	Model param.Field[CreateEmbeddingRequestModel] `json:"model,required"`
	// The number of dimensions the resulting output embeddings should have. Only
	// supported in `text-embedding-3` and later models.
	Dimensions param.Field[int64] `json:"dimensions"`
	// The format to return the embeddings in. Can be either `float` or
	// [`base64`](https://pypi.org/project/pybase64/).
	EncodingFormat param.Field[CreateEmbeddingRequestEncodingFormat] `json:"encoding_format"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor
	// and detect abuse.
	// [Learn more](https://platform.openai.com/docs/guides/safety-best-practices#end-user-ids).
	User param.Field[string] `json:"user"`
}

func (r CreateEmbeddingRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Input text to embed, encoded as a string or array of tokens. To embed multiple
// inputs in a single request, pass an array of strings or array of token arrays.
// The input must not exceed the max input tokens for the model (8192 tokens for
// all embedding models), cannot be an empty string, and any array must be 2048
// dimensions or less.
// [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken)
// for counting tokens. In addition to the per-input token limit, all embedding
// models enforce a maximum of 300,000 tokens summed across all inputs in a single
// request.
//
// Satisfied by [shared.UnionString],
// [CreateEmbeddingRequestInputCreateEmbeddingRequestInputArrayParam],
// [CreateEmbeddingRequestInputCreateEmbeddingRequestInputArrayParam],
// [CreateEmbeddingRequestInputCreateEmbeddingRequestInputArrayParam].
type CreateEmbeddingRequestInputUnionParam interface {
	ImplementsCreateEmbeddingRequestInputUnionParam()
}

type CreateEmbeddingRequestInputCreateEmbeddingRequestInputArrayParam []string

func (r CreateEmbeddingRequestInputCreateEmbeddingRequestInputArrayParam) ImplementsCreateEmbeddingRequestInputUnionParam() {
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

func (r CreateEmbeddingRequestModel) IsKnown() bool {
	switch r {
	case CreateEmbeddingRequestModelTextEmbeddingAda002, CreateEmbeddingRequestModelTextEmbedding3Small, CreateEmbeddingRequestModelTextEmbedding3Large:
		return true
	}
	return false
}

// The format to return the embeddings in. Can be either `float` or
// [`base64`](https://pypi.org/project/pybase64/).
type CreateEmbeddingRequestEncodingFormat string

const (
	CreateEmbeddingRequestEncodingFormatFloat  CreateEmbeddingRequestEncodingFormat = "float"
	CreateEmbeddingRequestEncodingFormatBase64 CreateEmbeddingRequestEncodingFormat = "base64"
)

func (r CreateEmbeddingRequestEncodingFormat) IsKnown() bool {
	switch r {
	case CreateEmbeddingRequestEncodingFormatFloat, CreateEmbeddingRequestEncodingFormatBase64:
		return true
	}
	return false
}

// Response from embeddings endpoint.
type CreateEmbeddingResponse struct {
	// List of embedding objects
	Data []CreateEmbeddingResponseData `json:"data,required"`
	// The model used for embeddings
	Model string `json:"model,required"`
	// Usage statistics (prompt_tokens, total_tokens)
	Usage map[string]int64 `json:"usage,required"`
	// Object type, always 'list'
	Object CreateEmbeddingResponseObject `json:"object"`
	JSON   createEmbeddingResponseJSON   `json:"-"`
}

// createEmbeddingResponseJSON contains the JSON metadata for the struct
// [CreateEmbeddingResponse]
type createEmbeddingResponseJSON struct {
	Data        apijson.Field
	Model       apijson.Field
	Usage       apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateEmbeddingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createEmbeddingResponseJSON) RawJSON() string {
	return r.raw
}

// Single embedding object.
type CreateEmbeddingResponseData struct {
	// The embedding vector (float array or base64 string)
	Embedding CreateEmbeddingResponseDataEmbeddingUnion `json:"embedding,required"`
	// Index of the embedding in the list
	Index int64 `json:"index,required"`
	// Object type, always 'embedding'
	Object CreateEmbeddingResponseDataObject `json:"object"`
	JSON   createEmbeddingResponseDataJSON   `json:"-"`
}

// createEmbeddingResponseDataJSON contains the JSON metadata for the struct
// [CreateEmbeddingResponseData]
type createEmbeddingResponseDataJSON struct {
	Embedding   apijson.Field
	Index       apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateEmbeddingResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createEmbeddingResponseDataJSON) RawJSON() string {
	return r.raw
}

// The embedding vector (float array or base64 string)
//
// Union satisfied by [CreateEmbeddingResponseDataEmbeddingArray] or
// [shared.UnionString].
type CreateEmbeddingResponseDataEmbeddingUnion interface {
	ImplementsCreateEmbeddingResponseDataEmbeddingUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CreateEmbeddingResponseDataEmbeddingUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreateEmbeddingResponseDataEmbeddingArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CreateEmbeddingResponseDataEmbeddingArray []float64

func (r CreateEmbeddingResponseDataEmbeddingArray) ImplementsCreateEmbeddingResponseDataEmbeddingUnion() {
}

// Object type, always 'embedding'
type CreateEmbeddingResponseDataObject string

const (
	CreateEmbeddingResponseDataObjectEmbedding CreateEmbeddingResponseDataObject = "embedding"
)

func (r CreateEmbeddingResponseDataObject) IsKnown() bool {
	switch r {
	case CreateEmbeddingResponseDataObjectEmbedding:
		return true
	}
	return false
}

// Object type, always 'list'
type CreateEmbeddingResponseObject string

const (
	CreateEmbeddingResponseObjectList CreateEmbeddingResponseObject = "list"
)

func (r CreateEmbeddingResponseObject) IsKnown() bool {
	switch r {
	case CreateEmbeddingResponseObjectList:
		return true
	}
	return false
}

type EmbeddingNewParams struct {
	// Schema for CreateEmbeddingRequest.
	//
	// Fields:
	//
	//   - input (required): str | Annotated[list[str], MinLen(1), MaxLen(2048),
	//     ArrayTitle("CreateEmbeddingRequestInputArray")] | Annotated[list[int],
	//     MinLen(1), MaxLen(2048), ArrayTitle("CreateEmbeddingRequestInputArray")] |
	//     Annotated[list[Annotated[list[int], MinLen(1),
	//     ArrayTitle("CreateEmbeddingRequestInputItemArray")]], MinLen(1), MaxLen(2048),
	//     ArrayTitle("CreateEmbeddingRequestInputArray")]
	//   - model (required): str | Literal["text-embedding-ada-002",
	//     "text-embedding-3-small", "text-embedding-3-large"]
	//   - encoding_format (optional): Literal["float", "base64"]
	//   - dimensions (optional): int
	//   - user (optional): str
	CreateEmbeddingRequest CreateEmbeddingRequestParam `json:"create_embedding_request,required"`
}

func (r EmbeddingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateEmbeddingRequest)
}
