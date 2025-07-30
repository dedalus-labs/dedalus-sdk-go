// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"net/http"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/respjson"
)

// RootService contains methods and other services that help with interacting with
// the dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRootService] method instead.
type RootService struct {
	Options []option.RequestOption
}

// NewRootService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRootService(opts ...option.RequestOption) (r RootService) {
	r = RootService{}
	r.Options = opts
	return
}

// Root
func (r *RootService) Get(ctx context.Context, opts ...option.RequestOption) (res *RootGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := ""
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response model for the root endpoint.
type RootGetResponse struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RootGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RootGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
