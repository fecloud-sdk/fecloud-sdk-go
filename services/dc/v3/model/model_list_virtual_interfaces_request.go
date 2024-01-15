package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListVirtualInterfacesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Fields *[]string `json:"fields,omitempty"`

	SortDir *[]ListVirtualInterfacesRequestSortDir `json:"sort_dir,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Status *[]string `json:"status,omitempty"`

	DirectConnectId *[]string `json:"direct_connect_id,omitempty"`

	VgwId *[]string `json:"vgw_id,omitempty"`
}

func (o ListVirtualInterfacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirtualInterfacesRequest struct{}"
	}

	return strings.Join([]string{"ListVirtualInterfacesRequest", string(data)}, " ")
}

type ListVirtualInterfacesRequestSortDir struct {
	value string
}

type ListVirtualInterfacesRequestSortDirEnum struct {
	ASC  ListVirtualInterfacesRequestSortDir
	DESC ListVirtualInterfacesRequestSortDir
}

func GetListVirtualInterfacesRequestSortDirEnum() ListVirtualInterfacesRequestSortDirEnum {
	return ListVirtualInterfacesRequestSortDirEnum{
		ASC: ListVirtualInterfacesRequestSortDir{
			value: "asc",
		},
		DESC: ListVirtualInterfacesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListVirtualInterfacesRequestSortDir) Value() string {
	return c.value
}

func (c ListVirtualInterfacesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVirtualInterfacesRequestSortDir) UnmarshalJSON(b []byte) error {
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
