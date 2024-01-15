package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateUserResponse struct {
	AccessKey *string `json:"access_key,omitempty"`

	SecretKey *string `json:"secret_key,omitempty"`

	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	Admin *bool `json:"admin,omitempty"`

	DefaultTopicPerm *UpdateUserResponseDefaultTopicPerm `json:"default_topic_perm,omitempty"`

	DefaultGroupPerm *UpdateUserResponseDefaultGroupPerm `json:"default_group_perm,omitempty"`

	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty"`

	GroupPerms     *[]UserGroupPerms `json:"group_perms,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserResponse", string(data)}, " ")
}

type UpdateUserResponseDefaultTopicPerm struct {
	value string
}

type UpdateUserResponseDefaultTopicPermEnum struct {
	PUB     UpdateUserResponseDefaultTopicPerm
	SUB     UpdateUserResponseDefaultTopicPerm
	PUB_SUB UpdateUserResponseDefaultTopicPerm
	DENY    UpdateUserResponseDefaultTopicPerm
}

func GetUpdateUserResponseDefaultTopicPermEnum() UpdateUserResponseDefaultTopicPermEnum {
	return UpdateUserResponseDefaultTopicPermEnum{
		PUB: UpdateUserResponseDefaultTopicPerm{
			value: "PUB",
		},
		SUB: UpdateUserResponseDefaultTopicPerm{
			value: "SUB",
		},
		PUB_SUB: UpdateUserResponseDefaultTopicPerm{
			value: "PUB|SUB",
		},
		DENY: UpdateUserResponseDefaultTopicPerm{
			value: "DENY",
		},
	}
}

func (c UpdateUserResponseDefaultTopicPerm) Value() string {
	return c.value
}

func (c UpdateUserResponseDefaultTopicPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateUserResponseDefaultTopicPerm) UnmarshalJSON(b []byte) error {
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

type UpdateUserResponseDefaultGroupPerm struct {
	value string
}

type UpdateUserResponseDefaultGroupPermEnum struct {
	SUB  UpdateUserResponseDefaultGroupPerm
	DENY UpdateUserResponseDefaultGroupPerm
}

func GetUpdateUserResponseDefaultGroupPermEnum() UpdateUserResponseDefaultGroupPermEnum {
	return UpdateUserResponseDefaultGroupPermEnum{
		SUB: UpdateUserResponseDefaultGroupPerm{
			value: "SUB",
		},
		DENY: UpdateUserResponseDefaultGroupPerm{
			value: "DENY",
		},
	}
}

func (c UpdateUserResponseDefaultGroupPerm) Value() string {
	return c.value
}

func (c UpdateUserResponseDefaultGroupPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateUserResponseDefaultGroupPerm) UnmarshalJSON(b []byte) error {
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
