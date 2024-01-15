package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowKafkaTagsResponse struct {
	Tags           *[]TagEntity `json:"tags,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowKafkaTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaTagsResponse", string(data)}, " ")
}
