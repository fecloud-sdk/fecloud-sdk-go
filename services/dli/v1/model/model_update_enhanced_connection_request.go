package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEnhancedConnectionRequest struct {
	ConnectionId string `json:"connection_id"`

	Body *UpdateEnhancedConnectionRequestBody `json:"body,omitempty"`
}

func (o UpdateEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateEnhancedConnectionRequest", string(data)}, " ")
}
