// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingOpportunisticOnionService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingOpportunisticOnionService] method instead.
type SettingOpportunisticOnionService struct {
	Options []option.RequestOption
}

// NewSettingOpportunisticOnionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingOpportunisticOnionService(opts ...option.RequestOption) (r *SettingOpportunisticOnionService) {
	r = &SettingOpportunisticOnionService{}
	r.Options = opts
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *SettingOpportunisticOnionService) Edit(ctx context.Context, params SettingOpportunisticOnionEditParams, opts ...option.RequestOption) (res *OpportunisticOnion, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticOnionEditResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *SettingOpportunisticOnionService) Get(ctx context.Context, query SettingOpportunisticOnionGetParams, opts ...option.RequestOption) (res *OpportunisticOnion, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticOnionGetResponseEnvelope
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type OpportunisticOnion struct {
	// ID of the zone setting.
	ID OpportunisticOnionID `json:"id,required"`
	// Current value of the zone setting.
	Value OpportunisticOnionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       opportunisticOnionJSON `json:"-"`
}

// opportunisticOnionJSON contains the JSON metadata for the struct
// [OpportunisticOnion]
type opportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r opportunisticOnionJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type OpportunisticOnionID string

const (
	OpportunisticOnionIDOpportunisticOnion OpportunisticOnionID = "opportunistic_onion"
)

func (r OpportunisticOnionID) IsKnown() bool {
	switch r {
	case OpportunisticOnionIDOpportunisticOnion:
		return true
	}
	return false
}

// Current value of the zone setting.
type OpportunisticOnionValue string

const (
	OpportunisticOnionValueOn  OpportunisticOnionValue = "on"
	OpportunisticOnionValueOff OpportunisticOnionValue = "off"
)

func (r OpportunisticOnionValue) IsKnown() bool {
	switch r {
	case OpportunisticOnionValueOn, OpportunisticOnionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type OpportunisticOnionEditable bool

const (
	OpportunisticOnionEditableTrue  OpportunisticOnionEditable = true
	OpportunisticOnionEditableFalse OpportunisticOnionEditable = false
)

func (r OpportunisticOnionEditable) IsKnown() bool {
	switch r {
	case OpportunisticOnionEditableTrue, OpportunisticOnionEditableFalse:
		return true
	}
	return false
}

type SettingOpportunisticOnionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingOpportunisticOnionEditParamsValue] `json:"value,required"`
}

func (r SettingOpportunisticOnionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingOpportunisticOnionEditParamsValue string

const (
	SettingOpportunisticOnionEditParamsValueOn  SettingOpportunisticOnionEditParamsValue = "on"
	SettingOpportunisticOnionEditParamsValueOff SettingOpportunisticOnionEditParamsValue = "off"
)

func (r SettingOpportunisticOnionEditParamsValue) IsKnown() bool {
	switch r {
	case SettingOpportunisticOnionEditParamsValueOn, SettingOpportunisticOnionEditParamsValueOff:
		return true
	}
	return false
}

type SettingOpportunisticOnionEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result OpportunisticOnion                                `json:"result"`
	JSON   settingOpportunisticOnionEditResponseEnvelopeJSON `json:"-"`
}

// settingOpportunisticOnionEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingOpportunisticOnionEditResponseEnvelope]
type settingOpportunisticOnionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticOnionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingOpportunisticOnionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOpportunisticOnionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result OpportunisticOnion                               `json:"result"`
	JSON   settingOpportunisticOnionGetResponseEnvelopeJSON `json:"-"`
}

// settingOpportunisticOnionGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingOpportunisticOnionGetResponseEnvelope]
type settingOpportunisticOnionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOpportunisticOnionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
