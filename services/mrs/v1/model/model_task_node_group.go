package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type TaskNodeGroup struct {
	NodeNum int32 `json:"node_num"`

	NodeSize string `json:"node_size"`

	DataVolumeType TaskNodeGroupDataVolumeType `json:"data_volume_type"`

	DataVolumeCount int32 `json:"data_volume_count"`

	DataVolumeSize int32 `json:"data_volume_size"`

	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy,omitempty"`
}

func (o TaskNodeGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskNodeGroup struct{}"
	}

	return strings.Join([]string{"TaskNodeGroup", string(data)}, " ")
}

type TaskNodeGroupDataVolumeType struct {
	value string
}

type TaskNodeGroupDataVolumeTypeEnum struct {
	SATA  TaskNodeGroupDataVolumeType
	SAS   TaskNodeGroupDataVolumeType
	SSD   TaskNodeGroupDataVolumeType
	GPSSD TaskNodeGroupDataVolumeType
}

func GetTaskNodeGroupDataVolumeTypeEnum() TaskNodeGroupDataVolumeTypeEnum {
	return TaskNodeGroupDataVolumeTypeEnum{
		SATA: TaskNodeGroupDataVolumeType{
			value: "SATA",
		},
		SAS: TaskNodeGroupDataVolumeType{
			value: "SAS",
		},
		SSD: TaskNodeGroupDataVolumeType{
			value: "SSD",
		},
		GPSSD: TaskNodeGroupDataVolumeType{
			value: "GPSSD",
		},
	}
}

func (c TaskNodeGroupDataVolumeType) Value() string {
	return c.value
}

func (c TaskNodeGroupDataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskNodeGroupDataVolumeType) UnmarshalJSON(b []byte) error {
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
