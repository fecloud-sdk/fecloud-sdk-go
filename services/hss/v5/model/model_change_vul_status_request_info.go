package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeVulStatusRequestInfo struct {
	OperateType string `json:"operate_type"`

	Remark *string `json:"remark,omitempty"`

	SelectType *string `json:"select_type,omitempty"`

	Type *string `json:"type,omitempty"`

	DataList []VulOperateInfo `json:"data_list"`

	HostDataList *[]HostVulOperateInfo `json:"host_data_list,omitempty"`
}

func (o ChangeVulStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVulStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeVulStatusRequestInfo", string(data)}, " ")
}
