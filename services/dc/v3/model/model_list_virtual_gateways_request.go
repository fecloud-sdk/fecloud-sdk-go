package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListVirtualGatewaysRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Fields *[]string `json:"fields,omitempty"`

	SortDir *[]ListVirtualGatewaysRequestSortDir `json:"sort_dir,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	Id *[]string `json:"id,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	VpcId *[]string `json:"vpc_id,omitempty"`
}

func (o ListVirtualGatewaysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirtualGatewaysRequest struct{}"
	}

	return strings.Join([]string{"ListVirtualGatewaysRequest", string(data)}, " ")
}

type ListVirtualGatewaysRequestSortDir struct {
	value string
}

type ListVirtualGatewaysRequestSortDirEnum struct {
	ASC  ListVirtualGatewaysRequestSortDir
	DESC ListVirtualGatewaysRequestSortDir
}

func GetListVirtualGatewaysRequestSortDirEnum() ListVirtualGatewaysRequestSortDirEnum {
	return ListVirtualGatewaysRequestSortDirEnum{
		ASC: ListVirtualGatewaysRequestSortDir{
			value: "asc",
		},
		DESC: ListVirtualGatewaysRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListVirtualGatewaysRequestSortDir) Value() string {
	return c.value
}

func (c ListVirtualGatewaysRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVirtualGatewaysRequestSortDir) UnmarshalJSON(b []byte) error {
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
