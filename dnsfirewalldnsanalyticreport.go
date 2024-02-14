// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DNSFirewallDNSAnalyticReportService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDNSFirewallDNSAnalyticReportService] method instead.
type DNSFirewallDNSAnalyticReportService struct {
	Options []option.RequestOption
	Bytimes *DNSFirewallDNSAnalyticReportBytimeService
}

// NewDNSFirewallDNSAnalyticReportService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDNSFirewallDNSAnalyticReportService(opts ...option.RequestOption) (r *DNSFirewallDNSAnalyticReportService) {
	r = &DNSFirewallDNSAnalyticReportService{}
	r.Options = opts
	r.Bytimes = NewDNSFirewallDNSAnalyticReportBytimeService(opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *DNSFirewallDNSAnalyticReportService) List(ctx context.Context, accountIdentifier string, identifier string, query DNSFirewallDNSAnalyticReportListParams, opts ...option.RequestOption) (res *DNSFirewallDNSAnalyticReportListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSFirewallDNSAnalyticReportListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSFirewallDNSAnalyticReportListResponse struct {
	// Array with one row per combination of dimension values.
	Data []DNSFirewallDNSAnalyticReportListResponseData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                                   `json:"min,required"`
	Query DNSFirewallDNSAnalyticReportListResponseQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                                  `json:"totals,required"`
	JSON   dnsFirewallDNSAnalyticReportListResponseJSON `json:"-"`
}

// dnsFirewallDNSAnalyticReportListResponseJSON contains the JSON metadata for the
// struct [DNSFirewallDNSAnalyticReportListResponse]
type dnsFirewallDNSAnalyticReportListResponseJSON struct {
	Data        apijson.Field
	DataLag     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	Query       apijson.Field
	Rows        apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDNSAnalyticReportListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDNSAnalyticReportListResponseData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is a single value.
	Metrics []float64                                        `json:"metrics,required"`
	JSON    dnsFirewallDNSAnalyticReportListResponseDataJSON `json:"-"`
}

// dnsFirewallDNSAnalyticReportListResponseDataJSON contains the JSON metadata for
// the struct [DNSFirewallDNSAnalyticReportListResponseData]
type dnsFirewallDNSAnalyticReportListResponseDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDNSAnalyticReportListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDNSAnalyticReportListResponseQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string                                          `json:"sort"`
	JSON dnsFirewallDNSAnalyticReportListResponseQueryJSON `json:"-"`
}

// dnsFirewallDNSAnalyticReportListResponseQueryJSON contains the JSON metadata for
// the struct [DNSFirewallDNSAnalyticReportListResponseQuery]
type dnsFirewallDNSAnalyticReportListResponseQueryJSON struct {
	Dimensions  apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	Until       apijson.Field
	Filters     apijson.Field
	Sort        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDNSAnalyticReportListResponseQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDNSAnalyticReportListParams struct {
	// A comma-separated list of dimensions to group results by.
	Dimensions param.Field[string] `query:"dimensions"`
	// Segmentation filter in 'attribute operator value' format.
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// A comma-separated list of metrics to query.
	Metrics param.Field[string] `query:"metrics"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// A comma-separated list of dimensions to sort by, where each dimension may be
	// prefixed by - (descending) or + (ascending).
	Sort param.Field[string] `query:"sort"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [DNSFirewallDNSAnalyticReportListParams]'s query parameters
// as `url.Values`.
func (r DNSFirewallDNSAnalyticReportListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DNSFirewallDNSAnalyticReportListResponseEnvelope struct {
	Errors   []DNSFirewallDNSAnalyticReportListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallDNSAnalyticReportListResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSFirewallDNSAnalyticReportListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSFirewallDNSAnalyticReportListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsFirewallDNSAnalyticReportListResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallDNSAnalyticReportListResponseEnvelopeJSON contains the JSON metadata
// for the struct [DNSFirewallDNSAnalyticReportListResponseEnvelope]
type dnsFirewallDNSAnalyticReportListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDNSAnalyticReportListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDNSAnalyticReportListResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    dnsFirewallDNSAnalyticReportListResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsFirewallDNSAnalyticReportListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DNSFirewallDNSAnalyticReportListResponseEnvelopeErrors]
type dnsFirewallDNSAnalyticReportListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDNSAnalyticReportListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDNSAnalyticReportListResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    dnsFirewallDNSAnalyticReportListResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsFirewallDNSAnalyticReportListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DNSFirewallDNSAnalyticReportListResponseEnvelopeMessages]
type dnsFirewallDNSAnalyticReportListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDNSAnalyticReportListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallDNSAnalyticReportListResponseEnvelopeSuccess bool

const (
	DNSFirewallDNSAnalyticReportListResponseEnvelopeSuccessTrue DNSFirewallDNSAnalyticReportListResponseEnvelopeSuccess = true
)