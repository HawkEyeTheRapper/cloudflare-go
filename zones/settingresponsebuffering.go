// File generated from our OpenAPI spec by Stainless.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// SettingResponseBufferingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingResponseBufferingService] method instead.
type SettingResponseBufferingService struct {
	Options []option.RequestOption
}

// NewSettingResponseBufferingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingResponseBufferingService(opts ...option.RequestOption) (r *SettingResponseBufferingService) {
	r = &SettingResponseBufferingService{}
	r.Options = opts
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *SettingResponseBufferingService) Edit(ctx context.Context, params SettingResponseBufferingEditParams, opts ...option.RequestOption) (res *SettingResponseBufferingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingResponseBufferingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/response_buffering", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *SettingResponseBufferingService) Get(ctx context.Context, query SettingResponseBufferingGetParams, opts ...option.RequestOption) (res *SettingResponseBufferingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingResponseBufferingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/response_buffering", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingResponseBufferingEditResponse struct {
	// ID of the zone setting.
	ID SettingResponseBufferingEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingResponseBufferingEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingResponseBufferingEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingResponseBufferingEditResponseJSON `json:"-"`
}

// settingResponseBufferingEditResponseJSON contains the JSON metadata for the
// struct [SettingResponseBufferingEditResponse]
type settingResponseBufferingEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingResponseBufferingEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingResponseBufferingEditResponseID string

const (
	SettingResponseBufferingEditResponseIDResponseBuffering SettingResponseBufferingEditResponseID = "response_buffering"
)

// Current value of the zone setting.
type SettingResponseBufferingEditResponseValue string

const (
	SettingResponseBufferingEditResponseValueOn  SettingResponseBufferingEditResponseValue = "on"
	SettingResponseBufferingEditResponseValueOff SettingResponseBufferingEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingResponseBufferingEditResponseEditable bool

const (
	SettingResponseBufferingEditResponseEditableTrue  SettingResponseBufferingEditResponseEditable = true
	SettingResponseBufferingEditResponseEditableFalse SettingResponseBufferingEditResponseEditable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingResponseBufferingGetResponse struct {
	// ID of the zone setting.
	ID SettingResponseBufferingGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingResponseBufferingGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingResponseBufferingGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingResponseBufferingGetResponseJSON `json:"-"`
}

// settingResponseBufferingGetResponseJSON contains the JSON metadata for the
// struct [SettingResponseBufferingGetResponse]
type settingResponseBufferingGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingResponseBufferingGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingResponseBufferingGetResponseID string

const (
	SettingResponseBufferingGetResponseIDResponseBuffering SettingResponseBufferingGetResponseID = "response_buffering"
)

// Current value of the zone setting.
type SettingResponseBufferingGetResponseValue string

const (
	SettingResponseBufferingGetResponseValueOn  SettingResponseBufferingGetResponseValue = "on"
	SettingResponseBufferingGetResponseValueOff SettingResponseBufferingGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingResponseBufferingGetResponseEditable bool

const (
	SettingResponseBufferingGetResponseEditableTrue  SettingResponseBufferingGetResponseEditable = true
	SettingResponseBufferingGetResponseEditableFalse SettingResponseBufferingGetResponseEditable = false
)

type SettingResponseBufferingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingResponseBufferingEditParamsValue] `json:"value,required"`
}

func (r SettingResponseBufferingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingResponseBufferingEditParamsValue string

const (
	SettingResponseBufferingEditParamsValueOn  SettingResponseBufferingEditParamsValue = "on"
	SettingResponseBufferingEditParamsValueOff SettingResponseBufferingEditParamsValue = "off"
)

type SettingResponseBufferingEditResponseEnvelope struct {
	Errors   []SettingResponseBufferingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingResponseBufferingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result SettingResponseBufferingEditResponse             `json:"result"`
	JSON   settingResponseBufferingEditResponseEnvelopeJSON `json:"-"`
}

// settingResponseBufferingEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingResponseBufferingEditResponseEnvelope]
type settingResponseBufferingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingResponseBufferingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingResponseBufferingEditResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingResponseBufferingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingResponseBufferingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingResponseBufferingEditResponseEnvelopeErrors]
type settingResponseBufferingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingResponseBufferingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingResponseBufferingEditResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingResponseBufferingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingResponseBufferingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingResponseBufferingEditResponseEnvelopeMessages]
type settingResponseBufferingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingResponseBufferingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingResponseBufferingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingResponseBufferingGetResponseEnvelope struct {
	Errors   []SettingResponseBufferingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingResponseBufferingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result SettingResponseBufferingGetResponse             `json:"result"`
	JSON   settingResponseBufferingGetResponseEnvelopeJSON `json:"-"`
}

// settingResponseBufferingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingResponseBufferingGetResponseEnvelope]
type settingResponseBufferingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingResponseBufferingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingResponseBufferingGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingResponseBufferingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingResponseBufferingGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingResponseBufferingGetResponseEnvelopeErrors]
type settingResponseBufferingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingResponseBufferingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingResponseBufferingGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingResponseBufferingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingResponseBufferingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingResponseBufferingGetResponseEnvelopeMessages]
type settingResponseBufferingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingResponseBufferingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
