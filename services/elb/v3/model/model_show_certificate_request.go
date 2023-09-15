package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowCertificateRequest Request Object
type ShowCertificateRequest struct {

	// 证书ID。
	CertificateId string `json:"certificate_id"`
}

func (o ShowCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificateRequest", string(data)}, " ")
}