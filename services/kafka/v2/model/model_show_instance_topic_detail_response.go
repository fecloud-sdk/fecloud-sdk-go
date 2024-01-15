package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceTopicDetailResponse struct {
	Topic *string `json:"topic,omitempty"`

	Partitions *[]ShowInstanceTopicDetailRespPartitions `json:"partitions,omitempty"`

	GroupSubscribed *[]string `json:"group_subscribed,omitempty"`
	HttpStatusCode  int       `json:"-"`
}

func (o ShowInstanceTopicDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopicDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopicDetailResponse", string(data)}, " ")
}
