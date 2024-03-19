// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ListBulkOperationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewListBulkOperationService] method
// instead.
type ListBulkOperationService struct {
	Options []option.RequestOption
}

// NewListBulkOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewListBulkOperationService(opts ...option.RequestOption) (r *ListBulkOperationService) {
	r = &ListBulkOperationService{}
	r.Options = opts
	return
}

// Gets the current status of an asynchronous operation on a list.
//
// The `status` property can have one of the following values: `pending`,
// `running`, `completed`, or `failed`. If the status is `failed`, the `error`
// property will contain a message describing the error.
func (r *ListBulkOperationService) Get(ctx context.Context, accountIdentifier string, operationID string, opts ...option.RequestOption) (res *[]ListBulkOperationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ListBulkOperationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/bulk_operations/%s", accountIdentifier, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ListBulkOperationGetResponse = interface{}

type ListBulkOperationGetResponseEnvelope struct {
	Errors   []ListBulkOperationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ListBulkOperationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []ListBulkOperationGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ListBulkOperationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    listBulkOperationGetResponseEnvelopeJSON    `json:"-"`
}

// listBulkOperationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ListBulkOperationGetResponseEnvelope]
type listBulkOperationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListBulkOperationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listBulkOperationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ListBulkOperationGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    listBulkOperationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// listBulkOperationGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ListBulkOperationGetResponseEnvelopeErrors]
type listBulkOperationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListBulkOperationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listBulkOperationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ListBulkOperationGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    listBulkOperationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// listBulkOperationGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ListBulkOperationGetResponseEnvelopeMessages]
type listBulkOperationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListBulkOperationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listBulkOperationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ListBulkOperationGetResponseEnvelopeSuccess bool

const (
	ListBulkOperationGetResponseEnvelopeSuccessTrue ListBulkOperationGetResponseEnvelopeSuccess = true
)
