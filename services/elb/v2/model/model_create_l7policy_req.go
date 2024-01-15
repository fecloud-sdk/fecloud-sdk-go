package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateL7policyReq struct {
	Name *string `json:"name,omitempty"`

	Action CreateL7policyReqAction `json:"action"`

	TenantId *string `json:"tenant_id,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Description *string `json:"description,omitempty"`

	ListenerId string `json:"listener_id"`

	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`

	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`

	RedirectUrl *string `json:"redirect_url,omitempty"`

	Position *int32 `json:"position,omitempty"`

	Rules *[]CreateL7ruleReqInPolicy `json:"rules,omitempty"`
}

func (o CreateL7policyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7policyReq struct{}"
	}

	return strings.Join([]string{"CreateL7policyReq", string(data)}, " ")
}

type CreateL7policyReqAction struct {
	value string
}

type CreateL7policyReqActionEnum struct {
	REDIRECT_TO_POOL     CreateL7policyReqAction
	REDIRECT_TO_LISTENER CreateL7policyReqAction
}

func GetCreateL7policyReqActionEnum() CreateL7policyReqActionEnum {
	return CreateL7policyReqActionEnum{
		REDIRECT_TO_POOL: CreateL7policyReqAction{
			value: "REDIRECT_TO_POOL",
		},
		REDIRECT_TO_LISTENER: CreateL7policyReqAction{
			value: "REDIRECT_TO_LISTENER",
		},
	}
}

func (c CreateL7policyReqAction) Value() string {
	return c.value
}

func (c CreateL7policyReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateL7policyReqAction) UnmarshalJSON(b []byte) error {
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
