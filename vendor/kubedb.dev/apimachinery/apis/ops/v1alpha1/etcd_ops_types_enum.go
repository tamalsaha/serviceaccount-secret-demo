// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package v1alpha1

import (
	"fmt"
	"strings"
)

const (
	// EtcdOpsRequestTypeUpdateVersion is a EtcdOpsRequestType of type UpdateVersion.
	EtcdOpsRequestTypeUpdateVersion EtcdOpsRequestType = "UpdateVersion"
	// EtcdOpsRequestTypeHorizontalScaling is a EtcdOpsRequestType of type HorizontalScaling.
	EtcdOpsRequestTypeHorizontalScaling EtcdOpsRequestType = "HorizontalScaling"
	// EtcdOpsRequestTypeVerticalScaling is a EtcdOpsRequestType of type VerticalScaling.
	EtcdOpsRequestTypeVerticalScaling EtcdOpsRequestType = "VerticalScaling"
	// EtcdOpsRequestTypeVolumeExpansion is a EtcdOpsRequestType of type VolumeExpansion.
	EtcdOpsRequestTypeVolumeExpansion EtcdOpsRequestType = "VolumeExpansion"
	// EtcdOpsRequestTypeRestart is a EtcdOpsRequestType of type Restart.
	EtcdOpsRequestTypeRestart EtcdOpsRequestType = "Restart"
	// EtcdOpsRequestTypeReconfigure is a EtcdOpsRequestType of type Reconfigure.
	EtcdOpsRequestTypeReconfigure EtcdOpsRequestType = "Reconfigure"
	// EtcdOpsRequestTypeReconfigureTLS is a EtcdOpsRequestType of type ReconfigureTLS.
	EtcdOpsRequestTypeReconfigureTLS EtcdOpsRequestType = "ReconfigureTLS"
)

var ErrInvalidEtcdOpsRequestType = fmt.Errorf("not a valid EtcdOpsRequestType, try [%s]", strings.Join(_EtcdOpsRequestTypeNames, ", "))

var _EtcdOpsRequestTypeNames = []string{
	string(EtcdOpsRequestTypeUpdateVersion),
	string(EtcdOpsRequestTypeHorizontalScaling),
	string(EtcdOpsRequestTypeVerticalScaling),
	string(EtcdOpsRequestTypeVolumeExpansion),
	string(EtcdOpsRequestTypeRestart),
	string(EtcdOpsRequestTypeReconfigure),
	string(EtcdOpsRequestTypeReconfigureTLS),
}

// EtcdOpsRequestTypeNames returns a list of possible string values of EtcdOpsRequestType.
func EtcdOpsRequestTypeNames() []string {
	tmp := make([]string, len(_EtcdOpsRequestTypeNames))
	copy(tmp, _EtcdOpsRequestTypeNames)
	return tmp
}

// EtcdOpsRequestTypeValues returns a list of the values for EtcdOpsRequestType
func EtcdOpsRequestTypeValues() []EtcdOpsRequestType {
	return []EtcdOpsRequestType{
		EtcdOpsRequestTypeUpdateVersion,
		EtcdOpsRequestTypeHorizontalScaling,
		EtcdOpsRequestTypeVerticalScaling,
		EtcdOpsRequestTypeVolumeExpansion,
		EtcdOpsRequestTypeRestart,
		EtcdOpsRequestTypeReconfigure,
		EtcdOpsRequestTypeReconfigureTLS,
	}
}

// String implements the Stringer interface.
func (x EtcdOpsRequestType) String() string {
	return string(x)
}

// String implements the Stringer interface.
func (x EtcdOpsRequestType) IsValid() bool {
	_, err := ParseEtcdOpsRequestType(string(x))
	return err == nil
}

var _EtcdOpsRequestTypeValue = map[string]EtcdOpsRequestType{
	"UpdateVersion":     EtcdOpsRequestTypeUpdateVersion,
	"HorizontalScaling": EtcdOpsRequestTypeHorizontalScaling,
	"VerticalScaling":   EtcdOpsRequestTypeVerticalScaling,
	"VolumeExpansion":   EtcdOpsRequestTypeVolumeExpansion,
	"Restart":           EtcdOpsRequestTypeRestart,
	"Reconfigure":       EtcdOpsRequestTypeReconfigure,
	"ReconfigureTLS":    EtcdOpsRequestTypeReconfigureTLS,
}

// ParseEtcdOpsRequestType attempts to convert a string to a EtcdOpsRequestType.
func ParseEtcdOpsRequestType(name string) (EtcdOpsRequestType, error) {
	if x, ok := _EtcdOpsRequestTypeValue[name]; ok {
		return x, nil
	}
	return EtcdOpsRequestType(""), fmt.Errorf("%s is %w", name, ErrInvalidEtcdOpsRequestType)
}

// MustParseEtcdOpsRequestType converts a string to a EtcdOpsRequestType, and panics if is not valid.
func MustParseEtcdOpsRequestType(name string) EtcdOpsRequestType {
	val, err := ParseEtcdOpsRequestType(name)
	if err != nil {
		panic(err)
	}
	return val
}
