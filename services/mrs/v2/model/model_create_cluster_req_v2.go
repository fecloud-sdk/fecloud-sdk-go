package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateClusterReqV2 struct {
	IsDecProject *bool `json:"is_dec_project,omitempty"`

	ClusterVersion string `json:"cluster_version"`

	ClusterName string `json:"cluster_name"`

	ClusterType string `json:"cluster_type"`

	Region string `json:"region"`

	VpcName string `json:"vpc_name"`

	SubnetId *string `json:"subnet_id,omitempty"`

	SubnetName string `json:"subnet_name"`

	Components string `json:"components"`

	ExternalDatasources *[]ClusterDataConnectorMap `json:"external_datasources,omitempty"`

	AvailabilityZone string `json:"availability_zone"`

	SecurityGroupsId *string `json:"security_groups_id,omitempty"`

	AutoCreateDefaultSecurityGroup *bool `json:"auto_create_default_security_group,omitempty"`

	SafeMode string `json:"safe_mode"`

	ManagerAdminPassword string `json:"manager_admin_password"`

	LoginMode string `json:"login_mode"`

	NodeRootPassword *string `json:"node_root_password,omitempty"`

	NodeKeypairName *string `json:"node_keypair_name,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	EipAddress *string `json:"eip_address,omitempty"`

	EipId *string `json:"eip_id,omitempty"`

	MrsEcsDefaultAgency *string `json:"mrs_ecs_default_agency,omitempty"`

	TemplateId *string `json:"template_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	LogCollection *CreateClusterReqV2LogCollection `json:"log_collection,omitempty"`

	NodeGroups []NodeGroupV2 `json:"node_groups"`

	BootstrapScripts *[]BootstrapScript `json:"bootstrap_scripts,omitempty"`

	AddJobs *[]AddJobsReqV11 `json:"add_jobs,omitempty"`

	LogUri *string `json:"log_uri,omitempty"`

	ComponentConfigs *[]ComponentConfig `json:"component_configs,omitempty"`
}

func (o CreateClusterReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterReqV2 struct{}"
	}

	return strings.Join([]string{"CreateClusterReqV2", string(data)}, " ")
}

type CreateClusterReqV2LogCollection struct {
	value int32
}

type CreateClusterReqV2LogCollectionEnum struct {
	E_0 CreateClusterReqV2LogCollection
	E_1 CreateClusterReqV2LogCollection
}

func GetCreateClusterReqV2LogCollectionEnum() CreateClusterReqV2LogCollectionEnum {
	return CreateClusterReqV2LogCollectionEnum{
		E_0: CreateClusterReqV2LogCollection{
			value: 0,
		}, E_1: CreateClusterReqV2LogCollection{
			value: 1,
		},
	}
}

func (c CreateClusterReqV2LogCollection) Value() int32 {
	return c.value
}

func (c CreateClusterReqV2LogCollection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV2LogCollection) UnmarshalJSON(b []byte) error {
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
