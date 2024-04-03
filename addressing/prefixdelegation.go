// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PrefixDelegationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPrefixDelegationService] method
// instead.
type PrefixDelegationService struct {
	Options []option.RequestOption
}

// NewPrefixDelegationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrefixDelegationService(opts ...option.RequestOption) (r *PrefixDelegationService) {
	r = &PrefixDelegationService{}
	r.Options = opts
	return
}

// Create a new account delegation for a given IP prefix.
func (r *PrefixDelegationService) New(ctx context.Context, prefixID string, params PrefixDelegationNewParams, opts ...option.RequestOption) (res *AddressingIpamDelegations, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixDelegationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all delegations for a given account IP prefix.
func (r *PrefixDelegationService) List(ctx context.Context, prefixID string, query PrefixDelegationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AddressingIpamDelegations], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", query.AccountID, prefixID)
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

// List all delegations for a given account IP prefix.
func (r *PrefixDelegationService) ListAutoPaging(ctx context.Context, prefixID string, query PrefixDelegationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AddressingIpamDelegations] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, prefixID, query, opts...))
}

// Delete an account delegation for a given IP prefix.
func (r *PrefixDelegationService) Delete(ctx context.Context, prefixID string, delegationID string, body PrefixDelegationDeleteParams, opts ...option.RequestOption) (res *PrefixDelegationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixDelegationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations/%s", body.AccountID, prefixID, delegationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingIpamDelegations struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                        `json:"parent_prefix_id"`
	JSON           addressingIpamDelegationsJSON `json:"-"`
}

// addressingIpamDelegationsJSON contains the JSON metadata for the struct
// [AddressingIpamDelegations]
type addressingIpamDelegationsJSON struct {
	ID                 apijson.Field
	CIDR               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AddressingIpamDelegations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingIpamDelegationsJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationDeleteResponse struct {
	// Delegation identifier tag.
	ID   string                             `json:"id"`
	JSON prefixDelegationDeleteResponseJSON `json:"-"`
}

// prefixDelegationDeleteResponseJSON contains the JSON metadata for the struct
// [PrefixDelegationDeleteResponse]
type prefixDelegationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR param.Field[string] `json:"cidr,required"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID param.Field[string] `json:"delegated_account_id,required"`
}

func (r PrefixDelegationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixDelegationNewResponseEnvelope struct {
	Errors   []PrefixDelegationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixDelegationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingIpamDelegations                     `json:"result,required"`
	// Whether the API call was successful
	Success PrefixDelegationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    prefixDelegationNewResponseEnvelopeJSON    `json:"-"`
}

// prefixDelegationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixDelegationNewResponseEnvelope]
type prefixDelegationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    prefixDelegationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixDelegationNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixDelegationNewResponseEnvelopeErrors]
type prefixDelegationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    prefixDelegationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixDelegationNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixDelegationNewResponseEnvelopeMessages]
type prefixDelegationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixDelegationNewResponseEnvelopeSuccess bool

const (
	PrefixDelegationNewResponseEnvelopeSuccessTrue PrefixDelegationNewResponseEnvelopeSuccess = true
)

func (r PrefixDelegationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixDelegationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixDelegationListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixDelegationDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixDelegationDeleteResponseEnvelope struct {
	Errors   []PrefixDelegationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixDelegationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   PrefixDelegationDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PrefixDelegationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    prefixDelegationDeleteResponseEnvelopeJSON    `json:"-"`
}

// prefixDelegationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixDelegationDeleteResponseEnvelope]
type prefixDelegationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    prefixDelegationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixDelegationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PrefixDelegationDeleteResponseEnvelopeErrors]
type prefixDelegationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    prefixDelegationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixDelegationDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PrefixDelegationDeleteResponseEnvelopeMessages]
type prefixDelegationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixDelegationDeleteResponseEnvelopeSuccess bool

const (
	PrefixDelegationDeleteResponseEnvelopeSuccessTrue PrefixDelegationDeleteResponseEnvelopeSuccess = true
)

func (r PrefixDelegationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixDelegationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}