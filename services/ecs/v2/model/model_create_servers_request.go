package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateServersRequest Request Object
type CreateServersRequest struct {
	Body *CreateServersRequestBody `json:"body,omitempty"`
}

func (o CreateServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServersRequest struct{}"
	}

	return strings.Join([]string{"CreateServersRequest", string(data)}, " ")
}
