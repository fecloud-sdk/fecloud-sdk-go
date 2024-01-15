package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateInstanceByEngineReq struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Engine CreateInstanceByEngineReqEngine `json:"engine"`

	EngineVersion CreateInstanceByEngineReqEngineVersion `json:"engine_version"`

	StorageSpace int32 `json:"storage_space"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId string `json:"security_group_id"`

	AvailableZones []string `json:"available_zones"`

	ProductId CreateInstanceByEngineReqProductId `json:"product_id"`

	SslEnable *bool `json:"ssl_enable,omitempty"`

	StorageSpecCode CreateInstanceByEngineReqStorageSpecCode `json:"storage_spec_code"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	EnableAcl *bool `json:"enable_acl,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	PublicipId *string `json:"publicip_id,omitempty"`

	BrokerNum int32 `json:"broker_num"`

	BssParam *BssParam `json:"bss_param,omitempty"`
}

func (o CreateInstanceByEngineReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceByEngineReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceByEngineReq", string(data)}, " ")
}

type CreateInstanceByEngineReqEngine struct {
	value string
}

type CreateInstanceByEngineReqEngineEnum struct {
	RELIABILITY CreateInstanceByEngineReqEngine
}

func GetCreateInstanceByEngineReqEngineEnum() CreateInstanceByEngineReqEngineEnum {
	return CreateInstanceByEngineReqEngineEnum{
		RELIABILITY: CreateInstanceByEngineReqEngine{
			value: "reliability",
		},
	}
}

func (c CreateInstanceByEngineReqEngine) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqEngine) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqEngineVersion struct {
	value string
}

type CreateInstanceByEngineReqEngineVersionEnum struct {
	E_4_8_0 CreateInstanceByEngineReqEngineVersion
}

func GetCreateInstanceByEngineReqEngineVersionEnum() CreateInstanceByEngineReqEngineVersionEnum {
	return CreateInstanceByEngineReqEngineVersionEnum{
		E_4_8_0: CreateInstanceByEngineReqEngineVersion{
			value: "4.8.0",
		},
	}
}

func (c CreateInstanceByEngineReqEngineVersion) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqEngineVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqEngineVersion) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqProductId struct {
	value string
}

type CreateInstanceByEngineReqProductIdEnum struct {
	C6_4U8G_CLUSTER_SMALL CreateInstanceByEngineReqProductId
	C6_4U8G_CLUSTER       CreateInstanceByEngineReqProductId
	C6_8U16G_CLUSTER      CreateInstanceByEngineReqProductId
	C6_12U24G_CLUSTER     CreateInstanceByEngineReqProductId
	C6_16U32G_CLUSTER     CreateInstanceByEngineReqProductId
}

func GetCreateInstanceByEngineReqProductIdEnum() CreateInstanceByEngineReqProductIdEnum {
	return CreateInstanceByEngineReqProductIdEnum{
		C6_4U8G_CLUSTER_SMALL: CreateInstanceByEngineReqProductId{
			value: "c6.4u8g.cluster.small",
		},
		C6_4U8G_CLUSTER: CreateInstanceByEngineReqProductId{
			value: "c6.4u8g.cluster",
		},
		C6_8U16G_CLUSTER: CreateInstanceByEngineReqProductId{
			value: "c6.8u16g.cluster",
		},
		C6_12U24G_CLUSTER: CreateInstanceByEngineReqProductId{
			value: "c6.12u24g.cluster",
		},
		C6_16U32G_CLUSTER: CreateInstanceByEngineReqProductId{
			value: "c6.16u32g.cluster",
		},
	}
}

func (c CreateInstanceByEngineReqProductId) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqProductId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqProductId) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqStorageSpecCode struct {
	value string
}

type CreateInstanceByEngineReqStorageSpecCodeEnum struct {
	DMS_PHYSICAL_STORAGE_HIGH_V2  CreateInstanceByEngineReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_ULTRA_V2 CreateInstanceByEngineReqStorageSpecCode
}

func GetCreateInstanceByEngineReqStorageSpecCodeEnum() CreateInstanceByEngineReqStorageSpecCodeEnum {
	return CreateInstanceByEngineReqStorageSpecCodeEnum{
		DMS_PHYSICAL_STORAGE_HIGH_V2: CreateInstanceByEngineReqStorageSpecCode{
			value: "dms.physical.storage.high.v2",
		},
		DMS_PHYSICAL_STORAGE_ULTRA_V2: CreateInstanceByEngineReqStorageSpecCode{
			value: "dms.physical.storage.ultra.v2",
		},
	}
}

func (c CreateInstanceByEngineReqStorageSpecCode) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqStorageSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqStorageSpecCode) UnmarshalJSON(b []byte) error {
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
