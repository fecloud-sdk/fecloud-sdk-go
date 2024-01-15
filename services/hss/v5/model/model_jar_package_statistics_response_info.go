package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JarPackageStatisticsResponseInfo struct {
	FileName *string `json:"file_name,omitempty"`

	Num *int32 `json:"num,omitempty"`
}

func (o JarPackageStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JarPackageStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"JarPackageStatisticsResponseInfo", string(data)}, " ")
}
