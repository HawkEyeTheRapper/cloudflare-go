// File generated from our OpenAPI spec by Stainless.

package origin_tls_client_auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// HostnameService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewHostnameService] method instead.
type HostnameService struct {
	Options      []option.RequestOption
	Certificates *HostnameCertificateService
}

// NewHostnameService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHostnameService(opts ...option.RequestOption) (r *HostnameService) {
	r = &HostnameService{}
	r.Options = opts
	r.Certificates = NewHostnameCertificateService(opts...)
	return
}

// Associate a hostname to a certificate and enable, disable or invalidate the
// association. If disabled, client certificate will not be sent to the hostname
// even if activated at the zone level. 100 maximum associations on a single
// certificate are allowed. Note: Use a null value for parameter _enabled_ to
// invalidate the association.
func (r *HostnameService) Update(ctx context.Context, params HostnameUpdateParams, opts ...option.RequestOption) (res *[]HostnameUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Hostname Status for Client Authentication
func (r *HostnameService) Get(ctx context.Context, hostname string, query HostnameGetParams, opts ...option.RequestOption) (res *HostnameGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/%s", query.ZoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HostnameUpdateResponse struct {
	// Identifier
	ID string `json:"id"`
	// Identifier
	CertID string `json:"cert_id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled bool `json:"enabled,nullable"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname string `json:"hostname"`
	// The hostname certificate's private key.
	PrivateKey string                     `json:"private_key"`
	JSON       hostnameUpdateResponseJSON `json:"-"`
}

// hostnameUpdateResponseJSON contains the JSON metadata for the struct
// [HostnameUpdateResponse]
type hostnameUpdateResponseJSON struct {
	ID          apijson.Field
	CertID      apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	Hostname    apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type HostnameGetResponse struct {
	// Identifier
	CertID string `json:"cert_id"`
	// Status of the certificate or the association.
	CertStatus HostnameGetResponseCertStatus `json:"cert_status"`
	// The time when the certificate was updated.
	CertUpdatedAt time.Time `json:"cert_updated_at" format:"date-time"`
	// The time when the certificate was uploaded.
	CertUploadedOn time.Time `json:"cert_uploaded_on" format:"date-time"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The time when the certificate was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled bool `json:"enabled,nullable"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname string `json:"hostname"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status HostnameGetResponseStatus `json:"status"`
	// The time when the certificate was updated.
	UpdatedAt time.Time               `json:"updated_at" format:"date-time"`
	JSON      hostnameGetResponseJSON `json:"-"`
}

// hostnameGetResponseJSON contains the JSON metadata for the struct
// [HostnameGetResponse]
type hostnameGetResponseJSON struct {
	CertID         apijson.Field
	CertStatus     apijson.Field
	CertUpdatedAt  apijson.Field
	CertUploadedOn apijson.Field
	Certificate    apijson.Field
	CreatedAt      apijson.Field
	Enabled        apijson.Field
	ExpiresOn      apijson.Field
	Hostname       apijson.Field
	Issuer         apijson.Field
	SerialNumber   apijson.Field
	Signature      apijson.Field
	Status         apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameGetResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate or the association.
type HostnameGetResponseCertStatus string

const (
	HostnameGetResponseCertStatusInitializing       HostnameGetResponseCertStatus = "initializing"
	HostnameGetResponseCertStatusPendingDeployment  HostnameGetResponseCertStatus = "pending_deployment"
	HostnameGetResponseCertStatusPendingDeletion    HostnameGetResponseCertStatus = "pending_deletion"
	HostnameGetResponseCertStatusActive             HostnameGetResponseCertStatus = "active"
	HostnameGetResponseCertStatusDeleted            HostnameGetResponseCertStatus = "deleted"
	HostnameGetResponseCertStatusDeploymentTimedOut HostnameGetResponseCertStatus = "deployment_timed_out"
	HostnameGetResponseCertStatusDeletionTimedOut   HostnameGetResponseCertStatus = "deletion_timed_out"
)

// Status of the certificate or the association.
type HostnameGetResponseStatus string

const (
	HostnameGetResponseStatusInitializing       HostnameGetResponseStatus = "initializing"
	HostnameGetResponseStatusPendingDeployment  HostnameGetResponseStatus = "pending_deployment"
	HostnameGetResponseStatusPendingDeletion    HostnameGetResponseStatus = "pending_deletion"
	HostnameGetResponseStatusActive             HostnameGetResponseStatus = "active"
	HostnameGetResponseStatusDeleted            HostnameGetResponseStatus = "deleted"
	HostnameGetResponseStatusDeploymentTimedOut HostnameGetResponseStatus = "deployment_timed_out"
	HostnameGetResponseStatusDeletionTimedOut   HostnameGetResponseStatus = "deletion_timed_out"
)

type HostnameUpdateParams struct {
	// Identifier
	ZoneID param.Field[string]                       `path:"zone_id,required"`
	Config param.Field[[]HostnameUpdateParamsConfig] `json:"config,required"`
}

func (r HostnameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameUpdateParamsConfig struct {
	// Certificate identifier tag.
	CertID param.Field[string] `json:"cert_id"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled param.Field[bool] `json:"enabled"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname param.Field[string] `json:"hostname"`
}

func (r HostnameUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameUpdateResponseEnvelope struct {
	Errors   []HostnameUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []HostnameUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    HostnameUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo HostnameUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       hostnameUpdateResponseEnvelopeJSON       `json:"-"`
}

// hostnameUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameUpdateResponseEnvelope]
type hostnameUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameUpdateResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    hostnameUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HostnameUpdateResponseEnvelopeErrors]
type hostnameUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameUpdateResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    hostnameUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HostnameUpdateResponseEnvelopeMessages]
type hostnameUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameUpdateResponseEnvelopeSuccess bool

const (
	HostnameUpdateResponseEnvelopeSuccessTrue HostnameUpdateResponseEnvelopeSuccess = true
)

type HostnameUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       hostnameUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// hostnameUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [HostnameUpdateResponseEnvelopeResultInfo]
type hostnameUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type HostnameGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameGetResponseEnvelope struct {
	Errors   []HostnameGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameGetResponseEnvelopeMessages `json:"messages,required"`
	Result   HostnameGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HostnameGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameGetResponseEnvelopeJSON    `json:"-"`
}

// hostnameGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameGetResponseEnvelope]
type hostnameGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    hostnameGetResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [HostnameGetResponseEnvelopeErrors]
type hostnameGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    hostnameGetResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HostnameGetResponseEnvelopeMessages]
type hostnameGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameGetResponseEnvelopeSuccess bool

const (
	HostnameGetResponseEnvelopeSuccessTrue HostnameGetResponseEnvelopeSuccess = true
)
