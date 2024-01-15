package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ClientNetworkRequestBody struct {
	ClientNetworkRanges []string `json:"client_network_ranges"`
}

func (o ClientNetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientNetworkRequestBody struct{}"
	}

	return strings.Join([]string{"ClientNetworkRequestBody", string(data)}, " ")
}
