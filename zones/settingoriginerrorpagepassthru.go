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

// SettingOriginErrorPagePassThruService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSettingOriginErrorPagePassThruService] method instead.
type SettingOriginErrorPagePassThruService struct {
	Options []option.RequestOption
}

// NewSettingOriginErrorPagePassThruService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingOriginErrorPagePassThruService(opts ...option.RequestOption) (r *SettingOriginErrorPagePassThruService) {
	r = &SettingOriginErrorPagePassThruService{}
	r.Options = opts
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *SettingOriginErrorPagePassThruService) Edit(ctx context.Context, params SettingOriginErrorPagePassThruEditParams, opts ...option.RequestOption) (res *SettingOriginErrorPagePassThruEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOriginErrorPagePassThruEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *SettingOriginErrorPagePassThruService) Get(ctx context.Context, query SettingOriginErrorPagePassThruGetParams, opts ...option.RequestOption) (res *SettingOriginErrorPagePassThruGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOriginErrorPagePassThruGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingOriginErrorPagePassThruEditResponse struct {
	// ID of the zone setting.
	ID SettingOriginErrorPagePassThruEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOriginErrorPagePassThruEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOriginErrorPagePassThruEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOriginErrorPagePassThruEditResponseJSON `json:"-"`
}

// settingOriginErrorPagePassThruEditResponseJSON contains the JSON metadata for
// the struct [SettingOriginErrorPagePassThruEditResponse]
type settingOriginErrorPagePassThruEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginErrorPagePassThruEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingOriginErrorPagePassThruEditResponseID string

const (
	SettingOriginErrorPagePassThruEditResponseIDOriginErrorPagePassThru SettingOriginErrorPagePassThruEditResponseID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type SettingOriginErrorPagePassThruEditResponseValue string

const (
	SettingOriginErrorPagePassThruEditResponseValueOn  SettingOriginErrorPagePassThruEditResponseValue = "on"
	SettingOriginErrorPagePassThruEditResponseValueOff SettingOriginErrorPagePassThruEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOriginErrorPagePassThruEditResponseEditable bool

const (
	SettingOriginErrorPagePassThruEditResponseEditableTrue  SettingOriginErrorPagePassThruEditResponseEditable = true
	SettingOriginErrorPagePassThruEditResponseEditableFalse SettingOriginErrorPagePassThruEditResponseEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingOriginErrorPagePassThruGetResponse struct {
	// ID of the zone setting.
	ID SettingOriginErrorPagePassThruGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOriginErrorPagePassThruGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOriginErrorPagePassThruGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOriginErrorPagePassThruGetResponseJSON `json:"-"`
}

// settingOriginErrorPagePassThruGetResponseJSON contains the JSON metadata for the
// struct [SettingOriginErrorPagePassThruGetResponse]
type settingOriginErrorPagePassThruGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginErrorPagePassThruGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingOriginErrorPagePassThruGetResponseID string

const (
	SettingOriginErrorPagePassThruGetResponseIDOriginErrorPagePassThru SettingOriginErrorPagePassThruGetResponseID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type SettingOriginErrorPagePassThruGetResponseValue string

const (
	SettingOriginErrorPagePassThruGetResponseValueOn  SettingOriginErrorPagePassThruGetResponseValue = "on"
	SettingOriginErrorPagePassThruGetResponseValueOff SettingOriginErrorPagePassThruGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOriginErrorPagePassThruGetResponseEditable bool

const (
	SettingOriginErrorPagePassThruGetResponseEditableTrue  SettingOriginErrorPagePassThruGetResponseEditable = true
	SettingOriginErrorPagePassThruGetResponseEditableFalse SettingOriginErrorPagePassThruGetResponseEditable = false
)

type SettingOriginErrorPagePassThruEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingOriginErrorPagePassThruEditParamsValue] `json:"value,required"`
}

func (r SettingOriginErrorPagePassThruEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingOriginErrorPagePassThruEditParamsValue string

const (
	SettingOriginErrorPagePassThruEditParamsValueOn  SettingOriginErrorPagePassThruEditParamsValue = "on"
	SettingOriginErrorPagePassThruEditParamsValueOff SettingOriginErrorPagePassThruEditParamsValue = "off"
)

type SettingOriginErrorPagePassThruEditResponseEnvelope struct {
	Errors   []SettingOriginErrorPagePassThruEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOriginErrorPagePassThruEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result SettingOriginErrorPagePassThruEditResponse             `json:"result"`
	JSON   settingOriginErrorPagePassThruEditResponseEnvelopeJSON `json:"-"`
}

// settingOriginErrorPagePassThruEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingOriginErrorPagePassThruEditResponseEnvelope]
type settingOriginErrorPagePassThruEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginErrorPagePassThruEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingOriginErrorPagePassThruEditResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    settingOriginErrorPagePassThruEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOriginErrorPagePassThruEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingOriginErrorPagePassThruEditResponseEnvelopeErrors]
type settingOriginErrorPagePassThruEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginErrorPagePassThruEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingOriginErrorPagePassThruEditResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    settingOriginErrorPagePassThruEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOriginErrorPagePassThruEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingOriginErrorPagePassThruEditResponseEnvelopeMessages]
type settingOriginErrorPagePassThruEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginErrorPagePassThruEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingOriginErrorPagePassThruGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOriginErrorPagePassThruGetResponseEnvelope struct {
	Errors   []SettingOriginErrorPagePassThruGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOriginErrorPagePassThruGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result SettingOriginErrorPagePassThruGetResponse             `json:"result"`
	JSON   settingOriginErrorPagePassThruGetResponseEnvelopeJSON `json:"-"`
}

// settingOriginErrorPagePassThruGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingOriginErrorPagePassThruGetResponseEnvelope]
type settingOriginErrorPagePassThruGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginErrorPagePassThruGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingOriginErrorPagePassThruGetResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    settingOriginErrorPagePassThruGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOriginErrorPagePassThruGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingOriginErrorPagePassThruGetResponseEnvelopeErrors]
type settingOriginErrorPagePassThruGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginErrorPagePassThruGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingOriginErrorPagePassThruGetResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    settingOriginErrorPagePassThruGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOriginErrorPagePassThruGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingOriginErrorPagePassThruGetResponseEnvelopeMessages]
type settingOriginErrorPagePassThruGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginErrorPagePassThruGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
