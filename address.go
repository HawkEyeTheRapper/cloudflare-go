// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AddressService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAddressService] method instead.
type AddressService struct {
	Options      []option.RequestOption
	Services     *AddressServiceService
	AddressMaps  *AddressAddressMapService
	LOADocuments *AddressLOADocumentService
	Prefixes     *AddressPrefixService
}

// NewAddressService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddressService(opts ...option.RequestOption) (r *AddressService) {
	r = &AddressService{}
	r.Options = opts
	r.Services = NewAddressServiceService(opts...)
	r.AddressMaps = NewAddressAddressMapService(opts...)
	r.LOADocuments = NewAddressLOADocumentService(opts...)
	r.Prefixes = NewAddressPrefixService(opts...)
	return
}
