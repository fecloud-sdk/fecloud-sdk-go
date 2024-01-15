package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateGrantRequest struct {
	Body *CreateGrantRequestBody `json:"body,omitempty"`
}

func (o CreateGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGrantRequest struct{}"
	}

	return strings.Join([]string{"CreateGrantRequest", string(data)}, " ")
}
