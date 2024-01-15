package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateConnectorReq struct {
	Specification *CreateConnectorReqSpecification `json:"specification,omitempty"`

	NodeCnt *string `json:"node_cnt,omitempty"`

	SpecCode string `json:"spec_code"`
}

func (o CreateConnectorReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectorReq struct{}"
	}

	return strings.Join([]string{"CreateConnectorReq", string(data)}, " ")
}

type CreateConnectorReqSpecification struct {
	value string
}

type CreateConnectorReqSpecificationEnum struct {
	E_100_MB  CreateConnectorReqSpecification
	E_300_MB  CreateConnectorReqSpecification
	E_600_MB  CreateConnectorReqSpecification
	E_1200_MB CreateConnectorReqSpecification
}

func GetCreateConnectorReqSpecificationEnum() CreateConnectorReqSpecificationEnum {
	return CreateConnectorReqSpecificationEnum{
		E_100_MB: CreateConnectorReqSpecification{
			value: "100MB",
		},
		E_300_MB: CreateConnectorReqSpecification{
			value: "300MB",
		},
		E_600_MB: CreateConnectorReqSpecification{
			value: "600MB",
		},
		E_1200_MB: CreateConnectorReqSpecification{
			value: "1200MB",
		},
	}
}

func (c CreateConnectorReqSpecification) Value() string {
	return c.value
}

func (c CreateConnectorReqSpecification) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConnectorReqSpecification) UnmarshalJSON(b []byte) error {
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
