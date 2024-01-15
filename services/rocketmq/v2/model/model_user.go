package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type User struct {
	AccessKey *string `json:"access_key,omitempty"`

	SecretKey *string `json:"secret_key,omitempty"`

	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	Admin *bool `json:"admin,omitempty"`

	DefaultTopicPerm *UserDefaultTopicPerm `json:"default_topic_perm,omitempty"`

	DefaultGroupPerm *UserDefaultGroupPerm `json:"default_group_perm,omitempty"`

	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty"`

	GroupPerms *[]UserGroupPerms `json:"group_perms,omitempty"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}

type UserDefaultTopicPerm struct {
	value string
}

type UserDefaultTopicPermEnum struct {
	PUB     UserDefaultTopicPerm
	SUB     UserDefaultTopicPerm
	PUB_SUB UserDefaultTopicPerm
	DENY    UserDefaultTopicPerm
}

func GetUserDefaultTopicPermEnum() UserDefaultTopicPermEnum {
	return UserDefaultTopicPermEnum{
		PUB: UserDefaultTopicPerm{
			value: "PUB",
		},
		SUB: UserDefaultTopicPerm{
			value: "SUB",
		},
		PUB_SUB: UserDefaultTopicPerm{
			value: "PUB|SUB",
		},
		DENY: UserDefaultTopicPerm{
			value: "DENY",
		},
	}
}

func (c UserDefaultTopicPerm) Value() string {
	return c.value
}

func (c UserDefaultTopicPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserDefaultTopicPerm) UnmarshalJSON(b []byte) error {
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

type UserDefaultGroupPerm struct {
	value string
}

type UserDefaultGroupPermEnum struct {
	SUB  UserDefaultGroupPerm
	DENY UserDefaultGroupPerm
}

func GetUserDefaultGroupPermEnum() UserDefaultGroupPermEnum {
	return UserDefaultGroupPermEnum{
		SUB: UserDefaultGroupPerm{
			value: "SUB",
		},
		DENY: UserDefaultGroupPerm{
			value: "DENY",
		},
	}
}

func (c UserDefaultGroupPerm) Value() string {
	return c.value
}

func (c UserDefaultGroupPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserDefaultGroupPerm) UnmarshalJSON(b []byte) error {
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
