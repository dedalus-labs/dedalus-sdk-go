// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalussdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/dedalus-sdk-go/internal/apijson"
	"github.com/stainless-sdks/dedalus-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/dedalus-sdk-go/option"
	"github.com/stainless-sdks/dedalus-sdk-go/packages/respjson"
)

// ModelService contains methods and other services that help with interacting with
// the dedalus-sdk API.
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
// Args: model_id: The ID of the model to retrieve (e.g., 'gpt-4',
// 'claude-3-5-sonnet-20241022') user: Authenticated user obtained from API key
// validation
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
//	model = client.models.retrieve("gpt-4")
//
//	print(f"Model: {model.id}")
//	print(f"Owner: {model.owned_by}")
//	```
//
//	Response:
//	```json
//	{
//	    "id": "gpt-4",
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
//	            "id": "gpt-4",
//	            "object": "model",
//	            "owned_by": "openai"
//	        },
//	        {
//	            "id": "claude-3-5-sonnet-20241022",
//	            "object": "model",
//	            "owned_by": "anthropic"
//	        }
//	    ]
//	}
//	```
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *ModelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Model information compatible with OpenAI API.
//
// Represents a language model available through the Dedalus API. Models are
// aggregated from multiple providers (OpenAI, Anthropic, etc.) and made available
// through a unified interface.
//
// Attributes: id: Unique model identifier (e.g., 'gpt-4',
// 'claude-3-5-sonnet-20241022') object: Always 'model' for compatibility with
// OpenAI API created: Unix timestamp when model was created (may be None)
// owned_by: Provider organization that owns the model root: Base model identifier
// if this is a fine-tuned variant parent: Parent model identifier for hierarchical
// relationships permission: Access permissions (reserved for future use)
//
// Example: { "id": "gpt-4", "object": "model", "created": 1687882411, "owned_by":
// "openai" }
type Model struct {
	// Model identifier
	ID string `json:"id,required"`
	// Creation timestamp
	Created int64 `json:"created,nullable"`
	// Object type
	Object string `json:"object"`
	// Model owner
	OwnedBy string `json:"owned_by"`
	// Parent model
	Parent string `json:"parent,nullable"`
	// Permissions
	Permission []map[string]any `json:"permission,nullable"`
	// Root model
	Root string `json:"root,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Object      respjson.Field
		OwnedBy     respjson.Field
		Parent      respjson.Field
		Permission  respjson.Field
		Root        respjson.Field
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
// Attributes: object: Always 'list' for compatibility with OpenAI API data: List
// of Model objects representing available models
//
// Example: { "object": "list", "data": [ { "id": "gpt-4", "object": "model",
// "owned_by": "openai" }, { "id": "claude-3-5-sonnet-20241022", "object": "model",
// "owned_by": "anthropic" } ] }
type ModelListResponse struct {
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
func (r ModelListResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
