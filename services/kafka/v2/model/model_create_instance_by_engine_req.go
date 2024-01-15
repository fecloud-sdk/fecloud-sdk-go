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

	EngineVersion string `json:"engine_version"`

	BrokerNum int32 `json:"broker_num"`

	StorageSpace int32 `json:"storage_space"`

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

	PublicipId *string `json:"publicip_id,omitempty"`

	SslEnable *bool `json:"ssl_enable,omitempty"`

	KafkaSecurityProtocol *string `json:"kafka_security_protocol,omitempty"`

	SaslEnabledMechanisms *[]CreateInstanceByEngineReqSaslEnabledMechanisms `json:"sasl_enabled_mechanisms,omitempty"`

	RetentionPolicy *CreateInstanceByEngineReqRetentionPolicy `json:"retention_policy,omitempty"`

	DiskEncryptedEnable *bool `json:"disk_encrypted_enable,omitempty"`

	DiskEncryptedKey *string `json:"disk_encrypted_key,omitempty"`

	ConnectorEnable *bool `json:"connector_enable,omitempty"`

	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty"`

	StorageSpecCode CreateInstanceByEngineReqStorageSpecCode `json:"storage_spec_code"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]TagEntity `json:"tags,omitempty"`

	ArchType *string `json:"arch_type,omitempty"`

	VpcClientPlain *bool `json:"vpc_client_plain,omitempty"`

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
	KAFKA CreateInstanceByEngineReqEngine
}

func GetCreateInstanceByEngineReqEngineEnum() CreateInstanceByEngineReqEngineEnum {
	return CreateInstanceByEngineReqEngineEnum{
		KAFKA: CreateInstanceByEngineReqEngine{
			value: "kafka",
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

type CreateInstanceByEngineReqSaslEnabledMechanisms struct {
	value string
}

type CreateInstanceByEngineReqSaslEnabledMechanismsEnum struct {
	PLAIN         CreateInstanceByEngineReqSaslEnabledMechanisms
	SCRAM_SHA_512 CreateInstanceByEngineReqSaslEnabledMechanisms
}

func GetCreateInstanceByEngineReqSaslEnabledMechanismsEnum() CreateInstanceByEngineReqSaslEnabledMechanismsEnum {
	return CreateInstanceByEngineReqSaslEnabledMechanismsEnum{
		PLAIN: CreateInstanceByEngineReqSaslEnabledMechanisms{
			value: "PLAIN",
		},
		SCRAM_SHA_512: CreateInstanceByEngineReqSaslEnabledMechanisms{
			value: "SCRAM-SHA-512",
		},
	}
}

func (c CreateInstanceByEngineReqSaslEnabledMechanisms) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqSaslEnabledMechanisms) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqSaslEnabledMechanisms) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqRetentionPolicy struct {
	value string
}

type CreateInstanceByEngineReqRetentionPolicyEnum struct {
	TIME_BASE      CreateInstanceByEngineReqRetentionPolicy
	PRODUCE_REJECT CreateInstanceByEngineReqRetentionPolicy
}

func GetCreateInstanceByEngineReqRetentionPolicyEnum() CreateInstanceByEngineReqRetentionPolicyEnum {
	return CreateInstanceByEngineReqRetentionPolicyEnum{
		TIME_BASE: CreateInstanceByEngineReqRetentionPolicy{
			value: "time_base",
		},
		PRODUCE_REJECT: CreateInstanceByEngineReqRetentionPolicy{
			value: "produce_reject",
		},
	}
}

func (c CreateInstanceByEngineReqRetentionPolicy) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqRetentionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqRetentionPolicy) UnmarshalJSON(b []byte) error {
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
