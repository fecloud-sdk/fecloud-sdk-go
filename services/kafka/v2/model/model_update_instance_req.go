package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateInstanceReq struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	MaintainBegin *string `json:"maintain_begin,omitempty"`

	MaintainEnd *string `json:"maintain_end,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	RetentionPolicy *UpdateInstanceReqRetentionPolicy `json:"retention_policy,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o UpdateInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceReq struct{}"
	}

	return strings.Join([]string{"UpdateInstanceReq", string(data)}, " ")
}

type UpdateInstanceReqRetentionPolicy struct {
	value string
}

type UpdateInstanceReqRetentionPolicyEnum struct {
	PRODUCE_REJECT UpdateInstanceReqRetentionPolicy
	TIME_BASE      UpdateInstanceReqRetentionPolicy
}

func GetUpdateInstanceReqRetentionPolicyEnum() UpdateInstanceReqRetentionPolicyEnum {
	return UpdateInstanceReqRetentionPolicyEnum{
		PRODUCE_REJECT: UpdateInstanceReqRetentionPolicy{
			value: "produce_reject",
		},
		TIME_BASE: UpdateInstanceReqRetentionPolicy{
			value: "time_base",
		},
	}
}

func (c UpdateInstanceReqRetentionPolicy) Value() string {
	return c.value
}

func (c UpdateInstanceReqRetentionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceReqRetentionPolicy) UnmarshalJSON(b []byte) error {
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
