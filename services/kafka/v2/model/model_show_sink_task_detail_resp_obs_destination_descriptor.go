package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSinkTaskDetailRespObsDestinationDescriptor struct {
	ConsumerStrategy *string `json:"consumer_strategy,omitempty"`

	DestinationFileType *string `json:"destination_file_type,omitempty"`

	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	ObsPath *string `json:"obs_path,omitempty"`

	PartitionFormat *string `json:"partition_format,omitempty"`

	RecordDelimiter *string `json:"record_delimiter,omitempty"`

	DeliverTimeInterval *int32 `json:"deliver_time_interval,omitempty"`

	ObsPartSize *int64 `json:"obs_part_size,omitempty"`
}

func (o ShowSinkTaskDetailRespObsDestinationDescriptor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSinkTaskDetailRespObsDestinationDescriptor struct{}"
	}

	return strings.Join([]string{"ShowSinkTaskDetailRespObsDestinationDescriptor", string(data)}, " ")
}
