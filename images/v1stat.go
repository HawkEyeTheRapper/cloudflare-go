// File generated from our OpenAPI spec by Stainless.

package images

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// V1StatService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewV1StatService] method instead.
type V1StatService struct {
	Options []option.RequestOption
}

// NewV1StatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1StatService(opts ...option.RequestOption) (r *V1StatService) {
	r = &V1StatService{}
	r.Options = opts
	return
}

// Fetch usage statistics details for Cloudflare Images.
func (r *V1StatService) Get(ctx context.Context, query V1StatGetParams, opts ...option.RequestOption) (res *V1StatGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V1StatGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/stats", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type V1StatGetResponse struct {
	Count V1StatGetResponseCount `json:"count"`
	JSON  v1StatGetResponseJSON  `json:"-"`
}

// v1StatGetResponseJSON contains the JSON metadata for the struct
// [V1StatGetResponse]
type v1StatGetResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1StatGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1StatGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1StatGetResponseCount struct {
	// Cloudflare Images allowed usage.
	Allowed float64 `json:"allowed"`
	// Cloudflare Images current usage.
	Current float64                    `json:"current"`
	JSON    v1StatGetResponseCountJSON `json:"-"`
}

// v1StatGetResponseCountJSON contains the JSON metadata for the struct
// [V1StatGetResponseCount]
type v1StatGetResponseCountJSON struct {
	Allowed     apijson.Field
	Current     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1StatGetResponseCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1StatGetResponseCountJSON) RawJSON() string {
	return r.raw
}

type V1StatGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1StatGetResponseEnvelope struct {
	Errors   []V1StatGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1StatGetResponseEnvelopeMessages `json:"messages,required"`
	Result   V1StatGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1StatGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1StatGetResponseEnvelopeJSON    `json:"-"`
}

// v1StatGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1StatGetResponseEnvelope]
type v1StatGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1StatGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1StatGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1StatGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    v1StatGetResponseEnvelopeErrorsJSON `json:"-"`
}

// v1StatGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1StatGetResponseEnvelopeErrors]
type v1StatGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1StatGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1StatGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1StatGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    v1StatGetResponseEnvelopeMessagesJSON `json:"-"`
}

// v1StatGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [V1StatGetResponseEnvelopeMessages]
type v1StatGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1StatGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1StatGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1StatGetResponseEnvelopeSuccess bool

const (
	V1StatGetResponseEnvelopeSuccessTrue V1StatGetResponseEnvelopeSuccess = true
)
