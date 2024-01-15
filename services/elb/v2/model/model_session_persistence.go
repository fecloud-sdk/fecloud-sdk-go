package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type SessionPersistence struct {
	Type SessionPersistenceType `json:"type"`

	CookieName *string `json:"cookie_name,omitempty"`

	PersistenceTimeout *int32 `json:"persistence_timeout,omitempty"`
}

func (o SessionPersistence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionPersistence struct{}"
	}

	return strings.Join([]string{"SessionPersistence", string(data)}, " ")
}

type SessionPersistenceType struct {
	value string
}

type SessionPersistenceTypeEnum struct {
	SOURCE_IP   SessionPersistenceType
	HTTP_COOKIE SessionPersistenceType
	APP_COOKIE  SessionPersistenceType
}

func GetSessionPersistenceTypeEnum() SessionPersistenceTypeEnum {
	return SessionPersistenceTypeEnum{
		SOURCE_IP: SessionPersistenceType{
			value: "SOURCE_IP",
		},
		HTTP_COOKIE: SessionPersistenceType{
			value: "HTTP_COOKIE",
		},
		APP_COOKIE: SessionPersistenceType{
			value: "APP_COOKIE",
		},
	}
}

func (c SessionPersistenceType) Value() string {
	return c.value
}

func (c SessionPersistenceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SessionPersistenceType) UnmarshalJSON(b []byte) error {
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
