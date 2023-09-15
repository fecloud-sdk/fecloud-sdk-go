package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateScalingConfigResponse Response Object
type CreateScalingConfigResponse struct {

	// 伸缩配置ID
	ScalingConfigurationId *string `json:"scaling_configuration_id,omitempty"`
	HttpStatusCode         int     `json:"-"`
}

func (o CreateScalingConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingConfigResponse", string(data)}, " ")
}
