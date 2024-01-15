package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ObsDestinationDescriptor struct {
	Topics string `json:"topics"`

	TopicsRegex *string `json:"topics_regex,omitempty"`

	ConsumerStrategy ObsDestinationDescriptorConsumerStrategy `json:"consumer_strategy"`

	DestinationFileType ObsDestinationDescriptorDestinationFileType `json:"destination_file_type"`

	AccessKey string `json:"access_key"`

	SecretKey string `json:"secret_key"`

	ObsBucketName string `json:"obs_bucket_name"`

	ObsPath *string `json:"obs_path,omitempty"`

	PartitionFormat *string `json:"partition_format,omitempty"`

	RecordDelimiter *string `json:"record_delimiter,omitempty"`

	DeliverTimeInterval int32 `json:"deliver_time_interval"`
}

func (o ObsDestinationDescriptor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsDestinationDescriptor struct{}"
	}

	return strings.Join([]string{"ObsDestinationDescriptor", string(data)}, " ")
}

type ObsDestinationDescriptorConsumerStrategy struct {
	value string
}

type ObsDestinationDescriptorConsumerStrategyEnum struct {
	LATEST   ObsDestinationDescriptorConsumerStrategy
	EARLIEST ObsDestinationDescriptorConsumerStrategy
}

func GetObsDestinationDescriptorConsumerStrategyEnum() ObsDestinationDescriptorConsumerStrategyEnum {
	return ObsDestinationDescriptorConsumerStrategyEnum{
		LATEST: ObsDestinationDescriptorConsumerStrategy{
			value: "latest",
		},
		EARLIEST: ObsDestinationDescriptorConsumerStrategy{
			value: "earliest",
		},
	}
}

func (c ObsDestinationDescriptorConsumerStrategy) Value() string {
	return c.value
}

func (c ObsDestinationDescriptorConsumerStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObsDestinationDescriptorConsumerStrategy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ObsDestinationDescriptorDestinationFileType struct {
	value string
}

type ObsDestinationDescriptorDestinationFileTypeEnum struct {
	TEXT ObsDestinationDescriptorDestinationFileType
}

func GetObsDestinationDescriptorDestinationFileTypeEnum() ObsDestinationDescriptorDestinationFileTypeEnum {
	return ObsDestinationDescriptorDestinationFileTypeEnum{
		TEXT: ObsDestinationDescriptorDestinationFileType{
			value: "TEXT",
		},
	}
}

func (c ObsDestinationDescriptorDestinationFileType) Value() string {
	return c.value
}

func (c ObsDestinationDescriptorDestinationFileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObsDestinationDescriptorDestinationFileType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
