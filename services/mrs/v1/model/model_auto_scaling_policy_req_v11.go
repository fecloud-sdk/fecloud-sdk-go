package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type AutoScalingPolicyReqV11 struct {
	NodeGroup AutoScalingPolicyReqV11NodeGroup `json:"node_group"`

	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy"`
}

func (o AutoScalingPolicyReqV11) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoScalingPolicyReqV11 struct{}"
	}

	return strings.Join([]string{"AutoScalingPolicyReqV11", string(data)}, " ")
}

type AutoScalingPolicyReqV11NodeGroup struct {
	value string
}

type AutoScalingPolicyReqV11NodeGroupEnum struct {
	TASK_NODE_DEFAULT_GROUP AutoScalingPolicyReqV11NodeGroup
}

func GetAutoScalingPolicyReqV11NodeGroupEnum() AutoScalingPolicyReqV11NodeGroupEnum {
	return AutoScalingPolicyReqV11NodeGroupEnum{
		TASK_NODE_DEFAULT_GROUP: AutoScalingPolicyReqV11NodeGroup{
			value: "task_node_default_group",
		},
	}
}

func (c AutoScalingPolicyReqV11NodeGroup) Value() string {
	return c.value
}

func (c AutoScalingPolicyReqV11NodeGroup) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoScalingPolicyReqV11NodeGroup) UnmarshalJSON(b []byte) error {
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
