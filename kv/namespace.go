// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// NamespaceService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewNamespaceService] method instead.
type NamespaceService struct {
	Options  []option.RequestOption
	Bulk     *NamespaceBulkService
	Keys     *NamespaceKeyService
	Metadata *NamespaceMetadataService
	Values   *NamespaceValueService
}

// NewNamespaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceService(opts ...option.RequestOption) (r *NamespaceService) {
	r = &NamespaceService{}
	r.Options = opts
	r.Bulk = NewNamespaceBulkService(opts...)
	r.Keys = NewNamespaceKeyService(opts...)
	r.Metadata = NewNamespaceMetadataService(opts...)
	r.Values = NewNamespaceValueService(opts...)
	return
}

// Creates a namespace under the given title. A `400` is returned if the account
// already owns a namespace with this title. A namespace must be explicitly deleted
// to be replaced.
func (r *NamespaceService) New(ctx context.Context, params NamespaceNewParams, opts ...option.RequestOption) (res *WorkersKVNamespace, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modifies a namespace's title.
func (r *NamespaceService) Update(ctx context.Context, namespaceID string, params NamespaceUpdateParams, opts ...option.RequestOption) (res *NamespaceUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", params.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the namespaces owned by an account.
func (r *NamespaceService) List(ctx context.Context, params NamespaceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WorkersKVNamespace], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", params.AccountID)
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

// Returns the namespaces owned by an account.
func (r *NamespaceService) ListAutoPaging(ctx context.Context, params NamespaceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WorkersKVNamespace] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes the namespace corresponding to the given ID.
func (r *NamespaceService) Delete(ctx context.Context, namespaceID string, body NamespaceDeleteParams, opts ...option.RequestOption) (res *NamespaceDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", body.AccountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersKVNamespace struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool                   `json:"supports_url_encoding"`
	JSON                workersKVNamespaceJSON `json:"-"`
}

// workersKVNamespaceJSON contains the JSON metadata for the struct
// [WorkersKVNamespace]
type workersKVNamespaceJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *WorkersKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [kv.NamespaceUpdateResponseUnknown] or [shared.UnionString].
type NamespaceUpdateResponse interface {
	ImplementsKVNamespaceUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [kv.NamespaceDeleteResponseUnknown] or [shared.UnionString].
type NamespaceDeleteResponse interface {
	ImplementsKVNamespaceDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NamespaceNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r NamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceNewResponseEnvelope struct {
	Errors   []NamespaceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceNewResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersKVNamespace                     `json:"result,required"`
	// Whether the API call was successful
	Success NamespaceNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    namespaceNewResponseEnvelopeJSON    `json:"-"`
}

// namespaceNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceNewResponseEnvelope]
type namespaceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    namespaceNewResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [NamespaceNewResponseEnvelopeErrors]
type namespaceNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    namespaceNewResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [NamespaceNewResponseEnvelopeMessages]
type namespaceNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceNewResponseEnvelopeSuccess bool

const (
	NamespaceNewResponseEnvelopeSuccessTrue NamespaceNewResponseEnvelopeSuccess = true
)

func (r NamespaceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r NamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceUpdateResponseEnvelope struct {
	Errors   []NamespaceUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   NamespaceUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NamespaceUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    namespaceUpdateResponseEnvelopeJSON    `json:"-"`
}

// namespaceUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceUpdateResponseEnvelope]
type namespaceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    namespaceUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [NamespaceUpdateResponseEnvelopeErrors]
type namespaceUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    namespaceUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [NamespaceUpdateResponseEnvelopeMessages]
type namespaceUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceUpdateResponseEnvelopeSuccess bool

const (
	NamespaceUpdateResponseEnvelopeSuccessTrue NamespaceUpdateResponseEnvelopeSuccess = true
)

func (r NamespaceUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to order namespaces.
	Direction param.Field[NamespaceListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[NamespaceListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [NamespaceListParams]'s query parameters as `url.Values`.
func (r NamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order namespaces.
type NamespaceListParamsDirection string

const (
	NamespaceListParamsDirectionAsc  NamespaceListParamsDirection = "asc"
	NamespaceListParamsDirectionDesc NamespaceListParamsDirection = "desc"
)

func (r NamespaceListParamsDirection) IsKnown() bool {
	switch r {
	case NamespaceListParamsDirectionAsc, NamespaceListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by.
type NamespaceListParamsOrder string

const (
	NamespaceListParamsOrderID    NamespaceListParamsOrder = "id"
	NamespaceListParamsOrderTitle NamespaceListParamsOrder = "title"
)

func (r NamespaceListParamsOrder) IsKnown() bool {
	switch r {
	case NamespaceListParamsOrderID, NamespaceListParamsOrderTitle:
		return true
	}
	return false
}

type NamespaceDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type NamespaceDeleteResponseEnvelope struct {
	Errors   []NamespaceDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   NamespaceDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NamespaceDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    namespaceDeleteResponseEnvelopeJSON    `json:"-"`
}

// namespaceDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceDeleteResponseEnvelope]
type namespaceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    namespaceDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [NamespaceDeleteResponseEnvelopeErrors]
type namespaceDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    namespaceDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [NamespaceDeleteResponseEnvelopeMessages]
type namespaceDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceDeleteResponseEnvelopeSuccess bool

const (
	NamespaceDeleteResponseEnvelopeSuccessTrue NamespaceDeleteResponseEnvelopeSuccess = true
)

func (r NamespaceDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}