// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// DeviceRevokeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeviceRevokeService] method
// instead.
type DeviceRevokeService struct {
	Options []option.RequestOption
}

// NewDeviceRevokeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceRevokeService(opts ...option.RequestOption) (r *DeviceRevokeService) {
	r = &DeviceRevokeService{}
	r.Options = opts
	return
}

// Revokes a list of devices.
func (r *DeviceRevokeService) New(ctx context.Context, params DeviceRevokeNewParams, opts ...option.RequestOption) (res *DeviceRevokeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceRevokeNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/revoke", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [zero_trust.DeviceRevokeNewResponseUnknown] or
// [shared.UnionString].
type DeviceRevokeNewResponse interface {
	ImplementsZeroTrustDeviceRevokeNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeviceRevokeNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DeviceRevokeNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of device ids to revoke.
	Body param.Field[[]string] `json:"body,required"`
}

func (r DeviceRevokeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DeviceRevokeNewResponseEnvelope struct {
	Errors   []DeviceRevokeNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceRevokeNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceRevokeNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceRevokeNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceRevokeNewResponseEnvelopeJSON    `json:"-"`
}

// deviceRevokeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceRevokeNewResponseEnvelope]
type deviceRevokeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRevokeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRevokeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DeviceRevokeNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    deviceRevokeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceRevokeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceRevokeNewResponseEnvelopeErrors]
type deviceRevokeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRevokeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRevokeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DeviceRevokeNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    deviceRevokeNewResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceRevokeNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceRevokeNewResponseEnvelopeMessages]
type deviceRevokeNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRevokeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRevokeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceRevokeNewResponseEnvelopeSuccess bool

const (
	DeviceRevokeNewResponseEnvelopeSuccessTrue DeviceRevokeNewResponseEnvelopeSuccess = true
)

func (r DeviceRevokeNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceRevokeNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}