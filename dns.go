// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DNSService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewDNSService] method instead.
type DNSService struct {
	Options   []option.RequestOption
	Records   *DNSRecordService
	Analytics *DNSAnalyticsService
	Firewall  *DNSFirewallService
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	r.Records = NewDNSRecordService(opts...)
	r.Analytics = NewDNSAnalyticsService(opts...)
	r.Firewall = NewDNSFirewallService(opts...)
	return
}
