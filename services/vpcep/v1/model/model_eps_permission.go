package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type EpsPermission struct {
	Id *string `json:"id,omitempty"`

	Permission *string `json:"permission,omitempty"`

	PermissionType *EpsPermissionPermissionType `json:"permission_type,omitempty"`

	Description *string `json:"description,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`
}

func (o EpsPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsPermission struct{}"
	}

	return strings.Join([]string{"EpsPermission", string(data)}, " ")
}

type EpsPermissionPermissionType struct {
	value string
}

type EpsPermissionPermissionTypeEnum struct {
	DOMAIN_ID EpsPermissionPermissionType
	ORG_PATH  EpsPermissionPermissionType
}

func GetEpsPermissionPermissionTypeEnum() EpsPermissionPermissionTypeEnum {
	return EpsPermissionPermissionTypeEnum{
		DOMAIN_ID: EpsPermissionPermissionType{
			value: "domainId",
		},
		ORG_PATH: EpsPermissionPermissionType{
			value: "orgPath",
		},
	}
}

func (c EpsPermissionPermissionType) Value() string {
	return c.value
}

func (c EpsPermissionPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EpsPermissionPermissionType) UnmarshalJSON(b []byte) error {
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
