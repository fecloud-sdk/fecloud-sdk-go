package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateHealthmonitorRequestBody struct {
	Healthmonitor *CreateHealthmonitorReq `json:"healthmonitor"`
}

func (o CreateHealthmonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthmonitorRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHealthmonitorRequestBody", string(data)}, " ")
}
