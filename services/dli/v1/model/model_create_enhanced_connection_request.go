package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateEnhancedConnectionRequest struct {
	Body *CreateEnhancedConnectionRequestBody `json:"body,omitempty"`
}

func (o CreateEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateEnhancedConnectionRequest", string(data)}, " ")
}
