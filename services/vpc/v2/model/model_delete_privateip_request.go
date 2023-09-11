package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateipRequest struct {
	PrivateipId string `json:"privateip_id"`
}

func (o DeletePrivateipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateipRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateipRequest", string(data)}, " ")
}
