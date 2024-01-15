package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ResetConsumeOffsetRequest struct {
	Engine ResetConsumeOffsetRequestEngine `json:"engine"`

	InstanceId string `json:"instance_id"`

	GroupId string `json:"group_id"`

	Body *ResetConsumeOffsetReq `json:"body,omitempty"`
}

func (o ResetConsumeOffsetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetRequest struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetRequest", string(data)}, " ")
}

type ResetConsumeOffsetRequestEngine struct {
	value string
}

type ResetConsumeOffsetRequestEngineEnum struct {
	RELIABILITY ResetConsumeOffsetRequestEngine
}

func GetResetConsumeOffsetRequestEngineEnum() ResetConsumeOffsetRequestEngineEnum {
	return ResetConsumeOffsetRequestEngineEnum{
		RELIABILITY: ResetConsumeOffsetRequestEngine{
			value: "reliability",
		},
	}
}

func (c ResetConsumeOffsetRequestEngine) Value() string {
	return c.value
}

func (c ResetConsumeOffsetRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResetConsumeOffsetRequestEngine) UnmarshalJSON(b []byte) error {
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
