// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// IamService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIamService] method instead.
type IamService struct {
	Options          []option.RequestOption
	PermissionGroups *IamPermissionGroupService
	ResourceGroups   *IamResourceGroupService
}

// NewIamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIamService(opts ...option.RequestOption) (r *IamService) {
	r = &IamService{}
	r.Options = opts
	r.PermissionGroups = NewIamPermissionGroupService(opts...)
	r.ResourceGroups = NewIamResourceGroupService(opts...)
	return
}