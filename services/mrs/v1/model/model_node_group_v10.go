package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type NodeGroupV10 struct {
	GroupName *string `json:"GroupName,omitempty"`

	NodeNum *int32 `json:"NodeNum,omitempty"`

	NodeSize *string `json:"NodeSize,omitempty"`

	NodeSpecId *string `json:"NodeSpecId,omitempty"`

	VmProductId *string `json:"VmProductId,omitempty"`

	VmSpecCode *string `json:"VmSpecCode,omitempty"`

	NodeProductId *string `json:"NodeProductId,omitempty"`

	RootVolumeSize *int32 `json:"RootVolumeSize,omitempty"`

	RootVolumeProductId *string `json:"RootVolumeProductId,omitempty"`

	RootVolumeType *string `json:"RootVolumeType,omitempty"`

	RootVolumeResourceSpecCode *string `json:"RootVolumeResourceSpecCode,omitempty"`

	RootVolumeResourceType *string `json:"RootVolumeResourceType,omitempty"`

	DataVolumeType *NodeGroupV10DataVolumeType `json:"DataVolumeType,omitempty"`

	DataVolumeCount *int32 `json:"DataVolumeCount,omitempty"`

	DataVolumeSize *int32 `json:"DataVolumeSize,omitempty"`

	DataVolumeProductId *string `json:"DataVolumeProductId,omitempty"`

	DataVolumeResourceSpecCode *string `json:"DataVolumeResourceSpecCode,omitempty"`

	DataVolumeResourceType *string `json:"DataVolumeResourceType,omitempty"`
}

func (o NodeGroupV10) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeGroupV10 struct{}"
	}

	return strings.Join([]string{"NodeGroupV10", string(data)}, " ")
}

type NodeGroupV10DataVolumeType struct {
	value string
}

type NodeGroupV10DataVolumeTypeEnum struct {
	SATA  NodeGroupV10DataVolumeType
	SAS   NodeGroupV10DataVolumeType
	SSD   NodeGroupV10DataVolumeType
	GPSSD NodeGroupV10DataVolumeType
}

func GetNodeGroupV10DataVolumeTypeEnum() NodeGroupV10DataVolumeTypeEnum {
	return NodeGroupV10DataVolumeTypeEnum{
		SATA: NodeGroupV10DataVolumeType{
			value: "SATA",
		},
		SAS: NodeGroupV10DataVolumeType{
			value: "SAS",
		},
		SSD: NodeGroupV10DataVolumeType{
			value: "SSD",
		},
		GPSSD: NodeGroupV10DataVolumeType{
			value: "GPSSD",
		},
	}
}

func (c NodeGroupV10DataVolumeType) Value() string {
	return c.value
}

func (c NodeGroupV10DataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeGroupV10DataVolumeType) UnmarshalJSON(b []byte) error {
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
