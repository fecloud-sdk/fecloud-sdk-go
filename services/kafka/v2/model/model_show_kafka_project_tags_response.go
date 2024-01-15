package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowKafkaProjectTagsResponse struct {
	Tags           *[]TagMultyValueEntity `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowKafkaProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaProjectTagsResponse", string(data)}, " ")
}
