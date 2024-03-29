package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DetachServerVolumeResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetachServerVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachServerVolumeResponse struct{}"
	}

	return strings.Join([]string{"DetachServerVolumeResponse", string(data)}, " ")
}
