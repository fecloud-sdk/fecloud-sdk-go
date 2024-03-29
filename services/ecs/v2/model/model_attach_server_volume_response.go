package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttachServerVolumeResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachServerVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachServerVolumeResponse struct{}"
	}

	return strings.Join([]string{"AttachServerVolumeResponse", string(data)}, " ")
}
