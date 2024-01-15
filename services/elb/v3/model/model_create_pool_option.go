package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreatePoolOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Description *string `json:"description,omitempty"`

	LbAlgorithm string `json:"lb_algorithm"`

	ListenerId *string `json:"listener_id,omitempty"`

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	Name *string `json:"name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Protocol string `json:"protocol"`

	SessionPersistence *CreatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	SlowStart *CreatePoolSlowStartOption `json:"slow_start,omitempty"`

	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	Type *string `json:"type,omitempty"`

	ProtectionStatus *CreatePoolOptionProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`

	AnyPortEnable *bool `json:"any_port_enable,omitempty"`
}

func (o CreatePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolOption struct{}"
	}

	return strings.Join([]string{"CreatePoolOption", string(data)}, " ")
}

type CreatePoolOptionProtectionStatus struct {
	value string
}

type CreatePoolOptionProtectionStatusEnum struct {
	NON_PROTECTION     CreatePoolOptionProtectionStatus
	CONSOLE_PROTECTION CreatePoolOptionProtectionStatus
}

func GetCreatePoolOptionProtectionStatusEnum() CreatePoolOptionProtectionStatusEnum {
	return CreatePoolOptionProtectionStatusEnum{
		NON_PROTECTION: CreatePoolOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CreatePoolOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CreatePoolOptionProtectionStatus) Value() string {
	return c.value
}

func (c CreatePoolOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePoolOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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
