package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type PermissionObject struct {
	Id *string `json:"id,omitempty"`

	Permission *string `json:"permission,omitempty"`

	PermissionType *PermissionObjectPermissionType `json:"permission_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`
}

func (o PermissionObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionObject struct{}"
	}

	return strings.Join([]string{"PermissionObject", string(data)}, " ")
}

type PermissionObjectPermissionType struct {
	value string
}

type PermissionObjectPermissionTypeEnum struct {
	DOMAIN_ID PermissionObjectPermissionType
	ORG_PATH  PermissionObjectPermissionType
}

func GetPermissionObjectPermissionTypeEnum() PermissionObjectPermissionTypeEnum {
	return PermissionObjectPermissionTypeEnum{
		DOMAIN_ID: PermissionObjectPermissionType{
			value: "domainId",
		},
		ORG_PATH: PermissionObjectPermissionType{
			value: "orgPath",
		},
	}
}

func (c PermissionObjectPermissionType) Value() string {
	return c.value
}

func (c PermissionObjectPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionObjectPermissionType) UnmarshalJSON(b []byte) error {
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
