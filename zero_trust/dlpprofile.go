// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// DLPProfileService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDLPProfileService] method instead.
type DLPProfileService struct {
	Options    []option.RequestOption
	Custom     *DLPProfileCustomService
	Predefined *DLPProfilePredefinedService
}

// NewDLPProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPProfileService(opts ...option.RequestOption) (r *DLPProfileService) {
	r = &DLPProfileService{}
	r.Options = opts
	r.Custom = NewDLPProfileCustomService(opts...)
	r.Predefined = NewDLPProfilePredefinedService(opts...)
	return
}

// Lists all DLP profiles in an account.
func (r *DLPProfileService) List(ctx context.Context, query DLPProfileListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPProfiles], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all DLP profiles in an account.
func (r *DLPProfileService) ListAutoPaging(ctx context.Context, query DLPProfileListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPProfiles] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Fetches a DLP profile by ID. Supports both predefined and custom profiles
func (r *DLPProfileService) Get(ctx context.Context, profileID string, query DLPProfileGetParams, opts ...option.RequestOption) (res *DLPProfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [zero_trust.DLPPredefinedProfile],
// [zero_trust.DLPCustomProfile] or [zero_trust.DLPProfilesDLPIntegrationProfile].
type DLPProfiles interface {
	implementsZeroTrustDLPProfiles()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfiles)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPPredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPCustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilesDLPIntegrationProfile{}),
		},
	)
}

type DLPProfilesDLPIntegrationProfile struct {
	// The ID for this profile
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfilesDLPIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfilesDLPIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      dlpProfilesDLPIntegrationProfileJSON `json:"-"`
}

// dlpProfilesDLPIntegrationProfileJSON contains the JSON metadata for the struct
// [DLPProfilesDLPIntegrationProfile]
type dlpProfilesDLPIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilesDLPIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilesDLPIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilesDLPIntegrationProfile) implementsZeroTrustDLPProfiles() {}

// An entry derived from an integration
type DLPProfilesDLPIntegrationProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                               `json:"profile_id"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      dlpProfilesDLPIntegrationProfileEntryJSON `json:"-"`
}

// dlpProfilesDLPIntegrationProfileEntryJSON contains the JSON metadata for the
// struct [DLPProfilesDLPIntegrationProfileEntry]
type dlpProfilesDLPIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilesDLPIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilesDLPIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type DLPProfilesDLPIntegrationProfileType string

const (
	DLPProfilesDLPIntegrationProfileTypeIntegration DLPProfilesDLPIntegrationProfileType = "integration"
)

func (r DLPProfilesDLPIntegrationProfileType) IsKnown() bool {
	switch r {
	case DLPProfilesDLPIntegrationProfileTypeIntegration:
		return true
	}
	return false
}

// Union satisfied by [zero_trust.DLPPredefinedProfile],
// [zero_trust.DLPCustomProfile] or
// [zero_trust.DLPProfileGetResponseDLPIntegrationProfile].
type DLPProfileGetResponse interface {
	implementsZeroTrustDLPProfileGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPPredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPCustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseDLPIntegrationProfile{}),
		},
	)
}

type DLPProfileGetResponseDLPIntegrationProfile struct {
	// The ID for this profile
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfileGetResponseDLPIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfileGetResponseDLPIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      dlpProfileGetResponseDLPIntegrationProfileJSON `json:"-"`
}

// dlpProfileGetResponseDLPIntegrationProfileJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseDLPIntegrationProfile]
type dlpProfileGetResponseDLPIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseDLPIntegrationProfile) implementsZeroTrustDLPProfileGetResponse() {}

// An entry derived from an integration
type DLPProfileGetResponseDLPIntegrationProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                         `json:"profile_id"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	JSON      dlpProfileGetResponseDLPIntegrationProfileEntryJSON `json:"-"`
}

// dlpProfileGetResponseDLPIntegrationProfileEntryJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseDLPIntegrationProfileEntry]
type dlpProfileGetResponseDLPIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type DLPProfileGetResponseDLPIntegrationProfileType string

const (
	DLPProfileGetResponseDLPIntegrationProfileTypeIntegration DLPProfileGetResponseDLPIntegrationProfileType = "integration"
)

func (r DLPProfileGetResponseDLPIntegrationProfileType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseDLPIntegrationProfileTypeIntegration:
		return true
	}
	return false
}

type DLPProfileListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileGetResponseEnvelope struct {
	Errors   []DLPProfileGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPProfileGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DLPProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpProfileGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseEnvelope]
type dlpProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dlpProfileGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpProfileGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeErrors]
type dlpProfileGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dlpProfileGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpProfileGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeMessages]
type dlpProfileGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileGetResponseEnvelopeSuccess bool

const (
	DLPProfileGetResponseEnvelopeSuccessTrue DLPProfileGetResponseEnvelopeSuccess = true
)

func (r DLPProfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}