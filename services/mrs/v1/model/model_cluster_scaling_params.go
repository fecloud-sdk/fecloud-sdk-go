package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ClusterScalingParams struct {
	OrderId *string `json:"order_id,omitempty"`

	ScaleType ClusterScalingParamsScaleType `json:"scale_type"`

	NodeId string `json:"node_id"`

	NodeGroup *string `json:"node_group,omitempty"`

	TaskNodeInfo *TaskNodeInfo `json:"task_node_info,omitempty"`

	Instances int32 `json:"instances"`

	SkipBootstrapScripts *string `json:"skip_bootstrap_scripts,omitempty"`

	ScaleWithoutStart *bool `json:"scale_without_start,omitempty"`

	ServerIds *[]string `json:"server_ids,omitempty"`
}

func (o ClusterScalingParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterScalingParams struct{}"
	}

	return strings.Join([]string{"ClusterScalingParams", string(data)}, " ")
}

type ClusterScalingParamsScaleType struct {
	value string
}

type ClusterScalingParamsScaleTypeEnum struct {
	SCALE_IN  ClusterScalingParamsScaleType
	SCALE_OUT ClusterScalingParamsScaleType
}

func GetClusterScalingParamsScaleTypeEnum() ClusterScalingParamsScaleTypeEnum {
	return ClusterScalingParamsScaleTypeEnum{
		SCALE_IN: ClusterScalingParamsScaleType{
			value: "scale_in",
		},
		SCALE_OUT: ClusterScalingParamsScaleType{
			value: "scale_out",
		},
	}
}

func (c ClusterScalingParamsScaleType) Value() string {
	return c.value
}

func (c ClusterScalingParamsScaleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterScalingParamsScaleType) UnmarshalJSON(b []byte) error {
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
