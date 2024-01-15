package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListHostedDirectConnectsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Fields *[]string `json:"fields,omitempty"`

	SortDir *[]ListHostedDirectConnectsRequestSortDir `json:"sort_dir,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	HostingId *[]string `json:"hosting_id,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Name *[]string `json:"name,omitempty"`
}

func (o ListHostedDirectConnectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostedDirectConnectsRequest struct{}"
	}

	return strings.Join([]string{"ListHostedDirectConnectsRequest", string(data)}, " ")
}

type ListHostedDirectConnectsRequestSortDir struct {
	value string
}

type ListHostedDirectConnectsRequestSortDirEnum struct {
	ASC  ListHostedDirectConnectsRequestSortDir
	DESC ListHostedDirectConnectsRequestSortDir
}

func GetListHostedDirectConnectsRequestSortDirEnum() ListHostedDirectConnectsRequestSortDirEnum {
	return ListHostedDirectConnectsRequestSortDirEnum{
		ASC: ListHostedDirectConnectsRequestSortDir{
			value: "asc",
		},
		DESC: ListHostedDirectConnectsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListHostedDirectConnectsRequestSortDir) Value() string {
	return c.value
}

func (c ListHostedDirectConnectsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostedDirectConnectsRequestSortDir) UnmarshalJSON(b []byte) error {
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
