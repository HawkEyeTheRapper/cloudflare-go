// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarEmailSecurityMaliciousListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Radar.Emails.Security.Malicious.List(context.TODO(), cloudflare.RadarEmailSecurityMaliciousListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityMaliciousListParamsAggInterval1h),
		Arc:         cloudflare.F([]cloudflare.RadarEmailSecurityMaliciousListParamsArc{cloudflare.RadarEmailSecurityMaliciousListParamsArcPass, cloudflare.RadarEmailSecurityMaliciousListParamsArcNone, cloudflare.RadarEmailSecurityMaliciousListParamsArcFail}),
		Asn:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityMaliciousListParamsDateRange{cloudflare.RadarEmailSecurityMaliciousListParamsDateRange1d, cloudflare.RadarEmailSecurityMaliciousListParamsDateRange2d, cloudflare.RadarEmailSecurityMaliciousListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityMaliciousListParamsDKIM{cloudflare.RadarEmailSecurityMaliciousListParamsDKIMPass, cloudflare.RadarEmailSecurityMaliciousListParamsDKIMNone, cloudflare.RadarEmailSecurityMaliciousListParamsDKIMFail}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailSecurityMaliciousListParamsDmarc{cloudflare.RadarEmailSecurityMaliciousListParamsDmarcPass, cloudflare.RadarEmailSecurityMaliciousListParamsDmarcNone, cloudflare.RadarEmailSecurityMaliciousListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityMaliciousListParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityMaliciousListParamsSPF{cloudflare.RadarEmailSecurityMaliciousListParamsSPFPass, cloudflare.RadarEmailSecurityMaliciousListParamsSPFNone, cloudflare.RadarEmailSecurityMaliciousListParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}