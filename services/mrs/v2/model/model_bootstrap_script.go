package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BootstrapScript struct {
	Name string `json:"name"`

	Uri string `json:"uri"`

	Parameters *string `json:"parameters,omitempty"`

	Nodes []string `json:"nodes"`

	ActiveMaster *bool `json:"active_master,omitempty"`

	FailAction BootstrapScriptFailAction `json:"fail_action"`

	BeforeComponentStart *bool `json:"before_component_start,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	State *BootstrapScriptState `json:"state,omitempty"`

	ActionStages *[]BootstrapScriptActionStages `json:"action_stages,omitempty"`
}

func (o BootstrapScript) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BootstrapScript struct{}"
	}

	return strings.Join([]string{"BootstrapScript", string(data)}, " ")
}

type BootstrapScriptFailAction struct {
	value string
}

type BootstrapScriptFailActionEnum struct {
	CONTINUE BootstrapScriptFailAction
	ERROROUT BootstrapScriptFailAction
}

func GetBootstrapScriptFailActionEnum() BootstrapScriptFailActionEnum {
	return BootstrapScriptFailActionEnum{
		CONTINUE: BootstrapScriptFailAction{
			value: "continue",
		},
		ERROROUT: BootstrapScriptFailAction{
			value: "errorout",
		},
	}
}

func (c BootstrapScriptFailAction) Value() string {
	return c.value
}

func (c BootstrapScriptFailAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BootstrapScriptFailAction) UnmarshalJSON(b []byte) error {
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

type BootstrapScriptState struct {
	value string
}

type BootstrapScriptStateEnum struct {
	PENDING     BootstrapScriptState
	IN_PROGRESS BootstrapScriptState
	SUCCESS     BootstrapScriptState
	FAILURE     BootstrapScriptState
}

func GetBootstrapScriptStateEnum() BootstrapScriptStateEnum {
	return BootstrapScriptStateEnum{
		PENDING: BootstrapScriptState{
			value: "PENDING",
		},
		IN_PROGRESS: BootstrapScriptState{
			value: "IN_PROGRESS",
		},
		SUCCESS: BootstrapScriptState{
			value: "SUCCESS",
		},
		FAILURE: BootstrapScriptState{
			value: "FAILURE",
		},
	}
}

func (c BootstrapScriptState) Value() string {
	return c.value
}

func (c BootstrapScriptState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BootstrapScriptState) UnmarshalJSON(b []byte) error {
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

type BootstrapScriptActionStages struct {
	value string
}

type BootstrapScriptActionStagesEnum struct {
	BEFORE_COMPONENT_FIRST_START BootstrapScriptActionStages
	AFTER_COMPONENT_FIRST_START  BootstrapScriptActionStages
	BEFORE_SCALE_IN              BootstrapScriptActionStages
	AFTER_SCALE_IN               BootstrapScriptActionStages
	BEFORE_SCALE_OUT             BootstrapScriptActionStages
	AFTER_SCALE_OUT              BootstrapScriptActionStages
}

func GetBootstrapScriptActionStagesEnum() BootstrapScriptActionStagesEnum {
	return BootstrapScriptActionStagesEnum{
		BEFORE_COMPONENT_FIRST_START: BootstrapScriptActionStages{
			value: "BEFORE_COMPONENT_FIRST_START",
		},
		AFTER_COMPONENT_FIRST_START: BootstrapScriptActionStages{
			value: "AFTER_COMPONENT_FIRST_START",
		},
		BEFORE_SCALE_IN: BootstrapScriptActionStages{
			value: "BEFORE_SCALE_IN",
		},
		AFTER_SCALE_IN: BootstrapScriptActionStages{
			value: "AFTER_SCALE_IN",
		},
		BEFORE_SCALE_OUT: BootstrapScriptActionStages{
			value: "BEFORE_SCALE_OUT",
		},
		AFTER_SCALE_OUT: BootstrapScriptActionStages{
			value: "AFTER_SCALE_OUT",
		},
	}
}

func (c BootstrapScriptActionStages) Value() string {
	return c.value
}

func (c BootstrapScriptActionStages) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BootstrapScriptActionStages) UnmarshalJSON(b []byte) error {
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
