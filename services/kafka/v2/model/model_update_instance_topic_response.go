package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicResponse", string(data)}, " ")
}
