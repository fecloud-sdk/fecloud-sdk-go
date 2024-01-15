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

	Status *string `json:"status,omitempty"`

	Description *string `json:"description,omitempty"`

	Type *ShowInstanceRespType `json:"type,omitempty"`

	Specification *string `json:"specification,omitempty"`

	EngineVersion *string `json:"engine_version,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	ChargingMode *int32 `json:"charging_mode,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	VpcName *string `json:"vpc_name,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	ProductId *string `json:"product_id,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	SecurityGroupName *string `json:"security_group_name,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	SubnetName *string `json:"subnet_name,omitempty"`

	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	UserId *string `json:"user_id,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	MaintainBegin *string `json:"maintain_begin,omitempty"`

	MaintainEnd *string `json:"maintain_end,omitempty"`

	EnableLogCollection *bool `json:"enable_log_collection,omitempty"`

	StorageSpace *int32 `json:"storage_space,omitempty"`

	UsedStorageSpace *int32 `json:"used_storage_space,omitempty"`

	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	PublicipId *string `json:"publicip_id,omitempty"`

	PublicipAddress *string `json:"publicip_address,omitempty"`

	SslEnable *bool `json:"ssl_enable,omitempty"`

	CrossVpcInfo *string `json:"cross_vpc_info,omitempty"`

	StorageResourceId *string `json:"storage_resource_id,omitempty"`

	StorageSpecCode *string `json:"storage_spec_code,omitempty"`

	ServiceType *string `json:"service_type,omitempty"`

	StorageType *string `json:"storage_type,omitempty"`

	ExtendTimes *int64 `json:"extend_times,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	SupportFeatures *string `json:"support_features,omitempty"`

	DiskEncrypted *bool `json:"disk_encrypted,omitempty"`

	CesVersion *string `json:"ces_version,omitempty"`

	NodeNum *int32 `json:"node_num,omitempty"`

	NewSpecBillingEnable *bool `json:"new_spec_billing_enable,omitempty"`

	EnableAcl *bool `json:"enable_acl,omitempty"`

	BrokerNum *int32 `json:"broker_num,omitempty"`

	NamesrvAddress *string `json:"namesrv_address,omitempty"`

	BrokerAddress *string `json:"broker_address,omitempty"`

	PublicNamesrvAddress *string `json:"public_namesrv_address,omitempty"`

	PublicBrokerAddress *string `json:"public_broker_address,omitempty"`

	GrpcAddress *string `json:"grpc_address,omitempty"`

	PublicGrpcAddress *string `json:"public_grpc_address,omitempty"`

	Tags *[]TagEntity `json:"tags,omitempty"`

	TotalStorageSpace *int32 `json:"total_storage_space,omitempty"`

	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`
}

func (o ShowInstanceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResp struct{}"
	}

	return strings.Join([]string{"ShowInstanceResp", string(data)}, " ")
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
