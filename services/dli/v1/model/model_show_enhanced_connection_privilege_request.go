package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowEnhancedConnectionPrivilegeRequest struct {
	ConnectionId string `json:"connection_id"`
}

func (o ShowEnhancedConnectionPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnhancedConnectionPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"ShowEnhancedConnectionPrivilegeRequest", string(data)}, " ")
}
