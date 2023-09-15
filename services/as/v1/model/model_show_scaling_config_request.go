package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowScalingConfigRequest Request Object
type ShowScalingConfigRequest struct {

	// 伸缩配置ID，查询唯一配置。
	ScalingConfigurationId string `json:"scaling_configuration_id"`
}

func (o ShowScalingConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowScalingConfigRequest", string(data)}, " ")
}
