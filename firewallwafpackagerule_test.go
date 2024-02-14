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

func TestFirewallWAFPackageRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewalls.WAF.Packages.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		cloudflare.FirewallWAFPackageRuleUpdateParams{
			Mode: cloudflare.F(cloudflare.FirewallWAFPackageRuleUpdateParamsModeOn),
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

func TestFirewallWAFPackageRuleGet(t *testing.T) {
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
	_, err := client.Firewalls.WAF.Packages.Rules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"a25a9a7e9c00afc1fb2e0245519d725b",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFirewallWAFPackageRuleWAFRulesListWAFRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewalls.WAF.Packages.Rules.WAFRulesListWAFRules(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		cloudflare.FirewallWAFPackageRuleWAFRulesListWAFRulesParams{
			Direction: cloudflare.F(cloudflare.FirewallWAFPackageRuleWAFRulesListWAFRulesParamsDirectionDesc),
			Match:     cloudflare.F(cloudflare.FirewallWAFPackageRuleWAFRulesListWAFRulesParamsMatchAny),
			Mode:      cloudflare.F(cloudflare.FirewallWAFPackageRuleWAFRulesListWAFRulesParamsModeChl),
			Order:     cloudflare.F(cloudflare.FirewallWAFPackageRuleWAFRulesListWAFRulesParamsOrderPriority),
			Page:      cloudflare.F(1.000000),
			PerPage:   cloudflare.F(5.000000),
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