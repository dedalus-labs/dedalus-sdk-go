// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
)

// Response model for the root endpoint of the Dedalus API.
type GetResponse struct {
	Message string          `json:"message,required"`
	JSON    getResponseJSON `json:"-"`
}

// getResponseJSON contains the JSON metadata for the struct [GetResponse]
type getResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getResponseJSON) RawJSON() string {
	return r.raw
}
