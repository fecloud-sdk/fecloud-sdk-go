package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListServiceDescribeDetailsResponse struct {
	Id *string `json:"id,omitempty"`

	ServiceName *string `json:"service_name,omitempty"`

	ServiceType *ListServiceDescribeDetailsResponseServiceType `json:"service_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	IsCharge *bool `json:"is_charge,omitempty"`

	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	EnablePolicy   *bool `json:"enable_policy,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListServiceDescribeDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceDescribeDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceDescribeDetailsResponse", string(data)}, " ")
}

type ListServiceDescribeDetailsResponseServiceType struct {
	value string
}

type ListServiceDescribeDetailsResponseServiceTypeEnum struct {
	INTERFACE ListServiceDescribeDetailsResponseServiceType
}

func GetListServiceDescribeDetailsResponseServiceTypeEnum() ListServiceDescribeDetailsResponseServiceTypeEnum {
	return ListServiceDescribeDetailsResponseServiceTypeEnum{
		INTERFACE: ListServiceDescribeDetailsResponseServiceType{
			value: "interface",
		},
	}
}

func (c ListServiceDescribeDetailsResponseServiceType) Value() string {
	return c.value
}

func (c ListServiceDescribeDetailsResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceDescribeDetailsResponseServiceType) UnmarshalJSON(b []byte) error {
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
