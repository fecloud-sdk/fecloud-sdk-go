package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateDirectConnect struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Bandwidth *int32 `json:"bandwidth,omitempty"`

	PeerLocation *string `json:"peer_location,omitempty"`

	Status *UpdateDirectConnectStatus `json:"status,omitempty"`

	ProviderStatus *UpdateDirectConnectProviderStatus `json:"provider_status,omitempty"`
}

func (o UpdateDirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDirectConnect struct{}"
	}

	return strings.Join([]string{"UpdateDirectConnect", string(data)}, " ")
}

type UpdateDirectConnectStatus struct {
	value string
}

type UpdateDirectConnectStatusEnum struct {
	PENDING_PAY UpdateDirectConnectStatus
}

func GetUpdateDirectConnectStatusEnum() UpdateDirectConnectStatusEnum {
	return UpdateDirectConnectStatusEnum{
		PENDING_PAY: UpdateDirectConnectStatus{
			value: "PENDING_PAY",
		},
	}
}

func (c UpdateDirectConnectStatus) Value() string {
	return c.value
}

func (c UpdateDirectConnectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDirectConnectStatus) UnmarshalJSON(b []byte) error {
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

type UpdateDirectConnectProviderStatus struct {
	value string
}

type UpdateDirectConnectProviderStatusEnum struct {
	ACTIVE UpdateDirectConnectProviderStatus
	DOWN   UpdateDirectConnectProviderStatus
}

func GetUpdateDirectConnectProviderStatusEnum() UpdateDirectConnectProviderStatusEnum {
	return UpdateDirectConnectProviderStatusEnum{
		ACTIVE: UpdateDirectConnectProviderStatus{
			value: "ACTIVE",
		},
		DOWN: UpdateDirectConnectProviderStatus{
			value: "DOWN",
		},
	}
}

func (c UpdateDirectConnectProviderStatus) Value() string {
	return c.value
}

func (c UpdateDirectConnectProviderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDirectConnectProviderStatus) UnmarshalJSON(b []byte) error {
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
