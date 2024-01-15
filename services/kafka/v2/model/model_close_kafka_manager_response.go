package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CloseKafkaManagerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CloseKafkaManagerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseKafkaManagerResponse struct{}"
	}

	return strings.Join([]string{"CloseKafkaManagerResponse", string(data)}, " ")
}
