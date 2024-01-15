package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UserTopicPerms struct {
	Name *string `json:"name,omitempty"`

	Perm *UserTopicPermsPerm `json:"perm,omitempty"`
}

func (o UserTopicPerms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserTopicPerms struct{}"
	}

	return strings.Join([]string{"UserTopicPerms", string(data)}, " ")
}

type UserTopicPermsPerm struct {
	value string
}

type UserTopicPermsPermEnum struct {
	PUB     UserTopicPermsPerm
	SUB     UserTopicPermsPerm
	PUB_SUB UserTopicPermsPerm
	DENY    UserTopicPermsPerm
}

func GetUserTopicPermsPermEnum() UserTopicPermsPermEnum {
	return UserTopicPermsPermEnum{
		PUB: UserTopicPermsPerm{
			value: "PUB",
		},
		SUB: UserTopicPermsPerm{
			value: "SUB",
		},
		PUB_SUB: UserTopicPermsPerm{
			value: "PUB|SUB",
		},
		DENY: UserTopicPermsPerm{
			value: "DENY",
		},
	}
}

func (c UserTopicPermsPerm) Value() string {
	return c.value
}

func (c UserTopicPermsPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserTopicPermsPerm) UnmarshalJSON(b []byte) error {
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
