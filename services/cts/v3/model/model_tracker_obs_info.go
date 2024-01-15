package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type TrackerObsInfo struct {
	BucketName *string `json:"bucket_name,omitempty"`

	FilePrefixName *string `json:"file_prefix_name,omitempty"`

	IsObsCreated *bool `json:"is_obs_created,omitempty"`

	CompressType *TrackerObsInfoCompressType `json:"compress_type,omitempty"`

	IsSortByService *bool `json:"is_sort_by_service,omitempty"`
}

func (o TrackerObsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerObsInfo struct{}"
	}

	return strings.Join([]string{"TrackerObsInfo", string(data)}, " ")
}

type TrackerObsInfoCompressType struct {
	value string
}

type TrackerObsInfoCompressTypeEnum struct {
	GZIP TrackerObsInfoCompressType
	JSON TrackerObsInfoCompressType
}

func GetTrackerObsInfoCompressTypeEnum() TrackerObsInfoCompressTypeEnum {
	return TrackerObsInfoCompressTypeEnum{
		GZIP: TrackerObsInfoCompressType{
			value: "gzip",
		},
		JSON: TrackerObsInfoCompressType{
			value: "json",
		},
	}
}

func (c TrackerObsInfoCompressType) Value() string {
	return c.value
}

func (c TrackerObsInfoCompressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrackerObsInfoCompressType) UnmarshalJSON(b []byte) error {
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
