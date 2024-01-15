package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetRecyclePolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetRecyclePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecyclePolicyResponse struct{}"
	}

	return strings.Join([]string{"SetRecyclePolicyResponse", string(data)}, " ")
}
