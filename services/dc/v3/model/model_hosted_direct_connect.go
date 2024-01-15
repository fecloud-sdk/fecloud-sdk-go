package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type HostedDirectConnect struct {
	Id *string `json:"id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Bandwidth *int32 `json:"bandwidth,omitempty"`

	Location *string `json:"location,omitempty"`

	PeerLocation *string `json:"peer_location,omitempty"`

	HostingId *string `json:"hosting_id,omitempty"`

	Provider *string `json:"provider,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Vlan *int32 `json:"vlan,omitempty"`

	Status *HostedDirectConnectStatus `json:"status,omitempty"`

	ApplyTime *string `json:"apply_time,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	ProviderStatus *HostedDirectConnectProviderStatus `json:"provider_status,omitempty"`
}

func (o HostedDirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostedDirectConnect struct{}"
	}

	return strings.Join([]string{"HostedDirectConnect", string(data)}, " ")
}

type HostedDirectConnectStatus struct {
	value string
}

type HostedDirectConnectStatusEnum struct {
	BUILD          HostedDirectConnectStatus
	PAID           HostedDirectConnectStatus
	APPLY          HostedDirectConnectStatus
	PENDING_SURVEY HostedDirectConnectStatus
	ACTIVE         HostedDirectConnectStatus
	DOWN           HostedDirectConnectStatus
	ERROR          HostedDirectConnectStatus
	PENDING_DELETE HostedDirectConnectStatus
	DELETED        HostedDirectConnectStatus
	DENY           HostedDirectConnectStatus
	PENDING_PAY    HostedDirectConnectStatus
	ORDERING       HostedDirectConnectStatus
	ACCEPT         HostedDirectConnectStatus
	REJECTED       HostedDirectConnectStatus
}

func GetHostedDirectConnectStatusEnum() HostedDirectConnectStatusEnum {
	return HostedDirectConnectStatusEnum{
		BUILD: HostedDirectConnectStatus{
			value: "BUILD",
		},
		PAID: HostedDirectConnectStatus{
			value: "PAID",
		},
		APPLY: HostedDirectConnectStatus{
			value: "APPLY",
		},
		PENDING_SURVEY: HostedDirectConnectStatus{
			value: "PENDING_SURVEY",
		},
		ACTIVE: HostedDirectConnectStatus{
			value: "ACTIVE",
		},
		DOWN: HostedDirectConnectStatus{
			value: "DOWN",
		},
		ERROR: HostedDirectConnectStatus{
			value: "ERROR",
		},
		PENDING_DELETE: HostedDirectConnectStatus{
			value: "PENDING_DELETE",
		},
		DELETED: HostedDirectConnectStatus{
			value: "DELETED",
		},
		DENY: HostedDirectConnectStatus{
			value: "DENY",
		},
		PENDING_PAY: HostedDirectConnectStatus{
			value: "PENDING_PAY",
		},
		ORDERING: HostedDirectConnectStatus{
			value: "ORDERING",
		},
		ACCEPT: HostedDirectConnectStatus{
			value: "ACCEPT",
		},
		REJECTED: HostedDirectConnectStatus{
			value: "REJECTED",
		},
	}
}

func (c HostedDirectConnectStatus) Value() string {
	return c.value
}

func (c HostedDirectConnectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostedDirectConnectStatus) UnmarshalJSON(b []byte) error {
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

type HostedDirectConnectProviderStatus struct {
	value string
}

type HostedDirectConnectProviderStatusEnum struct {
	ACTIVE HostedDirectConnectProviderStatus
	DOWN   HostedDirectConnectProviderStatus
}

func GetHostedDirectConnectProviderStatusEnum() HostedDirectConnectProviderStatusEnum {
	return HostedDirectConnectProviderStatusEnum{
		ACTIVE: HostedDirectConnectProviderStatus{
			value: "ACTIVE",
		},
		DOWN: HostedDirectConnectProviderStatus{
			value: "DOWN",
		},
	}
}

func (c HostedDirectConnectProviderStatus) Value() string {
	return c.value
}

func (c HostedDirectConnectProviderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostedDirectConnectProviderStatus) UnmarshalJSON(b []byte) error {
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
