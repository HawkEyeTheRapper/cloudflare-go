// File generated from our OpenAPI spec by Stainless.

package magic_network_monitoring

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ConfigFullService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewConfigFullService] method instead.
type ConfigFullService struct {
	Options []option.RequestOption
}

// NewConfigFullService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigFullService(opts ...option.RequestOption) (r *ConfigFullService) {
	r = &ConfigFullService{}
	r.Options = opts
	return
}

// Lists default sampling, router IPs, and rules for account.
func (r *ConfigFullService) Get(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *ConfigFullGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigFullGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config/full", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConfigFullGetResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                    `json:"name,required"`
	RouterIPs []string                  `json:"router_ips,required"`
	JSON      configFullGetResponseJSON `json:"-"`
}

// configFullGetResponseJSON contains the JSON metadata for the struct
// [ConfigFullGetResponse]
type configFullGetResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigFullGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configFullGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigFullGetResponseEnvelope struct {
	Errors   []ConfigFullGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigFullGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ConfigFullGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ConfigFullGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    configFullGetResponseEnvelopeJSON    `json:"-"`
}

// configFullGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigFullGetResponseEnvelope]
type configFullGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigFullGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configFullGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigFullGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    configFullGetResponseEnvelopeErrorsJSON `json:"-"`
}

// configFullGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConfigFullGetResponseEnvelopeErrors]
type configFullGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigFullGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configFullGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigFullGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    configFullGetResponseEnvelopeMessagesJSON `json:"-"`
}

// configFullGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConfigFullGetResponseEnvelopeMessages]
type configFullGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigFullGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configFullGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigFullGetResponseEnvelopeSuccess bool

const (
	ConfigFullGetResponseEnvelopeSuccessTrue ConfigFullGetResponseEnvelopeSuccess = true
)
