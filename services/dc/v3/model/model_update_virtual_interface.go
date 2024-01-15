package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateVirtualInterface struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Bandwidth *int32 `json:"bandwidth,omitempty"`

	RemoteEpGroup *[]string `json:"remote_ep_group,omitempty"`

	ServiceEpGroup *[]string `json:"service_ep_group,omitempty"`

	EnableBfd *bool `json:"enable_bfd,omitempty"`

	EnableNqa *bool `json:"enable_nqa,omitempty"`

	Status *UpdateVirtualInterfaceStatus `json:"status,omitempty"`
}

func (o UpdateVirtualInterface) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualInterface struct{}"
	}

	return strings.Join([]string{"UpdateVirtualInterface", string(data)}, " ")
}

type UpdateVirtualInterfaceStatus struct {
	value string
}

type UpdateVirtualInterfaceStatusEnum struct {
	ACCEPTED UpdateVirtualInterfaceStatus
	REJECTED UpdateVirtualInterfaceStatus
}

func GetUpdateVirtualInterfaceStatusEnum() UpdateVirtualInterfaceStatusEnum {
	return UpdateVirtualInterfaceStatusEnum{
		ACCEPTED: UpdateVirtualInterfaceStatus{
			value: "ACCEPTED",
		},
		REJECTED: UpdateVirtualInterfaceStatus{
			value: "REJECTED",
		},
	}
}

func (c UpdateVirtualInterfaceStatus) Value() string {
	return c.value
}

func (c UpdateVirtualInterfaceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVirtualInterfaceStatus) UnmarshalJSON(b []byte) error {
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
