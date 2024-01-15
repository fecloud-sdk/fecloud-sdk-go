package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryConnectionsResponse struct {
	ClientIp string `json:"client_ip"`

	Count int32 `json:"count"`
}

func (o QueryConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryConnectionsResponse struct{}"
	}

	return strings.Join([]string{"QueryConnectionsResponse", string(data)}, " ")
}
