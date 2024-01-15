package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DisassociateQueueFromEnhancedConnectionRequest struct {
	ConnectionId string `json:"connection_id"`

	Body *DisassociateQueueFromEnhancedConnectionRequestBody `json:"body,omitempty"`
}

func (o DisassociateQueueFromEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateQueueFromEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"DisassociateQueueFromEnhancedConnectionRequest", string(data)}, " ")
}
