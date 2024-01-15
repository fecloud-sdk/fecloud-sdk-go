package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type VersionObject struct {
	Status *VersionObjectStatus `json:"status,omitempty"`

	Id *VersionObjectId `json:"id,omitempty"`

	Updated *string `json:"updated,omitempty"`

	Version *string `json:"version,omitempty"`

	MinVersion *string `json:"min_version,omitempty"`

	Links *[]Link `json:"links,omitempty"`
}

func (o VersionObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionObject struct{}"
	}

	return strings.Join([]string{"VersionObject", string(data)}, " ")
}

type VersionObjectStatus struct {
	value string
}

type VersionObjectStatusEnum struct {
	CURRENT    VersionObjectStatus
	SUPPORT    VersionObjectStatus
	DEPRECATED VersionObjectStatus
}

func GetVersionObjectStatusEnum() VersionObjectStatusEnum {
	return VersionObjectStatusEnum{
		CURRENT: VersionObjectStatus{
			value: "CURRENT",
		},
		SUPPORT: VersionObjectStatus{
			value: "SUPPORT",
		},
		DEPRECATED: VersionObjectStatus{
			value: "DEPRECATED",
		},
	}
}

func (c VersionObjectStatus) Value() string {
	return c.value
}

func (c VersionObjectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionObjectStatus) UnmarshalJSON(b []byte) error {
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

type VersionObjectId struct {
	value string
}

type VersionObjectIdEnum struct {
	V1 VersionObjectId
}

func GetVersionObjectIdEnum() VersionObjectIdEnum {
	return VersionObjectIdEnum{
		V1: VersionObjectId{
			value: "v1",
		},
	}
}

func (c VersionObjectId) Value() string {
	return c.value
}

func (c VersionObjectId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionObjectId) UnmarshalJSON(b []byte) error {
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
