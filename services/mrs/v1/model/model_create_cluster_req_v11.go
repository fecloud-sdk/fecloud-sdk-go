package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateClusterReqV11 struct {
	ClusterVersion string `json:"cluster_version"`

	ClusterName string `json:"cluster_name"`

	MasterNodeNum int32 `json:"master_node_num"`

	CoreNodeNum int32 `json:"core_node_num"`

	DataCenter string `json:"data_center"`

	Vpc string `json:"vpc"`

	MasterNodeSize string `json:"master_node_size"`

	CoreNodeSize string `json:"core_node_size"`

	ComponentList []ComponentAmbV11 `json:"component_list"`

	AvailableZoneId string `json:"available_zone_id"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SubnetName string `json:"subnet_name"`

	SecurityGroupsId *string `json:"security_groups_id,omitempty"`

	AddJobs *[]AddJobsReqV11 `json:"add_jobs,omitempty"`

	VolumeSize *int32 `json:"volume_size,omitempty"`

	VolumeType *CreateClusterReqV11VolumeType `json:"volume_type,omitempty"`

	MasterDataVolumeType *CreateClusterReqV11MasterDataVolumeType `json:"master_data_volume_type,omitempty"`

	MasterDataVolumeSize *int32 `json:"master_data_volume_size,omitempty"`

	MasterDataVolumeCount *CreateClusterReqV11MasterDataVolumeCount `json:"master_data_volume_count,omitempty"`

	CoreDataVolumeType *CreateClusterReqV11CoreDataVolumeType `json:"core_data_volume_type,omitempty"`

	CoreDataVolumeSize *int32 `json:"core_data_volume_size,omitempty"`

	CoreDataVolumeCount *int32 `json:"core_data_volume_count,omitempty"`

	TaskNodeGroups *[]TaskNodeGroup `json:"task_node_groups,omitempty"`

	BootstrapScripts *[]BootstrapScript `json:"bootstrap_scripts,omitempty"`

	NodePublicCertName *string `json:"node_public_cert_name,omitempty"`

	ClusterAdminSecret *string `json:"cluster_admin_secret,omitempty"`

	ClusterMasterSecret string `json:"cluster_master_secret"`

	SafeMode CreateClusterReqV11SafeMode `json:"safe_mode"`

	ClusterType *CreateClusterReqV11ClusterType `json:"cluster_type,omitempty"`

	LogCollection *CreateClusterReqV11LogCollection `json:"log_collection,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	LoginMode *CreateClusterReqV11LoginMode `json:"login_mode,omitempty"`

	NodeGroups *[]NodeGroupV11 `json:"node_groups,omitempty"`
}

func (o CreateClusterReqV11) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterReqV11 struct{}"
	}

	return strings.Join([]string{"CreateClusterReqV11", string(data)}, " ")
}

type CreateClusterReqV11VolumeType struct {
	value string
}

type CreateClusterReqV11VolumeTypeEnum struct {
	SATA  CreateClusterReqV11VolumeType
	SAS   CreateClusterReqV11VolumeType
	SSD   CreateClusterReqV11VolumeType
	GPSSD CreateClusterReqV11VolumeType
}

func GetCreateClusterReqV11VolumeTypeEnum() CreateClusterReqV11VolumeTypeEnum {
	return CreateClusterReqV11VolumeTypeEnum{
		SATA: CreateClusterReqV11VolumeType{
			value: "SATA",
		},
		SAS: CreateClusterReqV11VolumeType{
			value: "SAS",
		},
		SSD: CreateClusterReqV11VolumeType{
			value: "SSD",
		},
		GPSSD: CreateClusterReqV11VolumeType{
			value: "GPSSD",
		},
	}
}

func (c CreateClusterReqV11VolumeType) Value() string {
	return c.value
}

func (c CreateClusterReqV11VolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11VolumeType) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqV11MasterDataVolumeType struct {
	value string
}

type CreateClusterReqV11MasterDataVolumeTypeEnum struct {
	SATA  CreateClusterReqV11MasterDataVolumeType
	SAS   CreateClusterReqV11MasterDataVolumeType
	SSD   CreateClusterReqV11MasterDataVolumeType
	GPSSD CreateClusterReqV11MasterDataVolumeType
}

func GetCreateClusterReqV11MasterDataVolumeTypeEnum() CreateClusterReqV11MasterDataVolumeTypeEnum {
	return CreateClusterReqV11MasterDataVolumeTypeEnum{
		SATA: CreateClusterReqV11MasterDataVolumeType{
			value: "SATA",
		},
		SAS: CreateClusterReqV11MasterDataVolumeType{
			value: "SAS",
		},
		SSD: CreateClusterReqV11MasterDataVolumeType{
			value: "SSD",
		},
		GPSSD: CreateClusterReqV11MasterDataVolumeType{
			value: "GPSSD",
		},
	}
}

func (c CreateClusterReqV11MasterDataVolumeType) Value() string {
	return c.value
}

func (c CreateClusterReqV11MasterDataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11MasterDataVolumeType) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqV11MasterDataVolumeCount struct {
	value int32
}

type CreateClusterReqV11MasterDataVolumeCountEnum struct {
	E_1 CreateClusterReqV11MasterDataVolumeCount
}

func GetCreateClusterReqV11MasterDataVolumeCountEnum() CreateClusterReqV11MasterDataVolumeCountEnum {
	return CreateClusterReqV11MasterDataVolumeCountEnum{
		E_1: CreateClusterReqV11MasterDataVolumeCount{
			value: 1,
		},
	}
}

func (c CreateClusterReqV11MasterDataVolumeCount) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11MasterDataVolumeCount) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11MasterDataVolumeCount) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqV11CoreDataVolumeType struct {
	value string
}

type CreateClusterReqV11CoreDataVolumeTypeEnum struct {
	SATA  CreateClusterReqV11CoreDataVolumeType
	SAS   CreateClusterReqV11CoreDataVolumeType
	SSD   CreateClusterReqV11CoreDataVolumeType
	GPSSD CreateClusterReqV11CoreDataVolumeType
}

func GetCreateClusterReqV11CoreDataVolumeTypeEnum() CreateClusterReqV11CoreDataVolumeTypeEnum {
	return CreateClusterReqV11CoreDataVolumeTypeEnum{
		SATA: CreateClusterReqV11CoreDataVolumeType{
			value: "SATA",
		},
		SAS: CreateClusterReqV11CoreDataVolumeType{
			value: "SAS",
		},
		SSD: CreateClusterReqV11CoreDataVolumeType{
			value: "SSD",
		},
		GPSSD: CreateClusterReqV11CoreDataVolumeType{
			value: "GPSSD",
		},
	}
}

func (c CreateClusterReqV11CoreDataVolumeType) Value() string {
	return c.value
}

func (c CreateClusterReqV11CoreDataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11CoreDataVolumeType) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqV11SafeMode struct {
	value int32
}

type CreateClusterReqV11SafeModeEnum struct {
	E_0 CreateClusterReqV11SafeMode
	E_1 CreateClusterReqV11SafeMode
}

func GetCreateClusterReqV11SafeModeEnum() CreateClusterReqV11SafeModeEnum {
	return CreateClusterReqV11SafeModeEnum{
		E_0: CreateClusterReqV11SafeMode{
			value: 0,
		}, E_1: CreateClusterReqV11SafeMode{
			value: 1,
		},
	}
}

func (c CreateClusterReqV11SafeMode) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11SafeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11SafeMode) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqV11ClusterType struct {
	value int32
}

type CreateClusterReqV11ClusterTypeEnum struct {
	E_0 CreateClusterReqV11ClusterType
	E_1 CreateClusterReqV11ClusterType
}

func GetCreateClusterReqV11ClusterTypeEnum() CreateClusterReqV11ClusterTypeEnum {
	return CreateClusterReqV11ClusterTypeEnum{
		E_0: CreateClusterReqV11ClusterType{
			value: 0,
		}, E_1: CreateClusterReqV11ClusterType{
			value: 1,
		},
	}
}

func (c CreateClusterReqV11ClusterType) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11ClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11ClusterType) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqV11LogCollection struct {
	value int32
}

type CreateClusterReqV11LogCollectionEnum struct {
	E_0 CreateClusterReqV11LogCollection
	E_1 CreateClusterReqV11LogCollection
}

func GetCreateClusterReqV11LogCollectionEnum() CreateClusterReqV11LogCollectionEnum {
	return CreateClusterReqV11LogCollectionEnum{
		E_0: CreateClusterReqV11LogCollection{
			value: 0,
		}, E_1: CreateClusterReqV11LogCollection{
			value: 1,
		},
	}
}

func (c CreateClusterReqV11LogCollection) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11LogCollection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11LogCollection) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqV11LoginMode struct {
	value int32
}

type CreateClusterReqV11LoginModeEnum struct {
	E_0 CreateClusterReqV11LoginMode
	E_1 CreateClusterReqV11LoginMode
}

func GetCreateClusterReqV11LoginModeEnum() CreateClusterReqV11LoginModeEnum {
	return CreateClusterReqV11LoginModeEnum{
		E_0: CreateClusterReqV11LoginMode{
			value: 0,
		}, E_1: CreateClusterReqV11LoginMode{
			value: 1,
		},
	}
}

func (c CreateClusterReqV11LoginMode) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11LoginMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11LoginMode) UnmarshalJSON(b []byte) error {
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
