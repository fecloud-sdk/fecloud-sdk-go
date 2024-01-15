package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowInstanceResp struct {
	Name *string `json:"name,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engine_version,omitempty"`

	Description *string `json:"description,omitempty"`

	Specification *string `json:"specification,omitempty"`

	StorageSpace *int32 `json:"storage_space,omitempty"`

	PartitionNum *string `json:"partition_num,omitempty"`

	UsedStorageSpace *int32 `json:"used_storage_space,omitempty"`

	ConnectAddress *string `json:"connect_address,omitempty"`

	Port *int32 `json:"port,omitempty"`

	Status *string `json:"status,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	ChargingMode *int32 `json:"charging_mode,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	VpcName *string `json:"vpc_name,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	SubnetName *string `json:"subnet_name,omitempty"`

	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	UserId *string `json:"user_id,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	AccessUser *string `json:"access_user,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	MaintainBegin *string `json:"maintain_begin,omitempty"`

	MaintainEnd *string `json:"maintain_end,omitempty"`

	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	ManagementConnectAddress *string `json:"management_connect_address,omitempty"`

	SslEnable *bool `json:"ssl_enable,omitempty"`

	KafkaSecurityProtocol *string `json:"kafka_security_protocol,omitempty"`

	SaslEnabledMechanisms *[]ShowInstanceRespSaslEnabledMechanisms `json:"sasl_enabled_mechanisms,omitempty"`

	SslTwoWayEnable *bool `json:"ssl_two_way_enable,omitempty"`

	CertReplaced *bool `json:"cert_replaced,omitempty"`

	PublicManagementConnectAddress *string `json:"public_management_connect_address,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	IsLogicalVolume *bool `json:"is_logical_volume,omitempty"`

	ExtendTimes *int32 `json:"extend_times,omitempty"`

	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty"`

	Type *ShowInstanceRespType `json:"type,omitempty"`

	ProductId *string `json:"product_id,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	SecurityGroupName *string `json:"security_group_name,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	TotalStorageSpace *int32 `json:"total_storage_space,omitempty"`

	PublicConnectAddress *string `json:"public_connect_address,omitempty"`

	StorageResourceId *string `json:"storage_resource_id,omitempty"`

	StorageSpecCode *string `json:"storage_spec_code,omitempty"`

	ServiceType *string `json:"service_type,omitempty"`

	StorageType *string `json:"storage_type,omitempty"`

	RetentionPolicy *ShowInstanceRespRetentionPolicy `json:"retention_policy,omitempty"`

	KafkaPublicStatus *string `json:"kafka_public_status,omitempty"`

	PublicBandwidth *int32 `json:"public_bandwidth,omitempty"`

	KafkaManagerEnable *bool `json:"kafka_manager_enable,omitempty"`

	KafkaManagerUser *string `json:"kafka_manager_user,omitempty"`

	EnableLogCollection *bool `json:"enable_log_collection,omitempty"`

	CrossVpcInfo *string `json:"cross_vpc_info,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	Ipv6ConnectAddresses *[]string `json:"ipv6_connect_addresses,omitempty"`

	ConnectorEnable *bool `json:"connector_enable,omitempty"`

	ConnectorId *string `json:"connector_id,omitempty"`

	RestEnable *bool `json:"rest_enable,omitempty"`

	RestConnectAddress *string `json:"rest_connect_address,omitempty"`

	PublicBoundwidth *int32 `json:"public_boundwidth,omitempty"`

	MessageQueryInstEnable *bool `json:"message_query_inst_enable,omitempty"`

	VpcClientPlain *bool `json:"vpc_client_plain,omitempty"`

	SupportFeatures *string `json:"support_features,omitempty"`

	TraceEnable *bool `json:"trace_enable,omitempty"`

	AgentEnable *bool `json:"agent_enable,omitempty"`

	PodConnectAddress *string `json:"pod_connect_address,omitempty"`

	DiskEncrypted *bool `json:"disk_encrypted,omitempty"`

	DiskEncryptedKey *string `json:"disk_encrypted_key,omitempty"`

	KafkaPrivateConnectAddress *string `json:"kafka_private_connect_address,omitempty"`

	CesVersion *string `json:"ces_version,omitempty"`

	PublicAccessEnabled *string `json:"public_access_enabled,omitempty"`

	NodeNum *int32 `json:"node_num,omitempty"`

	EnableAcl *bool `json:"enable_acl,omitempty"`

	NewSpecBillingEnable *bool `json:"new_spec_billing_enable,omitempty"`

	BrokerNum *int32 `json:"broker_num,omitempty"`

	Tags *[]TagEntity `json:"tags,omitempty"`

	DrEnable *bool `json:"dr_enable,omitempty"`
}

func (o ShowInstanceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResp struct{}"
	}

	return strings.Join([]string{"ShowInstanceResp", string(data)}, " ")
}

type ShowInstanceRespSaslEnabledMechanisms struct {
	value string
}

type ShowInstanceRespSaslEnabledMechanismsEnum struct {
	PLAIN         ShowInstanceRespSaslEnabledMechanisms
	SCRAM_SHA_512 ShowInstanceRespSaslEnabledMechanisms
}

func GetShowInstanceRespSaslEnabledMechanismsEnum() ShowInstanceRespSaslEnabledMechanismsEnum {
	return ShowInstanceRespSaslEnabledMechanismsEnum{
		PLAIN: ShowInstanceRespSaslEnabledMechanisms{
			value: "PLAIN",
		},
		SCRAM_SHA_512: ShowInstanceRespSaslEnabledMechanisms{
			value: "SCRAM-SHA-512",
		},
	}
}

func (c ShowInstanceRespSaslEnabledMechanisms) Value() string {
	return c.value
}

func (c ShowInstanceRespSaslEnabledMechanisms) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceRespSaslEnabledMechanisms) UnmarshalJSON(b []byte) error {
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

type ShowInstanceRespType struct {
	value string
}

type ShowInstanceRespTypeEnum struct {
	SINGLE  ShowInstanceRespType
	CLUSTER ShowInstanceRespType
}

func GetShowInstanceRespTypeEnum() ShowInstanceRespTypeEnum {
	return ShowInstanceRespTypeEnum{
		SINGLE: ShowInstanceRespType{
			value: "single",
		},
		CLUSTER: ShowInstanceRespType{
			value: "cluster",
		},
	}
}

func (c ShowInstanceRespType) Value() string {
	return c.value
}

func (c ShowInstanceRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceRespType) UnmarshalJSON(b []byte) error {
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

type ShowInstanceRespRetentionPolicy struct {
	value string
}

type ShowInstanceRespRetentionPolicyEnum struct {
	TIME_BASE      ShowInstanceRespRetentionPolicy
	PRODUCE_REJECT ShowInstanceRespRetentionPolicy
}

func GetShowInstanceRespRetentionPolicyEnum() ShowInstanceRespRetentionPolicyEnum {
	return ShowInstanceRespRetentionPolicyEnum{
		TIME_BASE: ShowInstanceRespRetentionPolicy{
			value: "time_base",
		},
		PRODUCE_REJECT: ShowInstanceRespRetentionPolicy{
			value: "produce_reject",
		},
	}
}

func (c ShowInstanceRespRetentionPolicy) Value() string {
	return c.value
}

func (c ShowInstanceRespRetentionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceRespRetentionPolicy) UnmarshalJSON(b []byte) error {
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
