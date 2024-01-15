package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ClusterDataConnectorMap struct {
	MapId *int32 `json:"map_id,omitempty"`

	ConnectorId *string `json:"connector_id,omitempty"`

	ComponentName *string `json:"component_name,omitempty"`

	RoleType *ClusterDataConnectorMapRoleType `json:"role_type,omitempty"`

	SourceType *ClusterDataConnectorMapSourceType `json:"source_type,omitempty"`

	ClusterId *string `json:"cluster_id,omitempty"`

	Status *int32 `json:"status,omitempty"`
}

func (o ClusterDataConnectorMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDataConnectorMap struct{}"
	}

	return strings.Join([]string{"ClusterDataConnectorMap", string(data)}, " ")
}

type ClusterDataConnectorMapRoleType struct {
	value string
}

type ClusterDataConnectorMapRoleTypeEnum struct {
	LOCAL_DB      ClusterDataConnectorMapRoleType
	RDS_POSTGRES  ClusterDataConnectorMapRoleType
	RDS_MYSQL     ClusterDataConnectorMapRoleType
	GAUSSDB_MYSQL ClusterDataConnectorMapRoleType
}

func GetClusterDataConnectorMapRoleTypeEnum() ClusterDataConnectorMapRoleTypeEnum {
	return ClusterDataConnectorMapRoleTypeEnum{
		LOCAL_DB: ClusterDataConnectorMapRoleType{
			value: "LOCAL_DB",
		},
		RDS_POSTGRES: ClusterDataConnectorMapRoleType{
			value: "RDS_POSTGRES",
		},
		RDS_MYSQL: ClusterDataConnectorMapRoleType{
			value: "RDS_MYSQL",
		},
		GAUSSDB_MYSQL: ClusterDataConnectorMapRoleType{
			value: "gaussdb-mysql",
		},
	}
}

func (c ClusterDataConnectorMapRoleType) Value() string {
	return c.value
}

func (c ClusterDataConnectorMapRoleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterDataConnectorMapRoleType) UnmarshalJSON(b []byte) error {
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

type ClusterDataConnectorMapSourceType struct {
	value string
}

type ClusterDataConnectorMapSourceTypeEnum struct {
	LOCAL_DB      ClusterDataConnectorMapSourceType
	RDS_POSTGRES  ClusterDataConnectorMapSourceType
	RDS_MYSQL     ClusterDataConnectorMapSourceType
	GAUSSDB_MYSQL ClusterDataConnectorMapSourceType
}

func GetClusterDataConnectorMapSourceTypeEnum() ClusterDataConnectorMapSourceTypeEnum {
	return ClusterDataConnectorMapSourceTypeEnum{
		LOCAL_DB: ClusterDataConnectorMapSourceType{
			value: "LOCAL_DB",
		},
		RDS_POSTGRES: ClusterDataConnectorMapSourceType{
			value: "RDS_POSTGRES",
		},
		RDS_MYSQL: ClusterDataConnectorMapSourceType{
			value: "RDS_MYSQL",
		},
		GAUSSDB_MYSQL: ClusterDataConnectorMapSourceType{
			value: "gaussdb-mysql",
		},
	}
}

func (c ClusterDataConnectorMapSourceType) Value() string {
	return c.value
}

func (c ClusterDataConnectorMapSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterDataConnectorMapSourceType) UnmarshalJSON(b []byte) error {
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
