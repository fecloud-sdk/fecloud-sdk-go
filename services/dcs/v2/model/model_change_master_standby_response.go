package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeMasterStandbyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeMasterStandbyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeMasterStandbyResponse struct{}"
	}

	return strings.Join([]string{"ChangeMasterStandbyResponse", string(data)}, " ")
}
