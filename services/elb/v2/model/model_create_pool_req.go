package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreatePoolReq struct {
	Protocol CreatePoolReqProtocol `json:"protocol"`

	LbAlgorithm string `json:"lb_algorithm"`

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	ListenerId *string `json:"listener_id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	SessionPersistence *SessionPersistence `json:"session_persistence,omitempty"`

	ProtectionStatus *CreatePoolReqProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o CreatePoolReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolReq struct{}"
	}

	return strings.Join([]string{"CreatePoolReq", string(data)}, " ")
}

type CreatePoolReqProtocol struct {
	value string
}

type CreatePoolReqProtocolEnum struct {
	UDP  CreatePoolReqProtocol
	TCP  CreatePoolReqProtocol
	HTTP CreatePoolReqProtocol
}

func GetCreatePoolReqProtocolEnum() CreatePoolReqProtocolEnum {
	return CreatePoolReqProtocolEnum{
		UDP: CreatePoolReqProtocol{
			value: "UDP",
		},
		TCP: CreatePoolReqProtocol{
			value: "TCP",
		},
		HTTP: CreatePoolReqProtocol{
			value: "HTTP",
		},
	}
}

func (c CreatePoolReqProtocol) Value() string {
	return c.value
}

func (c CreatePoolReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePoolReqProtocol) UnmarshalJSON(b []byte) error {
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

type CreatePoolReqProtectionStatus struct {
	value string
}

type CreatePoolReqProtectionStatusEnum struct {
	NON_PROTECTION     CreatePoolReqProtectionStatus
	CONSOLE_PROTECTION CreatePoolReqProtectionStatus
}

func GetCreatePoolReqProtectionStatusEnum() CreatePoolReqProtectionStatusEnum {
	return CreatePoolReqProtectionStatusEnum{
		NON_PROTECTION: CreatePoolReqProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CreatePoolReqProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CreatePoolReqProtectionStatus) Value() string {
	return c.value
}

func (c CreatePoolReqProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePoolReqProtectionStatus) UnmarshalJSON(b []byte) error {
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
