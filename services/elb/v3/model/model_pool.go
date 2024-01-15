package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type Pool struct {
	AdminStateUp bool `json:"admin_state_up"`

	Description string `json:"description"`

	HealthmonitorId string `json:"healthmonitor_id"`

	Id string `json:"id"`

	LbAlgorithm string `json:"lb_algorithm"`

	Listeners []ListenerRef `json:"listeners"`

	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`

	Members []MemberRef `json:"members"`

	Name string `json:"name"`

	ProjectId string `json:"project_id"`

	Protocol string `json:"protocol"`

	SessionPersistence *SessionPersistence `json:"session_persistence"`

	IpVersion *string `json:"ip_version,omitempty"`

	SlowStart *SlowStart `json:"slow_start"`

	MemberDeletionProtectionEnable bool `json:"member_deletion_protection_enable"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	VpcId string `json:"vpc_id"`

	Type string `json:"type"`

	ProtectionStatus *PoolProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`

	AnyPortEnable *bool `json:"any_port_enable,omitempty"`
}

func (o Pool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pool struct{}"
	}

	return strings.Join([]string{"Pool", string(data)}, " ")
}

type PoolProtectionStatus struct {
	value string
}

type PoolProtectionStatusEnum struct {
	NON_PROTECTION     PoolProtectionStatus
	CONSOLE_PROTECTION PoolProtectionStatus
}

func GetPoolProtectionStatusEnum() PoolProtectionStatusEnum {
	return PoolProtectionStatusEnum{
		NON_PROTECTION: PoolProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: PoolProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c PoolProtectionStatus) Value() string {
	return c.value
}

func (c PoolProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoolProtectionStatus) UnmarshalJSON(b []byte) error {
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
