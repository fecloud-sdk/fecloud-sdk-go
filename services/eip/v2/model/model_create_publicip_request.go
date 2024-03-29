package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePublicipRequest Request Object
type CreatePublicipRequest struct {
	Body *CreatePublicipRequestBody `json:"body,omitempty"`
}

func (o CreatePublicipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicipRequest struct{}"
	}

	return strings.Join([]string{"CreatePublicipRequest", string(data)}, " ")
}
