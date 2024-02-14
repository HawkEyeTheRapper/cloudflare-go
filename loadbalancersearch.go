// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// LoadBalancerSearchService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerSearchService] method
// instead.
type LoadBalancerSearchService struct {
	Options []option.RequestOption
}

// NewLoadBalancerSearchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerSearchService(opts ...option.RequestOption) (r *LoadBalancerSearchService) {
	r = &LoadBalancerSearchService{}
	r.Options = opts
	return
}

// Search for Load Balancing resources.
func (r *LoadBalancerSearchService) List(ctx context.Context, accountIdentifier string, query LoadBalancerSearchListParams, opts ...option.RequestOption) (res *[]LoadBalancerSearchListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerSearchListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/search", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerSearchListResponse = interface{}

type LoadBalancerSearchListParams struct {
	Page         param.Field[interface{}]                              `query:"page"`
	PerPage      param.Field[interface{}]                              `query:"per_page"`
	SearchParams param.Field[LoadBalancerSearchListParamsSearchParams] `query:"search_params"`
}

// URLQuery serializes [LoadBalancerSearchListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerSearchListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LoadBalancerSearchListParamsSearchParams struct {
	// Search query term.
	Query param.Field[string] `query:"query"`
	// The type of references to include ("\*" for all).
	References param.Field[LoadBalancerSearchListParamsSearchParamsReferences] `query:"references"`
}

// URLQuery serializes [LoadBalancerSearchListParamsSearchParams]'s query
// parameters as `url.Values`.
func (r LoadBalancerSearchListParamsSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of references to include ("\*" for all).
type LoadBalancerSearchListParamsSearchParamsReferences string

const (
	LoadBalancerSearchListParamsSearchParamsReferencesEmpty    LoadBalancerSearchListParamsSearchParamsReferences = ""
	LoadBalancerSearchListParamsSearchParamsReferencesStar     LoadBalancerSearchListParamsSearchParamsReferences = "*"
	LoadBalancerSearchListParamsSearchParamsReferencesReferral LoadBalancerSearchListParamsSearchParamsReferences = "referral"
	LoadBalancerSearchListParamsSearchParamsReferencesReferrer LoadBalancerSearchListParamsSearchParamsReferences = "referrer"
)

type LoadBalancerSearchListResponseEnvelope struct {
	Errors   []LoadBalancerSearchListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerSearchListResponseEnvelopeMessages `json:"messages,required"`
	Result   []LoadBalancerSearchListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerSearchListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerSearchListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerSearchListResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerSearchListResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerSearchListResponseEnvelope]
type loadBalancerSearchListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerSearchListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerSearchListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    loadBalancerSearchListResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerSearchListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerSearchListResponseEnvelopeErrors]
type loadBalancerSearchListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerSearchListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerSearchListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    loadBalancerSearchListResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerSearchListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerSearchListResponseEnvelopeMessages]
type loadBalancerSearchListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerSearchListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerSearchListResponseEnvelopeSuccess bool

const (
	LoadBalancerSearchListResponseEnvelopeSuccessTrue LoadBalancerSearchListResponseEnvelopeSuccess = true
)

type LoadBalancerSearchListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       loadBalancerSearchListResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerSearchListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [LoadBalancerSearchListResponseEnvelopeResultInfo]
type loadBalancerSearchListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerSearchListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}