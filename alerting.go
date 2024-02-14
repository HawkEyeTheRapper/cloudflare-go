// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AlertingService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAlertingService] method instead.
type AlertingService struct {
	Options []option.RequestOption
	V3s     *AlertingV3Service
	V3      *AlertingV3Service
}

// NewAlertingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlertingService(opts ...option.RequestOption) (r *AlertingService) {
	r = &AlertingService{}
	r.Options = opts
	r.V3s = NewAlertingV3Service(opts...)
	r.V3 = NewAlertingV3Service(opts...)
	return
}