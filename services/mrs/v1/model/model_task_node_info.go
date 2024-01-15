package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type TaskNodeInfo struct {
	NodeSize string `json:"node_size"`

	DataVolumeType TaskNodeInfoDataVolumeType `json:"data_volume_type"`

	DataVolumeCount int32 `json:"data_volume_count"`

	DataVolumeSize int32 `json:"data_volume_size"`
}

func (o TaskNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskNodeInfo struct{}"
	}

	return strings.Join([]string{"TaskNodeInfo", string(data)}, " ")
}

type TaskNodeInfoDataVolumeType struct {
	value string
}

type TaskNodeInfoDataVolumeTypeEnum struct {
	SATA  TaskNodeInfoDataVolumeType
	SAS   TaskNodeInfoDataVolumeType
	SSD   TaskNodeInfoDataVolumeType
	GPSSD TaskNodeInfoDataVolumeType
}

func GetTaskNodeInfoDataVolumeTypeEnum() TaskNodeInfoDataVolumeTypeEnum {
	return TaskNodeInfoDataVolumeTypeEnum{
		SATA: TaskNodeInfoDataVolumeType{
			value: "SATA",
		},
		SAS: TaskNodeInfoDataVolumeType{
			value: "SAS",
		},
		SSD: TaskNodeInfoDataVolumeType{
			value: "SSD",
		},
		GPSSD: TaskNodeInfoDataVolumeType{
			value: "GPSSD",
		},
	}
}

func (c TaskNodeInfoDataVolumeType) Value() string {
	return c.value
}

func (c TaskNodeInfoDataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskNodeInfoDataVolumeType) UnmarshalJSON(b []byte) error {
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
