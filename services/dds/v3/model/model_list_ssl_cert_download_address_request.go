package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSslCertDownloadAddressRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o ListSslCertDownloadAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSslCertDownloadAddressRequest struct{}"
	}

	return strings.Join([]string{"ListSslCertDownloadAddressRequest", string(data)}, " ")
}
