package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ConnectionEndpoints struct {
	Id *string `json:"id,omitempty"`

	MarkerId *int32 `json:"marker_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	Error *[]QueryError `json:"error,omitempty"`

	Status *ConnectionEndpointsStatus `json:"status,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o ConnectionEndpoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionEndpoints struct{}"
	}

	return strings.Join([]string{"ConnectionEndpoints", string(data)}, " ")
}

type ConnectionEndpointsStatus struct {
	value string
}

type ConnectionEndpointsStatusEnum struct {
	PENDING_ACCEPTANCE ConnectionEndpointsStatus
	CREATING           ConnectionEndpointsStatus
	ACCEPTED           ConnectionEndpointsStatus
	REJECTED           ConnectionEndpointsStatus
	FAILED             ConnectionEndpointsStatus
	DELETING           ConnectionEndpointsStatus
}

func GetConnectionEndpointsStatusEnum() ConnectionEndpointsStatusEnum {
	return ConnectionEndpointsStatusEnum{
		PENDING_ACCEPTANCE: ConnectionEndpointsStatus{
			value: "pendingAcceptance",
		},
		CREATING: ConnectionEndpointsStatus{
			value: "creating",
		},
		ACCEPTED: ConnectionEndpointsStatus{
			value: "accepted",
		},
		REJECTED: ConnectionEndpointsStatus{
			value: "rejected",
		},
		FAILED: ConnectionEndpointsStatus{
			value: "failed",
		},
		DELETING: ConnectionEndpointsStatus{
			value: "deleting",
		},
	}
}

func (c ConnectionEndpointsStatus) Value() string {
	return c.value
}

func (c ConnectionEndpointsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionEndpointsStatus) UnmarshalJSON(b []byte) error {
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
