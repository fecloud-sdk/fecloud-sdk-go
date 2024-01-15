package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdatePoolReq struct {
	LbAlgorithm *string `json:"lb_algorithm,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	SessionPersistence *SessionPersistence `json:"session_persistence,omitempty"`

	ProtectionStatus *UpdatePoolReqProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o UpdatePoolReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolReq struct{}"
	}

	return strings.Join([]string{"UpdatePoolReq", string(data)}, " ")
}

type UpdatePoolReqProtectionStatus struct {
	value string
}

type UpdatePoolReqProtectionStatusEnum struct {
	NON_PROTECTION     UpdatePoolReqProtectionStatus
	CONSOLE_PROTECTION UpdatePoolReqProtectionStatus
}

func GetUpdatePoolReqProtectionStatusEnum() UpdatePoolReqProtectionStatusEnum {
	return UpdatePoolReqProtectionStatusEnum{
		NON_PROTECTION: UpdatePoolReqProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdatePoolReqProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdatePoolReqProtectionStatus) Value() string {
	return c.value
}

func (c UpdatePoolReqProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePoolReqProtectionStatus) UnmarshalJSON(b []byte) error {
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
