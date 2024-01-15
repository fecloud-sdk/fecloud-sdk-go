package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowEnhancedConnectionPrivilegeResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	ConnectionId *string `json:"connection_id,omitempty"`

	Privileges     *[]ProjectPrivilege `json:"privileges,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowEnhancedConnectionPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnhancedConnectionPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"ShowEnhancedConnectionPrivilegeResponse", string(data)}, " ")
}
