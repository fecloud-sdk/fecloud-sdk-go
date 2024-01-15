package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListLinkAggregationGroupRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Fields *[]string `json:"fields,omitempty"`

	SortDir *[]ListLinkAggregationGroupRequestSortDir `json:"sort_dir,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	Id *[]string `json:"id,omitempty"`
}

func (o ListLinkAggregationGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLinkAggregationGroupRequest struct{}"
	}

	return strings.Join([]string{"ListLinkAggregationGroupRequest", string(data)}, " ")
}

type ListLinkAggregationGroupRequestSortDir struct {
	value string
}

type ListLinkAggregationGroupRequestSortDirEnum struct {
	ASC  ListLinkAggregationGroupRequestSortDir
	DESC ListLinkAggregationGroupRequestSortDir
}

func GetListLinkAggregationGroupRequestSortDirEnum() ListLinkAggregationGroupRequestSortDirEnum {
	return ListLinkAggregationGroupRequestSortDirEnum{
		ASC: ListLinkAggregationGroupRequestSortDir{
			value: "asc",
		},
		DESC: ListLinkAggregationGroupRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListLinkAggregationGroupRequestSortDir) Value() string {
	return c.value
}

func (c ListLinkAggregationGroupRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLinkAggregationGroupRequestSortDir) UnmarshalJSON(b []byte) error {
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
