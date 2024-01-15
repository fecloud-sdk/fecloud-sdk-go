package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListServiceConnectionsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Id *string `json:"id,omitempty"`

	MarkerId *string `json:"marker_id,omitempty"`

	Status *string `json:"status,omitempty"`

	SortKey *ListServiceConnectionsRequestSortKey `json:"sort_key,omitempty"`

	SortDir *ListServiceConnectionsRequestSortDir `json:"sort_dir,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListServiceConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListServiceConnectionsRequest", string(data)}, " ")
}

type ListServiceConnectionsRequestSortKey struct {
	value string
}

type ListServiceConnectionsRequestSortKeyEnum struct {
	CREATE_AT ListServiceConnectionsRequestSortKey
	UPDATE_AT ListServiceConnectionsRequestSortKey
}

func GetListServiceConnectionsRequestSortKeyEnum() ListServiceConnectionsRequestSortKeyEnum {
	return ListServiceConnectionsRequestSortKeyEnum{
		CREATE_AT: ListServiceConnectionsRequestSortKey{
			value: "create_at",
		},
		UPDATE_AT: ListServiceConnectionsRequestSortKey{
			value: "update_at",
		},
	}
}

func (c ListServiceConnectionsRequestSortKey) Value() string {
	return c.value
}

func (c ListServiceConnectionsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceConnectionsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListServiceConnectionsRequestSortDir struct {
	value string
}

type ListServiceConnectionsRequestSortDirEnum struct {
	ASC  ListServiceConnectionsRequestSortDir
	DESC ListServiceConnectionsRequestSortDir
}

func GetListServiceConnectionsRequestSortDirEnum() ListServiceConnectionsRequestSortDirEnum {
	return ListServiceConnectionsRequestSortDirEnum{
		ASC: ListServiceConnectionsRequestSortDir{
			value: "asc",
		},
		DESC: ListServiceConnectionsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListServiceConnectionsRequestSortDir) Value() string {
	return c.value
}

func (c ListServiceConnectionsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceConnectionsRequestSortDir) UnmarshalJSON(b []byte) error {
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
