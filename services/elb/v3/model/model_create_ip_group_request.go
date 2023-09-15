package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateIpGroupRequest Request Object
type CreateIpGroupRequest struct {
	Body *CreateIpGroupRequestBody `json:"body,omitempty"`
}

func (o CreateIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateIpGroupRequest", string(data)}, " ")
}
