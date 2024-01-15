package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListEndpointServiceRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	Id *string `json:"id,omitempty"`

	Status *ListEndpointServiceRequestStatus `json:"status,omitempty"`

	SortKey *ListEndpointServiceRequestSortKey `json:"sort_key,omitempty"`

	SortDir *ListEndpointServiceRequestSortDir `json:"sort_dir,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o ListEndpointServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointServiceRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointServiceRequest", string(data)}, " ")
}

type ListEndpointServiceRequestStatus struct {
	value string
}

type ListEndpointServiceRequestStatusEnum struct {
	CREATING  ListEndpointServiceRequestStatus
	AVAILABLE ListEndpointServiceRequestStatus
	FAILED    ListEndpointServiceRequestStatus
	DELETING  ListEndpointServiceRequestStatus
}

func GetListEndpointServiceRequestStatusEnum() ListEndpointServiceRequestStatusEnum {
	return ListEndpointServiceRequestStatusEnum{
		CREATING: ListEndpointServiceRequestStatus{
			value: "creating",
		},
		AVAILABLE: ListEndpointServiceRequestStatus{
			value: "available",
		},
		FAILED: ListEndpointServiceRequestStatus{
			value: "failed",
		},
		DELETING: ListEndpointServiceRequestStatus{
			value: "deleting",
		},
	}
}

func (c ListEndpointServiceRequestStatus) Value() string {
	return c.value
}

func (c ListEndpointServiceRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointServiceRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListEndpointServiceRequestSortKey struct {
	value string
}

type ListEndpointServiceRequestSortKeyEnum struct {
	CREATE_AT ListEndpointServiceRequestSortKey
	UPDATE_AT ListEndpointServiceRequestSortKey
}

func GetListEndpointServiceRequestSortKeyEnum() ListEndpointServiceRequestSortKeyEnum {
	return ListEndpointServiceRequestSortKeyEnum{
		CREATE_AT: ListEndpointServiceRequestSortKey{
			value: "create_at",
		},
		UPDATE_AT: ListEndpointServiceRequestSortKey{
			value: "update_at",
		},
	}
}

func (c ListEndpointServiceRequestSortKey) Value() string {
	return c.value
}

func (c ListEndpointServiceRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointServiceRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListEndpointServiceRequestSortDir struct {
	value string
}

type ListEndpointServiceRequestSortDirEnum struct {
	ASC  ListEndpointServiceRequestSortDir
	DESC ListEndpointServiceRequestSortDir
}

func GetListEndpointServiceRequestSortDirEnum() ListEndpointServiceRequestSortDirEnum {
	return ListEndpointServiceRequestSortDirEnum{
		ASC: ListEndpointServiceRequestSortDir{
			value: "asc",
		},
		DESC: ListEndpointServiceRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListEndpointServiceRequestSortDir) Value() string {
	return c.value
}

func (c ListEndpointServiceRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointServiceRequestSortDir) UnmarshalJSON(b []byte) error {
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
