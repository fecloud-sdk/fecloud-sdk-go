package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceAutoCreateTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceAutoCreateTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceAutoCreateTopicResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceAutoCreateTopicResponse", string(data)}, " ")
}
