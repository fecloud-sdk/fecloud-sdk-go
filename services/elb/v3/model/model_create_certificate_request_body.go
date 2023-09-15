package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateCertificateRequestBody This is a auto create Body Object
type CreateCertificateRequestBody struct {
	Certificate *CreateCertificateOption `json:"certificate"`
}

func (o CreateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequestBody", string(data)}, " ")
}
