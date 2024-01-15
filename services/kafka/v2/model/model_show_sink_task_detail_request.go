package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowSinkTaskDetailRequest struct {
	ConnectorId string `json:"connector_id"`

	TaskId string `json:"task_id"`

	TopicInfo *ShowSinkTaskDetailRequestTopicInfo `json:"topic-info,omitempty"`
}

func (o ShowSinkTaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSinkTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSinkTaskDetailRequest", string(data)}, " ")
}

type ShowSinkTaskDetailRequestTopicInfo struct {
	value string
}

type ShowSinkTaskDetailRequestTopicInfoEnum struct {
	TRUE  ShowSinkTaskDetailRequestTopicInfo
	FALSE ShowSinkTaskDetailRequestTopicInfo
}

func GetShowSinkTaskDetailRequestTopicInfoEnum() ShowSinkTaskDetailRequestTopicInfoEnum {
	return ShowSinkTaskDetailRequestTopicInfoEnum{
		TRUE: ShowSinkTaskDetailRequestTopicInfo{
			value: "true",
		},
		FALSE: ShowSinkTaskDetailRequestTopicInfo{
			value: "false",
		},
	}
}

func (c ShowSinkTaskDetailRequestTopicInfo) Value() string {
	return c.value
}

func (c ShowSinkTaskDetailRequestTopicInfo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSinkTaskDetailRequestTopicInfo) UnmarshalJSON(b []byte) error {
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
