package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSparkJobResponse struct {
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

	ReqBody *string `json:"req_body,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	Feature *string `json:"feature,omitempty"`

	SparkVersion   *string `json:"spark_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSparkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobResponse", string(data)}, " ")
}
