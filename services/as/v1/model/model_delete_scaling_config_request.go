package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteScalingConfigRequest Request Object
type DeleteScalingConfigRequest struct {

	// 伸缩配置ID。
	ScalingConfigurationId string `json:"scaling_configuration_id"`
}

func (o DeleteScalingConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingConfigRequest", string(data)}, " ")
}
