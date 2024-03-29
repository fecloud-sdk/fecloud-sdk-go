package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListAgentDimensionInfoRequest struct {
	ContentType string `json:"Content-Type"`

	InstanceId string `json:"instance_id"`

	DimName ListAgentDimensionInfoRequestDimName `json:"dim_name"`

	DimValue *string `json:"dim_value,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAgentDimensionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentDimensionInfoRequest struct{}"
	}

	return strings.Join([]string{"ListAgentDimensionInfoRequest", string(data)}, " ")
}

type ListAgentDimensionInfoRequestDimName struct {
	value string
}

type ListAgentDimensionInfoRequestDimNameEnum struct {
	MOUNT_POINT ListAgentDimensionInfoRequestDimName
	DISK        ListAgentDimensionInfoRequestDimName
	PROC        ListAgentDimensionInfoRequestDimName
	GPU         ListAgentDimensionInfoRequestDimName
	RAID        ListAgentDimensionInfoRequestDimName
}

func GetListAgentDimensionInfoRequestDimNameEnum() ListAgentDimensionInfoRequestDimNameEnum {
	return ListAgentDimensionInfoRequestDimNameEnum{
		MOUNT_POINT: ListAgentDimensionInfoRequestDimName{
			value: "mount_point",
		},
		DISK: ListAgentDimensionInfoRequestDimName{
			value: "disk",
		},
		PROC: ListAgentDimensionInfoRequestDimName{
			value: "proc",
		},
		GPU: ListAgentDimensionInfoRequestDimName{
			value: "gpu",
		},
		RAID: ListAgentDimensionInfoRequestDimName{
			value: "raid",
		},
	}
}

func (c ListAgentDimensionInfoRequestDimName) Value() string {
	return c.value
}

func (c ListAgentDimensionInfoRequestDimName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentDimensionInfoRequestDimName) UnmarshalJSON(b []byte) error {
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
