package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateOrDeleteKafkaTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteKafkaTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteKafkaTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteKafkaTagResponse", string(data)}, " ")
}
