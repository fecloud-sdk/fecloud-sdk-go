package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowTopicStatusResponse struct {
	MaxOffset *int32 `json:"max_offset,omitempty"`

	MinOffset *int32 `json:"min_offset,omitempty"`

	Brokers        *[]ShowTopicStatusRespBrokers `json:"brokers,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowTopicStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusResponse", string(data)}, " ")
}
