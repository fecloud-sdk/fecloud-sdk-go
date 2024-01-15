package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListJarPackageStatisticsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]JarPackageStatisticsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListJarPackageStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJarPackageStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListJarPackageStatisticsResponse", string(data)}, " ")
}
