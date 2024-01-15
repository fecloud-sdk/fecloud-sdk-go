package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type L7policyResp struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Rules []ResourceList `json:"rules"`

	Action L7policyRespAction `json:"action"`

	ProvisioningStatus string `json:"provisioning_status"`

	TenantId string `json:"tenant_id"`

	ProjectId string `json:"project_id"`

	AdminStateUp bool `json:"admin_state_up"`

	Description string `json:"description"`

	ListenerId string `json:"listener_id"`

	RedirectPoolId string `json:"redirect_pool_id"`

	RedirectListenerId string `json:"redirect_listener_id"`

	RedirectUrl string `json:"redirect_url"`

	Position int32 `json:"position"`
}

func (o L7policyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7policyResp struct{}"
	}

	return strings.Join([]string{"L7policyResp", string(data)}, " ")
}

type L7policyRespAction struct {
	value string
}

type L7policyRespActionEnum struct {
	REDIRECT_TO_POOL     L7policyRespAction
	REDIRECT_TO_LISTENER L7policyRespAction
}

func GetL7policyRespActionEnum() L7policyRespActionEnum {
	return L7policyRespActionEnum{
		REDIRECT_TO_POOL: L7policyRespAction{
			value: "REDIRECT_TO_POOL",
		},
		REDIRECT_TO_LISTENER: L7policyRespAction{
			value: "REDIRECT_TO_LISTENER",
		},
	}
}

func (c L7policyRespAction) Value() string {
	return c.value
}

func (c L7policyRespAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7policyRespAction) UnmarshalJSON(b []byte) error {
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
