package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type PoolResp struct {
	Id string `json:"id"`

	ProjectId string `json:"project_id"`

	TenantId string `json:"tenant_id"`

	Name string `json:"name"`

	Description string `json:"description"`

	AdminStateUp bool `json:"admin_state_up"`

	Loadbalancers []ResourceList `json:"loadbalancers"`

	Listeners []ResourceList `json:"listeners"`

	Members []ResourceList `json:"members"`

	HealthmonitorId string `json:"healthmonitor_id"`

	SessionPersistence *SessionPersistence `json:"session_persistence"`

	Protocol PoolRespProtocol `json:"protocol"`

	LbAlgorithm PoolRespLbAlgorithm `json:"lb_algorithm"`

	ProtectionStatus PoolRespProtectionStatus `json:"protection_status"`

	ProtectionReason string `json:"protection_reason"`
}

func (o PoolResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolResp struct{}"
	}

	return strings.Join([]string{"PoolResp", string(data)}, " ")
}

type PoolRespProtocol struct {
	value string
}

type PoolRespProtocolEnum struct {
	UDP  PoolRespProtocol
	TCP  PoolRespProtocol
	HTTP PoolRespProtocol
}

func GetPoolRespProtocolEnum() PoolRespProtocolEnum {
	return PoolRespProtocolEnum{
		UDP: PoolRespProtocol{
			value: "UDP",
		},
		TCP: PoolRespProtocol{
			value: "TCP",
		},
		HTTP: PoolRespProtocol{
			value: "HTTP",
		},
	}
}

func (c PoolRespProtocol) Value() string {
	return c.value
}

func (c PoolRespProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoolRespProtocol) UnmarshalJSON(b []byte) error {
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

type PoolRespLbAlgorithm struct {
	value string
}

type PoolRespLbAlgorithmEnum struct {
	ROUND_ROBIN       PoolRespLbAlgorithm
	LEAST_CONNECTIONS PoolRespLbAlgorithm
	SOURCE_IP         PoolRespLbAlgorithm
}

func GetPoolRespLbAlgorithmEnum() PoolRespLbAlgorithmEnum {
	return PoolRespLbAlgorithmEnum{
		ROUND_ROBIN: PoolRespLbAlgorithm{
			value: "ROUND_ROBIN",
		},
		LEAST_CONNECTIONS: PoolRespLbAlgorithm{
			value: "LEAST_CONNECTIONS",
		},
		SOURCE_IP: PoolRespLbAlgorithm{
			value: "SOURCE_IP",
		},
	}
}

func (c PoolRespLbAlgorithm) Value() string {
	return c.value
}

func (c PoolRespLbAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoolRespLbAlgorithm) UnmarshalJSON(b []byte) error {
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

type PoolRespProtectionStatus struct {
	value string
}

type PoolRespProtectionStatusEnum struct {
	NON_PROTECTION     PoolRespProtectionStatus
	CONSOLE_PROTECTION PoolRespProtectionStatus
}

func GetPoolRespProtectionStatusEnum() PoolRespProtectionStatusEnum {
	return PoolRespProtectionStatusEnum{
		NON_PROTECTION: PoolRespProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: PoolRespProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c PoolRespProtectionStatus) Value() string {
	return c.value
}

func (c PoolRespProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoolRespProtectionStatus) UnmarshalJSON(b []byte) error {
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
