// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarQualityIqiService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarQualityIqiService] method
// instead.
type RadarQualityIqiService struct {
	Options          []option.RequestOption
	TimeseriesGroups *RadarQualityIqiTimeseriesGroupService
}

// NewRadarQualityIqiService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarQualityIqiService(opts ...option.RequestOption) (r *RadarQualityIqiService) {
	r = &RadarQualityIqiService{}
	r.Options = opts
	r.TimeseriesGroups = NewRadarQualityIqiTimeseriesGroupService(opts...)
	return
}

// Get a summary (percentiles) of bandwidth, latency or DNS response time from the
// Radar Internet Quality Index (IQI).
func (r *RadarQualityIqiService) Get(ctx context.Context, query RadarQualityIqiGetParams, opts ...option.RequestOption) (res *RadarQualityIqiGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualityIqiGetResponseEnvelope
	path := "radar/quality/iqi/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarQualityIqiGetResponse struct {
	Meta     RadarQualityIqiGetResponseMeta     `json:"meta,required"`
	Summary0 RadarQualityIqiGetResponseSummary0 `json:"summary_0,required"`
	JSON     radarQualityIqiGetResponseJSON     `json:"-"`
}

// radarQualityIqiGetResponseJSON contains the JSON metadata for the struct
// [RadarQualityIqiGetResponse]
type radarQualityIqiGetResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiGetResponseMeta struct {
	DateRange      []RadarQualityIqiGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                       `json:"lastUpdated,required"`
	Normalization  string                                       `json:"normalization,required"`
	ConfidenceInfo RadarQualityIqiGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualityIqiGetResponseMetaJSON           `json:"-"`
}

// radarQualityIqiGetResponseMetaJSON contains the JSON metadata for the struct
// [RadarQualityIqiGetResponseMeta]
type radarQualityIqiGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualityIqiGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      radarQualityIqiGetResponseMetaDateRangeJSON `json:"-"`
}

// radarQualityIqiGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarQualityIqiGetResponseMetaDateRange]
type radarQualityIqiGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiGetResponseMetaConfidenceInfo struct {
	Annotations []RadarQualityIqiGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                    `json:"level"`
	JSON        radarQualityIqiGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualityIqiGetResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [RadarQualityIqiGetResponseMetaConfidenceInfo]
type radarQualityIqiGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                     `json:"dataSource,required"`
	Description     string                                                     `json:"description,required"`
	EventType       string                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                  `json:"startTime" format:"date-time"`
	JSON            radarQualityIqiGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualityIqiGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [RadarQualityIqiGetResponseMetaConfidenceInfoAnnotation]
type radarQualityIqiGetResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarQualityIqiGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiGetResponseSummary0 struct {
	P25  string                                 `json:"p25,required"`
	P50  string                                 `json:"p50,required"`
	P75  string                                 `json:"p75,required"`
	JSON radarQualityIqiGetResponseSummary0JSON `json:"-"`
}

// radarQualityIqiGetResponseSummary0JSON contains the JSON metadata for the struct
// [RadarQualityIqiGetResponseSummary0]
type radarQualityIqiGetResponseSummary0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiGetParams struct {
	// Which metric to return: bandwidth, latency or DNS response time.
	Metric param.Field[RadarQualityIqiGetParamsMetric] `query:"metric,required"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarQualityIqiGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualityIqiGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualityIqiGetParams]'s query parameters as
// `url.Values`.
func (r RadarQualityIqiGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which metric to return: bandwidth, latency or DNS response time.
type RadarQualityIqiGetParamsMetric string

const (
	RadarQualityIqiGetParamsMetricBandwidth RadarQualityIqiGetParamsMetric = "BANDWIDTH"
	RadarQualityIqiGetParamsMetricDNS       RadarQualityIqiGetParamsMetric = "DNS"
	RadarQualityIqiGetParamsMetricLatency   RadarQualityIqiGetParamsMetric = "LATENCY"
)

type RadarQualityIqiGetParamsDateRange string

const (
	RadarQualityIqiGetParamsDateRange1d         RadarQualityIqiGetParamsDateRange = "1d"
	RadarQualityIqiGetParamsDateRange2d         RadarQualityIqiGetParamsDateRange = "2d"
	RadarQualityIqiGetParamsDateRange7d         RadarQualityIqiGetParamsDateRange = "7d"
	RadarQualityIqiGetParamsDateRange14d        RadarQualityIqiGetParamsDateRange = "14d"
	RadarQualityIqiGetParamsDateRange28d        RadarQualityIqiGetParamsDateRange = "28d"
	RadarQualityIqiGetParamsDateRange12w        RadarQualityIqiGetParamsDateRange = "12w"
	RadarQualityIqiGetParamsDateRange24w        RadarQualityIqiGetParamsDateRange = "24w"
	RadarQualityIqiGetParamsDateRange52w        RadarQualityIqiGetParamsDateRange = "52w"
	RadarQualityIqiGetParamsDateRange1dControl  RadarQualityIqiGetParamsDateRange = "1dControl"
	RadarQualityIqiGetParamsDateRange2dControl  RadarQualityIqiGetParamsDateRange = "2dControl"
	RadarQualityIqiGetParamsDateRange7dControl  RadarQualityIqiGetParamsDateRange = "7dControl"
	RadarQualityIqiGetParamsDateRange14dControl RadarQualityIqiGetParamsDateRange = "14dControl"
	RadarQualityIqiGetParamsDateRange28dControl RadarQualityIqiGetParamsDateRange = "28dControl"
	RadarQualityIqiGetParamsDateRange12wControl RadarQualityIqiGetParamsDateRange = "12wControl"
	RadarQualityIqiGetParamsDateRange24wControl RadarQualityIqiGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarQualityIqiGetParamsFormat string

const (
	RadarQualityIqiGetParamsFormatJson RadarQualityIqiGetParamsFormat = "JSON"
	RadarQualityIqiGetParamsFormatCsv  RadarQualityIqiGetParamsFormat = "CSV"
)

type RadarQualityIqiGetResponseEnvelope struct {
	Result  RadarQualityIqiGetResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarQualityIqiGetResponseEnvelopeJSON `json:"-"`
}

// radarQualityIqiGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarQualityIqiGetResponseEnvelope]
type radarQualityIqiGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
