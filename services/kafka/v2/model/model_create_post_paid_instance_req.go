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

	EngineVersion string `json:"engine_version"`

	Specification *CreatePostPaidInstanceReqSpecification `json:"specification,omitempty"`

	BrokerNum *int32 `json:"broker_num,omitempty"`

	StorageSpace int32 `json:"storage_space"`

	PartitionNum *CreatePostPaidInstanceReqPartitionNum `json:"partition_num,omitempty"`

	AccessUser *string `json:"access_user,omitempty"`

	Password *string `json:"password,omitempty"`

	VpcId string `json:"vpc_id"`

	SecurityGroupId string `json:"security_group_id"`

	SubnetId string `json:"subnet_id"`

	AvailableZones []string `json:"available_zones"`

	ProductId string `json:"product_id"`

	KafkaManagerUser *string `json:"kafka_manager_user,omitempty"`

	KafkaManagerPassword *string `json:"kafka_manager_password,omitempty"`

	MaintainBegin *string `json:"maintain_begin,omitempty"`

	MaintainEnd *string `json:"maintain_end,omitempty"`

	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	PublicBandwidth *int32 `json:"public_bandwidth,omitempty"`

	PublicipId *string `json:"publicip_id,omitempty"`

	SslEnable *bool `json:"ssl_enable,omitempty"`

	KafkaSecurityProtocol *string `json:"kafka_security_protocol,omitempty"`

	SaslEnabledMechanisms *[]CreatePostPaidInstanceReqSaslEnabledMechanisms `json:"sasl_enabled_mechanisms,omitempty"`

	RetentionPolicy *CreatePostPaidInstanceReqRetentionPolicy `json:"retention_policy,omitempty"`

	DiskEncryptedEnable *bool `json:"disk_encrypted_enable,omitempty"`

	DiskEncryptedKey *string `json:"disk_encrypted_key,omitempty"`

	ConnectorEnable *bool `json:"connector_enable,omitempty"`

	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty"`

	StorageSpecCode CreatePostPaidInstanceReqStorageSpecCode `json:"storage_spec_code"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]TagEntity `json:"tags,omitempty"`
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
	KAFKA CreatePostPaidInstanceReqEngine
}

func GetCreatePostPaidInstanceReqEngineEnum() CreatePostPaidInstanceReqEngineEnum {
	return CreatePostPaidInstanceReqEngineEnum{
		KAFKA: CreatePostPaidInstanceReqEngine{
			value: "kafka",
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

type CreatePostPaidInstanceReqSpecification struct {
	value string
}

type CreatePostPaidInstanceReqSpecificationEnum struct {
	E_100_MB          CreatePostPaidInstanceReqSpecification
	E_300_MB          CreatePostPaidInstanceReqSpecification
	E_600_MB          CreatePostPaidInstanceReqSpecification
	E_1200_MB         CreatePostPaidInstanceReqSpecification
	C6_2U4G_CLUSTER   CreatePostPaidInstanceReqSpecification
	C6_4U8G_CLUSTER   CreatePostPaidInstanceReqSpecification
	C6_8U16G_CLUSTER  CreatePostPaidInstanceReqSpecification
	C6_12U24G_CLUSTER CreatePostPaidInstanceReqSpecification
	C6_16U32G_CLUSTER CreatePostPaidInstanceReqSpecification
}

func GetCreatePostPaidInstanceReqSpecificationEnum() CreatePostPaidInstanceReqSpecificationEnum {
	return CreatePostPaidInstanceReqSpecificationEnum{
		E_100_MB: CreatePostPaidInstanceReqSpecification{
			value: "100MB",
		},
		E_300_MB: CreatePostPaidInstanceReqSpecification{
			value: "300MB",
		},
		E_600_MB: CreatePostPaidInstanceReqSpecification{
			value: "600MB",
		},
		E_1200_MB: CreatePostPaidInstanceReqSpecification{
			value: "1200MB",
		},
		C6_2U4G_CLUSTER: CreatePostPaidInstanceReqSpecification{
			value: "c6.2u4g.cluster",
		},
		C6_4U8G_CLUSTER: CreatePostPaidInstanceReqSpecification{
			value: "c6.4u8g.cluster",
		},
		C6_8U16G_CLUSTER: CreatePostPaidInstanceReqSpecification{
			value: "c6.8u16g.cluster",
		},
		C6_12U24G_CLUSTER: CreatePostPaidInstanceReqSpecification{
			value: "c6.12u24g.cluster",
		},
		C6_16U32G_CLUSTER: CreatePostPaidInstanceReqSpecification{
			value: "c6.16u32g.cluster",
		},
	}
}

func (c CreatePostPaidInstanceReqSpecification) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqSpecification) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqSpecification) UnmarshalJSON(b []byte) error {
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

type CreatePostPaidInstanceReqPartitionNum struct {
	value int32
}

type CreatePostPaidInstanceReqPartitionNumEnum struct {
	E_250  CreatePostPaidInstanceReqPartitionNum
	E_300  CreatePostPaidInstanceReqPartitionNum
	E_500  CreatePostPaidInstanceReqPartitionNum
	E_900  CreatePostPaidInstanceReqPartitionNum
	E_1000 CreatePostPaidInstanceReqPartitionNum
	E_1500 CreatePostPaidInstanceReqPartitionNum
	E_1800 CreatePostPaidInstanceReqPartitionNum
	E_2000 CreatePostPaidInstanceReqPartitionNum
}

func GetCreatePostPaidInstanceReqPartitionNumEnum() CreatePostPaidInstanceReqPartitionNumEnum {
	return CreatePostPaidInstanceReqPartitionNumEnum{
		E_250: CreatePostPaidInstanceReqPartitionNum{
			value: 250,
		}, E_300: CreatePostPaidInstanceReqPartitionNum{
			value: 300,
		}, E_500: CreatePostPaidInstanceReqPartitionNum{
			value: 500,
		}, E_900: CreatePostPaidInstanceReqPartitionNum{
			value: 900,
		}, E_1000: CreatePostPaidInstanceReqPartitionNum{
			value: 1000,
		}, E_1500: CreatePostPaidInstanceReqPartitionNum{
			value: 1500,
		}, E_1800: CreatePostPaidInstanceReqPartitionNum{
			value: 1800,
		}, E_2000: CreatePostPaidInstanceReqPartitionNum{
			value: 2000,
		},
	}
}

func (c CreatePostPaidInstanceReqPartitionNum) Value() int32 {
	return c.value
}

func (c CreatePostPaidInstanceReqPartitionNum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqPartitionNum) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreatePostPaidInstanceReqSaslEnabledMechanisms struct {
	value string
}

type CreatePostPaidInstanceReqSaslEnabledMechanismsEnum struct {
	PLAIN         CreatePostPaidInstanceReqSaslEnabledMechanisms
	SCRAM_SHA_512 CreatePostPaidInstanceReqSaslEnabledMechanisms
}

func GetCreatePostPaidInstanceReqSaslEnabledMechanismsEnum() CreatePostPaidInstanceReqSaslEnabledMechanismsEnum {
	return CreatePostPaidInstanceReqSaslEnabledMechanismsEnum{
		PLAIN: CreatePostPaidInstanceReqSaslEnabledMechanisms{
			value: "PLAIN",
		},
		SCRAM_SHA_512: CreatePostPaidInstanceReqSaslEnabledMechanisms{
			value: "SCRAM-SHA-512",
		},
	}
}

func (c CreatePostPaidInstanceReqSaslEnabledMechanisms) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqSaslEnabledMechanisms) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqSaslEnabledMechanisms) UnmarshalJSON(b []byte) error {
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

type CreatePostPaidInstanceReqRetentionPolicy struct {
	value string
}

type CreatePostPaidInstanceReqRetentionPolicyEnum struct {
	TIME_BASE      CreatePostPaidInstanceReqRetentionPolicy
	PRODUCE_REJECT CreatePostPaidInstanceReqRetentionPolicy
}

func GetCreatePostPaidInstanceReqRetentionPolicyEnum() CreatePostPaidInstanceReqRetentionPolicyEnum {
	return CreatePostPaidInstanceReqRetentionPolicyEnum{
		TIME_BASE: CreatePostPaidInstanceReqRetentionPolicy{
			value: "time_base",
		},
		PRODUCE_REJECT: CreatePostPaidInstanceReqRetentionPolicy{
			value: "produce_reject",
		},
	}
}

func (c CreatePostPaidInstanceReqRetentionPolicy) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceReqRetentionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceReqRetentionPolicy) UnmarshalJSON(b []byte) error {
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
	DMS_PHYSICAL_STORAGE_NORMAL   CreatePostPaidInstanceReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_HIGH     CreatePostPaidInstanceReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_ULTRA    CreatePostPaidInstanceReqStorageSpecCode
}

func GetCreatePostPaidInstanceReqStorageSpecCodeEnum() CreatePostPaidInstanceReqStorageSpecCodeEnum {
	return CreatePostPaidInstanceReqStorageSpecCodeEnum{
		DMS_PHYSICAL_STORAGE_HIGH_V2: CreatePostPaidInstanceReqStorageSpecCode{
			value: "dms.physical.storage.high.v2",
		},
		DMS_PHYSICAL_STORAGE_ULTRA_V2: CreatePostPaidInstanceReqStorageSpecCode{
			value: "dms.physical.storage.ultra.v2",
		},
		DMS_PHYSICAL_STORAGE_NORMAL: CreatePostPaidInstanceReqStorageSpecCode{
			value: "dms.physical.storage.normal",
		},
		DMS_PHYSICAL_STORAGE_HIGH: CreatePostPaidInstanceReqStorageSpecCode{
			value: "dms.physical.storage.high",
		},
		DMS_PHYSICAL_STORAGE_ULTRA: CreatePostPaidInstanceReqStorageSpecCode{
			value: "dms.physical.storage.ultra",
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
