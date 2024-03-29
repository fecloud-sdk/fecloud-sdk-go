package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type AgentDimension struct {
	Name *AgentDimensionName `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`

	OriginValue *string `json:"origin_value,omitempty"`
}

func (o AgentDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentDimension struct{}"
	}

	return strings.Join([]string{"AgentDimension", string(data)}, " ")
}

type AgentDimensionName struct {
	value string
}

type AgentDimensionNameEnum struct {
	MOUNT_POINT AgentDimensionName
	DISK        AgentDimensionName
	PROC        AgentDimensionName
	GPU         AgentDimensionName
	RAID        AgentDimensionName
}

func GetAgentDimensionNameEnum() AgentDimensionNameEnum {
	return AgentDimensionNameEnum{
		MOUNT_POINT: AgentDimensionName{
			value: "mount_point",
		},
		DISK: AgentDimensionName{
			value: "disk",
		},
		PROC: AgentDimensionName{
			value: "proc",
		},
		GPU: AgentDimensionName{
			value: "gpu",
		},
		RAID: AgentDimensionName{
			value: "raid",
		},
	}
}

func (c AgentDimensionName) Value() string {
	return c.value
}

func (c AgentDimensionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgentDimensionName) UnmarshalJSON(b []byte) error {
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
