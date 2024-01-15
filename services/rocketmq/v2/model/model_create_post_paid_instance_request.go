package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePostPaidInstanceRequest struct {
	Body *CreatePostPaidInstanceReq `json:"body,omitempty"`
}

func (o CreatePostPaidInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceRequest", string(data)}, " ")
}
