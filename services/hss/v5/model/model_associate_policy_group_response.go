package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AssociatePolicyGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociatePolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"AssociatePolicyGroupResponse", string(data)}, " ")
}
