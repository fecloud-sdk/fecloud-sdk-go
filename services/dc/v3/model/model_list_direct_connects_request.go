package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListDirectConnectsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Fields *[]string `json:"fields,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	SortDir *[]ListDirectConnectsRequestSortDir `json:"sort_dir,omitempty"`

	HostingId *[]string `json:"hosting_id,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Name *[]string `json:"name,omitempty"`
}

func (o ListDirectConnectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectConnectsRequest struct{}"
	}

	return strings.Join([]string{"ListDirectConnectsRequest", string(data)}, " ")
}

type ListDirectConnectsRequestSortDir struct {
	value string
}

type ListDirectConnectsRequestSortDirEnum struct {
	ASC  ListDirectConnectsRequestSortDir
	DESC ListDirectConnectsRequestSortDir
}

func GetListDirectConnectsRequestSortDirEnum() ListDirectConnectsRequestSortDirEnum {
	return ListDirectConnectsRequestSortDirEnum{
		ASC: ListDirectConnectsRequestSortDir{
			value: "asc",
		},
		DESC: ListDirectConnectsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListDirectConnectsRequestSortDir) Value() string {
	return c.value
}

func (c ListDirectConnectsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDirectConnectsRequestSortDir) UnmarshalJSON(b []byte) error {
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
