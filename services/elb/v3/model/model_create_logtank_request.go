package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateLogtankRequest Request Object
type CreateLogtankRequest struct {
	Body *CreateLogtankRequestBody `json:"body,omitempty"`
}

func (o CreateLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankRequest struct{}"
	}

	return strings.Join([]string{"CreateLogtankRequest", string(data)}, " ")
}
