package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type AddOrRemoveServicePermissionsResponse struct {
	Permissions *[]string `json:"permissions,omitempty"`

	PermissionType *AddOrRemoveServicePermissionsResponsePermissionType `json:"permission_type,omitempty"`
	HttpStatusCode int                                                  `json:"-"`
}

func (o AddOrRemoveServicePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrRemoveServicePermissionsResponse struct{}"
	}

	return strings.Join([]string{"AddOrRemoveServicePermissionsResponse", string(data)}, " ")
}

type AddOrRemoveServicePermissionsResponsePermissionType struct {
	value string
}

type AddOrRemoveServicePermissionsResponsePermissionTypeEnum struct {
	DOMAIN_ID AddOrRemoveServicePermissionsResponsePermissionType
	ORG_PATH  AddOrRemoveServicePermissionsResponsePermissionType
}

func GetAddOrRemoveServicePermissionsResponsePermissionTypeEnum() AddOrRemoveServicePermissionsResponsePermissionTypeEnum {
	return AddOrRemoveServicePermissionsResponsePermissionTypeEnum{
		DOMAIN_ID: AddOrRemoveServicePermissionsResponsePermissionType{
			value: "domainId",
		},
		ORG_PATH: AddOrRemoveServicePermissionsResponsePermissionType{
			value: "orgPath",
		},
	}
}

func (c AddOrRemoveServicePermissionsResponsePermissionType) Value() string {
	return c.value
}

func (c AddOrRemoveServicePermissionsResponsePermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddOrRemoveServicePermissionsResponsePermissionType) UnmarshalJSON(b []byte) error {
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
