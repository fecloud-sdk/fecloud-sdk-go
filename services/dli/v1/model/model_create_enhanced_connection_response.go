package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateEnhancedConnectionResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	ConnectionId   *string `json:"connection_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEnhancedConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnhancedConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateEnhancedConnectionResponse", string(data)}, " ")
}
