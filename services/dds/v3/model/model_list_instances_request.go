package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListInstancesRequest struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Mode *ListInstancesRequestMode `json:"mode,omitempty"`

	DatastoreType *ListInstancesRequestDatastoreType `json:"datastore_type,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Tags *string `json:"tags,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}

type ListInstancesRequestMode struct {
	value string
}

type ListInstancesRequestModeEnum struct {
	SHARDING    ListInstancesRequestMode
	REPLICA_SET ListInstancesRequestMode
	SINGLE      ListInstancesRequestMode
}

func GetListInstancesRequestModeEnum() ListInstancesRequestModeEnum {
	return ListInstancesRequestModeEnum{
		SHARDING: ListInstancesRequestMode{
			value: "Sharding",
		},
		REPLICA_SET: ListInstancesRequestMode{
			value: "ReplicaSet",
		},
		SINGLE: ListInstancesRequestMode{
			value: "Single",
		},
	}
}

func (c ListInstancesRequestMode) Value() string {
	return c.value
}

func (c ListInstancesRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestMode) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestDatastoreType struct {
	value string
}

type ListInstancesRequestDatastoreTypeEnum struct {
	DDS_COMMUNITY ListInstancesRequestDatastoreType
	DDS_ENHANCED  ListInstancesRequestDatastoreType
}

func GetListInstancesRequestDatastoreTypeEnum() ListInstancesRequestDatastoreTypeEnum {
	return ListInstancesRequestDatastoreTypeEnum{
		DDS_COMMUNITY: ListInstancesRequestDatastoreType{
			value: "DDS-Community",
		},
		DDS_ENHANCED: ListInstancesRequestDatastoreType{
			value: "DDS-Enhanced",
		},
	}
}

func (c ListInstancesRequestDatastoreType) Value() string {
	return c.value
}

func (c ListInstancesRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestDatastoreType) UnmarshalJSON(b []byte) error {
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
