package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRestoreCollectionsRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	DbName string `json:"db_name"`

	RestoreTime string `json:"restore_time"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRestoreCollectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreCollectionsRequest struct{}"
	}

	return strings.Join([]string{"ListRestoreCollectionsRequest", string(data)}, " ")
}
