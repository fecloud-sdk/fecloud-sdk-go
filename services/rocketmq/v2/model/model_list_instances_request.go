package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListInstancesRequest struct {
	Engine *ListInstancesRequestEngine `json:"engine,omitempty"`

	Name *string `json:"name,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Status *ListInstancesRequestStatus `json:"status,omitempty"`

	IncludeFailure *ListInstancesRequestIncludeFailure `json:"include_failure,omitempty"`

	ExactMatchName *ListInstancesRequestExactMatchName `json:"exact_match_name,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}

type ListInstancesRequestEngine struct {
	value string
}

type ListInstancesRequestEngineEnum struct {
	RELIABILITY ListInstancesRequestEngine
}

func GetListInstancesRequestEngineEnum() ListInstancesRequestEngineEnum {
	return ListInstancesRequestEngineEnum{
		RELIABILITY: ListInstancesRequestEngine{
			value: "reliability",
		},
	}
}

func (c ListInstancesRequestEngine) Value() string {
	return c.value
}

func (c ListInstancesRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestEngine) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestStatus struct {
	value string
}

type ListInstancesRequestStatusEnum struct {
	CREATING        ListInstancesRequestStatus
	RUNNING         ListInstancesRequestStatus
	FAULTY          ListInstancesRequestStatus
	RESTARTING      ListInstancesRequestStatus
	RESIZING        ListInstancesRequestStatus
	RESIZING_FAILED ListInstancesRequestStatus
	FROZEN          ListInstancesRequestStatus
}

func GetListInstancesRequestStatusEnum() ListInstancesRequestStatusEnum {
	return ListInstancesRequestStatusEnum{
		CREATING: ListInstancesRequestStatus{
			value: "CREATING",
		},
		RUNNING: ListInstancesRequestStatus{
			value: "RUNNING",
		},
		FAULTY: ListInstancesRequestStatus{
			value: "FAULTY",
		},
		RESTARTING: ListInstancesRequestStatus{
			value: "RESTARTING",
		},
		RESIZING: ListInstancesRequestStatus{
			value: "RESIZING",
		},
		RESIZING_FAILED: ListInstancesRequestStatus{
			value: "RESIZING FAILED",
		},
		FROZEN: ListInstancesRequestStatus{
			value: "FROZEN",
		},
	}
}

func (c ListInstancesRequestStatus) Value() string {
	return c.value
}

func (c ListInstancesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestIncludeFailure struct {
	value string
}

type ListInstancesRequestIncludeFailureEnum struct {
	TRUE  ListInstancesRequestIncludeFailure
	FALSE ListInstancesRequestIncludeFailure
}

func GetListInstancesRequestIncludeFailureEnum() ListInstancesRequestIncludeFailureEnum {
	return ListInstancesRequestIncludeFailureEnum{
		TRUE: ListInstancesRequestIncludeFailure{
			value: "true",
		},
		FALSE: ListInstancesRequestIncludeFailure{
			value: "false",
		},
	}
}

func (c ListInstancesRequestIncludeFailure) Value() string {
	return c.value
}

func (c ListInstancesRequestIncludeFailure) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestIncludeFailure) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestExactMatchName struct {
	value string
}

type ListInstancesRequestExactMatchNameEnum struct {
	TRUE  ListInstancesRequestExactMatchName
	FALSE ListInstancesRequestExactMatchName
}

func GetListInstancesRequestExactMatchNameEnum() ListInstancesRequestExactMatchNameEnum {
	return ListInstancesRequestExactMatchNameEnum{
		TRUE: ListInstancesRequestExactMatchName{
			value: "true",
		},
		FALSE: ListInstancesRequestExactMatchName{
			value: "false",
		},
	}
}

func (c ListInstancesRequestExactMatchName) Value() string {
	return c.value
}

func (c ListInstancesRequestExactMatchName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestExactMatchName) UnmarshalJSON(b []byte) error {
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
