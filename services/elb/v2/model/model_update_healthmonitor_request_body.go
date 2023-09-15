package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateHealthmonitorRequestBody This is a auto create Body Object
type UpdateHealthmonitorRequestBody struct {
	Healthmonitor *UpdateHealthmonitorReq `json:"healthmonitor"`
}

func (o UpdateHealthmonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthmonitorRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHealthmonitorRequestBody", string(data)}, " ")
}
