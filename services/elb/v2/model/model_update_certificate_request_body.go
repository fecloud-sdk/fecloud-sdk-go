package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateCertificateRequestBody struct {
	Certificate *string `json:"certificate,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	Description *string `json:"description,omitempty"`

	Domain *string `json:"domain,omitempty"`

	Name *string `json:"name,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

func (o UpdateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequestBody", string(data)}, " ")
}
