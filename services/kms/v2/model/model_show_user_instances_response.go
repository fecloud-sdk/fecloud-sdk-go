package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowUserInstancesResponse struct {
	InstanceNum    *int32 `json:"instance_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowUserInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowUserInstancesResponse", string(data)}, " ")
}
