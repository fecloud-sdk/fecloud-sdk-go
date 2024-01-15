package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePoolRequestBody struct {
	Pool *UpdatePoolReq `json:"pool"`
}

func (o UpdatePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePoolRequestBody", string(data)}, " ")
}
