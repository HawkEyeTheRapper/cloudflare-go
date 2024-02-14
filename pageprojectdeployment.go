// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// PageProjectDeploymentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPageProjectDeploymentService]
// method instead.
type PageProjectDeploymentService struct {
	Options   []option.RequestOption
	Histories *PageProjectDeploymentHistoryService
	Retries   *PageProjectDeploymentRetryService
	Rollbacks *PageProjectDeploymentRollbackService
}

// NewPageProjectDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageProjectDeploymentService(opts ...option.RequestOption) (r *PageProjectDeploymentService) {
	r = &PageProjectDeploymentService{}
	r.Options = opts
	r.Histories = NewPageProjectDeploymentHistoryService(opts...)
	r.Retries = NewPageProjectDeploymentRetryService(opts...)
	r.Rollbacks = NewPageProjectDeploymentRollbackService(opts...)
	return
}

// Start a new deployment from production. The repository and account must have
// already been authorized on the Cloudflare Pages dashboard.
func (r *PageProjectDeploymentService) New(ctx context.Context, accountID string, projectName string, body PageProjectDeploymentNewParams, opts ...option.RequestOption) (res *PageProjectDeploymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of project deployments.
func (r *PageProjectDeploymentService) List(ctx context.Context, accountID string, projectName string, opts ...option.RequestOption) (res *[]PageProjectDeploymentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a deployment.
func (r *PageProjectDeploymentService) Delete(ctx context.Context, accountID string, projectName string, deploymentID string, opts ...option.RequestOption) (res *PageProjectDeploymentDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetch information about a deployment.
func (r *PageProjectDeploymentService) Get(ctx context.Context, accountID string, projectName string, deploymentID string, opts ...option.RequestOption) (res *PageProjectDeploymentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageProjectDeploymentNewResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger PageProjectDeploymentNewResponseDeploymentTrigger `json:"deployment_trigger"`
	// A dict of env variables to build this deploy.
	EnvVars interface{} `json:"env_vars"`
	// Type of deploy.
	Environment string `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped   bool        `json:"is_skipped"`
	LatestStage interface{} `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string      `json:"short_id"`
	Source  interface{} `json:"source"`
	// List of past stages.
	Stages []PageProjectDeploymentNewResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                               `json:"url"`
	JSON pageProjectDeploymentNewResponseJSON `json:"-"`
}

// pageProjectDeploymentNewResponseJSON contains the JSON metadata for the struct
// [PageProjectDeploymentNewResponse]
type pageProjectDeploymentNewResponseJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PageProjectDeploymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type PageProjectDeploymentNewResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata PageProjectDeploymentNewResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                `json:"type"`
	JSON pageProjectDeploymentNewResponseDeploymentTriggerJSON `json:"-"`
}

// pageProjectDeploymentNewResponseDeploymentTriggerJSON contains the JSON metadata
// for the struct [PageProjectDeploymentNewResponseDeploymentTrigger]
type pageProjectDeploymentNewResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentNewResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type PageProjectDeploymentNewResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                        `json:"commit_message"`
	JSON          pageProjectDeploymentNewResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// pageProjectDeploymentNewResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct
// [PageProjectDeploymentNewResponseDeploymentTriggerMetadata]
type pageProjectDeploymentNewResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageProjectDeploymentNewResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type PageProjectDeploymentNewResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                    `json:"status"`
	JSON   pageProjectDeploymentNewResponseStageJSON `json:"-"`
}

// pageProjectDeploymentNewResponseStageJSON contains the JSON metadata for the
// struct [PageProjectDeploymentNewResponseStage]
type pageProjectDeploymentNewResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentNewResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentListResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger PageProjectDeploymentListResponseDeploymentTrigger `json:"deployment_trigger"`
	// A dict of env variables to build this deploy.
	EnvVars interface{} `json:"env_vars"`
	// Type of deploy.
	Environment string `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped   bool        `json:"is_skipped"`
	LatestStage interface{} `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string      `json:"short_id"`
	Source  interface{} `json:"source"`
	// List of past stages.
	Stages []PageProjectDeploymentListResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                                `json:"url"`
	JSON pageProjectDeploymentListResponseJSON `json:"-"`
}

// pageProjectDeploymentListResponseJSON contains the JSON metadata for the struct
// [PageProjectDeploymentListResponse]
type pageProjectDeploymentListResponseJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type PageProjectDeploymentListResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata PageProjectDeploymentListResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                 `json:"type"`
	JSON pageProjectDeploymentListResponseDeploymentTriggerJSON `json:"-"`
}

// pageProjectDeploymentListResponseDeploymentTriggerJSON contains the JSON
// metadata for the struct [PageProjectDeploymentListResponseDeploymentTrigger]
type pageProjectDeploymentListResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type PageProjectDeploymentListResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                         `json:"commit_message"`
	JSON          pageProjectDeploymentListResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// pageProjectDeploymentListResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct
// [PageProjectDeploymentListResponseDeploymentTriggerMetadata]
type pageProjectDeploymentListResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type PageProjectDeploymentListResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                     `json:"status"`
	JSON   pageProjectDeploymentListResponseStageJSON `json:"-"`
}

// pageProjectDeploymentListResponseStageJSON contains the JSON metadata for the
// struct [PageProjectDeploymentListResponseStage]
type pageProjectDeploymentListResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentDeleteResponse = interface{}

type PageProjectDeploymentGetResponse struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases     []interface{} `json:"aliases,nullable"`
	BuildConfig interface{}   `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger PageProjectDeploymentGetResponseDeploymentTrigger `json:"deployment_trigger"`
	// A dict of env variables to build this deploy.
	EnvVars interface{} `json:"env_vars"`
	// Type of deploy.
	Environment string `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped   bool        `json:"is_skipped"`
	LatestStage interface{} `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string      `json:"short_id"`
	Source  interface{} `json:"source"`
	// List of past stages.
	Stages []PageProjectDeploymentGetResponseStage `json:"stages"`
	// The live URL to view this deployment.
	URL  string                               `json:"url"`
	JSON pageProjectDeploymentGetResponseJSON `json:"-"`
}

// pageProjectDeploymentGetResponseJSON contains the JSON metadata for the struct
// [PageProjectDeploymentGetResponse]
type pageProjectDeploymentGetResponseJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PageProjectDeploymentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Info about what caused the deployment.
type PageProjectDeploymentGetResponseDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata PageProjectDeploymentGetResponseDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type string                                                `json:"type"`
	JSON pageProjectDeploymentGetResponseDeploymentTriggerJSON `json:"-"`
}

// pageProjectDeploymentGetResponseDeploymentTriggerJSON contains the JSON metadata
// for the struct [PageProjectDeploymentGetResponseDeploymentTrigger]
type pageProjectDeploymentGetResponseDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentGetResponseDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional info about the trigger.
type PageProjectDeploymentGetResponseDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                        `json:"commit_message"`
	JSON          pageProjectDeploymentGetResponseDeploymentTriggerMetadataJSON `json:"-"`
}

// pageProjectDeploymentGetResponseDeploymentTriggerMetadataJSON contains the JSON
// metadata for the struct
// [PageProjectDeploymentGetResponseDeploymentTriggerMetadata]
type pageProjectDeploymentGetResponseDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageProjectDeploymentGetResponseDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the deployment.
type PageProjectDeploymentGetResponseStage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name string `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status string                                    `json:"status"`
	JSON   pageProjectDeploymentGetResponseStageJSON `json:"-"`
}

// pageProjectDeploymentGetResponseStageJSON contains the JSON metadata for the
// struct [PageProjectDeploymentGetResponseStage]
type pageProjectDeploymentGetResponseStageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentGetResponseStage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentNewParams struct {
	// The branch to build the new deployment from. The `HEAD` of the branch will be
	// used. If omitted, the production branch will be used by default.
	Branch param.Field[string] `json:"branch"`
}

func (r PageProjectDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageProjectDeploymentNewResponseEnvelope struct {
	Errors   []PageProjectDeploymentNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDeploymentNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentNewResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDeploymentNewResponseEnvelope]
type pageProjectDeploymentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    pageProjectDeploymentNewResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [PageProjectDeploymentNewResponseEnvelopeErrors]
type pageProjectDeploymentNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    pageProjectDeploymentNewResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PageProjectDeploymentNewResponseEnvelopeMessages]
type pageProjectDeploymentNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentNewResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentNewResponseEnvelopeSuccessTrue PageProjectDeploymentNewResponseEnvelopeSuccess = true
)

type PageProjectDeploymentListResponseEnvelope struct {
	Errors   []PageProjectDeploymentListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PageProjectDeploymentListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success    PageProjectDeploymentListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PageProjectDeploymentListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pageProjectDeploymentListResponseEnvelopeJSON       `json:"-"`
}

// pageProjectDeploymentListResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDeploymentListResponseEnvelope]
type pageProjectDeploymentListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    pageProjectDeploymentListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [PageProjectDeploymentListResponseEnvelopeErrors]
type pageProjectDeploymentListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    pageProjectDeploymentListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PageProjectDeploymentListResponseEnvelopeMessages]
type pageProjectDeploymentListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentListResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentListResponseEnvelopeSuccessTrue PageProjectDeploymentListResponseEnvelopeSuccess = true
)

type PageProjectDeploymentListResponseEnvelopeResultInfo struct {
	Count      interface{}                                             `json:"count"`
	Page       interface{}                                             `json:"page"`
	PerPage    interface{}                                             `json:"per_page"`
	TotalCount interface{}                                             `json:"total_count"`
	JSON       pageProjectDeploymentListResponseEnvelopeResultInfoJSON `json:"-"`
}

// pageProjectDeploymentListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [PageProjectDeploymentListResponseEnvelopeResultInfo]
type pageProjectDeploymentListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentGetResponseEnvelope struct {
	Errors   []PageProjectDeploymentGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDeploymentGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentGetResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDeploymentGetResponseEnvelope]
type pageProjectDeploymentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    pageProjectDeploymentGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [PageProjectDeploymentGetResponseEnvelopeErrors]
type pageProjectDeploymentGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    pageProjectDeploymentGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PageProjectDeploymentGetResponseEnvelopeMessages]
type pageProjectDeploymentGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentGetResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentGetResponseEnvelopeSuccessTrue PageProjectDeploymentGetResponseEnvelopeSuccess = true
)