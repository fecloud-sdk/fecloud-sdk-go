package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateLoadbalancerReq struct {
	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	VipSubnetId string `json:"vip_subnet_id"`

	VipAddress *string `json:"vip_address,omitempty"`

	Provider *CreateLoadbalancerReqProvider `json:"provider,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ProtectionStatus *CreateLoadbalancerReqProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`

	ChargeMode *string `json:"charge_mode,omitempty"`

	PrepaidOptions *PrepaidCreateOption `json:"prepaid_options,omitempty"`
}

func (o CreateLoadbalancerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerReq struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerReq", string(data)}, " ")
}

type CreateLoadbalancerReqProvider struct {
	value string
}

type CreateLoadbalancerReqProviderEnum struct {
	VLB CreateLoadbalancerReqProvider
}

func GetCreateLoadbalancerReqProviderEnum() CreateLoadbalancerReqProviderEnum {
	return CreateLoadbalancerReqProviderEnum{
		VLB: CreateLoadbalancerReqProvider{
			value: "vlb",
		},
	}
}

func (c CreateLoadbalancerReqProvider) Value() string {
	return c.value
}

func (c CreateLoadbalancerReqProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadbalancerReqProvider) UnmarshalJSON(b []byte) error {
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

type CreateLoadbalancerReqProtectionStatus struct {
	value string
}

type CreateLoadbalancerReqProtectionStatusEnum struct {
	NON_PROTECTION     CreateLoadbalancerReqProtectionStatus
	CONSOLE_PROTECTION CreateLoadbalancerReqProtectionStatus
}

func GetCreateLoadbalancerReqProtectionStatusEnum() CreateLoadbalancerReqProtectionStatusEnum {
	return CreateLoadbalancerReqProtectionStatusEnum{
		NON_PROTECTION: CreateLoadbalancerReqProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CreateLoadbalancerReqProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CreateLoadbalancerReqProtectionStatus) Value() string {
	return c.value
}

func (c CreateLoadbalancerReqProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadbalancerReqProtectionStatus) UnmarshalJSON(b []byte) error {
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
