// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting

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

// DestinationPagerdutyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDestinationPagerdutyService]
// method instead.
type DestinationPagerdutyService struct {
	Options []option.RequestOption
}

// NewDestinationPagerdutyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDestinationPagerdutyService(opts ...option.RequestOption) (r *DestinationPagerdutyService) {
	r = &DestinationPagerdutyService{}
	r.Options = opts
	return
}

// Creates a new token for integrating with PagerDuty.
func (r *DestinationPagerdutyService) New(ctx context.Context, body DestinationPagerdutyNewParams, opts ...option.RequestOption) (res *DestinationPagerdutyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationPagerdutyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes all the PagerDuty Services connected to the account.
func (r *DestinationPagerdutyService) Delete(ctx context.Context, body DestinationPagerdutyDeleteParams, opts ...option.RequestOption) (res *DestinationPagerdutyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationPagerdutyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a list of all configured PagerDuty services.
func (r *DestinationPagerdutyService) Get(ctx context.Context, query DestinationPagerdutyGetParams, opts ...option.RequestOption) (res *[]AaaPagerduty, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationPagerdutyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Links PagerDuty with the account using the integration token.
func (r *DestinationPagerdutyService) Link(ctx context.Context, tokenID string, query DestinationPagerdutyLinkParams, opts ...option.RequestOption) (res *DestinationPagerdutyLinkResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationPagerdutyLinkResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect/%s", query.AccountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AaaPagerduty struct {
	// UUID
	ID string `json:"id"`
	// The name of the pagerduty service.
	Name string           `json:"name"`
	JSON aaaPagerdutyJSON `json:"-"`
}

// aaaPagerdutyJSON contains the JSON metadata for the struct [AaaPagerduty]
type aaaPagerdutyJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaPagerduty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaPagerdutyJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyNewResponse struct {
	// token in form of UUID
	ID   string                              `json:"id"`
	JSON destinationPagerdutyNewResponseJSON `json:"-"`
}

// destinationPagerdutyNewResponseJSON contains the JSON metadata for the struct
// [DestinationPagerdutyNewResponse]
type destinationPagerdutyNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyNewResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [alerting.DestinationPagerdutyDeleteResponseUnknown],
// [alerting.DestinationPagerdutyDeleteResponseArray] or [shared.UnionString].
type DestinationPagerdutyDeleteResponse interface {
	ImplementsAlertingDestinationPagerdutyDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DestinationPagerdutyDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DestinationPagerdutyDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DestinationPagerdutyDeleteResponseArray []interface{}

func (r DestinationPagerdutyDeleteResponseArray) ImplementsAlertingDestinationPagerdutyDeleteResponse() {
}

type DestinationPagerdutyLinkResponse struct {
	// UUID
	ID   string                               `json:"id"`
	JSON destinationPagerdutyLinkResponseJSON `json:"-"`
}

// destinationPagerdutyLinkResponseJSON contains the JSON metadata for the struct
// [DestinationPagerdutyLinkResponse]
type destinationPagerdutyLinkResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyLinkResponseJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationPagerdutyNewResponseEnvelope struct {
	Errors   []DestinationPagerdutyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationPagerdutyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DestinationPagerdutyNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DestinationPagerdutyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    destinationPagerdutyNewResponseEnvelopeJSON    `json:"-"`
}

// destinationPagerdutyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationPagerdutyNewResponseEnvelope]
type destinationPagerdutyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    destinationPagerdutyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationPagerdutyNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DestinationPagerdutyNewResponseEnvelopeErrors]
type destinationPagerdutyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    destinationPagerdutyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationPagerdutyNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DestinationPagerdutyNewResponseEnvelopeMessages]
type destinationPagerdutyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationPagerdutyNewResponseEnvelopeSuccess bool

const (
	DestinationPagerdutyNewResponseEnvelopeSuccessTrue DestinationPagerdutyNewResponseEnvelopeSuccess = true
)

func (r DestinationPagerdutyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationPagerdutyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationPagerdutyDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationPagerdutyDeleteResponseEnvelope struct {
	Errors   []DestinationPagerdutyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationPagerdutyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   DestinationPagerdutyDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DestinationPagerdutyDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DestinationPagerdutyDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       destinationPagerdutyDeleteResponseEnvelopeJSON       `json:"-"`
}

// destinationPagerdutyDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [DestinationPagerdutyDeleteResponseEnvelope]
type destinationPagerdutyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    destinationPagerdutyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationPagerdutyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DestinationPagerdutyDeleteResponseEnvelopeErrors]
type destinationPagerdutyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    destinationPagerdutyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationPagerdutyDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DestinationPagerdutyDeleteResponseEnvelopeMessages]
type destinationPagerdutyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationPagerdutyDeleteResponseEnvelopeSuccess bool

const (
	DestinationPagerdutyDeleteResponseEnvelopeSuccessTrue DestinationPagerdutyDeleteResponseEnvelopeSuccess = true
)

func (r DestinationPagerdutyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationPagerdutyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationPagerdutyDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       destinationPagerdutyDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// destinationPagerdutyDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DestinationPagerdutyDeleteResponseEnvelopeResultInfo]
type destinationPagerdutyDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationPagerdutyGetResponseEnvelope struct {
	Errors   []DestinationPagerdutyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationPagerdutyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []AaaPagerduty                                    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DestinationPagerdutyGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DestinationPagerdutyGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       destinationPagerdutyGetResponseEnvelopeJSON       `json:"-"`
}

// destinationPagerdutyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationPagerdutyGetResponseEnvelope]
type destinationPagerdutyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    destinationPagerdutyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationPagerdutyGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DestinationPagerdutyGetResponseEnvelopeErrors]
type destinationPagerdutyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    destinationPagerdutyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationPagerdutyGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DestinationPagerdutyGetResponseEnvelopeMessages]
type destinationPagerdutyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationPagerdutyGetResponseEnvelopeSuccess bool

const (
	DestinationPagerdutyGetResponseEnvelopeSuccessTrue DestinationPagerdutyGetResponseEnvelopeSuccess = true
)

func (r DestinationPagerdutyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationPagerdutyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationPagerdutyGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       destinationPagerdutyGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// destinationPagerdutyGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DestinationPagerdutyGetResponseEnvelopeResultInfo]
type destinationPagerdutyGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyLinkParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationPagerdutyLinkResponseEnvelope struct {
	Errors   []DestinationPagerdutyLinkResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationPagerdutyLinkResponseEnvelopeMessages `json:"messages,required"`
	Result   DestinationPagerdutyLinkResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DestinationPagerdutyLinkResponseEnvelopeSuccess `json:"success,required"`
	JSON    destinationPagerdutyLinkResponseEnvelopeJSON    `json:"-"`
}

// destinationPagerdutyLinkResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationPagerdutyLinkResponseEnvelope]
type destinationPagerdutyLinkResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyLinkResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyLinkResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyLinkResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    destinationPagerdutyLinkResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationPagerdutyLinkResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DestinationPagerdutyLinkResponseEnvelopeErrors]
type destinationPagerdutyLinkResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyLinkResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyLinkResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyLinkResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    destinationPagerdutyLinkResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationPagerdutyLinkResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DestinationPagerdutyLinkResponseEnvelopeMessages]
type destinationPagerdutyLinkResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyLinkResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyLinkResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationPagerdutyLinkResponseEnvelopeSuccess bool

const (
	DestinationPagerdutyLinkResponseEnvelopeSuccessTrue DestinationPagerdutyLinkResponseEnvelopeSuccess = true
)

func (r DestinationPagerdutyLinkResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationPagerdutyLinkResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}