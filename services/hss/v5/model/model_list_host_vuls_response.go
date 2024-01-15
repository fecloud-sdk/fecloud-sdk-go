package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostVulsResponse struct {
	TotalNum *int64 `json:"total_num,omitempty"`

	DataList       *[]HostVulInfo `json:"data_list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListHostVulsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostVulsResponse struct{}"
	}

	return strings.Join([]string{"ListHostVulsResponse", string(data)}, " ")
}
