package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateIpGroupRequestBody This is a auto create Body Object
type CreateIpGroupRequestBody struct {
	Ipgroup *CreateIpGroupOption `json:"ipgroup"`
}

func (o CreateIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIpGroupRequestBody", string(data)}, " ")
}
