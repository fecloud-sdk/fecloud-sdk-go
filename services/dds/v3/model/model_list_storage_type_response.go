package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListStorageTypeResponse struct {
	StorageType *[]Storage `json:"storage_type,omitempty"`

	DssPoolInfo    *[]DssPoolInfo `json:"dss_pool_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListStorageTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageTypeResponse struct{}"
	}

	return strings.Join([]string{"ListStorageTypeResponse", string(data)}, " ")
}
