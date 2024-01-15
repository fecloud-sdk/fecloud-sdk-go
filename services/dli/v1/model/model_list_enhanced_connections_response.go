package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListEnhancedConnectionsResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Connections *[]EnhancedConnectionInfo `json:"connections,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnhancedConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnhancedConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListEnhancedConnectionsResponse", string(data)}, " ")
}
