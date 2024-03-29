package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ModifyParamRequest struct {
	Value string `json:"value"`
}

func (o ModifyParamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyParamRequest struct{}"
	}

	return strings.Join([]string{"ModifyParamRequest", string(data)}, " ")
}
