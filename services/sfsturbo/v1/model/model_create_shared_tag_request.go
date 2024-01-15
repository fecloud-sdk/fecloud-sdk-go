package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSharedTagRequest struct {
	ContentType string `json:"Content-Type"`

	ShareId string `json:"share_id"`

	Body *CreateSharedTagRequestBody `json:"body,omitempty"`
}

func (o CreateSharedTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedTagRequest struct{}"
	}

	return strings.Join([]string{"CreateSharedTagRequest", string(data)}, " ")
}
