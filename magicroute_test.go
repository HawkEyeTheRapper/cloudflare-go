// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestMagicRouteUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Magics.Routes.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.MagicRouteUpdateParams{
			Nexthop:     cloudflare.F("203.0.113.1"),
			Prefix:      cloudflare.F("192.0.2.0/24"),
			Priority:    cloudflare.F(int64(0)),
			Description: cloudflare.F("New route for new prefix 203.0.113.1"),
			Scope: cloudflare.F(cloudflare.MagicRouteUpdateParamsScope{
				ColoNames:   cloudflare.F([]string{"den01", "den01", "den01"}),
				ColoRegions: cloudflare.F([]string{"APAC", "APAC", "APAC"}),
			}),
			Weight: cloudflare.F(int64(0)),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMagicRouteDelete(t *testing.T) {
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
	_, err := client.Magics.Routes.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMagicRouteGet(t *testing.T) {
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
	_, err := client.Magics.Routes.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMagicRouteMagicStaticRoutesNewRoutes(t *testing.T) {
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
	_, err := client.Magics.Routes.MagicStaticRoutesNewRoutes(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.MagicRouteMagicStaticRoutesNewRoutesParams{
			Body: cloudflare.F[any](map[string]interface{}{}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMagicRouteMagicStaticRoutesListRoutes(t *testing.T) {
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
	_, err := client.Magics.Routes.MagicStaticRoutesListRoutes(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMagicRouteMagicStaticRoutesUpdateManyRoutes(t *testing.T) {
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
	_, err := client.Magics.Routes.MagicStaticRoutesUpdateManyRoutes(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.MagicRouteMagicStaticRoutesUpdateManyRoutesParams{
			Routes: cloudflare.F([]cloudflare.MagicRouteMagicStaticRoutesUpdateManyRoutesParamsRoute{{
				Description: cloudflare.F("New route for new prefix 203.0.113.1"),
				Nexthop:     cloudflare.F("203.0.113.1"),
				Prefix:      cloudflare.F("192.0.2.0/24"),
				Priority:    cloudflare.F(int64(0)),
				Scope: cloudflare.F(cloudflare.MagicRouteMagicStaticRoutesUpdateManyRoutesParamsRoutesScope{
					ColoNames:   cloudflare.F([]string{"den01", "den01", "den01"}),
					ColoRegions: cloudflare.F([]string{"APAC", "APAC", "APAC"}),
				}),
				Weight: cloudflare.F(int64(0)),
			}, {
				Description: cloudflare.F("New route for new prefix 203.0.113.1"),
				Nexthop:     cloudflare.F("203.0.113.1"),
				Prefix:      cloudflare.F("192.0.2.0/24"),
				Priority:    cloudflare.F(int64(0)),
				Scope: cloudflare.F(cloudflare.MagicRouteMagicStaticRoutesUpdateManyRoutesParamsRoutesScope{
					ColoNames:   cloudflare.F([]string{"den01", "den01", "den01"}),
					ColoRegions: cloudflare.F([]string{"APAC", "APAC", "APAC"}),
				}),
				Weight: cloudflare.F(int64(0)),
			}, {
				Description: cloudflare.F("New route for new prefix 203.0.113.1"),
				Nexthop:     cloudflare.F("203.0.113.1"),
				Prefix:      cloudflare.F("192.0.2.0/24"),
				Priority:    cloudflare.F(int64(0)),
				Scope: cloudflare.F(cloudflare.MagicRouteMagicStaticRoutesUpdateManyRoutesParamsRoutesScope{
					ColoNames:   cloudflare.F([]string{"den01", "den01", "den01"}),
					ColoRegions: cloudflare.F([]string{"APAC", "APAC", "APAC"}),
				}),
				Weight: cloudflare.F(int64(0)),
			}}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}