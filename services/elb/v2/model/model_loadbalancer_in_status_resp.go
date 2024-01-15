package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type LoadbalancerInStatusResp struct {
	Name string `json:"name"`

	Id string `json:"id"`

	Listeners []ListenersInStatusResp `json:"listeners"`

	Pools []PoolsInStatusResp `json:"pools"`

	OperatingStatus LoadbalancerInStatusRespOperatingStatus `json:"operating_status"`

	ProvisioningStatus LoadbalancerInStatusRespProvisioningStatus `json:"provisioning_status"`
}

func (o LoadbalancerInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadbalancerInStatusResp struct{}"
	}

	return strings.Join([]string{"LoadbalancerInStatusResp", string(data)}, " ")
}

type LoadbalancerInStatusRespOperatingStatus struct {
	value string
}

type LoadbalancerInStatusRespOperatingStatusEnum struct {
	ONLINE     LoadbalancerInStatusRespOperatingStatus
	OFFLINE    LoadbalancerInStatusRespOperatingStatus
	DEGRADED   LoadbalancerInStatusRespOperatingStatus
	DISABLED   LoadbalancerInStatusRespOperatingStatus
	NO_MONITOR LoadbalancerInStatusRespOperatingStatus
}

func GetLoadbalancerInStatusRespOperatingStatusEnum() LoadbalancerInStatusRespOperatingStatusEnum {
	return LoadbalancerInStatusRespOperatingStatusEnum{
		ONLINE: LoadbalancerInStatusRespOperatingStatus{
			value: "ONLINE",
		},
		OFFLINE: LoadbalancerInStatusRespOperatingStatus{
			value: "OFFLINE",
		},
		DEGRADED: LoadbalancerInStatusRespOperatingStatus{
			value: "DEGRADED",
		},
		DISABLED: LoadbalancerInStatusRespOperatingStatus{
			value: "DISABLED",
		},
		NO_MONITOR: LoadbalancerInStatusRespOperatingStatus{
			value: "NO_MONITOR",
		},
	}
}

func (c LoadbalancerInStatusRespOperatingStatus) Value() string {
	return c.value
}

func (c LoadbalancerInStatusRespOperatingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerInStatusRespOperatingStatus) UnmarshalJSON(b []byte) error {
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

type LoadbalancerInStatusRespProvisioningStatus struct {
	value string
}

type LoadbalancerInStatusRespProvisioningStatusEnum struct {
	ACTIVE         LoadbalancerInStatusRespProvisioningStatus
	PENDING_CREATE LoadbalancerInStatusRespProvisioningStatus
	ERROR          LoadbalancerInStatusRespProvisioningStatus
}

func GetLoadbalancerInStatusRespProvisioningStatusEnum() LoadbalancerInStatusRespProvisioningStatusEnum {
	return LoadbalancerInStatusRespProvisioningStatusEnum{
		ACTIVE: LoadbalancerInStatusRespProvisioningStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: LoadbalancerInStatusRespProvisioningStatus{
			value: "PENDING_CREATE",
		},
		ERROR: LoadbalancerInStatusRespProvisioningStatus{
			value: "ERROR",
		},
	}
}

func (c LoadbalancerInStatusRespProvisioningStatus) Value() string {
	return c.value
}

func (c LoadbalancerInStatusRespProvisioningStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerInStatusRespProvisioningStatus) UnmarshalJSON(b []byte) error {
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
