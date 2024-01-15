package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type OpenProxyRequest struct {
	FlavorId *string `json:"flavor_id,omitempty"`

	NodeNum *int32 `json:"node_num,omitempty"`
}

func (o OpenProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenProxyRequest struct{}"
	}

	return strings.Join([]string{"OpenProxyRequest", string(data)}, " ")
}
