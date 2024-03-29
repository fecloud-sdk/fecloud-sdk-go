package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeleteCertificateResponse", string(data)}, " ")
}
