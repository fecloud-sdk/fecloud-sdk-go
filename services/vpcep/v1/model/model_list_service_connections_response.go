package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListServiceConnectionsResponse struct {
	Connections *[]ConnectionEndpoints `json:"connections,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListServiceConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceConnectionsResponse", string(data)}, " ")
}
