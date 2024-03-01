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

func TestZeroTrustAccessIdentityProviderNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.ZeroTrust.Access.IdentityProviders.New(context.TODO(), cloudflare.ZeroTrustAccessIdentityProviderNewParams{
		Config: cloudflare.F(cloudflare.ZeroTrustAccessIdentityProviderNewParamsConfig{
			ClientID:                 cloudflare.F("<your client id>"),
			ClientSecret:             cloudflare.F("<your client secret>"),
			Claims:                   cloudflare.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
			EmailClaimName:           cloudflare.F("custom_claim_name"),
			ConditionalAccessEnabled: cloudflare.F(true),
			DirectoryID:              cloudflare.F("<your azure directory uuid>"),
			SupportGroups:            cloudflare.F(true),
			CentrifyAccount:          cloudflare.F("https://abc123.my.centrify.com/"),
			CentrifyAppID:            cloudflare.F("exampleapp"),
			AppsDomain:               cloudflare.F("mycompany.com"),
			AuthURL:                  cloudflare.F("https://accounts.google.com/o/oauth2/auth"),
			CertsURL:                 cloudflare.F("https://www.googleapis.com/oauth2/v3/certs"),
			Scopes:                   cloudflare.F([]string{"openid", "email", "profile"}),
			TokenURL:                 cloudflare.F("https://accounts.google.com/o/oauth2/token"),
			AuthorizationServerID:    cloudflare.F("aus9o8wzkhckw9TLa0h7z"),
			OktaAccount:              cloudflare.F("https://dev-abc123.oktapreview.com"),
			OneloginAccount:          cloudflare.F("https://mycompany.onelogin.com"),
			PingEnvID:                cloudflare.F("342b5660-0c32-4936-a5a4-ce21fae57b0a"),
			Attributes:               cloudflare.F([]string{"group", "department_code", "divison"}),
			EmailAttributeName:       cloudflare.F("Email"),
			HeaderAttributes: cloudflare.F([]cloudflare.ZeroTrustAccessIdentityProviderNewParamsConfigHeaderAttribute{{
				AttributeName: cloudflare.F("string"),
				HeaderName:    cloudflare.F("string"),
			}, {
				AttributeName: cloudflare.F("string"),
				HeaderName:    cloudflare.F("string"),
			}, {
				AttributeName: cloudflare.F("string"),
				HeaderName:    cloudflare.F("string"),
			}}),
			IdpPublicCerts: cloudflare.F([]string{"string", "string", "string"}),
			IssuerURL:      cloudflare.F("https://whoami.com"),
			SignRequest:    cloudflare.F(true),
			SSOTargetURL:   cloudflare.F("https://edgeaccess.org/idp/saml/login"),
		}),
		Name:      cloudflare.F("Widget Corps IDP"),
		Type:      cloudflare.F(cloudflare.ZeroTrustAccessIdentityProviderNewParamsTypeOnetimepin),
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
		ScimConfig: cloudflare.F(cloudflare.ZeroTrustAccessIdentityProviderNewParamsScimConfig{
			Enabled:                cloudflare.F(true),
			GroupMemberDeprovision: cloudflare.F(true),
			SeatDeprovision:        cloudflare.F(true),
			Secret:                 cloudflare.F("string"),
			UserDeprovision:        cloudflare.F(true),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustAccessIdentityProviderUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.ZeroTrust.Access.IdentityProviders.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustAccessIdentityProviderUpdateParams{
			Config: cloudflare.F(cloudflare.ZeroTrustAccessIdentityProviderUpdateParamsConfig{
				ClientID:                 cloudflare.F("<your client id>"),
				ClientSecret:             cloudflare.F("<your client secret>"),
				Claims:                   cloudflare.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
				EmailClaimName:           cloudflare.F("custom_claim_name"),
				ConditionalAccessEnabled: cloudflare.F(true),
				DirectoryID:              cloudflare.F("<your azure directory uuid>"),
				SupportGroups:            cloudflare.F(true),
				CentrifyAccount:          cloudflare.F("https://abc123.my.centrify.com/"),
				CentrifyAppID:            cloudflare.F("exampleapp"),
				AppsDomain:               cloudflare.F("mycompany.com"),
				AuthURL:                  cloudflare.F("https://accounts.google.com/o/oauth2/auth"),
				CertsURL:                 cloudflare.F("https://www.googleapis.com/oauth2/v3/certs"),
				Scopes:                   cloudflare.F([]string{"openid", "email", "profile"}),
				TokenURL:                 cloudflare.F("https://accounts.google.com/o/oauth2/token"),
				AuthorizationServerID:    cloudflare.F("aus9o8wzkhckw9TLa0h7z"),
				OktaAccount:              cloudflare.F("https://dev-abc123.oktapreview.com"),
				OneloginAccount:          cloudflare.F("https://mycompany.onelogin.com"),
				PingEnvID:                cloudflare.F("342b5660-0c32-4936-a5a4-ce21fae57b0a"),
				Attributes:               cloudflare.F([]string{"group", "department_code", "divison"}),
				EmailAttributeName:       cloudflare.F("Email"),
				HeaderAttributes: cloudflare.F([]cloudflare.ZeroTrustAccessIdentityProviderUpdateParamsConfigHeaderAttribute{{
					AttributeName: cloudflare.F("string"),
					HeaderName:    cloudflare.F("string"),
				}, {
					AttributeName: cloudflare.F("string"),
					HeaderName:    cloudflare.F("string"),
				}, {
					AttributeName: cloudflare.F("string"),
					HeaderName:    cloudflare.F("string"),
				}}),
				IdpPublicCerts: cloudflare.F([]string{"string", "string", "string"}),
				IssuerURL:      cloudflare.F("https://whoami.com"),
				SignRequest:    cloudflare.F(true),
				SSOTargetURL:   cloudflare.F("https://edgeaccess.org/idp/saml/login"),
			}),
			Name:      cloudflare.F("Widget Corps IDP"),
			Type:      cloudflare.F(cloudflare.ZeroTrustAccessIdentityProviderUpdateParamsTypeOnetimepin),
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
			ScimConfig: cloudflare.F(cloudflare.ZeroTrustAccessIdentityProviderUpdateParamsScimConfig{
				Enabled:                cloudflare.F(true),
				GroupMemberDeprovision: cloudflare.F(true),
				SeatDeprovision:        cloudflare.F(true),
				Secret:                 cloudflare.F("string"),
				UserDeprovision:        cloudflare.F(true),
			}),
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

func TestZeroTrustAccessIdentityProviderListWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.ZeroTrust.Access.IdentityProviders.List(context.TODO(), cloudflare.ZeroTrustAccessIdentityProviderListParams{
		AccountID: cloudflare.F("string"),
		ZoneID:    cloudflare.F("string"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustAccessIdentityProviderDeleteWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.ZeroTrust.Access.IdentityProviders.Delete(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustAccessIdentityProviderDeleteParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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

func TestZeroTrustAccessIdentityProviderGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.ZeroTrust.Access.IdentityProviders.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustAccessIdentityProviderGetParams{
			AccountID: cloudflare.F("string"),
			ZoneID:    cloudflare.F("string"),
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