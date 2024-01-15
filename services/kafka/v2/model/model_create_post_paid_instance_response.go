package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePostPaidInstanceResponse struct {
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostPaidInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceResponse", string(data)}, " ")
}
