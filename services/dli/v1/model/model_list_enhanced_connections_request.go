package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListEnhancedConnectionsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Name *string `json:"name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Status *string `json:"status,omitempty"`

	Tags *string `json:"tags,omitempty"`
}

func (o ListEnhancedConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnhancedConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListEnhancedConnectionsRequest", string(data)}, " ")
}
