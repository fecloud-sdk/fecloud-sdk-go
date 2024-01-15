package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ScaleScript struct {
	Name string `json:"name"`

	Uri string `json:"uri"`

	Parameters *string `json:"parameters,omitempty"`

	Nodes []string `json:"nodes"`

	ActiveMaster *bool `json:"active_master,omitempty"`

	FailAction ScaleScriptFailAction `json:"fail_action"`

	ActionStage ScaleScriptActionStage `json:"action_stage"`
}

func (o ScaleScript) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleScript struct{}"
	}

	return strings.Join([]string{"ScaleScript", string(data)}, " ")
}

type ScaleScriptFailAction struct {
	value string
}

type ScaleScriptFailActionEnum struct {
	CONTINUE ScaleScriptFailAction
	ERROROUT ScaleScriptFailAction
}

func GetScaleScriptFailActionEnum() ScaleScriptFailActionEnum {
	return ScaleScriptFailActionEnum{
		CONTINUE: ScaleScriptFailAction{
			value: "continue",
		},
		ERROROUT: ScaleScriptFailAction{
			value: "errorout",
		},
	}
}

func (c ScaleScriptFailAction) Value() string {
	return c.value
}

func (c ScaleScriptFailAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScaleScriptFailAction) UnmarshalJSON(b []byte) error {
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

type ScaleScriptActionStage struct {
	value string
}

type ScaleScriptActionStageEnum struct {
	BEFORE_SCALE_OUT ScaleScriptActionStage
	BEFORE_SCALE_IN  ScaleScriptActionStage
	AFTER_SCALE_OUT  ScaleScriptActionStage
	AFTER_SCALE_IN   ScaleScriptActionStage
}

func GetScaleScriptActionStageEnum() ScaleScriptActionStageEnum {
	return ScaleScriptActionStageEnum{
		BEFORE_SCALE_OUT: ScaleScriptActionStage{
			value: "before_scale_out",
		},
		BEFORE_SCALE_IN: ScaleScriptActionStage{
			value: "before_scale_in",
		},
		AFTER_SCALE_OUT: ScaleScriptActionStage{
			value: "after_scale_out",
		},
		AFTER_SCALE_IN: ScaleScriptActionStage{
			value: "after_scale_in",
		},
	}
}

func (c ScaleScriptActionStage) Value() string {
	return c.value
}

func (c ScaleScriptActionStage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScaleScriptActionStage) UnmarshalJSON(b []byte) error {
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
