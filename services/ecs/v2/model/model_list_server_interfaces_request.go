package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListServerInterfacesRequest Request Object
type ListServerInterfacesRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ListServerInterfacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerInterfacesRequest struct{}"
	}

	return strings.Join([]string{"ListServerInterfacesRequest", string(data)}, " ")
}
