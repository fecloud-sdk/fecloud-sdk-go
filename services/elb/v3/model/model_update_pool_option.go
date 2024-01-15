package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdatePoolOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Description *string `json:"description,omitempty"`

	LbAlgorithm *string `json:"lb_algorithm,omitempty"`

	Name *string `json:"name,omitempty"`

	SessionPersistence *UpdatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	SlowStart *UpdatePoolSlowStartOption `json:"slow_start,omitempty"`

	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	Type *string `json:"type,omitempty"`

	ProtectionStatus *UpdatePoolOptionProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o UpdatePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolOption struct{}"
	}

	return strings.Join([]string{"UpdatePoolOption", string(data)}, " ")
}

type UpdatePoolOptionProtectionStatus struct {
	value string
}

type UpdatePoolOptionProtectionStatusEnum struct {
	NON_PROTECTION     UpdatePoolOptionProtectionStatus
	CONSOLE_PROTECTION UpdatePoolOptionProtectionStatus
}

func GetUpdatePoolOptionProtectionStatusEnum() UpdatePoolOptionProtectionStatusEnum {
	return UpdatePoolOptionProtectionStatusEnum{
		NON_PROTECTION: UpdatePoolOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdatePoolOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdatePoolOptionProtectionStatus) Value() string {
	return c.value
}

func (c UpdatePoolOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePoolOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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
