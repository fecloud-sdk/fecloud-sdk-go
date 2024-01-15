package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListConsumerGroupOfTopicResponse struct {
	Groups         *[]string `json:"groups,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListConsumerGroupOfTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumerGroupOfTopicResponse struct{}"
	}

	return strings.Join([]string{"ListConsumerGroupOfTopicResponse", string(data)}, " ")
}
