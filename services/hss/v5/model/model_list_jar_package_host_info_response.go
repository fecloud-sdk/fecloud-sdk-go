package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListJarPackageHostInfoResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]JarPackageHostInfo `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListJarPackageHostInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJarPackageHostInfoResponse struct{}"
	}

	return strings.Join([]string{"ListJarPackageHostInfoResponse", string(data)}, " ")
}
