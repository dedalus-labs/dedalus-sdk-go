// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"github.com/dedalus-labs/dedalus-sdk-go/option"
)

// PrivateService contains methods and other services that help with interacting
// with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrivateService] method instead.
type PrivateService struct {
	Options []option.RequestOption
}

// NewPrivateService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPrivateService(opts ...option.RequestOption) (r *PrivateService) {
	r = &PrivateService{}
	r.Options = opts
	return
}
