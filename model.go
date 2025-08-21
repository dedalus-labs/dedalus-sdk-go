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
func (r *ModelService) Get(ctx context.Context, modelID string, opts ...option.RequestOption) (res *DedalusModel, err error) {
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

// Extended model with configuration capabilities.
//
// Inherits basic metadata from Model and adds configuration fields that can be
// used when creating chat completions. This allows bundling model selection with
// model-specific parameters.
//
// Use this when you want to:
//
// - Pre-configure model parameters
// - Pass model-specific settings
// - Use intelligent routing with attributes
//
// Example: model = DedalusModel( id="gpt-4", temperature=0.7, max_tokens=1000,
// attributes={"intelligence": 0.9, "cost": 0.8} )
//
//	completion = client.chat.completions.create(
//	    model=model,  # Pass the configured model
//	    messages=[...]
//	)
type DedalusModel struct {
	// Model identifier
	ID string `json:"id,required"`
	// [Dedalus] Custom attributes for intelligent model routing (e.g., intelligence,
	// speed, creativity, cost).
	Attributes map[string]float64 `json:"attributes,nullable"`
	// Unix timestamp of model creation
	Created int64 `json:"created"`
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
	// Model name (alias for id)
	Name string `json:"name,nullable"`
	// Object type, always 'model'
	Object string `json:"object"`
	// Organization that owns this model
	OwnedBy string `json:"owned_by"`
	// Whether to enable parallel function calling.
	ParallelToolCalls bool `json:"parallel_tool_calls,nullable"`
	// Penalize new tokens based on whether they appear in the text so far.
	PresencePenalty float64 `json:"presence_penalty,nullable"`
	// Format for the model output (e.g., {'type': 'json_object'}).
	ResponseFormat map[string]any `json:"response_format,nullable"`
	// Seed for deterministic sampling.
	Seed int64 `json:"seed,nullable"`
	// Latency tier for the request (e.g., 'auto', 'default').
	ServiceTier string `json:"service_tier,nullable"`
	// Up to 4 sequences where the API will stop generating further tokens.
	Stop DedalusModelStopUnion `json:"stop,nullable"`
	// Whether to stream back partial progress.
	Stream bool `json:"stream,nullable"`
	// Options for streaming responses.
	StreamOptions map[string]any `json:"stream_options,nullable"`
	// Sampling temperature (0 to 2). Higher values make output more random.
	Temperature float64 `json:"temperature,nullable"`
	// Controls which tool is called by the model.
	ToolChoice DedalusModelToolChoiceUnion `json:"tool_choice,nullable"`
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
		PresencePenalty     respjson.Field
		ResponseFormat      respjson.Field
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
func (r DedalusModel) RawJSON() string { return r.JSON.raw }
func (r *DedalusModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DedalusModel to a DedalusModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DedalusModelParam.Overrides()
func (r DedalusModel) ToParam() DedalusModelParam {
	return param.Override[DedalusModelParam](json.RawMessage(r.RawJSON()))
}

// DedalusModelStopUnion contains all possible properties and values from [string],
// [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type DedalusModelStopUnion struct {
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

func (u DedalusModelStopUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DedalusModelStopUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DedalusModelStopUnion) RawJSON() string { return u.JSON.raw }

func (r *DedalusModelStopUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DedalusModelToolChoiceUnion contains all possible properties and values from
// [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfDedalusModelToolChoiceMapItem]
type DedalusModelToolChoiceUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDedalusModelToolChoiceMapItem any `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfDedalusModelToolChoiceMapItem respjson.Field
		raw                             string
	} `json:"-"`
}

func (u DedalusModelToolChoiceUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DedalusModelToolChoiceUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DedalusModelToolChoiceUnion) RawJSON() string { return u.JSON.raw }

func (r *DedalusModelToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Extended model with configuration capabilities.
//
// Inherits basic metadata from Model and adds configuration fields that can be
// used when creating chat completions. This allows bundling model selection with
// model-specific parameters.
//
// Use this when you want to:
//
// - Pre-configure model parameters
// - Pass model-specific settings
// - Use intelligent routing with attributes
//
// Example: model = DedalusModel( id="gpt-4", temperature=0.7, max_tokens=1000,
// attributes={"intelligence": 0.9, "cost": 0.8} )
//
//	completion = client.chat.completions.create(
//	    model=model,  # Pass the configured model
//	    messages=[...]
//	)
//
// The property ID is required.
type DedalusModelParam struct {
	// Model identifier
	ID string `json:"id,required"`
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
	// Model name (alias for id)
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether to enable parallel function calling.
	ParallelToolCalls param.Opt[bool] `json:"parallel_tool_calls,omitzero"`
	// Penalize new tokens based on whether they appear in the text so far.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
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
	// Unix timestamp of model creation
	Created param.Opt[int64] `json:"created,omitzero"`
	// Object type, always 'model'
	Object param.Opt[string] `json:"object,omitzero"`
	// Organization that owns this model
	OwnedBy param.Opt[string] `json:"owned_by,omitzero"`
	// [Dedalus] Custom attributes for intelligent model routing (e.g., intelligence,
	// speed, creativity, cost).
	Attributes map[string]float64 `json:"attributes,omitzero"`
	// Modify the likelihood of specified tokens appearing.
	LogitBias map[string]float64 `json:"logit_bias,omitzero"`
	// [Dedalus] Additional metadata for request tracking and debugging.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Format for the model output (e.g., {'type': 'json_object'}).
	ResponseFormat map[string]any `json:"response_format,omitzero"`
	// Up to 4 sequences where the API will stop generating further tokens.
	Stop DedalusModelStopUnionParam `json:"stop,omitzero"`
	// Options for streaming responses.
	StreamOptions map[string]any `json:"stream_options,omitzero"`
	// Controls which tool is called by the model.
	ToolChoice DedalusModelToolChoiceUnionParam `json:"tool_choice,omitzero"`
	// List of tools the model may call.
	Tools []map[string]any `json:"tools,omitzero"`
	paramObj
}

func (r DedalusModelParam) MarshalJSON() (data []byte, err error) {
	type shadow DedalusModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DedalusModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DedalusModelStopUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u DedalusModelStopUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *DedalusModelStopUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DedalusModelStopUnionParam) asAny() any {
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
type DedalusModelToolChoiceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u DedalusModelToolChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap)
}
func (u *DedalusModelToolChoiceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DedalusModelToolChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Model metadata following OpenAI's exact structure.
//
// This is a read-only representation of available models returned by GET
// /v1/models. Contains only essential metadata, no configuration fields.
//
// Attributes: id: Model identifier (e.g., 'gpt-4', 'claude-3-5-sonnet') created:
// Unix timestamp when model was created object: Always 'model' for OpenAI
// compatibility owned_by: Organization that owns the model
//
// Example: { "id": "gpt-4", "created": 1687882411, "object": "model", "owned_by":
// "openai" }
type Model struct {
	// Model identifier
	ID string `json:"id,required"`
	// Unix timestamp of model creation
	Created int64 `json:"created"`
	// Object type, always 'model'
	Object string `json:"object"`
	// Organization that owns this model
	OwnedBy string `json:"owned_by"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Object      respjson.Field
		OwnedBy     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
