package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// SwitchConfigurationResponse Response Object
type SwitchConfigurationResponse struct {

	// 应用参数模板的异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchConfigurationResponse struct{}"
	}

	return strings.Join([]string{"SwitchConfigurationResponse", string(data)}, " ")
}
