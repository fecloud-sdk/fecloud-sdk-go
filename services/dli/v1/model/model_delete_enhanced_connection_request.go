package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteEnhancedConnectionRequest struct {
	ConnectionId string `json:"connection_id"`
}

func (o DeleteEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteEnhancedConnectionRequest", string(data)}, " ")
}
