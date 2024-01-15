package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type MembersInStatusResp struct {
	Id string `json:"id"`

	Address string `json:"address"`

	ProtocolPort int32 `json:"protocol_port"`

	OperatingStatus MembersInStatusRespOperatingStatus `json:"operating_status"`

	ProvisioningStatus string `json:"provisioning_status"`
}

func (o MembersInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MembersInStatusResp struct{}"
	}

	return strings.Join([]string{"MembersInStatusResp", string(data)}, " ")
}

type MembersInStatusRespOperatingStatus struct {
	value string
}

type MembersInStatusRespOperatingStatusEnum struct {
	ONLINE     MembersInStatusRespOperatingStatus
	OFFLINE    MembersInStatusRespOperatingStatus
	DISABLED   MembersInStatusRespOperatingStatus
	NO_MONITOR MembersInStatusRespOperatingStatus
}

func GetMembersInStatusRespOperatingStatusEnum() MembersInStatusRespOperatingStatusEnum {
	return MembersInStatusRespOperatingStatusEnum{
		ONLINE: MembersInStatusRespOperatingStatus{
			value: "ONLINE",
		},
		OFFLINE: MembersInStatusRespOperatingStatus{
			value: "OFFLINE",
		},
		DISABLED: MembersInStatusRespOperatingStatus{
			value: "DISABLED",
		},
		NO_MONITOR: MembersInStatusRespOperatingStatus{
			value: "NO_MONITOR",
		},
	}
}

func (c MembersInStatusRespOperatingStatus) Value() string {
	return c.value
}

func (c MembersInStatusRespOperatingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MembersInStatusRespOperatingStatus) UnmarshalJSON(b []byte) error {
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
