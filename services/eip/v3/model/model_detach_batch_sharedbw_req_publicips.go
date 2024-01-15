package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DetachBatchSharedbwReqPublicips struct {

	// - 功能说明：EIP ID
	PublicipId string `json:"publicip_id"`

	Bandwidth *DetachSharedbwDict `json:"bandwidth"`
}

func (o DetachBatchSharedbwReqPublicips) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachBatchSharedbwReqPublicips struct{}"
	}

	return strings.Join([]string{"DetachBatchSharedbwReqPublicips", string(data)}, " ")
}
