package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowEnhancedConnectionRequest struct {
	ConnectionId string `json:"connection_id"`
}

func (o ShowEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowEnhancedConnectionRequest", string(data)}, " ")
}
