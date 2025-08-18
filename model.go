// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/respjson"
)

// ModelService contains methods and other services that help with interacting with
// the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.Options = opts
	return
}

// Get information about a specific model.
//
// Returns detailed information about a specific model by ID. The model must be
// available to your API key's configured providers.
//
// Args: model_id: The ID of the model to retrieve (e.g., 'openai/gpt-4',
// 'anthropic/claude-3-5-sonnet-20241022') user: Authenticated user obtained from
// API key validation
//
// Returns: DedalusModel: Information about the requested model
//
// Raises: HTTPException: - 401 if authentication fails - 404 if model not found or
// not accessible with current API key - 500 if internal error occurs
//
// Requires: Valid API key with 'read' scope permission
//
// Example: ```python import dedalus_labs
//
//	client = dedalus_labs.Client(api_key="your-api-key")
//	model = client.models.retrieve("openai/gpt-4")
//
//	print(f"Model: {model.id}")
//	print(f"Owner: {model.owned_by}")
//	```
//
//	Response:
//	```json
//	{
//	    "id": "openai/gpt-4",
//	    "object": "model",
//	    "created": 1687882411,
//	    "owned_by": "openai"
//	}
//	```
func (r *ModelService) Get(ctx context.Context, modelID string, opts ...option.RequestOption) (res *Model, err error) {
	opts = append(r.Options[:], opts...)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("v1/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List available models.
//
// Returns a list of available models from all configured providers. Models are
// filtered based on provider availability and API key configuration. Only models
// from providers with valid API keys are returned.
//
// Args: user: Authenticated user obtained from API key validation
//
// Returns: ModelsResponse: Object containing list of available models
//
// Raises: HTTPException: - 401 if authentication fails - 500 if internal error
// occurs during model listing
//
// Requires: Valid API key with 'read' scope permission
//
// Example: ```python import dedalus_labs
//
//	client = dedalus_labs.Client(api_key="your-api-key")
//	models = client.models.list()
//
//	for model in models.data:
//	    print(f"Model: {model.id} (Owner: {model.owned_by})")
//	```
//
//	Response:
//	```json
//	{
//	    "object": "list",
//	    "data": [
//	        {
//	            "id": "openai/gpt-4",
//	            "object": "model",
//	            "owned_by": "openai"
//	        },
//	        {
//	            "id": "anthropic/claude-3-5-sonnet-20241022",
//	            "object": "model",
//	            "owned_by": "anthropic"
//	        }
//	    ]
//	}
//	```
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *ModelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Model information and configuration for Dedalus API.
//
// Represents both model metadata (for listings) and configuration options (for
// chat completions). When used in GET /v1/models, only metadata fields are
// populated. When used in chat requests, configuration fields control model
// behavior.
//
// Attributes: id: Provider-prefixed model identifier (e.g., 'openai/gpt-4',
// 'anthropic/claude-3-5-sonnet') name: Alias for id (used in chat requests)
// object: Always 'model' for OpenAI compatibility created: Unix timestamp when
// model was created owned_by: Provider organization that owns the model root: Base
// model identifier if this is a fine-tuned variant parent: Parent model identifier
// for hierarchical relationships permission: Access permissions (reserved for
// future use)
//
//	Configuration fields (used in chat requests):
//	temperature: Sampling temperature (0 to 2)
//	top_p: Nucleus sampling parameter (0 to 1)
//	frequency_penalty: Frequency penalty (-2 to 2)
//	presence_penalty: Presence penalty (-2 to 2)
//	max_tokens: Maximum tokens to generate
//	attributes: Custom attributes for model routing
//	metadata: Additional metadata
//
// Example: Listing response: { "id": "openai/gpt-4o", "object": "model",
// "created": 1687882411, "owned_by": "openai" }
//
//	Chat request:
//	{
//	    "name": "openai/gpt-4o",
//	    "temperature": 0.7,
//	    "max_tokens": 1000,
//	    "attributes": {"intelligence": 0.9, "cost": 0.8}
//	}
type Model struct {
	// Model identifier
	ID string `json:"id,nullable"`
	// [Dedalus] Custom attributes for intelligent model routing (e.g., intelligence,
	// speed, creativity, cost).
	Attributes map[string]float64 `json:"attributes,nullable"`
	// Creation timestamp
	Created int64 `json:"created,nullable"`
	// Penalize new tokens based on their frequency in the text so far.
	FrequencyPenalty float64 `json:"frequency_penalty,nullable"`
	// Modify the likelihood of specified tokens appearing.
	LogitBias map[string]float64 `json:"logit_bias,nullable"`
	// Whether to return log probabilities of the output tokens.
	Logprobs bool `json:"logprobs,nullable"`
	// An upper bound for the number of tokens that can be generated for a completion.
	MaxCompletionTokens int64 `json:"max_completion_tokens,nullable"`
	// Maximum number of tokens to generate.
	MaxTokens int64 `json:"max_tokens,nullable"`
	// [Dedalus] Additional metadata for request tracking and debugging.
	Metadata map[string]string `json:"metadata,nullable"`
	// Number of completions to generate for each prompt.
	N int64 `json:"n,nullable"`
	// Model name (alias for id in chat requests)
	Name string `json:"name,nullable"`
	// Object type
	Object string `json:"object"`
	// Model owner
	OwnedBy string `json:"owned_by"`
	// Whether to enable parallel function calling.
	ParallelToolCalls bool `json:"parallel_tool_calls,nullable"`
	// Parent model
	Parent string `json:"parent,nullable"`
	// Permissions
	Permission []map[string]any `json:"permission,nullable"`
	// Penalize new tokens based on whether they appear in the text so far.
	PresencePenalty float64 `json:"presence_penalty,nullable"`
	// Format for the model output (e.g., {'type': 'json_object'}).
	ResponseFormat map[string]any `json:"response_format,nullable"`
	// Root model
	Root string `json:"root,nullable"`
	// Seed for deterministic sampling.
	Seed int64 `json:"seed,nullable"`
	// Latency tier for the request (e.g., 'auto', 'default').
	ServiceTier string `json:"service_tier,nullable"`
	// Up to 4 sequences where the API will stop generating further tokens.
	Stop ModelStopUnion `json:"stop,nullable"`
	// Whether to stream back partial progress.
	Stream bool `json:"stream,nullable"`
	// Options for streaming responses.
	StreamOptions map[string]any `json:"stream_options,nullable"`
	// Sampling temperature (0 to 2). Higher values make output more random.
	Temperature float64 `json:"temperature,nullable"`
	// Controls which tool is called by the model.
	ToolChoice ModelToolChoiceUnion `json:"tool_choice,nullable"`
	// List of tools the model may call.
	Tools []map[string]any `json:"tools,nullable"`
	// Number of most likely tokens to return at each token position.
	TopLogprobs int64 `json:"top_logprobs,nullable"`
	// Nucleus sampling parameter. Alternative to temperature.
	TopP float64 `json:"top_p,nullable"`
	// A unique identifier representing your end-user.
	User string `json:"user,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Attributes          respjson.Field
		Created             respjson.Field
		FrequencyPenalty    respjson.Field
		LogitBias           respjson.Field
		Logprobs            respjson.Field
		MaxCompletionTokens respjson.Field
		MaxTokens           respjson.Field
		Metadata            respjson.Field
		N                   respjson.Field
		Name                respjson.Field
		Object              respjson.Field
		OwnedBy             respjson.Field
		ParallelToolCalls   respjson.Field
		Parent              respjson.Field
		Permission          respjson.Field
		PresencePenalty     respjson.Field
		ResponseFormat      respjson.Field
		Root                respjson.Field
		Seed                respjson.Field
		ServiceTier         respjson.Field
		Stop                respjson.Field
		Stream              respjson.Field
		StreamOptions       respjson.Field
		Temperature         respjson.Field
		ToolChoice          respjson.Field
		Tools               respjson.Field
		TopLogprobs         respjson.Field
		TopP                respjson.Field
		User                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Model to a ModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ModelParam.Overrides()
func (r Model) ToParam() ModelParam {
	return param.Override[ModelParam](json.RawMessage(r.RawJSON()))
}

// ModelStopUnion contains all possible properties and values from [string],
// [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type ModelStopUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u ModelStopUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelStopUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModelStopUnion) RawJSON() string { return u.JSON.raw }

func (r *ModelStopUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ModelToolChoiceUnion contains all possible properties and values from [string],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfModelToolChoiceMapItem]
type ModelToolChoiceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfModelToolChoiceMapItem any `json:",inline"`
	JSON                     struct {
		OfString                 respjson.Field
		OfModelToolChoiceMapItem respjson.Field
		raw                      string
	} `json:"-"`
}

func (u ModelToolChoiceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelToolChoiceUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModelToolChoiceUnion) RawJSON() string { return u.JSON.raw }

func (r *ModelToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model information and configuration for Dedalus API.
//
// Represents both model metadata (for listings) and configuration options (for
// chat completions). When used in GET /v1/models, only metadata fields are
// populated. When used in chat requests, configuration fields control model
// behavior.
//
// Attributes: id: Provider-prefixed model identifier (e.g., 'openai/gpt-4',
// 'anthropic/claude-3-5-sonnet') name: Alias for id (used in chat requests)
// object: Always 'model' for OpenAI compatibility created: Unix timestamp when
// model was created owned_by: Provider organization that owns the model root: Base
// model identifier if this is a fine-tuned variant parent: Parent model identifier
// for hierarchical relationships permission: Access permissions (reserved for
// future use)
//
//	Configuration fields (used in chat requests):
//	temperature: Sampling temperature (0 to 2)
//	top_p: Nucleus sampling parameter (0 to 1)
//	frequency_penalty: Frequency penalty (-2 to 2)
//	presence_penalty: Presence penalty (-2 to 2)
//	max_tokens: Maximum tokens to generate
//	attributes: Custom attributes for model routing
//	metadata: Additional metadata
//
// Example: Listing response: { "id": "openai/gpt-4o", "object": "model",
// "created": 1687882411, "owned_by": "openai" }
//
//	Chat request:
//	{
//	    "name": "openai/gpt-4o",
//	    "temperature": 0.7,
//	    "max_tokens": 1000,
//	    "attributes": {"intelligence": 0.9, "cost": 0.8}
//	}
type ModelParam struct {
	// Model identifier
	ID param.Opt[string] `json:"id,omitzero"`
	// Creation timestamp
	Created param.Opt[int64] `json:"created,omitzero"`
	// Penalize new tokens based on their frequency in the text so far.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// Whether to return log probabilities of the output tokens.
	Logprobs param.Opt[bool] `json:"logprobs,omitzero"`
	// An upper bound for the number of tokens that can be generated for a completion.
	MaxCompletionTokens param.Opt[int64] `json:"max_completion_tokens,omitzero"`
	// Maximum number of tokens to generate.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// Number of completions to generate for each prompt.
	N param.Opt[int64] `json:"n,omitzero"`
	// Model name (alias for id in chat requests)
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether to enable parallel function calling.
	ParallelToolCalls param.Opt[bool] `json:"parallel_tool_calls,omitzero"`
	// Parent model
	Parent param.Opt[string] `json:"parent,omitzero"`
	// Penalize new tokens based on whether they appear in the text so far.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// Root model
	Root param.Opt[string] `json:"root,omitzero"`
	// Seed for deterministic sampling.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Latency tier for the request (e.g., 'auto', 'default').
	ServiceTier param.Opt[string] `json:"service_tier,omitzero"`
	// Whether to stream back partial progress.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// Sampling temperature (0 to 2). Higher values make output more random.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Number of most likely tokens to return at each token position.
	TopLogprobs param.Opt[int64] `json:"top_logprobs,omitzero"`
	// Nucleus sampling parameter. Alternative to temperature.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// A unique identifier representing your end-user.
	User param.Opt[string] `json:"user,omitzero"`
	// Object type
	Object param.Opt[string] `json:"object,omitzero"`
	// Model owner
	OwnedBy param.Opt[string] `json:"owned_by,omitzero"`
	// [Dedalus] Custom attributes for intelligent model routing (e.g., intelligence,
	// speed, creativity, cost).
	Attributes map[string]float64 `json:"attributes,omitzero"`
	// Modify the likelihood of specified tokens appearing.
	LogitBias map[string]float64 `json:"logit_bias,omitzero"`
	// [Dedalus] Additional metadata for request tracking and debugging.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Permissions
	Permission []map[string]any `json:"permission,omitzero"`
	// Format for the model output (e.g., {'type': 'json_object'}).
	ResponseFormat map[string]any `json:"response_format,omitzero"`
	// Up to 4 sequences where the API will stop generating further tokens.
	Stop ModelStopUnionParam `json:"stop,omitzero"`
	// Options for streaming responses.
	StreamOptions map[string]any `json:"stream_options,omitzero"`
	// Controls which tool is called by the model.
	ToolChoice ModelToolChoiceUnionParam `json:"tool_choice,omitzero"`
	// List of tools the model may call.
	Tools []map[string]any `json:"tools,omitzero"`
	paramObj
}

func (r ModelParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ModelStopUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ModelStopUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ModelStopUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ModelStopUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ModelToolChoiceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u ModelToolChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap)
}
func (u *ModelToolChoiceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ModelToolChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Response containing list of available models.
//
// Returns all models available to the authenticated user based on their API key
// permissions and configured providers.
//
// Attributes: object: Always 'list' for compatibility with OpenAI API data: list
// of Model objects representing available models
//
// Example: { "object": "list", "data": [ { "id": "openai/gpt-4", "object":
// "model", "owned_by": "openai" }, { "id": "anthropic/claude-3-5-sonnet-20241022",
// "object": "model", "owned_by": "anthropic" } ] }
type ModelsResponse struct {
	// List of models
	Data []Model `json:"data,required"`
	// Object type
	Object string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
