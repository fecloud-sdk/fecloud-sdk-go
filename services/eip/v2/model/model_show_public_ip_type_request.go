package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowPublicIpTypeRequest Request Object
type ShowPublicIpTypeRequest struct {
}

func (o ShowPublicIpTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicIpTypeRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicIpTypeRequest", string(data)}, " ")
}
