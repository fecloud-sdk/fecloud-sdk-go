package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeletePublicipRequest Request Object
type DeletePublicipRequest struct {

	// 弹性公网IP唯一标识
	PublicipId string `json:"publicip_id"`
}

func (o DeletePublicipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicipRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicipRequest", string(data)}, " ")
}
