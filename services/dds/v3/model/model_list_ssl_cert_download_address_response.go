package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSslCertDownloadAddressResponse struct {
	Certs          *[]CertInfo `json:"certs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListSslCertDownloadAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSslCertDownloadAddressResponse struct{}"
	}

	return strings.Join([]string{"ListSslCertDownloadAddressResponse", string(data)}, " ")
}
