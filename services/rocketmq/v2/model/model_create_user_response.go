package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateUserResponse struct {
	AccessKey *string `json:"access_key,omitempty"`

	SecretKey *string `json:"secret_key,omitempty"`

	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	Admin *bool `json:"admin,omitempty"`

	DefaultTopicPerm *CreateUserResponseDefaultTopicPerm `json:"default_topic_perm,omitempty"`

	DefaultGroupPerm *CreateUserResponseDefaultGroupPerm `json:"default_group_perm,omitempty"`

	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty"`

	GroupPerms     *[]UserGroupPerms `json:"group_perms,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserResponse struct{}"
	}

	return strings.Join([]string{"CreateUserResponse", string(data)}, " ")
}

type CreateUserResponseDefaultTopicPerm struct {
	value string
}

type CreateUserResponseDefaultTopicPermEnum struct {
	PUB     CreateUserResponseDefaultTopicPerm
	SUB     CreateUserResponseDefaultTopicPerm
	PUB_SUB CreateUserResponseDefaultTopicPerm
	DENY    CreateUserResponseDefaultTopicPerm
}

func GetCreateUserResponseDefaultTopicPermEnum() CreateUserResponseDefaultTopicPermEnum {
	return CreateUserResponseDefaultTopicPermEnum{
		PUB: CreateUserResponseDefaultTopicPerm{
			value: "PUB",
		},
		SUB: CreateUserResponseDefaultTopicPerm{
			value: "SUB",
		},
		PUB_SUB: CreateUserResponseDefaultTopicPerm{
			value: "PUB|SUB",
		},
		DENY: CreateUserResponseDefaultTopicPerm{
			value: "DENY",
		},
	}
}

func (c CreateUserResponseDefaultTopicPerm) Value() string {
	return c.value
}

func (c CreateUserResponseDefaultTopicPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateUserResponseDefaultTopicPerm) UnmarshalJSON(b []byte) error {
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

type CreateUserResponseDefaultGroupPerm struct {
	value string
}

type CreateUserResponseDefaultGroupPermEnum struct {
	SUB  CreateUserResponseDefaultGroupPerm
	DENY CreateUserResponseDefaultGroupPerm
}

func GetCreateUserResponseDefaultGroupPermEnum() CreateUserResponseDefaultGroupPermEnum {
	return CreateUserResponseDefaultGroupPermEnum{
		SUB: CreateUserResponseDefaultGroupPerm{
			value: "SUB",
		},
		DENY: CreateUserResponseDefaultGroupPerm{
			value: "DENY",
		},
	}
}

func (c CreateUserResponseDefaultGroupPerm) Value() string {
	return c.value
}

func (c CreateUserResponseDefaultGroupPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateUserResponseDefaultGroupPerm) UnmarshalJSON(b []byte) error {
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
