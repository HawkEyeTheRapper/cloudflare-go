// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// FirewallAccessRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallAccessRuleService] method
// instead.
type FirewallAccessRuleService struct {
	Options []option.RequestOption
}

// NewFirewallAccessRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallAccessRuleService(opts ...option.RequestOption) (r *FirewallAccessRuleService) {
	r = &FirewallAccessRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for all zones owned by the current user.
//
// Note: To create an IP Access rule that applies to a specific zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *FirewallAccessRuleService) New(ctx context.Context, body FirewallAccessRuleNewParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAccessRuleNewResponseEnvelope
	path := "user/firewall/access_rules/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches IP Access rules of the user. You can filter the results using several
// optional parameters.
func (r *FirewallAccessRuleService) List(ctx context.Context, query FirewallAccessRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[FirewallRule], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/firewall/access_rules/rules"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Fetches IP Access rules of the user. You can filter the results using several
// optional parameters.
func (r *FirewallAccessRuleService) ListAutoPaging(ctx context.Context, query FirewallAccessRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[FirewallRule] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Deletes an IP Access rule at the user level.
//
// Note: Deleting a user-level rule will affect all zones owned by the user.
func (r *FirewallAccessRuleService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *FirewallAccessRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAccessRuleDeleteResponseEnvelope
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an IP Access rule defined at the user level. You can only update the
// rule action (`mode` parameter) and notes.
func (r *FirewallAccessRuleService) Edit(ctx context.Context, identifier string, body FirewallAccessRuleEditParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAccessRuleEditResponseEnvelope
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallRule struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []FirewallRuleAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration FirewallRuleConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode FirewallRuleMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string           `json:"notes"`
	JSON  firewallRuleJSON `json:"-"`
}

// firewallRuleJSON contains the JSON metadata for the struct [FirewallRule]
type firewallRuleJSON struct {
	ID            apijson.Field
	AllowedModes  apijson.Field
	Configuration apijson.Field
	Mode          apijson.Field
	CreatedOn     apijson.Field
	ModifiedOn    apijson.Field
	Notes         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type FirewallRuleAllowedMode string

const (
	FirewallRuleAllowedModeBlock            FirewallRuleAllowedMode = "block"
	FirewallRuleAllowedModeChallenge        FirewallRuleAllowedMode = "challenge"
	FirewallRuleAllowedModeWhitelist        FirewallRuleAllowedMode = "whitelist"
	FirewallRuleAllowedModeJsChallenge      FirewallRuleAllowedMode = "js_challenge"
	FirewallRuleAllowedModeManagedChallenge FirewallRuleAllowedMode = "managed_challenge"
)

func (r FirewallRuleAllowedMode) IsKnown() bool {
	switch r {
	case FirewallRuleAllowedModeBlock, FirewallRuleAllowedModeChallenge, FirewallRuleAllowedModeWhitelist, FirewallRuleAllowedModeJsChallenge, FirewallRuleAllowedModeManagedChallenge:
		return true
	}
	return false
}

// The rule configuration.
//
// Union satisfied by [user.FirewallRuleConfigurationLegacyJhsIPConfiguration],
// [user.FirewallRuleConfigurationLegacyJhsIPV6Configuration],
// [user.FirewallRuleConfigurationLegacyJhsCIDRConfiguration],
// [user.FirewallRuleConfigurationLegacyJhsASNConfiguration] or
// [user.FirewallRuleConfigurationLegacyJhsCountryConfiguration].
type FirewallRuleConfiguration interface {
	implementsUserFirewallRuleConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallRuleConfiguration)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleConfigurationLegacyJhsIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleConfigurationLegacyJhsIPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleConfigurationLegacyJhsCIDRConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleConfigurationLegacyJhsASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleConfigurationLegacyJhsCountryConfiguration{}),
		},
	)
}

type FirewallRuleConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallRuleConfigurationLegacyJhsIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                `json:"value"`
	JSON  firewallRuleConfigurationLegacyJhsIPConfigurationJSON `json:"-"`
}

// firewallRuleConfigurationLegacyJhsIPConfigurationJSON contains the JSON metadata
// for the struct [FirewallRuleConfigurationLegacyJhsIPConfiguration]
type firewallRuleConfigurationLegacyJhsIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleConfigurationLegacyJhsIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConfigurationLegacyJhsIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleConfigurationLegacyJhsIPConfiguration) implementsUserFirewallRuleConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallRuleConfigurationLegacyJhsIPConfigurationTarget string

const (
	FirewallRuleConfigurationLegacyJhsIPConfigurationTargetIP FirewallRuleConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

func (r FirewallRuleConfigurationLegacyJhsIPConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationLegacyJhsIPConfigurationTargetIP:
		return true
	}
	return false
}

type FirewallRuleConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target FirewallRuleConfigurationLegacyJhsIPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                  `json:"value"`
	JSON  firewallRuleConfigurationLegacyJhsIPV6ConfigurationJSON `json:"-"`
}

// firewallRuleConfigurationLegacyJhsIPV6ConfigurationJSON contains the JSON
// metadata for the struct [FirewallRuleConfigurationLegacyJhsIPV6Configuration]
type firewallRuleConfigurationLegacyJhsIPV6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleConfigurationLegacyJhsIPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConfigurationLegacyJhsIPV6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallRuleConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallRuleConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	FirewallRuleConfigurationLegacyJhsIPV6ConfigurationTargetIp6 FirewallRuleConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

func (r FirewallRuleConfigurationLegacyJhsIPV6ConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationLegacyJhsIPV6ConfigurationTargetIp6:
		return true
	}
	return false
}

type FirewallRuleConfigurationLegacyJhsCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target FirewallRuleConfigurationLegacyJhsCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                  `json:"value"`
	JSON  firewallRuleConfigurationLegacyJhsCIDRConfigurationJSON `json:"-"`
}

// firewallRuleConfigurationLegacyJhsCIDRConfigurationJSON contains the JSON
// metadata for the struct [FirewallRuleConfigurationLegacyJhsCIDRConfiguration]
type firewallRuleConfigurationLegacyJhsCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleConfigurationLegacyJhsCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConfigurationLegacyJhsCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleConfigurationLegacyJhsCIDRConfiguration) implementsUserFirewallRuleConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallRuleConfigurationLegacyJhsCIDRConfigurationTarget string

const (
	FirewallRuleConfigurationLegacyJhsCIDRConfigurationTargetIPRange FirewallRuleConfigurationLegacyJhsCIDRConfigurationTarget = "ip_range"
)

func (r FirewallRuleConfigurationLegacyJhsCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationLegacyJhsCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type FirewallRuleConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target FirewallRuleConfigurationLegacyJhsASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                 `json:"value"`
	JSON  firewallRuleConfigurationLegacyJhsASNConfigurationJSON `json:"-"`
}

// firewallRuleConfigurationLegacyJhsASNConfigurationJSON contains the JSON
// metadata for the struct [FirewallRuleConfigurationLegacyJhsASNConfiguration]
type firewallRuleConfigurationLegacyJhsASNConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleConfigurationLegacyJhsASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConfigurationLegacyJhsASNConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleConfigurationLegacyJhsASNConfiguration) implementsUserFirewallRuleConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallRuleConfigurationLegacyJhsASNConfigurationTarget string

const (
	FirewallRuleConfigurationLegacyJhsASNConfigurationTargetASN FirewallRuleConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

func (r FirewallRuleConfigurationLegacyJhsASNConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationLegacyJhsASNConfigurationTargetASN:
		return true
	}
	return false
}

type FirewallRuleConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target FirewallRuleConfigurationLegacyJhsCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                     `json:"value"`
	JSON  firewallRuleConfigurationLegacyJhsCountryConfigurationJSON `json:"-"`
}

// firewallRuleConfigurationLegacyJhsCountryConfigurationJSON contains the JSON
// metadata for the struct [FirewallRuleConfigurationLegacyJhsCountryConfiguration]
type firewallRuleConfigurationLegacyJhsCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleConfigurationLegacyJhsCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConfigurationLegacyJhsCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallRuleConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallRuleConfigurationLegacyJhsCountryConfigurationTarget string

const (
	FirewallRuleConfigurationLegacyJhsCountryConfigurationTargetCountry FirewallRuleConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

func (r FirewallRuleConfigurationLegacyJhsCountryConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationLegacyJhsCountryConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type FirewallRuleMode string

const (
	FirewallRuleModeBlock            FirewallRuleMode = "block"
	FirewallRuleModeChallenge        FirewallRuleMode = "challenge"
	FirewallRuleModeWhitelist        FirewallRuleMode = "whitelist"
	FirewallRuleModeJsChallenge      FirewallRuleMode = "js_challenge"
	FirewallRuleModeManagedChallenge FirewallRuleMode = "managed_challenge"
)

func (r FirewallRuleMode) IsKnown() bool {
	switch r {
	case FirewallRuleModeBlock, FirewallRuleModeChallenge, FirewallRuleModeWhitelist, FirewallRuleModeJsChallenge, FirewallRuleModeManagedChallenge:
		return true
	}
	return false
}

type FirewallAccessRuleDeleteResponse struct {
	// The unique identifier of the IP Access rule.
	ID   string                               `json:"id"`
	JSON firewallAccessRuleDeleteResponseJSON `json:"-"`
}

// firewallAccessRuleDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallAccessRuleDeleteResponse]
type firewallAccessRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[FirewallAccessRuleNewParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallAccessRuleNewParamsMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r FirewallAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
//
// Satisfied by
// [user.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration],
// [user.FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration],
// [user.FirewallAccessRuleNewParamsConfigurationLegacyJhsCIDRConfiguration],
// [user.FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration],
// [user.FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration].
type FirewallAccessRuleNewParamsConfiguration interface {
	implementsUserFirewallAccessRuleNewParamsConfiguration()
}

type FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget = "ip"
)

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationLegacyJhsIPConfigurationTargetIP:
		return true
	}
	return false
}

type FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6Configuration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6 FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget = "ip6"
)

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationLegacyJhsIPV6ConfigurationTargetIp6:
		return true
	}
	return false
}

type FirewallAccessRuleNewParamsConfigurationLegacyJhsCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCIDRConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCIDRConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallAccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTargetIPRange FirewallAccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTarget = "ip_range"
)

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationLegacyJhsCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTargetASN FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget = "asn"
)

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationLegacyJhsASNConfigurationTargetASN:
		return true
	}
	return false
}

type FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfiguration) implementsUserFirewallAccessRuleNewParamsConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget string

const (
	FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTargetCountry FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget = "country"
)

func (r FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsConfigurationLegacyJhsCountryConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type FirewallAccessRuleNewParamsMode string

const (
	FirewallAccessRuleNewParamsModeBlock            FirewallAccessRuleNewParamsMode = "block"
	FirewallAccessRuleNewParamsModeChallenge        FirewallAccessRuleNewParamsMode = "challenge"
	FirewallAccessRuleNewParamsModeWhitelist        FirewallAccessRuleNewParamsMode = "whitelist"
	FirewallAccessRuleNewParamsModeJsChallenge      FirewallAccessRuleNewParamsMode = "js_challenge"
	FirewallAccessRuleNewParamsModeManagedChallenge FirewallAccessRuleNewParamsMode = "managed_challenge"
)

func (r FirewallAccessRuleNewParamsMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewParamsModeBlock, FirewallAccessRuleNewParamsModeChallenge, FirewallAccessRuleNewParamsModeWhitelist, FirewallAccessRuleNewParamsModeJsChallenge, FirewallAccessRuleNewParamsModeManagedChallenge:
		return true
	}
	return false
}

type FirewallAccessRuleNewResponseEnvelope struct {
	Errors   []FirewallAccessRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallAccessRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallRule                                    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallAccessRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallAccessRuleNewResponseEnvelopeJSON    `json:"-"`
}

// firewallAccessRuleNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallAccessRuleNewResponseEnvelope]
type firewallAccessRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    firewallAccessRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallAccessRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallAccessRuleNewResponseEnvelopeErrors]
type firewallAccessRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    firewallAccessRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallAccessRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallAccessRuleNewResponseEnvelopeMessages]
type firewallAccessRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleNewResponseEnvelopeSuccess bool

const (
	FirewallAccessRuleNewResponseEnvelopeSuccessTrue FirewallAccessRuleNewResponseEnvelopeSuccess = true
)

func (r FirewallAccessRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallAccessRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FirewallAccessRuleListParams struct {
	// The direction used to sort returned rules.
	Direction     param.Field[FirewallAccessRuleListParamsDirection]     `query:"direction"`
	EgsPagination param.Field[FirewallAccessRuleListParamsEgsPagination] `query:"egs-pagination"`
	Filters       param.Field[FirewallAccessRuleListParamsFilters]       `query:"filters"`
	// The field used to sort returned rules.
	Order param.Field[FirewallAccessRuleListParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallAccessRuleListParams]'s query parameters as
// `url.Values`.
func (r FirewallAccessRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rules.
type FirewallAccessRuleListParamsDirection string

const (
	FirewallAccessRuleListParamsDirectionAsc  FirewallAccessRuleListParamsDirection = "asc"
	FirewallAccessRuleListParamsDirectionDesc FirewallAccessRuleListParamsDirection = "desc"
)

func (r FirewallAccessRuleListParamsDirection) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListParamsDirectionAsc, FirewallAccessRuleListParamsDirectionDesc:
		return true
	}
	return false
}

type FirewallAccessRuleListParamsEgsPagination struct {
	Json param.Field[FirewallAccessRuleListParamsEgsPaginationJson] `query:"json"`
}

// URLQuery serializes [FirewallAccessRuleListParamsEgsPagination]'s query
// parameters as `url.Values`.
func (r FirewallAccessRuleListParamsEgsPagination) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallAccessRuleListParamsEgsPaginationJson struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallAccessRuleListParamsEgsPaginationJson]'s query
// parameters as `url.Values`.
func (r FirewallAccessRuleListParamsEgsPaginationJson) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallAccessRuleListParamsFilters struct {
	// The target to search in existing rules.
	ConfigurationTarget param.Field[FirewallAccessRuleListParamsFiltersConfigurationTarget] `query:"configuration.target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	ConfigurationValue param.Field[string] `query:"configuration.value"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[FirewallAccessRuleListParamsFiltersMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallAccessRuleListParamsFiltersMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
}

// URLQuery serializes [FirewallAccessRuleListParamsFilters]'s query parameters as
// `url.Values`.
func (r FirewallAccessRuleListParamsFilters) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The target to search in existing rules.
type FirewallAccessRuleListParamsFiltersConfigurationTarget string

const (
	FirewallAccessRuleListParamsFiltersConfigurationTargetIP      FirewallAccessRuleListParamsFiltersConfigurationTarget = "ip"
	FirewallAccessRuleListParamsFiltersConfigurationTargetIPRange FirewallAccessRuleListParamsFiltersConfigurationTarget = "ip_range"
	FirewallAccessRuleListParamsFiltersConfigurationTargetASN     FirewallAccessRuleListParamsFiltersConfigurationTarget = "asn"
	FirewallAccessRuleListParamsFiltersConfigurationTargetCountry FirewallAccessRuleListParamsFiltersConfigurationTarget = "country"
)

func (r FirewallAccessRuleListParamsFiltersConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListParamsFiltersConfigurationTargetIP, FirewallAccessRuleListParamsFiltersConfigurationTargetIPRange, FirewallAccessRuleListParamsFiltersConfigurationTargetASN, FirewallAccessRuleListParamsFiltersConfigurationTargetCountry:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type FirewallAccessRuleListParamsFiltersMatch string

const (
	FirewallAccessRuleListParamsFiltersMatchAny FirewallAccessRuleListParamsFiltersMatch = "any"
	FirewallAccessRuleListParamsFiltersMatchAll FirewallAccessRuleListParamsFiltersMatch = "all"
)

func (r FirewallAccessRuleListParamsFiltersMatch) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListParamsFiltersMatchAny, FirewallAccessRuleListParamsFiltersMatchAll:
		return true
	}
	return false
}

// The action to apply to a matched request.
type FirewallAccessRuleListParamsFiltersMode string

const (
	FirewallAccessRuleListParamsFiltersModeBlock            FirewallAccessRuleListParamsFiltersMode = "block"
	FirewallAccessRuleListParamsFiltersModeChallenge        FirewallAccessRuleListParamsFiltersMode = "challenge"
	FirewallAccessRuleListParamsFiltersModeWhitelist        FirewallAccessRuleListParamsFiltersMode = "whitelist"
	FirewallAccessRuleListParamsFiltersModeJsChallenge      FirewallAccessRuleListParamsFiltersMode = "js_challenge"
	FirewallAccessRuleListParamsFiltersModeManagedChallenge FirewallAccessRuleListParamsFiltersMode = "managed_challenge"
)

func (r FirewallAccessRuleListParamsFiltersMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListParamsFiltersModeBlock, FirewallAccessRuleListParamsFiltersModeChallenge, FirewallAccessRuleListParamsFiltersModeWhitelist, FirewallAccessRuleListParamsFiltersModeJsChallenge, FirewallAccessRuleListParamsFiltersModeManagedChallenge:
		return true
	}
	return false
}

// The field used to sort returned rules.
type FirewallAccessRuleListParamsOrder string

const (
	FirewallAccessRuleListParamsOrderConfigurationTarget FirewallAccessRuleListParamsOrder = "configuration.target"
	FirewallAccessRuleListParamsOrderConfigurationValue  FirewallAccessRuleListParamsOrder = "configuration.value"
	FirewallAccessRuleListParamsOrderMode                FirewallAccessRuleListParamsOrder = "mode"
)

func (r FirewallAccessRuleListParamsOrder) IsKnown() bool {
	switch r {
	case FirewallAccessRuleListParamsOrderConfigurationTarget, FirewallAccessRuleListParamsOrderConfigurationValue, FirewallAccessRuleListParamsOrderMode:
		return true
	}
	return false
}

type FirewallAccessRuleDeleteResponseEnvelope struct {
	Errors   []FirewallAccessRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallAccessRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallAccessRuleDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallAccessRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallAccessRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// firewallAccessRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallAccessRuleDeleteResponseEnvelope]
type firewallAccessRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    firewallAccessRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallAccessRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallAccessRuleDeleteResponseEnvelopeErrors]
type firewallAccessRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    firewallAccessRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallAccessRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallAccessRuleDeleteResponseEnvelopeMessages]
type firewallAccessRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleDeleteResponseEnvelopeSuccess bool

const (
	FirewallAccessRuleDeleteResponseEnvelopeSuccessTrue FirewallAccessRuleDeleteResponseEnvelopeSuccess = true
)

func (r FirewallAccessRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallAccessRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FirewallAccessRuleEditParams struct {
	// The action to apply to a matched request.
	Mode param.Field[FirewallAccessRuleEditParamsMode] `json:"mode"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r FirewallAccessRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to apply to a matched request.
type FirewallAccessRuleEditParamsMode string

const (
	FirewallAccessRuleEditParamsModeBlock            FirewallAccessRuleEditParamsMode = "block"
	FirewallAccessRuleEditParamsModeChallenge        FirewallAccessRuleEditParamsMode = "challenge"
	FirewallAccessRuleEditParamsModeWhitelist        FirewallAccessRuleEditParamsMode = "whitelist"
	FirewallAccessRuleEditParamsModeJsChallenge      FirewallAccessRuleEditParamsMode = "js_challenge"
	FirewallAccessRuleEditParamsModeManagedChallenge FirewallAccessRuleEditParamsMode = "managed_challenge"
)

func (r FirewallAccessRuleEditParamsMode) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditParamsModeBlock, FirewallAccessRuleEditParamsModeChallenge, FirewallAccessRuleEditParamsModeWhitelist, FirewallAccessRuleEditParamsModeJsChallenge, FirewallAccessRuleEditParamsModeManagedChallenge:
		return true
	}
	return false
}

type FirewallAccessRuleEditResponseEnvelope struct {
	Errors   []FirewallAccessRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallAccessRuleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallRule                                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallAccessRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallAccessRuleEditResponseEnvelopeJSON    `json:"-"`
}

// firewallAccessRuleEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallAccessRuleEditResponseEnvelope]
type firewallAccessRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleEditResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallAccessRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallAccessRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallAccessRuleEditResponseEnvelopeErrors]
type firewallAccessRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallAccessRuleEditResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    firewallAccessRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallAccessRuleEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallAccessRuleEditResponseEnvelopeMessages]
type firewallAccessRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAccessRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAccessRuleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAccessRuleEditResponseEnvelopeSuccess bool

const (
	FirewallAccessRuleEditResponseEnvelopeSuccessTrue FirewallAccessRuleEditResponseEnvelopeSuccess = true
)

func (r FirewallAccessRuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallAccessRuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}