package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteCertificateRequest struct {
	CertificateId string `json:"certificate_id"`
}

func (o DeleteCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertificateRequest struct{}"
	}

	return strings.Join([]string{"DeleteCertificateRequest", string(data)}, " ")
}
