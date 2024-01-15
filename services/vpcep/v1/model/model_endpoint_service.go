package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type EndpointService struct {
	Id *string `json:"id,omitempty"`

	Owner *string `json:"owner,omitempty"`

	ServiceName *string `json:"service_name,omitempty"`

	ServiceType *EndpointServiceServiceType `json:"service_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	IsCharge *bool `json:"is_charge,omitempty"`

	EnablePolicy *bool `json:"enable_policy,omitempty"`
}

func (o EndpointService) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointService struct{}"
	}

	return strings.Join([]string{"EndpointService", string(data)}, " ")
}

type EndpointServiceServiceType struct {
	value string
}

type EndpointServiceServiceTypeEnum struct {
	INTERFACE EndpointServiceServiceType
	GATEWAY   EndpointServiceServiceType
}

func GetEndpointServiceServiceTypeEnum() EndpointServiceServiceTypeEnum {
	return EndpointServiceServiceTypeEnum{
		INTERFACE: EndpointServiceServiceType{
			value: "interface",
		},
		GATEWAY: EndpointServiceServiceType{
			value: "gateway",
		},
	}
}

func (c EndpointServiceServiceType) Value() string {
	return c.value
}

func (c EndpointServiceServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointServiceServiceType) UnmarshalJSON(b []byte) error {
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
