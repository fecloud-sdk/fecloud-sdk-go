package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreatePostPaidInstanceReq struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Engine CreatePostPaidInstanceReqEngine `json:"engine"`

	EngineVersion CreatePostPaidInstanceReqEngineVersion `json:"engine_version"`

	StorageSpace int32 `json:"storage_space"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId string `json:"security_group_id"`

	AvailableZones []string `json:"available_zones"`

	ProductId CreatePostPaidInstanceReqProductId `json:"product_id"`

	SslEnable *bool `json:"ssl_enable,omitempty"`

	StorageSpecCode CreatePostPaidInstanceReqStorageSpecCode `json:"storage_spec_code"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	EnableAcl *bool `json:"enable_acl,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	PublicipId *string `json:"publicip_id,omitempty"`

	BrokerNum int32 `json:"broker_num"`
}

func (o CreatePostPaidInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceReq struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceReq", string(data)}, " ")
}

type CreatePostPaidInstanceReqEngine struct {
	value string
}

type CreatePostPaidInstanceReqEngineEnum struct {
	RELIABILITY CreatePostPaidInstanceReqEngine
}

func GetCreatePostPaidInstanceReqEngineEnum() CreatePostPaidInstanceReqEngineEnum {
	return CreatePostPaidInstanceReqEngineEnum{
		RELIABILITY: CreatePostPaidInstanceReqEngine{
			value: "reliability",
		},
	}
}

func (c CreatePostPaidInstanceReqEngine) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqEngine) UnmarshalJSON(b []byte) error {
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

type CreatePostPaidInstanceReqEngineVersion struct {
	value string
}

type CreatePostPaidInstanceReqEngineVersionEnum struct {
	E_4_8_0 CreatePostPaidInstanceReqEngineVersion
}

func GetCreatePostPaidInstanceReqEngineVersionEnum() CreatePostPaidInstanceReqEngineVersionEnum {
	return CreatePostPaidInstanceReqEngineVersionEnum{
		E_4_8_0: CreatePostPaidInstanceReqEngineVersion{
			value: "4.8.0",
		},
	}
}

func (c CreatePostPaidInstanceReqEngineVersion) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqEngineVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqEngineVersion) UnmarshalJSON(b []byte) error {
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

type CreatePostPaidInstanceReqProductId struct {
	value string
}

type CreatePostPaidInstanceReqProductIdEnum struct {
	C6_4U8G_CLUSTER   CreatePostPaidInstanceReqProductId
	C6_8U16G_CLUSTER  CreatePostPaidInstanceReqProductId
	C6_12U24G_CLUSTER CreatePostPaidInstanceReqProductId
	C6_16U32G_CLUSTER CreatePostPaidInstanceReqProductId
}

func GetCreatePostPaidInstanceReqProductIdEnum() CreatePostPaidInstanceReqProductIdEnum {
	return CreatePostPaidInstanceReqProductIdEnum{
		C6_4U8G_CLUSTER: CreatePostPaidInstanceReqProductId{
			value: "c6.4u8g.cluster",
		},
		C6_8U16G_CLUSTER: CreatePostPaidInstanceReqProductId{
			value: "c6.8u16g.cluster",
		},
		C6_12U24G_CLUSTER: CreatePostPaidInstanceReqProductId{
			value: "c6.12u24g.cluster",
		},
		C6_16U32G_CLUSTER: CreatePostPaidInstanceReqProductId{
			value: "c6.16u32g.cluster",
		},
	}
}

func (c CreatePostPaidInstanceReqProductId) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqProductId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqProductId) UnmarshalJSON(b []byte) error {
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

type CreatePostPaidInstanceReqStorageSpecCode struct {
	value string
}

type CreatePostPaidInstanceReqStorageSpecCodeEnum struct {
	DMS_PHYSICAL_STORAGE_HIGH_V2  CreatePostPaidInstanceReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_ULTRA_V2 CreatePostPaidInstanceReqStorageSpecCode
}

func GetCreatePostPaidInstanceReqStorageSpecCodeEnum() CreatePostPaidInstanceReqStorageSpecCodeEnum {
	return CreatePostPaidInstanceReqStorageSpecCodeEnum{
		DMS_PHYSICAL_STORAGE_HIGH_V2: CreatePostPaidInstanceReqStorageSpecCode{
			value: "dms.physical.storage.high.v2",
		},
		DMS_PHYSICAL_STORAGE_ULTRA_V2: CreatePostPaidInstanceReqStorageSpecCode{
			value: "dms.physical.storage.ultra.v2",
		},
	}
}

func (c CreatePostPaidInstanceReqStorageSpecCode) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqStorageSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqStorageSpecCode) UnmarshalJSON(b []byte) error {
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
