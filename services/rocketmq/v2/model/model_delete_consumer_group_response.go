package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteConsumerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteConsumerGroupResponse", string(data)}, " ")
}
