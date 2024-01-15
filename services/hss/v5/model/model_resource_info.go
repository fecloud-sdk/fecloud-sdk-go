package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ResourceInfo struct {
	HostId *string `json:"host_id,omitempty"`

	HistoryBackupStatus *ResourceInfoHistoryBackupStatus `json:"history_backup_status,omitempty"`
}

func (o ResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInfo struct{}"
	}

	return strings.Join([]string{"ResourceInfo", string(data)}, " ")
}

type ResourceInfoHistoryBackupStatus struct {
	value string
}

type ResourceInfoHistoryBackupStatusEnum struct {
	OPENED ResourceInfoHistoryBackupStatus
	CLOSED ResourceInfoHistoryBackupStatus
}

func GetResourceInfoHistoryBackupStatusEnum() ResourceInfoHistoryBackupStatusEnum {
	return ResourceInfoHistoryBackupStatusEnum{
		OPENED: ResourceInfoHistoryBackupStatus{
			value: "opened",
		},
		CLOSED: ResourceInfoHistoryBackupStatus{
			value: "closed",
		},
	}
}

func (c ResourceInfoHistoryBackupStatus) Value() string {
	return c.value
}

func (c ResourceInfoHistoryBackupStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceInfoHistoryBackupStatus) UnmarshalJSON(b []byte) error {
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
