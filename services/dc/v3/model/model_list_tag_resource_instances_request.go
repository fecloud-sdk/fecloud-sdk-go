package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListTagResourceInstancesRequest struct {
	ResourceType ListTagResourceInstancesRequestResourceType `json:"resource_type"`

	Body *ListTagResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListTagResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListTagResourceInstancesRequest", string(data)}, " ")
}

type ListTagResourceInstancesRequestResourceType struct {
	value string
}

type ListTagResourceInstancesRequestResourceTypeEnum struct {
	DC_DIRECTCONNECT ListTagResourceInstancesRequestResourceType
	DC_VGW           ListTagResourceInstancesRequestResourceType
	DC_VIF           ListTagResourceInstancesRequestResourceType
}

func GetListTagResourceInstancesRequestResourceTypeEnum() ListTagResourceInstancesRequestResourceTypeEnum {
	return ListTagResourceInstancesRequestResourceTypeEnum{
		DC_DIRECTCONNECT: ListTagResourceInstancesRequestResourceType{
			value: "dc-directconnect",
		},
		DC_VGW: ListTagResourceInstancesRequestResourceType{
			value: "dc-vgw",
		},
		DC_VIF: ListTagResourceInstancesRequestResourceType{
			value: "dc-vif",
		},
	}
}

func (c ListTagResourceInstancesRequestResourceType) Value() string {
	return c.value
}

func (c ListTagResourceInstancesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagResourceInstancesRequestResourceType) UnmarshalJSON(b []byte) error {
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
