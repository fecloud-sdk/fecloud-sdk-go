package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QuerySessionResponse struct {
	Id string `json:"id"`

	Active bool `json:"active"`

	Operation string `json:"operation"`

	Type string `json:"type"`

	CostTime string `json:"cost_time"`

	PlanSummary string `json:"plan_summary"`

	Host string `json:"host"`

	Client string `json:"client"`

	Description string `json:"description"`

	Namespace string `json:"namespace"`
}

func (o QuerySessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySessionResponse struct{}"
	}

	return strings.Join([]string{"QuerySessionResponse", string(data)}, " ")
}
