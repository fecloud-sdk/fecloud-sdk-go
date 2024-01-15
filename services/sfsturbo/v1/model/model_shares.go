package model

import (
	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/sdktime"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"
	"strings"
)

type Shares struct {
	ActionProgress *ActionProgress `json:"action_progress,omitempty"`

	Version *string `json:"version,omitempty"`

	AvailCapacity *string `json:"avail_capacity,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	AzName *string `json:"az_name,omitempty"`

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	CryptKeyId *string `json:"crypt_key_id,omitempty"`

	ExpandType *string `json:"expand_type,omitempty"`

	ExportLocation *string `json:"export_location,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	PayModel *SharesPayModel `json:"pay_model,omitempty"`

	Region *string `json:"region,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	ShareProto *string `json:"share_proto,omitempty"`

	ShareType *string `json:"share_type,omitempty"`

	Size *string `json:"size,omitempty"`

	Status *string `json:"status,omitempty"`

	SubStatus *string `json:"sub_status,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`
}

func (o Shares) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Shares struct{}"
	}

	return strings.Join([]string{"Shares", string(data)}, " ")
}

type SharesPayModel struct {
	value string
}

type SharesPayModelEnum struct {
	E_0 SharesPayModel
	E_1 SharesPayModel
}

func GetSharesPayModelEnum() SharesPayModelEnum {
	return SharesPayModelEnum{
		E_0: SharesPayModel{
			value: "0",
		},
		E_1: SharesPayModel{
			value: "1",
		},
	}
}

func (c SharesPayModel) Value() string {
	return c.value
}

func (c SharesPayModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SharesPayModel) UnmarshalJSON(b []byte) error {
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
