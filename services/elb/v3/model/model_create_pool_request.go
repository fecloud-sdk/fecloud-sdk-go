package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePoolRequest Request Object
type CreatePoolRequest struct {
	Body *CreatePoolRequestBody `json:"body,omitempty"`
}

func (o CreatePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolRequest struct{}"
	}

	return strings.Join([]string{"CreatePoolRequest", string(data)}, " ")
}
