package model

import (
	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/sdktime"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"
	"strings"
)

type DirectConnect struct {
	Id *string `json:"id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	PortType *DirectConnectPortType `json:"port_type,omitempty"`

	Bandwidth *int32 `json:"bandwidth,omitempty"`

	Location *string `json:"location,omitempty"`

	PeerLocation *string `json:"peer_location,omitempty"`

	DeviceId *string `json:"device_id,omitempty"`

	Type *DirectConnectType `json:"type,omitempty"`

	HostingId *string `json:"hosting_id,omitempty"`

	ChargeMode *DirectConnectChargeMode `json:"charge_mode,omitempty"`

	Provider *string `json:"provider,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Vlan *int32 `json:"vlan,omitempty"`

	Status *DirectConnectStatus `json:"status,omitempty"`

	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	ProviderStatus *DirectConnectProviderStatus `json:"provider_status,omitempty"`

	PeerPortType *string `json:"peer_port_type,omitempty"`

	PeerProvider *string `json:"peer_provider,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	ProductId *string `json:"product_id,omitempty"`

	SpecCode *string `json:"spec_code,omitempty"`

	PeriodType *int32 `json:"period_type,omitempty"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	VgwType *DirectConnectVgwType `json:"vgw_type,omitempty"`

	LagId *string `json:"lag_id,omitempty"`

	SignedAgreementStatus *DirectConnectSignedAgreementStatus `json:"signed_agreement_status,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o DirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectConnect struct{}"
	}

	return strings.Join([]string{"DirectConnect", string(data)}, " ")
}

type DirectConnectPortType struct {
	value string
}

type DirectConnectPortTypeEnum struct {
	E_1_G   DirectConnectPortType
	E_10_G  DirectConnectPortType
	E_40_G  DirectConnectPortType
	E_100_G DirectConnectPortType
}

func GetDirectConnectPortTypeEnum() DirectConnectPortTypeEnum {
	return DirectConnectPortTypeEnum{
		E_1_G: DirectConnectPortType{
			value: "1G",
		},
		E_10_G: DirectConnectPortType{
			value: "10G",
		},
		E_40_G: DirectConnectPortType{
			value: "40G",
		},
		E_100_G: DirectConnectPortType{
			value: "100G",
		},
	}
}

func (c DirectConnectPortType) Value() string {
	return c.value
}

func (c DirectConnectPortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectPortType) UnmarshalJSON(b []byte) error {
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

type DirectConnectType struct {
	value string
}

type DirectConnectTypeEnum struct {
	STANDARD DirectConnectType
	HOSTING  DirectConnectType
	HOSTED   DirectConnectType
}

func GetDirectConnectTypeEnum() DirectConnectTypeEnum {
	return DirectConnectTypeEnum{
		STANDARD: DirectConnectType{
			value: "standard",
		},
		HOSTING: DirectConnectType{
			value: "hosting",
		},
		HOSTED: DirectConnectType{
			value: "hosted",
		},
	}
}

func (c DirectConnectType) Value() string {
	return c.value
}

func (c DirectConnectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectType) UnmarshalJSON(b []byte) error {
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

type DirectConnectChargeMode struct {
	value string
}

type DirectConnectChargeModeEnum struct {
	PREPAYMENT DirectConnectChargeMode
	BANDWIDTH  DirectConnectChargeMode
	TRAFFIC    DirectConnectChargeMode
}

func GetDirectConnectChargeModeEnum() DirectConnectChargeModeEnum {
	return DirectConnectChargeModeEnum{
		PREPAYMENT: DirectConnectChargeMode{
			value: "prepayment",
		},
		BANDWIDTH: DirectConnectChargeMode{
			value: "bandwidth",
		},
		TRAFFIC: DirectConnectChargeMode{
			value: "traffic",
		},
	}
}

func (c DirectConnectChargeMode) Value() string {
	return c.value
}

func (c DirectConnectChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectChargeMode) UnmarshalJSON(b []byte) error {
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

type DirectConnectStatus struct {
	value string
}

type DirectConnectStatusEnum struct {
	BUILD          DirectConnectStatus
	PAID           DirectConnectStatus
	APPLY          DirectConnectStatus
	PENDING_SURVEY DirectConnectStatus
	ACTIVE         DirectConnectStatus
	DOWN           DirectConnectStatus
	ERROR          DirectConnectStatus
	PENDING_DELETE DirectConnectStatus
	DELETED        DirectConnectStatus
	DENY           DirectConnectStatus
	PENDING_PAY    DirectConnectStatus
	ORDERING       DirectConnectStatus
	ACCEPT         DirectConnectStatus
	REJECTED       DirectConnectStatus
}

func GetDirectConnectStatusEnum() DirectConnectStatusEnum {
	return DirectConnectStatusEnum{
		BUILD: DirectConnectStatus{
			value: "BUILD",
		},
		PAID: DirectConnectStatus{
			value: "PAID",
		},
		APPLY: DirectConnectStatus{
			value: "APPLY",
		},
		PENDING_SURVEY: DirectConnectStatus{
			value: "PENDING_SURVEY",
		},
		ACTIVE: DirectConnectStatus{
			value: "ACTIVE",
		},
		DOWN: DirectConnectStatus{
			value: "DOWN",
		},
		ERROR: DirectConnectStatus{
			value: "ERROR",
		},
		PENDING_DELETE: DirectConnectStatus{
			value: "PENDING_DELETE",
		},
		DELETED: DirectConnectStatus{
			value: "DELETED",
		},
		DENY: DirectConnectStatus{
			value: "DENY",
		},
		PENDING_PAY: DirectConnectStatus{
			value: "PENDING_PAY",
		},
		ORDERING: DirectConnectStatus{
			value: "ORDERING",
		},
		ACCEPT: DirectConnectStatus{
			value: "ACCEPT",
		},
		REJECTED: DirectConnectStatus{
			value: "REJECTED",
		},
	}
}

func (c DirectConnectStatus) Value() string {
	return c.value
}

func (c DirectConnectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectStatus) UnmarshalJSON(b []byte) error {
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

type DirectConnectProviderStatus struct {
	value string
}

type DirectConnectProviderStatusEnum struct {
	ACTIVE DirectConnectProviderStatus
	DOWN   DirectConnectProviderStatus
}

func GetDirectConnectProviderStatusEnum() DirectConnectProviderStatusEnum {
	return DirectConnectProviderStatusEnum{
		ACTIVE: DirectConnectProviderStatus{
			value: "ACTIVE",
		},
		DOWN: DirectConnectProviderStatus{
			value: "DOWN",
		},
	}
}

func (c DirectConnectProviderStatus) Value() string {
	return c.value
}

func (c DirectConnectProviderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectProviderStatus) UnmarshalJSON(b []byte) error {
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

type DirectConnectVgwType struct {
	value string
}

type DirectConnectVgwTypeEnum struct {
	DEFAULT DirectConnectVgwType
}

func GetDirectConnectVgwTypeEnum() DirectConnectVgwTypeEnum {
	return DirectConnectVgwTypeEnum{
		DEFAULT: DirectConnectVgwType{
			value: "default",
		},
	}
}

func (c DirectConnectVgwType) Value() string {
	return c.value
}

func (c DirectConnectVgwType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectVgwType) UnmarshalJSON(b []byte) error {
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

type DirectConnectSignedAgreementStatus struct {
	value string
}

type DirectConnectSignedAgreementStatusEnum struct {
	SIGNED DirectConnectSignedAgreementStatus
}

func GetDirectConnectSignedAgreementStatusEnum() DirectConnectSignedAgreementStatusEnum {
	return DirectConnectSignedAgreementStatusEnum{
		SIGNED: DirectConnectSignedAgreementStatus{
			value: "signed",
		},
	}
}

func (c DirectConnectSignedAgreementStatus) Value() string {
	return c.value
}

func (c DirectConnectSignedAgreementStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectConnectSignedAgreementStatus) UnmarshalJSON(b []byte) error {
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
