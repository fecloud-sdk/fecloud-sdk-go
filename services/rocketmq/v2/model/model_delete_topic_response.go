package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicResponse struct{}"
	}

	return strings.Join([]string{"DeleteTopicResponse", string(data)}, " ")
}
