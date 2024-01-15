package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceConsumerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConsumerGroupResponse", string(data)}, " ")
}
