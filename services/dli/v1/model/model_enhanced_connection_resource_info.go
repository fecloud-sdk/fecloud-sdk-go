package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EnhancedConnectionResourceInfo struct {
	PeerId *string `json:"peer_id,omitempty"`

	Status *string `json:"status,omitempty"`

	Name *string `json:"name,omitempty"`

	ErrMsg *string `json:"err_msg,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o EnhancedConnectionResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnhancedConnectionResourceInfo struct{}"
	}

	return strings.Join([]string{"EnhancedConnectionResourceInfo", string(data)}, " ")
}
