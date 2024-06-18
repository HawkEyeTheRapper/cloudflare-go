// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// BGPLeakService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPLeakService] method instead.
type BGPLeakService struct {
	Options []option.RequestOption
	Events  *BGPLeakEventService
}

// NewBGPLeakService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBGPLeakService(opts ...option.RequestOption) (r *BGPLeakService) {
	r = &BGPLeakService{}
	r.Options = opts
	r.Events = NewBGPLeakEventService(opts...)
	return
}
