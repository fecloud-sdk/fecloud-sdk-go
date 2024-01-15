package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListConsumeGroupAccessPolicyRequest struct {
	Engine ListConsumeGroupAccessPolicyRequestEngine `json:"engine"`

	InstanceId string `json:"instance_id"`

	GroupId string `json:"group_id"`

	Offset *string `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`
}

func (o ListConsumeGroupAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumeGroupAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListConsumeGroupAccessPolicyRequest", string(data)}, " ")
}

type ListConsumeGroupAccessPolicyRequestEngine struct {
	value string
}

type ListConsumeGroupAccessPolicyRequestEngineEnum struct {
	RELIABILITY ListConsumeGroupAccessPolicyRequestEngine
}

func GetListConsumeGroupAccessPolicyRequestEngineEnum() ListConsumeGroupAccessPolicyRequestEngineEnum {
	return ListConsumeGroupAccessPolicyRequestEngineEnum{
		RELIABILITY: ListConsumeGroupAccessPolicyRequestEngine{
			value: "reliability",
		},
	}
}

func (c ListConsumeGroupAccessPolicyRequestEngine) Value() string {
	return c.value
}

func (c ListConsumeGroupAccessPolicyRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConsumeGroupAccessPolicyRequestEngine) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
