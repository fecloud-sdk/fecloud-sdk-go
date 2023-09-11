package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AttachEipResponse Response Object
type AttachEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipResponse struct{}"
	}

	return strings.Join([]string{"AttachEipResponse", string(data)}, " ")
}
