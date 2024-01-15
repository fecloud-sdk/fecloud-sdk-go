package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateOrDeleteRocketmqTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteRocketmqTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteRocketmqTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteRocketmqTagResponse", string(data)}, " ")
}
