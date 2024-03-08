// File generated from our OpenAPI spec by Stainless.

package durable_objects

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// NamespaceObjectService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNamespaceObjectService] method
// instead.
type NamespaceObjectService struct {
	Options []option.RequestOption
}

// NewNamespaceObjectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceObjectService(opts ...option.RequestOption) (r *NamespaceObjectService) {
	r = &NamespaceObjectService{}
	r.Options = opts
	return
}

// Returns the Durable Objects in a given namespace.
func (r *NamespaceObjectService) List(ctx context.Context, id string, params NamespaceObjectListParams, opts ...option.RequestOption) (res *[]NamespaceObjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceObjectListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces/%s/objects", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NamespaceObjectListResponse struct {
	// ID of the Durable Object.
	ID string `json:"id"`
	// Whether the Durable Object has stored data.
	HasStoredData bool                            `json:"hasStoredData"`
	JSON          namespaceObjectListResponseJSON `json:"-"`
}

// namespaceObjectListResponseJSON contains the JSON metadata for the struct
// [NamespaceObjectListResponse]
type namespaceObjectListResponseJSON struct {
	ID            apijson.Field
	HasStoredData apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *NamespaceObjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceObjectListResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceObjectListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records. A valid value for the cursor can be obtained from the
	// cursors object in the result_info structure.
	Cursor param.Field[string] `query:"cursor"`
	// The number of objects to return. The cursor attribute may be used to iterate
	// over the next batch of objects if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [NamespaceObjectListParams]'s query parameters as
// `url.Values`.
func (r NamespaceObjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NamespaceObjectListResponseEnvelope struct {
	Errors   []NamespaceObjectListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceObjectListResponseEnvelopeMessages `json:"messages,required"`
	Result   []NamespaceObjectListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    NamespaceObjectListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo NamespaceObjectListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       namespaceObjectListResponseEnvelopeJSON       `json:"-"`
}

// namespaceObjectListResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceObjectListResponseEnvelope]
type namespaceObjectListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceObjectListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceObjectListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceObjectListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    namespaceObjectListResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceObjectListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [NamespaceObjectListResponseEnvelopeErrors]
type namespaceObjectListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceObjectListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceObjectListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceObjectListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    namespaceObjectListResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceObjectListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [NamespaceObjectListResponseEnvelopeMessages]
type namespaceObjectListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceObjectListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceObjectListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceObjectListResponseEnvelopeSuccess bool

const (
	NamespaceObjectListResponseEnvelopeSuccessTrue NamespaceObjectListResponseEnvelopeSuccess = true
)

type NamespaceObjectListResponseEnvelopeResultInfo struct {
	// Total results returned based on your list parameters.
	Count float64 `json:"count"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records. A valid value for the cursor can be obtained from the
	// cursors object in the result_info structure.
	Cursor string `json:"cursor"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       namespaceObjectListResponseEnvelopeResultInfoJSON `json:"-"`
}

// namespaceObjectListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [NamespaceObjectListResponseEnvelopeResultInfo]
type namespaceObjectListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceObjectListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceObjectListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
