package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateTopicReq struct {
	ReadQueueNum float32 `json:"read_queue_num,omitempty"`

	WriteQueueNum float32 `json:"write_queue_num,omitempty"`

	Permission *UpdateTopicReqPermission `json:"permission,omitempty"`
}

func (o UpdateTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicReq struct{}"
	}

	return strings.Join([]string{"UpdateTopicReq", string(data)}, " ")
}

type UpdateTopicReqPermission struct {
	value string
}

type UpdateTopicReqPermissionEnum struct {
	SUB UpdateTopicReqPermission
	PUB UpdateTopicReqPermission
	ALL UpdateTopicReqPermission
}

func GetUpdateTopicReqPermissionEnum() UpdateTopicReqPermissionEnum {
	return UpdateTopicReqPermissionEnum{
		SUB: UpdateTopicReqPermission{
			value: "sub",
		},
		PUB: UpdateTopicReqPermission{
			value: "pub",
		},
		ALL: UpdateTopicReqPermission{
			value: "all",
		},
	}
}

func (c UpdateTopicReqPermission) Value() string {
	return c.value
}

func (c UpdateTopicReqPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTopicReqPermission) UnmarshalJSON(b []byte) error {
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
