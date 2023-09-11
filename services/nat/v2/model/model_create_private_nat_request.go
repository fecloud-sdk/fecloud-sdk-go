package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePrivateNatRequest Request Object
type CreatePrivateNatRequest struct {
	Body *CreatePrivateNatRequestBody `json:"body,omitempty"`
}

func (o CreatePrivateNatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatRequest", string(data)}, " ")
}
