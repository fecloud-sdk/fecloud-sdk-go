package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowPolicyAndInstanceQuotaRequest Request Object
type ShowPolicyAndInstanceQuotaRequest struct {

	// 伸缩组ID。
	ScalingGroupId string `json:"scaling_group_id"`
}

func (o ShowPolicyAndInstanceQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyAndInstanceQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyAndInstanceQuotaRequest", string(data)}, " ")
}
