// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// MemberService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMemberService] method instead.
type MemberService struct {
	Options []option.RequestOption
}

// NewMemberService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMemberService(opts ...option.RequestOption) (r *MemberService) {
	r = &MemberService{}
	r.Options = opts
	return
}

// Add a user to the list of members for this account.
func (r *MemberService) New(ctx context.Context, params MemberNewParams, opts ...option.RequestOption) (res *MemberWithCode, err error) {
	opts = append(r.Options[:], opts...)
	var env MemberNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify an account member.
func (r *MemberService) Update(ctx context.Context, memberID string, params MemberUpdateParams, opts ...option.RequestOption) (res *Member, err error) {
	opts = append(r.Options[:], opts...)
	var env MemberUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", params.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all members of an account.
func (r *MemberService) List(ctx context.Context, params MemberListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[MemberListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%v/members", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List all members of an account.
func (r *MemberService) ListAutoPaging(ctx context.Context, params MemberListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[MemberListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Remove a member from an account.
func (r *MemberService) Delete(ctx context.Context, memberID string, body MemberDeleteParams, opts ...option.RequestOption) (res *MemberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MemberDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", body.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific member of an account.
func (r *MemberService) Get(ctx context.Context, memberID string, query MemberGetParams, opts ...option.RequestOption) (res *Member, err error) {
	opts = append(r.Options[:], opts...)
	var env MemberGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", query.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Member struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []MemberRole `json:"roles,required"`
	Status interface{}  `json:"status,required"`
	User   MemberUser   `json:"user,required"`
	JSON   memberJSON   `json:"-"`
}

// memberJSON contains the JSON metadata for the struct [Member]
type memberJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Member) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberJSON) RawJSON() string {
	return r.raw
}

type MemberRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string           `json:"name,required"`
	Permissions MemberPermission `json:"permissions,required"`
	JSON        memberRoleJSON   `json:"-"`
}

// memberRoleJSON contains the JSON metadata for the struct [MemberRole]
type memberRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberRoleJSON) RawJSON() string {
	return r.raw
}

type MemberUser struct {
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Identifier
	ID string `json:"id"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// User's last name
	LastName string `json:"last_name,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool           `json:"two_factor_authentication_enabled"`
	JSON                           memberUserJSON `json:"-"`
}

// memberUserJSON contains the JSON metadata for the struct [MemberUser]
type memberUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *MemberUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUserJSON) RawJSON() string {
	return r.raw
}

type MemberPermission struct {
	Analytics    PermissionGrant      `json:"analytics"`
	Billing      PermissionGrant      `json:"billing"`
	CachePurge   PermissionGrant      `json:"cache_purge"`
	DNS          PermissionGrant      `json:"dns"`
	DNSRecords   PermissionGrant      `json:"dns_records"`
	Lb           PermissionGrant      `json:"lb"`
	Logs         PermissionGrant      `json:"logs"`
	Organization PermissionGrant      `json:"organization"`
	SSL          PermissionGrant      `json:"ssl"`
	WAF          PermissionGrant      `json:"waf"`
	ZoneSettings PermissionGrant      `json:"zone_settings"`
	Zones        PermissionGrant      `json:"zones"`
	JSON         memberPermissionJSON `json:"-"`
}

// memberPermissionJSON contains the JSON metadata for the struct
// [MemberPermission]
type memberPermissionJSON struct {
	Analytics    apijson.Field
	Billing      apijson.Field
	CachePurge   apijson.Field
	DNS          apijson.Field
	DNSRecords   apijson.Field
	Lb           apijson.Field
	Logs         apijson.Field
	Organization apijson.Field
	SSL          apijson.Field
	WAF          apijson.Field
	ZoneSettings apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MemberPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberPermissionJSON) RawJSON() string {
	return r.raw
}

type MemberPermissionParam struct {
	Analytics    param.Field[PermissionGrantParam] `json:"analytics"`
	Billing      param.Field[PermissionGrantParam] `json:"billing"`
	CachePurge   param.Field[PermissionGrantParam] `json:"cache_purge"`
	DNS          param.Field[PermissionGrantParam] `json:"dns"`
	DNSRecords   param.Field[PermissionGrantParam] `json:"dns_records"`
	Lb           param.Field[PermissionGrantParam] `json:"lb"`
	Logs         param.Field[PermissionGrantParam] `json:"logs"`
	Organization param.Field[PermissionGrantParam] `json:"organization"`
	SSL          param.Field[PermissionGrantParam] `json:"ssl"`
	WAF          param.Field[PermissionGrantParam] `json:"waf"`
	ZoneSettings param.Field[PermissionGrantParam] `json:"zone_settings"`
	Zones        param.Field[PermissionGrantParam] `json:"zones"`
}

func (r MemberPermissionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberWithCode struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []MemberWithCodeRole `json:"roles,required"`
	Status interface{}          `json:"status,required"`
	User   MemberWithCodeUser   `json:"user,required"`
	// The unique activation code for the account membership.
	Code string             `json:"code"`
	JSON memberWithCodeJSON `json:"-"`
}

// memberWithCodeJSON contains the JSON metadata for the struct [MemberWithCode]
type memberWithCodeJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberWithCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberWithCodeJSON) RawJSON() string {
	return r.raw
}

type MemberWithCodeRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                 `json:"name,required"`
	Permissions MemberPermission       `json:"permissions,required"`
	JSON        memberWithCodeRoleJSON `json:"-"`
}

// memberWithCodeRoleJSON contains the JSON metadata for the struct
// [MemberWithCodeRole]
type memberWithCodeRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberWithCodeRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberWithCodeRoleJSON) RawJSON() string {
	return r.raw
}

type MemberWithCodeUser struct {
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Identifier
	ID string `json:"id"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// User's last name
	LastName string `json:"last_name,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool                   `json:"two_factor_authentication_enabled"`
	JSON                           memberWithCodeUserJSON `json:"-"`
}

// memberWithCodeUserJSON contains the JSON metadata for the struct
// [MemberWithCodeUser]
type memberWithCodeUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *MemberWithCodeUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberWithCodeUserJSON) RawJSON() string {
	return r.raw
}

type MemberListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Member Name.
	Name string `json:"name,required,nullable"`
	// Roles assigned to this Member.
	Roles []Role `json:"roles,required"`
	// A member's status in the organization.
	Status MemberListResponseStatus `json:"status,required"`
	JSON   memberListResponseJSON   `json:"-"`
}

// memberListResponseJSON contains the JSON metadata for the struct
// [MemberListResponse]
type memberListResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberListResponseJSON) RawJSON() string {
	return r.raw
}

// A member's status in the organization.
type MemberListResponseStatus string

const (
	MemberListResponseStatusAccepted MemberListResponseStatus = "accepted"
	MemberListResponseStatusInvited  MemberListResponseStatus = "invited"
)

func (r MemberListResponseStatus) IsKnown() bool {
	switch r {
	case MemberListResponseStatusAccepted, MemberListResponseStatusInvited:
		return true
	}
	return false
}

type MemberDeleteResponse struct {
	// Identifier
	ID   string                   `json:"id,required"`
	JSON memberDeleteResponseJSON `json:"-"`
}

// memberDeleteResponseJSON contains the JSON metadata for the struct
// [MemberDeleteResponse]
type memberDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MemberNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of roles associated with this member.
	Roles  param.Field[[]string]              `json:"roles,required"`
	Status param.Field[MemberNewParamsStatus] `json:"status"`
}

func (r MemberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberNewParamsStatus string

const (
	MemberNewParamsStatusAccepted MemberNewParamsStatus = "accepted"
	MemberNewParamsStatusPending  MemberNewParamsStatus = "pending"
)

func (r MemberNewParamsStatus) IsKnown() bool {
	switch r {
	case MemberNewParamsStatusAccepted, MemberNewParamsStatusPending:
		return true
	}
	return false
}

type MemberNewResponseEnvelope struct {
	Errors   []MemberNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MemberWithCode                      `json:"result,required"`
	// Whether the API call was successful
	Success MemberNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    memberNewResponseEnvelopeJSON    `json:"-"`
}

// memberNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelope]
type memberNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberNewResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    memberNewResponseEnvelopeErrorsJSON `json:"-"`
}

// memberNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelopeErrors]
type memberNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberNewResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    memberNewResponseEnvelopeMessagesJSON `json:"-"`
}

// memberNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelopeMessages]
type memberNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MemberNewResponseEnvelopeSuccess bool

const (
	MemberNewResponseEnvelopeSuccessTrue MemberNewResponseEnvelopeSuccess = true
)

func (r MemberNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MemberUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// Roles assigned to this member.
	Roles param.Field[[]MemberUpdateParamsRole] `json:"roles,required"`
}

func (r MemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRole struct {
	// Role identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r MemberUpdateParamsRole) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateResponseEnvelope struct {
	Errors   []MemberUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   Member                                 `json:"result,required"`
	// Whether the API call was successful
	Success MemberUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    memberUpdateResponseEnvelopeJSON    `json:"-"`
}

// memberUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberUpdateResponseEnvelope]
type memberUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    memberUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// memberUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberUpdateResponseEnvelopeErrors]
type memberUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    memberUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// memberUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MemberUpdateResponseEnvelopeMessages]
type memberUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MemberUpdateResponseEnvelopeSuccess bool

const (
	MemberUpdateResponseEnvelopeSuccessTrue MemberUpdateResponseEnvelopeSuccess = true
)

func (r MemberUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MemberListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// Direction to order results.
	Direction param.Field[MemberListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[MemberListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A member's status in the account.
	Status param.Field[MemberListParamsStatus] `query:"status"`
}

// URLQuery serializes [MemberListParams]'s query parameters as `url.Values`.
func (r MemberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type MemberListParamsDirection string

const (
	MemberListParamsDirectionAsc  MemberListParamsDirection = "asc"
	MemberListParamsDirectionDesc MemberListParamsDirection = "desc"
)

func (r MemberListParamsDirection) IsKnown() bool {
	switch r {
	case MemberListParamsDirectionAsc, MemberListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by.
type MemberListParamsOrder string

const (
	MemberListParamsOrderUserFirstName MemberListParamsOrder = "user.first_name"
	MemberListParamsOrderUserLastName  MemberListParamsOrder = "user.last_name"
	MemberListParamsOrderUserEmail     MemberListParamsOrder = "user.email"
	MemberListParamsOrderStatus        MemberListParamsOrder = "status"
)

func (r MemberListParamsOrder) IsKnown() bool {
	switch r {
	case MemberListParamsOrderUserFirstName, MemberListParamsOrderUserLastName, MemberListParamsOrderUserEmail, MemberListParamsOrderStatus:
		return true
	}
	return false
}

// A member's status in the account.
type MemberListParamsStatus string

const (
	MemberListParamsStatusAccepted MemberListParamsStatus = "accepted"
	MemberListParamsStatusPending  MemberListParamsStatus = "pending"
	MemberListParamsStatusRejected MemberListParamsStatus = "rejected"
)

func (r MemberListParamsStatus) IsKnown() bool {
	switch r {
	case MemberListParamsStatusAccepted, MemberListParamsStatusPending, MemberListParamsStatusRejected:
		return true
	}
	return false
}

type MemberDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type MemberDeleteResponseEnvelope struct {
	Errors   []MemberDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MemberDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MemberDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    memberDeleteResponseEnvelopeJSON    `json:"-"`
}

// memberDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberDeleteResponseEnvelope]
type memberDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberDeleteResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    memberDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// memberDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberDeleteResponseEnvelopeErrors]
type memberDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberDeleteResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    memberDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// memberDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MemberDeleteResponseEnvelopeMessages]
type memberDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MemberDeleteResponseEnvelopeSuccess bool

const (
	MemberDeleteResponseEnvelopeSuccessTrue MemberDeleteResponseEnvelopeSuccess = true
)

func (r MemberDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MemberGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type MemberGetResponseEnvelope struct {
	Errors   []MemberGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberGetResponseEnvelopeMessages `json:"messages,required"`
	Result   Member                              `json:"result,required"`
	// Whether the API call was successful
	Success MemberGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    memberGetResponseEnvelopeJSON    `json:"-"`
}

// memberGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelope]
type memberGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    memberGetResponseEnvelopeErrorsJSON `json:"-"`
}

// memberGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelopeErrors]
type memberGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    memberGetResponseEnvelopeMessagesJSON `json:"-"`
}

// memberGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelopeMessages]
type memberGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MemberGetResponseEnvelopeSuccess bool

const (
	MemberGetResponseEnvelopeSuccessTrue MemberGetResponseEnvelopeSuccess = true
)

func (r MemberGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}