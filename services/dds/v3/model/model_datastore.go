package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type Datastore struct {
	Type DatastoreType `json:"type"`

	Version string `json:"version"`

	StorageEngine DatastoreStorageEngine `json:"storage_engine"`
}

func (o Datastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datastore struct{}"
	}

	return strings.Join([]string{"Datastore", string(data)}, " ")
}

type DatastoreType struct {
	value string
}

type DatastoreTypeEnum struct {
	DDS_COMMUNITY DatastoreType
}

func GetDatastoreTypeEnum() DatastoreTypeEnum {
	return DatastoreTypeEnum{
		DDS_COMMUNITY: DatastoreType{
			value: "DDS-Community",
		},
	}
}

func (c DatastoreType) Value() string {
	return c.value
}

func (c DatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatastoreType) UnmarshalJSON(b []byte) error {
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

type DatastoreStorageEngine struct {
	value string
}

type DatastoreStorageEngineEnum struct {
	WIRED_TIGER DatastoreStorageEngine
	ROCKS_DB    DatastoreStorageEngine
}

func GetDatastoreStorageEngineEnum() DatastoreStorageEngineEnum {
	return DatastoreStorageEngineEnum{
		WIRED_TIGER: DatastoreStorageEngine{
			value: "wiredTiger",
		},
		ROCKS_DB: DatastoreStorageEngine{
			value: "rocksDB",
		},
	}
}

func (c DatastoreStorageEngine) Value() string {
	return c.value
}

func (c DatastoreStorageEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatastoreStorageEngine) UnmarshalJSON(b []byte) error {
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
