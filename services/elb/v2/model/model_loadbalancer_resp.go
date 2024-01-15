package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type LoadbalancerResp struct {
	Id string `json:"id"`

	TenantId string `json:"tenant_id"`

	Name string `json:"name"`

	Description string `json:"description"`

	VipSubnetId string `json:"vip_subnet_id"`

	VipPortId string `json:"vip_port_id"`

	VipAddress string `json:"vip_address"`

	Listeners []ResourceList `json:"listeners"`

	Pools []ResourceList `json:"pools"`

	Provider string `json:"provider"`

	OperatingStatus LoadbalancerRespOperatingStatus `json:"operating_status"`

	ProvisioningStatus LoadbalancerRespProvisioningStatus `json:"provisioning_status"`

	AdminStateUp bool `json:"admin_state_up"`

	CreatedAt string `json:"created_at"`

	UpdatedAt string `json:"updated_at"`

	EnterpriseProjectId string `json:"enterprise_project_id"`

	ProjectId string `json:"project_id"`

	Tags []string `json:"tags"`

	Publicips []PublicIpInfo `json:"publicips"`

	ChargeMode string `json:"charge_mode"`

	FrozenScene *string `json:"frozen_scene,omitempty"`
}

func (o LoadbalancerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadbalancerResp struct{}"
	}

	return strings.Join([]string{"LoadbalancerResp", string(data)}, " ")
}

type LoadbalancerRespOperatingStatus struct {
	value string
}

type LoadbalancerRespOperatingStatusEnum struct {
	ONLINE     LoadbalancerRespOperatingStatus
	OFFLINE    LoadbalancerRespOperatingStatus
	DEGRADED   LoadbalancerRespOperatingStatus
	DISABLED   LoadbalancerRespOperatingStatus
	NO_MONITOR LoadbalancerRespOperatingStatus
}

func GetLoadbalancerRespOperatingStatusEnum() LoadbalancerRespOperatingStatusEnum {
	return LoadbalancerRespOperatingStatusEnum{
		ONLINE: LoadbalancerRespOperatingStatus{
			value: "ONLINE",
		},
		OFFLINE: LoadbalancerRespOperatingStatus{
			value: "OFFLINE",
		},
		DEGRADED: LoadbalancerRespOperatingStatus{
			value: "DEGRADED",
		},
		DISABLED: LoadbalancerRespOperatingStatus{
			value: "DISABLED",
		},
		NO_MONITOR: LoadbalancerRespOperatingStatus{
			value: "NO_MONITOR",
		},
	}
}

func (c LoadbalancerRespOperatingStatus) Value() string {
	return c.value
}

func (c LoadbalancerRespOperatingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerRespOperatingStatus) UnmarshalJSON(b []byte) error {
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

type LoadbalancerRespProvisioningStatus struct {
	value string
}

type LoadbalancerRespProvisioningStatusEnum struct {
	ACTIVE         LoadbalancerRespProvisioningStatus
	PENDING_CREATE LoadbalancerRespProvisioningStatus
	ERROR          LoadbalancerRespProvisioningStatus
}

func GetLoadbalancerRespProvisioningStatusEnum() LoadbalancerRespProvisioningStatusEnum {
	return LoadbalancerRespProvisioningStatusEnum{
		ACTIVE: LoadbalancerRespProvisioningStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: LoadbalancerRespProvisioningStatus{
			value: "PENDING_CREATE",
		},
		ERROR: LoadbalancerRespProvisioningStatus{
			value: "ERROR",
		},
	}
}

func (c LoadbalancerRespProvisioningStatus) Value() string {
	return c.value
}

func (c LoadbalancerRespProvisioningStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerRespProvisioningStatus) UnmarshalJSON(b []byte) error {
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
