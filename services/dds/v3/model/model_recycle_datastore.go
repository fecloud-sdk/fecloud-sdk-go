package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type RecycleDatastore struct {
	Type RecycleDatastoreType `json:"type"`

	Version string `json:"version"`
}

func (o RecycleDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleDatastore struct{}"
	}

	return strings.Join([]string{"RecycleDatastore", string(data)}, " ")
}

type RecycleDatastoreType struct {
	value string
}

type RecycleDatastoreTypeEnum struct {
	DDS_COMMUNITY RecycleDatastoreType
}

func GetRecycleDatastoreTypeEnum() RecycleDatastoreTypeEnum {
	return RecycleDatastoreTypeEnum{
		DDS_COMMUNITY: RecycleDatastoreType{
			value: "DDS-Community",
		},
	}
}

func (c RecycleDatastoreType) Value() string {
	return c.value
}

func (c RecycleDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleDatastoreType) UnmarshalJSON(b []byte) error {
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
