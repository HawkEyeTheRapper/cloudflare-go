// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// GatewayLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayLocationService] method
// instead.
type GatewayLocationService struct {
	Options []option.RequestOption
}

// NewGatewayLocationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayLocationService(opts ...option.RequestOption) (r *GatewayLocationService) {
	r = &GatewayLocationService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway location.
func (r *GatewayLocationService) New(ctx context.Context, params GatewayLocationNewParams, opts ...option.RequestOption) (res *ZeroTrustGatewayLocations, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLocationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway location.
func (r *GatewayLocationService) Update(ctx context.Context, locationID interface{}, params GatewayLocationUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewayLocations, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLocationUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", params.AccountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Zero Trust Gateway locations for an account.
func (r *GatewayLocationService) List(ctx context.Context, query GatewayLocationListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayLocations, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLocationListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a configured Zero Trust Gateway location.
func (r *GatewayLocationService) Delete(ctx context.Context, locationID interface{}, body GatewayLocationDeleteParams, opts ...option.RequestOption) (res *GatewayLocationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLocationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", body.AccountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway location.
func (r *GatewayLocationService) Get(ctx context.Context, locationID interface{}, query GatewayLocationGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayLocations, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLocationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", query.AccountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayLocations struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []ZeroTrustGatewayLocationsNetwork `json:"networks"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayLocationsJSON      `json:"-"`
}

// zeroTrustGatewayLocationsJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayLocations]
type zeroTrustGatewayLocationsJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayLocationsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayLocationsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                               `json:"network,required"`
	JSON    zeroTrustGatewayLocationsNetworkJSON `json:"-"`
}

// zeroTrustGatewayLocationsNetworkJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayLocationsNetwork]
type zeroTrustGatewayLocationsNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLocationsNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayLocationsNetworkJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.GatewayLocationDeleteResponseUnknown] or
// [shared.UnionString].
type GatewayLocationDeleteResponse interface {
	ImplementsZeroTrustGatewayLocationDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GatewayLocationDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type GatewayLocationNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]GatewayLocationNewParamsNetwork] `json:"networks"`
}

func (r GatewayLocationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewParamsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationNewParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationNewResponseEnvelope struct {
	Errors   []GatewayLocationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLocationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayLocations                    `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLocationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLocationNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayLocationNewResponseEnvelope]
type gatewayLocationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    gatewayLocationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLocationNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayLocationNewResponseEnvelopeErrors]
type gatewayLocationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    gatewayLocationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLocationNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayLocationNewResponseEnvelopeMessages]
type gatewayLocationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayLocationNewResponseEnvelopeSuccess bool

const (
	GatewayLocationNewResponseEnvelopeSuccessTrue GatewayLocationNewResponseEnvelopeSuccess = true
)

type GatewayLocationUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]GatewayLocationUpdateParamsNetwork] `json:"networks"`
}

func (r GatewayLocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateParamsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r GatewayLocationUpdateParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLocationUpdateResponseEnvelope struct {
	Errors   []GatewayLocationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLocationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayLocations                       `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLocationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLocationUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayLocationUpdateResponseEnvelope]
type gatewayLocationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    gatewayLocationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLocationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [GatewayLocationUpdateResponseEnvelopeErrors]
type gatewayLocationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    gatewayLocationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLocationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayLocationUpdateResponseEnvelopeMessages]
type gatewayLocationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayLocationUpdateResponseEnvelopeSuccess bool

const (
	GatewayLocationUpdateResponseEnvelopeSuccessTrue GatewayLocationUpdateResponseEnvelopeSuccess = true
)

type GatewayLocationListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayLocationListResponseEnvelope struct {
	Errors   []GatewayLocationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLocationListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayLocations                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayLocationListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayLocationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayLocationListResponseEnvelopeJSON       `json:"-"`
}

// gatewayLocationListResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayLocationListResponseEnvelope]
type gatewayLocationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    gatewayLocationListResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLocationListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayLocationListResponseEnvelopeErrors]
type gatewayLocationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    gatewayLocationListResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLocationListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayLocationListResponseEnvelopeMessages]
type gatewayLocationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayLocationListResponseEnvelopeSuccess bool

const (
	GatewayLocationListResponseEnvelopeSuccessTrue GatewayLocationListResponseEnvelopeSuccess = true
)

type GatewayLocationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       gatewayLocationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayLocationListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [GatewayLocationListResponseEnvelopeResultInfo]
type gatewayLocationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayLocationDeleteResponseEnvelope struct {
	Errors   []GatewayLocationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLocationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayLocationDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLocationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLocationDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayLocationDeleteResponseEnvelope]
type gatewayLocationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    gatewayLocationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLocationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [GatewayLocationDeleteResponseEnvelopeErrors]
type gatewayLocationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    gatewayLocationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLocationDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayLocationDeleteResponseEnvelopeMessages]
type gatewayLocationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayLocationDeleteResponseEnvelopeSuccess bool

const (
	GatewayLocationDeleteResponseEnvelopeSuccessTrue GatewayLocationDeleteResponseEnvelopeSuccess = true
)

type GatewayLocationGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayLocationGetResponseEnvelope struct {
	Errors   []GatewayLocationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLocationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayLocations                    `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLocationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLocationGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayLocationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayLocationGetResponseEnvelope]
type gatewayLocationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    gatewayLocationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLocationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayLocationGetResponseEnvelopeErrors]
type gatewayLocationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayLocationGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    gatewayLocationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLocationGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayLocationGetResponseEnvelopeMessages]
type gatewayLocationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLocationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLocationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayLocationGetResponseEnvelopeSuccess bool

const (
	GatewayLocationGetResponseEnvelopeSuccessTrue GatewayLocationGetResponseEnvelopeSuccess = true
)
