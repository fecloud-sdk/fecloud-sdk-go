package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateipRequest struct {
	Body *CreatePrivateipRequestBody `json:"body,omitempty"`
}

func (o CreatePrivateipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateipRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateipRequest", string(data)}, " ")
}
