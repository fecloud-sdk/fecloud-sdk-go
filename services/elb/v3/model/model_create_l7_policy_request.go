package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateL7PolicyRequest Request Object
type CreateL7PolicyRequest struct {
	Body *CreateL7PolicyRequestBody `json:"body,omitempty"`
}

func (o CreateL7PolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7PolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyRequest", string(data)}, " ")
}
