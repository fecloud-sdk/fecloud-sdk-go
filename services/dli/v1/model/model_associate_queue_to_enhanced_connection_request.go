package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AssociateQueueToEnhancedConnectionRequest struct {
	ConnectionId string `json:"connection_id"`

	Body *AssociateQueueToEnhancedConnectionRequestBody `json:"body,omitempty"`
}

func (o AssociateQueueToEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateQueueToEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"AssociateQueueToEnhancedConnectionRequest", string(data)}, " ")
}
