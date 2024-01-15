package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateNatResponse struct {
	Gateway *PrivateNat `json:"gateway,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPrivateNatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateNatResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateNatResponse", string(data)}, " ")
}
