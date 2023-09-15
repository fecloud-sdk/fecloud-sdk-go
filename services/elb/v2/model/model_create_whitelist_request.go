package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateWhitelistRequest Request Object
type CreateWhitelistRequest struct {
	Body *CreateWhitelistRequestBody `json:"body,omitempty"`
}

func (o CreateWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhitelistRequest struct{}"
	}

	return strings.Join([]string{"CreateWhitelistRequest", string(data)}, " ")
}
