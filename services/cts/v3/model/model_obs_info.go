package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ObsInfo struct {
	BucketName *string `json:"bucket_name,omitempty"`

	FilePrefixName *string `json:"file_prefix_name,omitempty"`

	IsObsCreated *bool `json:"is_obs_created,omitempty"`

	CompressType *ObsInfoCompressType `json:"compress_type,omitempty"`

	IsSortByService *bool `json:"is_sort_by_service,omitempty"`
}

func (o ObsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsInfo struct{}"
	}

	return strings.Join([]string{"ObsInfo", string(data)}, " ")
}

type ObsInfoCompressType struct {
	value string
}

type ObsInfoCompressTypeEnum struct {
	GZIP ObsInfoCompressType
	JSON ObsInfoCompressType
}

func GetObsInfoCompressTypeEnum() ObsInfoCompressTypeEnum {
	return ObsInfoCompressTypeEnum{
		GZIP: ObsInfoCompressType{
			value: "gzip",
		},
		JSON: ObsInfoCompressType{
			value: "json",
		},
	}
}

func (c ObsInfoCompressType) Value() string {
	return c.value
}

func (c ObsInfoCompressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObsInfoCompressType) UnmarshalJSON(b []byte) error {
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
