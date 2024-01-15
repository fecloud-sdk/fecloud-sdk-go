package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateServerGroupRequest struct {
	Body *CreateServerGroupRequestBody `json:"body,omitempty"`
}

func (o CreateServerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateServerGroupRequest", string(data)}, " ")
}
