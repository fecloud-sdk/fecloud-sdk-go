package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AssociateQueueToEnhancedConnectionResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateQueueToEnhancedConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateQueueToEnhancedConnectionResponse struct{}"
	}

	return strings.Join([]string{"AssociateQueueToEnhancedConnectionResponse", string(data)}, " ")
}
