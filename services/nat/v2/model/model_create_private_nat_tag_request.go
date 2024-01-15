package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateNatTagRequest struct {
	ResourceId string `json:"resource_id"`

	Body *CreateResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreatePrivateNatTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatTagRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatTagRequest", string(data)}, " ")
}
