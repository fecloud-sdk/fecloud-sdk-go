package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DiskusageEntity struct {
	BrokerName *string `json:"broker_name,omitempty"`

	DataDiskSize *string `json:"data_disk_size,omitempty"`

	DataDiskUse *string `json:"data_disk_use,omitempty"`

	DataDiskFree *string `json:"data_disk_free,omitempty"`

	DataDiskUsePercentage *string `json:"data_disk_use_percentage,omitempty"`

	Status *string `json:"status,omitempty"`

	TopicList *[]DiskusageTopicEntity `json:"topic_list,omitempty"`
}

func (o DiskusageEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskusageEntity struct{}"
	}

	return strings.Join([]string{"DiskusageEntity", string(data)}, " ")
}
