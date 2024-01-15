package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListBackupsRequest struct {
	InstanceId *string `json:"instance_id,omitempty"`

	BackupId *string `json:"backup_id,omitempty"`

	BackupType *ListBackupsRequestBackupType `json:"backup_type,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	BeginTime *string `json:"begin_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Mode *ListBackupsRequestMode `json:"mode,omitempty"`
}

func (o ListBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsRequest struct{}"
	}

	return strings.Join([]string{"ListBackupsRequest", string(data)}, " ")
}

type ListBackupsRequestBackupType struct {
	value string
}

type ListBackupsRequestBackupTypeEnum struct {
	AUTO        ListBackupsRequestBackupType
	MANUAL      ListBackupsRequestBackupType
	INCREMENTAL ListBackupsRequestBackupType
}

func GetListBackupsRequestBackupTypeEnum() ListBackupsRequestBackupTypeEnum {
	return ListBackupsRequestBackupTypeEnum{
		AUTO: ListBackupsRequestBackupType{
			value: "Auto",
		},
		MANUAL: ListBackupsRequestBackupType{
			value: "Manual",
		},
		INCREMENTAL: ListBackupsRequestBackupType{
			value: "Incremental",
		},
	}
}

func (c ListBackupsRequestBackupType) Value() string {
	return c.value
}

func (c ListBackupsRequestBackupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestBackupType) UnmarshalJSON(b []byte) error {
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

type ListBackupsRequestMode struct {
	value string
}

type ListBackupsRequestModeEnum struct {
	SHARDING    ListBackupsRequestMode
	REPLICA_SET ListBackupsRequestMode
	SINGLE      ListBackupsRequestMode
}

func GetListBackupsRequestModeEnum() ListBackupsRequestModeEnum {
	return ListBackupsRequestModeEnum{
		SHARDING: ListBackupsRequestMode{
			value: "Sharding",
		},
		REPLICA_SET: ListBackupsRequestMode{
			value: "ReplicaSet",
		},
		SINGLE: ListBackupsRequestMode{
			value: "Single",
		},
	}
}

func (c ListBackupsRequestMode) Value() string {
	return c.value
}

func (c ListBackupsRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestMode) UnmarshalJSON(b []byte) error {
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
