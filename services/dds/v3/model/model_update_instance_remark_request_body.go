package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceRemarkRequestBody struct {
	Remark string `json:"remark"`
}

func (o UpdateInstanceRemarkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRemarkRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRemarkRequestBody", string(data)}, " ")
}
