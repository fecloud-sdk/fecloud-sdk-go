package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListPrivateNatsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	Description *[]string `json:"description,omitempty"`

	Spec *[]ListPrivateNatsRequestSpec `json:"spec,omitempty"`

	Status *[]ListPrivateNatsRequestStatus `json:"status,omitempty"`

	VpcId *[]string `json:"vpc_id,omitempty"`

	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListPrivateNatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateNatsRequest", string(data)}, " ")
}

type ListPrivateNatsRequestSpec struct {
	value string
}

type ListPrivateNatsRequestSpecEnum struct {
	SMALL       ListPrivateNatsRequestSpec
	MEDIUM      ListPrivateNatsRequestSpec
	LARGE       ListPrivateNatsRequestSpec
	EXTRA_LARGE ListPrivateNatsRequestSpec
}

func GetListPrivateNatsRequestSpecEnum() ListPrivateNatsRequestSpecEnum {
	return ListPrivateNatsRequestSpecEnum{
		SMALL: ListPrivateNatsRequestSpec{
			value: "Small",
		},
		MEDIUM: ListPrivateNatsRequestSpec{
			value: "Medium",
		},
		LARGE: ListPrivateNatsRequestSpec{
			value: "Large",
		},
		EXTRA_LARGE: ListPrivateNatsRequestSpec{
			value: "Extra-large",
		},
	}
}

func (c ListPrivateNatsRequestSpec) Value() string {
	return c.value
}

func (c ListPrivateNatsRequestSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateNatsRequestSpec) UnmarshalJSON(b []byte) error {
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

type ListPrivateNatsRequestStatus struct {
	value string
}

type ListPrivateNatsRequestStatusEnum struct {
	ACTIVE ListPrivateNatsRequestStatus
	FROZEN ListPrivateNatsRequestStatus
}

func GetListPrivateNatsRequestStatusEnum() ListPrivateNatsRequestStatusEnum {
	return ListPrivateNatsRequestStatusEnum{
		ACTIVE: ListPrivateNatsRequestStatus{
			value: "ACTIVE",
		},
		FROZEN: ListPrivateNatsRequestStatus{
			value: "FROZEN",
		},
	}
}

func (c ListPrivateNatsRequestStatus) Value() string {
	return c.value
}

func (c ListPrivateNatsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateNatsRequestStatus) UnmarshalJSON(b []byte) error {
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
