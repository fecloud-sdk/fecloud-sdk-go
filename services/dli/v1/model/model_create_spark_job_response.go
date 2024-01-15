package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSparkJobResponse struct {
	Id *string `json:"id,omitempty"`

	State *string `json:"state,omitempty"`

	AppId *string `json:"appId,omitempty"`

	Log *[]string `json:"log,omitempty"`

	ScType *string `json:"sc_type,omitempty"`

	ClusterName *string `json:"cluster_name,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	Name *string `json:"name,omitempty"`

	Owner *string `json:"owner,omitempty"`

	ProxyUser *string `json:"proxyUser,omitempty"`

	Kind *string `json:"kind,omitempty"`

	Queue *string `json:"queue,omitempty"`

	Image *string `json:"image,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	Duration       *int64 `json:"duration,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateSparkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSparkJobResponse struct{}"
	}

	return strings.Join([]string{"CreateSparkJobResponse", string(data)}, " ")
}
