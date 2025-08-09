// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
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
func NewModelService(opts ...option.RequestOption) (r *ModelService) {
	r = &ModelService{}
	r.Options = opts
	return
}

// Get information about a specific model.
//
// Returns detailed information about a specific model by ID. The model must be
// available to your API key's configured providers.
func (r *ModelService) Get(ctx context.Context, modelID string, opts ...option.RequestOption) (res *ModelInfo, err error) {
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
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *ModelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ModelInfo struct {
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
	Permission []interface{} `json:"permission,nullable"`
	// Root model
	Root string        `json:"root,nullable"`
	JSON modelInfoJSON `json:"-"`
}

// modelInfoJSON contains the JSON metadata for the struct [ModelInfo]
type modelInfoJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Object      apijson.Field
	OwnedBy     apijson.Field
	Parent      apijson.Field
	Permission  apijson.Field
	Root        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelInfoJSON) RawJSON() string {
	return r.raw
}

type ModelsResponse struct {
	// List of models
	Data []ModelInfo `json:"data,required"`
	// Object type
	Object string             `json:"object"`
	JSON   modelsResponseJSON `json:"-"`
}

// modelsResponseJSON contains the JSON metadata for the struct [ModelsResponse]
type modelsResponseJSON struct {
	Data        apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelsResponseJSON) RawJSON() string {
	return r.raw
}
