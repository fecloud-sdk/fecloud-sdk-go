package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateLoadbalancerReq struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ProtectionStatus *UpdateLoadbalancerReqProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o UpdateLoadbalancerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerReq struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerReq", string(data)}, " ")
}

type UpdateLoadbalancerReqProtectionStatus struct {
	value string
}

type UpdateLoadbalancerReqProtectionStatusEnum struct {
	NON_PROTECTION     UpdateLoadbalancerReqProtectionStatus
	CONSOLE_PROTECTION UpdateLoadbalancerReqProtectionStatus
}

func GetUpdateLoadbalancerReqProtectionStatusEnum() UpdateLoadbalancerReqProtectionStatusEnum {
	return UpdateLoadbalancerReqProtectionStatusEnum{
		NON_PROTECTION: UpdateLoadbalancerReqProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdateLoadbalancerReqProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdateLoadbalancerReqProtectionStatus) Value() string {
	return c.value
}

func (c UpdateLoadbalancerReqProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLoadbalancerReqProtectionStatus) UnmarshalJSON(b []byte) error {
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
