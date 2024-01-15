package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopicResponse", string(data)}, " ")
}
