package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AcceptOrRejectEndpointResponse struct {
	Connections    *[]ConnectionEndpoints `json:"connections,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o AcceptOrRejectEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptOrRejectEndpointResponse struct{}"
	}

	return strings.Join([]string{"AcceptOrRejectEndpointResponse", string(data)}, " ")
}
