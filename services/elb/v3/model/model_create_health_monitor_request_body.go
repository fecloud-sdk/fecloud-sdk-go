package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateHealthMonitorRequestBody This is a auto create Body Object
type CreateHealthMonitorRequestBody struct {
	Healthmonitor *CreateHealthMonitorOption `json:"healthmonitor"`
}

func (o CreateHealthMonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthMonitorRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHealthMonitorRequestBody", string(data)}, " ")
}
