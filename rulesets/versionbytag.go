// File generated from our OpenAPI spec by Stainless.

package rulesets

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// VersionByTagService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewVersionByTagService] method
// instead.
type VersionByTagService struct {
	Options []option.RequestOption
}

// NewVersionByTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVersionByTagService(opts ...option.RequestOption) (r *VersionByTagService) {
	r = &VersionByTagService{}
	r.Options = opts
	return
}

// Fetches the rules of a managed account ruleset version for a given tag.
func (r *VersionByTagService) Get(ctx context.Context, rulesetID string, rulesetVersion string, ruleTag string, query VersionByTagGetParams, opts ...option.RequestOption) (res *VersionByTagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VersionByTagGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s/by_tag/%s", query.AccountID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A result.
type VersionByTagGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind VersionByTagGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase VersionByTagGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []VersionByTagGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                      `json:"description"`
	JSON        versionByTagGetResponseJSON `json:"-"`
}

// versionByTagGetResponseJSON contains the JSON metadata for the struct
// [VersionByTagGetResponse]
type versionByTagGetResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type VersionByTagGetResponseKind string

const (
	VersionByTagGetResponseKindManaged VersionByTagGetResponseKind = "managed"
	VersionByTagGetResponseKindCustom  VersionByTagGetResponseKind = "custom"
	VersionByTagGetResponseKindRoot    VersionByTagGetResponseKind = "root"
	VersionByTagGetResponseKindZone    VersionByTagGetResponseKind = "zone"
)

// The phase of the ruleset.
type VersionByTagGetResponsePhase string

const (
	VersionByTagGetResponsePhaseDDOSL4                         VersionByTagGetResponsePhase = "ddos_l4"
	VersionByTagGetResponsePhaseDDOSL7                         VersionByTagGetResponsePhase = "ddos_l7"
	VersionByTagGetResponsePhaseHTTPConfigSettings             VersionByTagGetResponsePhase = "http_config_settings"
	VersionByTagGetResponsePhaseHTTPCustomErrors               VersionByTagGetResponsePhase = "http_custom_errors"
	VersionByTagGetResponsePhaseHTTPLogCustomFields            VersionByTagGetResponsePhase = "http_log_custom_fields"
	VersionByTagGetResponsePhaseHTTPRatelimit                  VersionByTagGetResponsePhase = "http_ratelimit"
	VersionByTagGetResponsePhaseHTTPRequestCacheSettings       VersionByTagGetResponsePhase = "http_request_cache_settings"
	VersionByTagGetResponsePhaseHTTPRequestDynamicRedirect     VersionByTagGetResponsePhase = "http_request_dynamic_redirect"
	VersionByTagGetResponsePhaseHTTPRequestFirewallCustom      VersionByTagGetResponsePhase = "http_request_firewall_custom"
	VersionByTagGetResponsePhaseHTTPRequestFirewallManaged     VersionByTagGetResponsePhase = "http_request_firewall_managed"
	VersionByTagGetResponsePhaseHTTPRequestLateTransform       VersionByTagGetResponsePhase = "http_request_late_transform"
	VersionByTagGetResponsePhaseHTTPRequestOrigin              VersionByTagGetResponsePhase = "http_request_origin"
	VersionByTagGetResponsePhaseHTTPRequestRedirect            VersionByTagGetResponsePhase = "http_request_redirect"
	VersionByTagGetResponsePhaseHTTPRequestSanitize            VersionByTagGetResponsePhase = "http_request_sanitize"
	VersionByTagGetResponsePhaseHTTPRequestSbfm                VersionByTagGetResponsePhase = "http_request_sbfm"
	VersionByTagGetResponsePhaseHTTPRequestSelectConfiguration VersionByTagGetResponsePhase = "http_request_select_configuration"
	VersionByTagGetResponsePhaseHTTPRequestTransform           VersionByTagGetResponsePhase = "http_request_transform"
	VersionByTagGetResponsePhaseHTTPResponseCompression        VersionByTagGetResponsePhase = "http_response_compression"
	VersionByTagGetResponsePhaseHTTPResponseFirewallManaged    VersionByTagGetResponsePhase = "http_response_firewall_managed"
	VersionByTagGetResponsePhaseHTTPResponseHeadersTransform   VersionByTagGetResponsePhase = "http_response_headers_transform"
	VersionByTagGetResponsePhaseMagicTransit                   VersionByTagGetResponsePhase = "magic_transit"
	VersionByTagGetResponsePhaseMagicTransitIDsManaged         VersionByTagGetResponsePhase = "magic_transit_ids_managed"
	VersionByTagGetResponsePhaseMagicTransitManaged            VersionByTagGetResponsePhase = "magic_transit_managed"
)

// Union satisfied by [rulesets.VersionByTagGetResponseRulesRulesetsBlockRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsExecuteRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsLogRule] or
// [rulesets.VersionByTagGetResponseRulesRulesetsSkipRule].
type VersionByTagGetResponseRule interface {
	implementsRulesetsVersionByTagGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionByTagGetResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type VersionByTagGetResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging VersionByTagGetResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON versionByTagGetResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseRulesRulesetsBlockRule]
type versionByTagGetResponseRulesRulesetsBlockRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsBlockRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsBlockRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsBlockRuleActionBlock VersionByTagGetResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response VersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     versionByTagGetResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// versionByTagGetResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsBlockRuleActionParameters]
type versionByTagGetResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type VersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                     `json:"status_code,required"`
	JSON       versionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponse]
type versionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type VersionByTagGetResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                     `json:"enabled,required"`
	JSON    versionByTagGetResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct [VersionByTagGetResponseRulesRulesetsBlockRuleLogging]
type versionByTagGetResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type VersionByTagGetResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging VersionByTagGetResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON versionByTagGetResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata
// for the struct [VersionByTagGetResponseRulesRulesetsExecuteRule]
type versionByTagGetResponseRulesRulesetsExecuteRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsExecuteRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsExecuteRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionExecute VersionByTagGetResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsExecuteRuleActionParameters]
type versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                         `json:"public_key,required"`
	JSON      versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    VersionByTagGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type VersionByTagGetResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                       `json:"enabled,required"`
	JSON    versionByTagGetResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct [VersionByTagGetResponseRulesRulesetsExecuteRuleLogging]
type versionByTagGetResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type VersionByTagGetResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsLogRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging VersionByTagGetResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON versionByTagGetResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsLogRuleJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseRulesRulesetsLogRule]
type versionByTagGetResponseRulesRulesetsLogRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsLogRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsLogRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsLogRuleActionLog VersionByTagGetResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type VersionByTagGetResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                   `json:"enabled,required"`
	JSON    versionByTagGetResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsLogRuleLoggingJSON contains the JSON
// metadata for the struct [VersionByTagGetResponseRulesRulesetsLogRuleLogging]
type versionByTagGetResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type VersionByTagGetResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging VersionByTagGetResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON versionByTagGetResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseRulesRulesetsSkipRule]
type versionByTagGetResponseRulesRulesetsSkipRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsSkipRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsSkipRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsSkipRuleActionSkip VersionByTagGetResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                         `json:"rulesets"`
	JSON     versionByTagGetResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsSkipRuleActionParameters]
type versionByTagGetResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductBic           VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductHot           VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock       VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductWAF           VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent VersionByTagGetResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type VersionByTagGetResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                    `json:"enabled,required"`
	JSON    versionByTagGetResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct [VersionByTagGetResponseRulesRulesetsSkipRuleLogging]
type versionByTagGetResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type VersionByTagGetParams struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

// A response object.
type VersionByTagGetResponseEnvelope struct {
	// A list of error messages.
	Errors []VersionByTagGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []VersionByTagGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result VersionByTagGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success VersionByTagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    versionByTagGetResponseEnvelopeJSON    `json:"-"`
}

// versionByTagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionByTagGetResponseEnvelope]
type versionByTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionByTagGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionByTagGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   versionByTagGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// versionByTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VersionByTagGetResponseEnvelopeErrors]
type versionByTagGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionByTagGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    versionByTagGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// versionByTagGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseEnvelopeErrorsSource]
type versionByTagGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionByTagGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionByTagGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   versionByTagGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// versionByTagGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VersionByTagGetResponseEnvelopeMessages]
type versionByTagGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionByTagGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    versionByTagGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// versionByTagGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseEnvelopeMessagesSource]
type versionByTagGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VersionByTagGetResponseEnvelopeSuccess bool

const (
	VersionByTagGetResponseEnvelopeSuccessTrue VersionByTagGetResponseEnvelopeSuccess = true
)
