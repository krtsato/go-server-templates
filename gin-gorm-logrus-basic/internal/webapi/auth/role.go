package auth

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Role アカウント権限 識別型
type Role int

// Roles Role スライス
type Roles []Role

// Role 種類
const (
	UnknownRole Role = iota
	AdminRole
	WriterRole
	ReaderRole
)

var roleValueMap = map[Role]string{
	AdminRole:  "adminRole",
	WriterRole: "writerRole",
	ReaderRole: "readerRole",
}

// String Role 文字列を返却
func (r Role) String() string {
	return roleValueMap[r]
}

func applyRole(str string) Role {
	for k, v := range roleValueMap {
		if strings.EqualFold(v, str) {
			return k
		}
	}
	return UnknownRole
}

// UnmarshalJSON for JSON
func (r *Role) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("fail to unmarshal role: %v", data)
	}
	*r = applyRole(str)
	return nil
}

func (rs Roles) contains(r Role) bool {
	for _, v := range rs {
		if v == r {
			return true
		}
	}
	return false
}

// HasRole アカウント権限に AdminRole, WriterRole, ReaderRole が含まれる真偽
func (rs Roles) HasRole(accountRoles ...Role) bool {
	for _, r := range accountRoles {
		if rs.contains(r) {
			return true
		}
	}
	return false
}
