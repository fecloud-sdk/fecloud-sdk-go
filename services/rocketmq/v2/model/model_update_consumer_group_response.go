package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateConsumerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateConsumerGroupResponse", string(data)}, " ")
}
