// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
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

// Retrieve a model.
//
// Retrieve detailed information about a specific model, including its
// capabilities, provider, and supported features.
//
// Args: model_id: The ID of the model to retrieve (e.g., 'openai/gpt-4',
// 'anthropic/claude-3-5-sonnet-20241022') user: Authenticated user obtained from
// API key validation
//
// Returns: Model: Information about the requested model
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
	opts = slices.Concat(r.Options, opts)
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
// Retrieve the complete list of models available to your organization, including
// models from OpenAI, Anthropic, Google, xAI, Mistral, Fireworks, and DeepSeek.
//
// Returns: ListModelsResponse: List of available models across all supported
// providers
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *ListModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response for /v1/models endpoint.
type ListModelsResponse struct {
	// List of available models
	Data []Model `json:"data,required"`
	// Response object type
	//
	// Any of "list".
	Object ListModelsResponseObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object type
type ListModelsResponseObject string

const (
	ListModelsResponseObjectList ListModelsResponseObject = "list"
)

// Unified model metadata across all providers.
//
// Combines provider-specific schemas into a single, consistent format. Fields that
// aren't available from a provider are set to None.
type Model struct {
	// Unique model identifier with provider prefix (e.g., 'openai/gpt-4')
	ID string `json:"id,required"`
	// When the model was released (RFC 3339)
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Provider that hosts this model
	//
	// Any of "openai", "anthropic", "google", "xai", "mistral", "groq", "fireworks",
	// "deepseek".
	Provider ModelProvider `json:"provider,required"`
	// Normalized model capabilities across all providers.
	Capabilities ModelCapabilities `json:"capabilities,nullable"`
	// Provider-declared default parameters for model generation.
	Defaults ModelDefaults `json:"defaults,nullable"`
	// Model description
	Description string `json:"description,nullable"`
	// Human-readable model name
	DisplayName string `json:"display_name,nullable"`
	// Provider-specific generation method names (None = not declared)
	ProviderDeclaredGenerationMethods []string `json:"provider_declared_generation_methods,nullable"`
	// Raw provider-specific metadata
	ProviderInfo map[string]any `json:"provider_info,nullable"`
	// Model version identifier
	Version string `json:"version,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                respjson.Field
		CreatedAt                         respjson.Field
		Provider                          respjson.Field
		Capabilities                      respjson.Field
		Defaults                          respjson.Field
		Description                       respjson.Field
		DisplayName                       respjson.Field
		ProviderDeclaredGenerationMethods respjson.Field
		ProviderInfo                      respjson.Field
		Version                           respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provider that hosts this model
type ModelProvider string

const (
	ModelProviderOpenAI    ModelProvider = "openai"
	ModelProviderAnthropic ModelProvider = "anthropic"
	ModelProviderGoogle    ModelProvider = "google"
	ModelProviderXai       ModelProvider = "xai"
	ModelProviderMistral   ModelProvider = "mistral"
	ModelProviderGroq      ModelProvider = "groq"
	ModelProviderFireworks ModelProvider = "fireworks"
	ModelProviderDeepseek  ModelProvider = "deepseek"
)

// Normalized model capabilities across all providers.
type ModelCapabilities struct {
	// Supports audio processing
	Audio bool `json:"audio,nullable"`
	// Supports image generation
	ImageGeneration bool `json:"image_generation,nullable"`
	// Maximum input tokens
	InputTokenLimit int64 `json:"input_token_limit,nullable"`
	// Maximum output tokens
	OutputTokenLimit int64 `json:"output_token_limit,nullable"`
	// Supports streaming responses
	Streaming bool `json:"streaming,nullable"`
	// Supports structured JSON output
	StructuredOutput bool `json:"structured_output,nullable"`
	// Supports text generation
	Text bool `json:"text,nullable"`
	// Supports extended thinking/reasoning
	Thinking bool `json:"thinking,nullable"`
	// Supports function/tool calling
	Tools bool `json:"tools,nullable"`
	// Supports image understanding
	Vision bool `json:"vision,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audio            respjson.Field
		ImageGeneration  respjson.Field
		InputTokenLimit  respjson.Field
		OutputTokenLimit respjson.Field
		Streaming        respjson.Field
		StructuredOutput respjson.Field
		Text             respjson.Field
		Thinking         respjson.Field
		Tools            respjson.Field
		Vision           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelCapabilities) RawJSON() string { return r.JSON.raw }
func (r *ModelCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provider-declared default parameters for model generation.
type ModelDefaults struct {
	// Default maximum output tokens
	MaxOutputTokens int64 `json:"max_output_tokens,nullable"`
	// Default temperature setting
	Temperature float64 `json:"temperature,nullable"`
	// Default top_k setting
	TopK int64 `json:"top_k,nullable"`
	// Default top_p setting
	TopP float64 `json:"top_p,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxOutputTokens respjson.Field
		Temperature     respjson.Field
		TopK            respjson.Field
		TopP            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelDefaults) RawJSON() string { return r.JSON.raw }
func (r *ModelDefaults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
