// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ACLService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewACLService] method instead.
type ACLService struct {
	Options []option.RequestOption
}

// NewACLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewACLService(opts ...option.RequestOption) (r *ACLService) {
	r = &ACLService{}
	r.Options = opts
	return
}

// Create ACL.
func (r *ACLService) New(ctx context.Context, params ACLNewParams, opts ...option.RequestOption) (res *SecondaryDNSACL, err error) {
	opts = append(r.Options[:], opts...)
	var env ACLNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify ACL.
func (r *ACLService) Update(ctx context.Context, aclID string, params ACLUpdateParams, opts ...option.RequestOption) (res *SecondaryDNSACL, err error) {
	opts = append(r.Options[:], opts...)
	var env ACLUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", params.AccountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List ACLs.
func (r *ACLService) List(ctx context.Context, query ACLListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SecondaryDNSACL], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls", query.AccountID)
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

// List ACLs.
func (r *ACLService) ListAutoPaging(ctx context.Context, query ACLListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SecondaryDNSACL] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete ACL.
func (r *ACLService) Delete(ctx context.Context, aclID string, body ACLDeleteParams, opts ...option.RequestOption) (res *ACLDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ACLDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", body.AccountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get ACL.
func (r *ACLService) Get(ctx context.Context, aclID string, query ACLGetParams, opts ...option.RequestOption) (res *SecondaryDNSACL, err error) {
	opts = append(r.Options[:], opts...)
	var env ACLGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", query.AccountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSACL struct {
	ID string `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string              `json:"name,required"`
	JSON secondaryDnsaclJSON `json:"-"`
}

// secondaryDnsaclJSON contains the JSON metadata for the struct [SecondaryDNSACL]
type secondaryDnsaclJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDnsaclJSON) RawJSON() string {
	return r.raw
}

type ACLDeleteResponse struct {
	ID   string                `json:"id"`
	JSON aclDeleteResponseJSON `json:"-"`
}

// aclDeleteResponseJSON contains the JSON metadata for the struct
// [ACLDeleteResponse]
type aclDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ACLNewParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ACLNewResponseEnvelope struct {
	Errors   []ACLNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ACLNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSACL                  `json:"result,required"`
	// Whether the API call was successful
	Success ACLNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    aclNewResponseEnvelopeJSON    `json:"-"`
}

// aclNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ACLNewResponseEnvelope]
type aclNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ACLNewResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    aclNewResponseEnvelopeErrorsJSON `json:"-"`
}

// aclNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ACLNewResponseEnvelopeErrors]
type aclNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ACLNewResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    aclNewResponseEnvelopeMessagesJSON `json:"-"`
}

// aclNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ACLNewResponseEnvelopeMessages]
type aclNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ACLNewResponseEnvelopeSuccess bool

const (
	ACLNewResponseEnvelopeSuccessTrue ACLNewResponseEnvelopeSuccess = true
)

func (r ACLNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ACLNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ACLUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange param.Field[string] `json:"ip_range,required"`
	// The name of the acl.
	Name param.Field[string] `json:"name,required"`
}

func (r ACLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACLUpdateResponseEnvelope struct {
	Errors   []ACLUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ACLUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSACL                     `json:"result,required"`
	// Whether the API call was successful
	Success ACLUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    aclUpdateResponseEnvelopeJSON    `json:"-"`
}

// aclUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ACLUpdateResponseEnvelope]
type aclUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ACLUpdateResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    aclUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// aclUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ACLUpdateResponseEnvelopeErrors]
type aclUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ACLUpdateResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    aclUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// aclUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ACLUpdateResponseEnvelopeMessages]
type aclUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ACLUpdateResponseEnvelopeSuccess bool

const (
	ACLUpdateResponseEnvelopeSuccessTrue ACLUpdateResponseEnvelopeSuccess = true
)

func (r ACLUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ACLUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ACLListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ACLDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ACLDeleteResponseEnvelope struct {
	Errors   []ACLDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ACLDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ACLDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ACLDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    aclDeleteResponseEnvelopeJSON    `json:"-"`
}

// aclDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ACLDeleteResponseEnvelope]
type aclDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ACLDeleteResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    aclDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// aclDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ACLDeleteResponseEnvelopeErrors]
type aclDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ACLDeleteResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    aclDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// aclDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ACLDeleteResponseEnvelopeMessages]
type aclDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ACLDeleteResponseEnvelopeSuccess bool

const (
	ACLDeleteResponseEnvelopeSuccessTrue ACLDeleteResponseEnvelopeSuccess = true
)

func (r ACLDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ACLDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ACLGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ACLGetResponseEnvelope struct {
	Errors   []ACLGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ACLGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSACL                  `json:"result,required"`
	// Whether the API call was successful
	Success ACLGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    aclGetResponseEnvelopeJSON    `json:"-"`
}

// aclGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ACLGetResponseEnvelope]
type aclGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ACLGetResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    aclGetResponseEnvelopeErrorsJSON `json:"-"`
}

// aclGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ACLGetResponseEnvelopeErrors]
type aclGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ACLGetResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    aclGetResponseEnvelopeMessagesJSON `json:"-"`
}

// aclGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ACLGetResponseEnvelopeMessages]
type aclGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ACLGetResponseEnvelopeSuccess bool

const (
	ACLGetResponseEnvelopeSuccessTrue ACLGetResponseEnvelopeSuccess = true
)

func (r ACLGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ACLGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}