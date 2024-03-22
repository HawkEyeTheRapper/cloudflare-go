// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// DestinationWebhookService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDestinationWebhookService] method
// instead.
type DestinationWebhookService struct {
	Options []option.RequestOption
}

// NewDestinationWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDestinationWebhookService(opts ...option.RequestOption) (r *DestinationWebhookService) {
	r = &DestinationWebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook destination.
func (r *DestinationWebhookService) New(ctx context.Context, params DestinationWebhookNewParams, opts ...option.RequestOption) (res *DestinationWebhookNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationWebhookNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a webhook destination.
func (r *DestinationWebhookService) Update(ctx context.Context, webhookID string, params DestinationWebhookUpdateParams, opts ...option.RequestOption) (res *DestinationWebhookUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationWebhookUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", params.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a list of all configured webhook destinations.
func (r *DestinationWebhookService) List(ctx context.Context, query DestinationWebhookListParams, opts ...option.RequestOption) (res *[]AaaWebhooks, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationWebhookListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured webhook destination.
func (r *DestinationWebhookService) Delete(ctx context.Context, webhookID string, body DestinationWebhookDeleteParams, opts ...option.RequestOption) (res *DestinationWebhookDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationWebhookDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", body.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for a single webhooks destination.
func (r *DestinationWebhookService) Get(ctx context.Context, webhookID string, query DestinationWebhookGetParams, opts ...option.RequestOption) (res *AaaWebhooks, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationWebhookGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", query.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AaaWebhooks struct {
	// The unique identifier of a webhook
	ID string `json:"id"`
	// Timestamp of when the webhook destination was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of the last time an attempt to dispatch a notification to this webhook
	// failed.
	LastFailure time.Time `json:"last_failure" format:"date-time"`
	// Timestamp of the last time Cloudflare was able to successfully dispatch a
	// notification using this webhook.
	LastSuccess time.Time `json:"last_success" format:"date-time"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name string `json:"name"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret string `json:"secret"`
	// Type of webhook endpoint.
	Type AaaWebhooksType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string          `json:"url"`
	JSON aaaWebhooksJSON `json:"-"`
}

// aaaWebhooksJSON contains the JSON metadata for the struct [AaaWebhooks]
type aaaWebhooksJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	LastFailure apijson.Field
	LastSuccess apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaWebhooks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaWebhooksJSON) RawJSON() string {
	return r.raw
}

// Type of webhook endpoint.
type AaaWebhooksType string

const (
	AaaWebhooksTypeSlack   AaaWebhooksType = "slack"
	AaaWebhooksTypeGeneric AaaWebhooksType = "generic"
	AaaWebhooksTypeGchat   AaaWebhooksType = "gchat"
)

func (r AaaWebhooksType) IsKnown() bool {
	switch r {
	case AaaWebhooksTypeSlack, AaaWebhooksTypeGeneric, AaaWebhooksTypeGchat:
		return true
	}
	return false
}

type DestinationWebhookNewResponse struct {
	// UUID
	ID   string                            `json:"id"`
	JSON destinationWebhookNewResponseJSON `json:"-"`
}

// destinationWebhookNewResponseJSON contains the JSON metadata for the struct
// [DestinationWebhookNewResponse]
type destinationWebhookNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookNewResponseJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookUpdateResponse struct {
	// UUID
	ID   string                               `json:"id"`
	JSON destinationWebhookUpdateResponseJSON `json:"-"`
}

// destinationWebhookUpdateResponseJSON contains the JSON metadata for the struct
// [DestinationWebhookUpdateResponse]
type destinationWebhookUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [alerting.DestinationWebhookDeleteResponseUnknown],
// [alerting.DestinationWebhookDeleteResponseArray] or [shared.UnionString].
type DestinationWebhookDeleteResponse interface {
	ImplementsAlertingDestinationWebhookDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DestinationWebhookDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DestinationWebhookDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DestinationWebhookDeleteResponseArray []interface{}

func (r DestinationWebhookDeleteResponseArray) ImplementsAlertingDestinationWebhookDeleteResponse() {}

type DestinationWebhookNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r DestinationWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DestinationWebhookNewResponseEnvelope struct {
	Errors   []DestinationWebhookNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationWebhookNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DestinationWebhookNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DestinationWebhookNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    destinationWebhookNewResponseEnvelopeJSON    `json:"-"`
}

// destinationWebhookNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookNewResponseEnvelope]
type destinationWebhookNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    destinationWebhookNewResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationWebhookNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DestinationWebhookNewResponseEnvelopeErrors]
type destinationWebhookNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    destinationWebhookNewResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationWebhookNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DestinationWebhookNewResponseEnvelopeMessages]
type destinationWebhookNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookNewResponseEnvelopeSuccess bool

const (
	DestinationWebhookNewResponseEnvelopeSuccessTrue DestinationWebhookNewResponseEnvelopeSuccess = true
)

func (r DestinationWebhookNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationWebhookUpdateParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r DestinationWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DestinationWebhookUpdateResponseEnvelope struct {
	Errors   []DestinationWebhookUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationWebhookUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DestinationWebhookUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DestinationWebhookUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    destinationWebhookUpdateResponseEnvelopeJSON    `json:"-"`
}

// destinationWebhookUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookUpdateResponseEnvelope]
type destinationWebhookUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    destinationWebhookUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationWebhookUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DestinationWebhookUpdateResponseEnvelopeErrors]
type destinationWebhookUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    destinationWebhookUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationWebhookUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DestinationWebhookUpdateResponseEnvelopeMessages]
type destinationWebhookUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookUpdateResponseEnvelopeSuccess bool

const (
	DestinationWebhookUpdateResponseEnvelopeSuccessTrue DestinationWebhookUpdateResponseEnvelopeSuccess = true
)

func (r DestinationWebhookUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationWebhookListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationWebhookListResponseEnvelope struct {
	Errors   []DestinationWebhookListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationWebhookListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AaaWebhooks                                    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DestinationWebhookListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DestinationWebhookListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       destinationWebhookListResponseEnvelopeJSON       `json:"-"`
}

// destinationWebhookListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookListResponseEnvelope]
type destinationWebhookListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    destinationWebhookListResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationWebhookListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DestinationWebhookListResponseEnvelopeErrors]
type destinationWebhookListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    destinationWebhookListResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationWebhookListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DestinationWebhookListResponseEnvelopeMessages]
type destinationWebhookListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookListResponseEnvelopeSuccess bool

const (
	DestinationWebhookListResponseEnvelopeSuccessTrue DestinationWebhookListResponseEnvelopeSuccess = true
)

func (r DestinationWebhookListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationWebhookListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       destinationWebhookListResponseEnvelopeResultInfoJSON `json:"-"`
}

// destinationWebhookListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DestinationWebhookListResponseEnvelopeResultInfo]
type destinationWebhookListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationWebhookDeleteResponseEnvelope struct {
	Errors   []DestinationWebhookDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationWebhookDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   DestinationWebhookDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DestinationWebhookDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DestinationWebhookDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       destinationWebhookDeleteResponseEnvelopeJSON       `json:"-"`
}

// destinationWebhookDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookDeleteResponseEnvelope]
type destinationWebhookDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    destinationWebhookDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationWebhookDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DestinationWebhookDeleteResponseEnvelopeErrors]
type destinationWebhookDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    destinationWebhookDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationWebhookDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DestinationWebhookDeleteResponseEnvelopeMessages]
type destinationWebhookDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookDeleteResponseEnvelopeSuccess bool

const (
	DestinationWebhookDeleteResponseEnvelopeSuccessTrue DestinationWebhookDeleteResponseEnvelopeSuccess = true
)

func (r DestinationWebhookDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationWebhookDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       destinationWebhookDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// destinationWebhookDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DestinationWebhookDeleteResponseEnvelopeResultInfo]
type destinationWebhookDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationWebhookGetResponseEnvelope struct {
	Errors   []DestinationWebhookGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationWebhookGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AaaWebhooks                                     `json:"result,required"`
	// Whether the API call was successful
	Success DestinationWebhookGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    destinationWebhookGetResponseEnvelopeJSON    `json:"-"`
}

// destinationWebhookGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookGetResponseEnvelope]
type destinationWebhookGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    destinationWebhookGetResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationWebhookGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DestinationWebhookGetResponseEnvelopeErrors]
type destinationWebhookGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    destinationWebhookGetResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationWebhookGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DestinationWebhookGetResponseEnvelopeMessages]
type destinationWebhookGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookGetResponseEnvelopeSuccess bool

const (
	DestinationWebhookGetResponseEnvelopeSuccessTrue DestinationWebhookGetResponseEnvelopeSuccess = true
)

func (r DestinationWebhookGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}