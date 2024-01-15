package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEnhancedConnectionResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEnhancedConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnhancedConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateEnhancedConnectionResponse", string(data)}, " ")
}
