package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type AddOrRemoveServicePermissionsRequestBody struct {
	Permissions []string `json:"permissions"`

	PermissionType *AddOrRemoveServicePermissionsRequestBodyPermissionType `json:"permission_type,omitempty"`

	Action AddOrRemoveServicePermissionsRequestBodyAction `json:"action"`
}

func (o AddOrRemoveServicePermissionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrRemoveServicePermissionsRequestBody struct{}"
	}

	return strings.Join([]string{"AddOrRemoveServicePermissionsRequestBody", string(data)}, " ")
}

type AddOrRemoveServicePermissionsRequestBodyPermissionType struct {
	value string
}

type AddOrRemoveServicePermissionsRequestBodyPermissionTypeEnum struct {
	DOMAIN_ID AddOrRemoveServicePermissionsRequestBodyPermissionType
	ORG_PATH  AddOrRemoveServicePermissionsRequestBodyPermissionType
}

func GetAddOrRemoveServicePermissionsRequestBodyPermissionTypeEnum() AddOrRemoveServicePermissionsRequestBodyPermissionTypeEnum {
	return AddOrRemoveServicePermissionsRequestBodyPermissionTypeEnum{
		DOMAIN_ID: AddOrRemoveServicePermissionsRequestBodyPermissionType{
			value: "domainId",
		},
		ORG_PATH: AddOrRemoveServicePermissionsRequestBodyPermissionType{
			value: "orgPath",
		},
	}
}

func (c AddOrRemoveServicePermissionsRequestBodyPermissionType) Value() string {
	return c.value
}

func (c AddOrRemoveServicePermissionsRequestBodyPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddOrRemoveServicePermissionsRequestBodyPermissionType) UnmarshalJSON(b []byte) error {
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

type AddOrRemoveServicePermissionsRequestBodyAction struct {
	value string
}

type AddOrRemoveServicePermissionsRequestBodyActionEnum struct {
	ADD    AddOrRemoveServicePermissionsRequestBodyAction
	REMOVE AddOrRemoveServicePermissionsRequestBodyAction
}

func GetAddOrRemoveServicePermissionsRequestBodyActionEnum() AddOrRemoveServicePermissionsRequestBodyActionEnum {
	return AddOrRemoveServicePermissionsRequestBodyActionEnum{
		ADD: AddOrRemoveServicePermissionsRequestBodyAction{
			value: "add",
		},
		REMOVE: AddOrRemoveServicePermissionsRequestBodyAction{
			value: "remove",
		},
	}
}

func (c AddOrRemoveServicePermissionsRequestBodyAction) Value() string {
	return c.value
}

func (c AddOrRemoveServicePermissionsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddOrRemoveServicePermissionsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
