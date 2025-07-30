// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalussdk

import (
	"github.com/stainless-sdks/dedalus-sdk-go/internal/apijson"
	"github.com/stainless-sdks/dedalus-sdk-go/packages/respjson"
)

// Response model for the root endpoint.
type GetRootResponse struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetRootResponse) RawJSON() string { return r.JSON.raw }
func (r *GetRootResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
