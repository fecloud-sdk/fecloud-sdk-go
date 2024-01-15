package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryDbParamsResp struct {
	Params *[]Params `json:"params,omitempty"`
}

func (o QueryDbParamsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDbParamsResp struct{}"
	}

	return strings.Join([]string{"QueryDbParamsResp", string(data)}, " ")
}
