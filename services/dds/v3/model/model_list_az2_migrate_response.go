package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAz2MigrateResponse struct {
	AzList         *[]Az2Migrate `json:"az_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListAz2MigrateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAz2MigrateResponse struct{}"
	}

	return strings.Join([]string{"ListAz2MigrateResponse", string(data)}, " ")
}
