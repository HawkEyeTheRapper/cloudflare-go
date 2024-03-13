// File generated from our OpenAPI spec by Stainless.

package subscriptions_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/subscriptions"
)

func TestSubscriptionNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Subscriptions.New(
		context.TODO(),
		"506e3185e9c882d175a2d0cb0093d9f2",
		subscriptions.SubscriptionNewParams{
			App: cloudflare.F(subscriptions.SubscriptionNewParamsApp{
				InstallID: cloudflare.F("string"),
			}),
			ComponentValues: cloudflare.F([]subscriptions.SubscriptionNewParamsComponentValue{{
				Default: cloudflare.F(5.000000),
				Name:    cloudflare.F("page_rules"),
				Price:   cloudflare.F(5.000000),
				Value:   cloudflare.F(20.000000),
			}, {
				Default: cloudflare.F(5.000000),
				Name:    cloudflare.F("page_rules"),
				Price:   cloudflare.F(5.000000),
				Value:   cloudflare.F(20.000000),
			}, {
				Default: cloudflare.F(5.000000),
				Name:    cloudflare.F("page_rules"),
				Price:   cloudflare.F(5.000000),
				Value:   cloudflare.F(20.000000),
			}}),
			Frequency: cloudflare.F(subscriptions.SubscriptionNewParamsFrequencyMonthly),
			RatePlan: cloudflare.F(subscriptions.SubscriptionNewParamsRatePlan{
				Currency:          cloudflare.F("USD"),
				ExternallyManaged: cloudflare.F(false),
				ID:                cloudflare.F[any]("free"),
				IsContract:        cloudflare.F(false),
				PublicName:        cloudflare.F("Business Plan"),
				Scope:             cloudflare.F("zone"),
				Sets:              cloudflare.F([]string{"string", "string", "string"}),
			}),
			Zone: cloudflare.F(subscriptions.SubscriptionNewParamsZone{}),
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

func TestSubscriptionUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Subscriptions.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"506e3185e9c882d175a2d0cb0093d9f2",
		subscriptions.SubscriptionUpdateParams{
			App: cloudflare.F(subscriptions.SubscriptionUpdateParamsApp{
				InstallID: cloudflare.F("string"),
			}),
			ComponentValues: cloudflare.F([]subscriptions.SubscriptionUpdateParamsComponentValue{{
				Default: cloudflare.F(5.000000),
				Name:    cloudflare.F("page_rules"),
				Price:   cloudflare.F(5.000000),
				Value:   cloudflare.F(20.000000),
			}, {
				Default: cloudflare.F(5.000000),
				Name:    cloudflare.F("page_rules"),
				Price:   cloudflare.F(5.000000),
				Value:   cloudflare.F(20.000000),
			}, {
				Default: cloudflare.F(5.000000),
				Name:    cloudflare.F("page_rules"),
				Price:   cloudflare.F(5.000000),
				Value:   cloudflare.F(20.000000),
			}}),
			Frequency: cloudflare.F(subscriptions.SubscriptionUpdateParamsFrequencyMonthly),
			RatePlan: cloudflare.F(subscriptions.SubscriptionUpdateParamsRatePlan{
				Currency:          cloudflare.F("USD"),
				ExternallyManaged: cloudflare.F(false),
				ID:                cloudflare.F[any]("free"),
				IsContract:        cloudflare.F(false),
				PublicName:        cloudflare.F("Business Plan"),
				Scope:             cloudflare.F("zone"),
				Sets:              cloudflare.F([]string{"string", "string", "string"}),
			}),
			Zone: cloudflare.F(subscriptions.SubscriptionUpdateParamsZone{}),
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

func TestSubscriptionList(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Subscriptions.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionDelete(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Subscriptions.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"506e3185e9c882d175a2d0cb0093d9f2",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionGet(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Subscriptions.Get(context.TODO(), "506e3185e9c882d175a2d0cb0093d9f2")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}