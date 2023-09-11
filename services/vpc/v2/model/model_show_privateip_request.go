package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateipRequest struct {
	PrivateipId string `json:"privateip_id"`
}

func (o ShowPrivateipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateipRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateipRequest", string(data)}, " ")
}
