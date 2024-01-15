package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRestoreCollectionsResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Collections    *[]string `json:"collections,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRestoreCollectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreCollectionsResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreCollectionsResponse", string(data)}, " ")
}
