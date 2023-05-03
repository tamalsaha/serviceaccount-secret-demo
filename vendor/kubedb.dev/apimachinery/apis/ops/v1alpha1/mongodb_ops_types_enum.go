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
	// MongoDBOpsRequestTypeUpgrade is a MongoDBOpsRequestType of type Upgrade.
	MongoDBOpsRequestTypeUpgrade MongoDBOpsRequestType = "Upgrade"
	// MongoDBOpsRequestTypeUpdateVersion is a MongoDBOpsRequestType of type UpdateVersion.
	MongoDBOpsRequestTypeUpdateVersion MongoDBOpsRequestType = "UpdateVersion"
	// MongoDBOpsRequestTypeHorizontalScaling is a MongoDBOpsRequestType of type HorizontalScaling.
	MongoDBOpsRequestTypeHorizontalScaling MongoDBOpsRequestType = "HorizontalScaling"
	// MongoDBOpsRequestTypeVerticalScaling is a MongoDBOpsRequestType of type VerticalScaling.
	MongoDBOpsRequestTypeVerticalScaling MongoDBOpsRequestType = "VerticalScaling"
	// MongoDBOpsRequestTypeVolumeExpansion is a MongoDBOpsRequestType of type VolumeExpansion.
	MongoDBOpsRequestTypeVolumeExpansion MongoDBOpsRequestType = "VolumeExpansion"
	// MongoDBOpsRequestTypeRestart is a MongoDBOpsRequestType of type Restart.
	MongoDBOpsRequestTypeRestart MongoDBOpsRequestType = "Restart"
	// MongoDBOpsRequestTypeReconfigure is a MongoDBOpsRequestType of type Reconfigure.
	MongoDBOpsRequestTypeReconfigure MongoDBOpsRequestType = "Reconfigure"
	// MongoDBOpsRequestTypeReconfigureTLS is a MongoDBOpsRequestType of type ReconfigureTLS.
	MongoDBOpsRequestTypeReconfigureTLS MongoDBOpsRequestType = "ReconfigureTLS"
	// MongoDBOpsRequestTypeReprovision is a MongoDBOpsRequestType of type Reprovision.
	MongoDBOpsRequestTypeReprovision MongoDBOpsRequestType = "Reprovision"
)

var ErrInvalidMongoDBOpsRequestType = fmt.Errorf("not a valid MongoDBOpsRequestType, try [%s]", strings.Join(_MongoDBOpsRequestTypeNames, ", "))

var _MongoDBOpsRequestTypeNames = []string{
	string(MongoDBOpsRequestTypeUpgrade),
	string(MongoDBOpsRequestTypeUpdateVersion),
	string(MongoDBOpsRequestTypeHorizontalScaling),
	string(MongoDBOpsRequestTypeVerticalScaling),
	string(MongoDBOpsRequestTypeVolumeExpansion),
	string(MongoDBOpsRequestTypeRestart),
	string(MongoDBOpsRequestTypeReconfigure),
	string(MongoDBOpsRequestTypeReconfigureTLS),
	string(MongoDBOpsRequestTypeReprovision),
}

// MongoDBOpsRequestTypeNames returns a list of possible string values of MongoDBOpsRequestType.
func MongoDBOpsRequestTypeNames() []string {
	tmp := make([]string, len(_MongoDBOpsRequestTypeNames))
	copy(tmp, _MongoDBOpsRequestTypeNames)
	return tmp
}

// MongoDBOpsRequestTypeValues returns a list of the values for MongoDBOpsRequestType
func MongoDBOpsRequestTypeValues() []MongoDBOpsRequestType {
	return []MongoDBOpsRequestType{
		MongoDBOpsRequestTypeUpgrade,
		MongoDBOpsRequestTypeUpdateVersion,
		MongoDBOpsRequestTypeHorizontalScaling,
		MongoDBOpsRequestTypeVerticalScaling,
		MongoDBOpsRequestTypeVolumeExpansion,
		MongoDBOpsRequestTypeRestart,
		MongoDBOpsRequestTypeReconfigure,
		MongoDBOpsRequestTypeReconfigureTLS,
		MongoDBOpsRequestTypeReprovision,
	}
}

// String implements the Stringer interface.
func (x MongoDBOpsRequestType) String() string {
	return string(x)
}

// String implements the Stringer interface.
func (x MongoDBOpsRequestType) IsValid() bool {
	_, err := ParseMongoDBOpsRequestType(string(x))
	return err == nil
}

var _MongoDBOpsRequestTypeValue = map[string]MongoDBOpsRequestType{
	"Upgrade":           MongoDBOpsRequestTypeUpgrade,
	"UpdateVersion":     MongoDBOpsRequestTypeUpdateVersion,
	"HorizontalScaling": MongoDBOpsRequestTypeHorizontalScaling,
	"VerticalScaling":   MongoDBOpsRequestTypeVerticalScaling,
	"VolumeExpansion":   MongoDBOpsRequestTypeVolumeExpansion,
	"Restart":           MongoDBOpsRequestTypeRestart,
	"Reconfigure":       MongoDBOpsRequestTypeReconfigure,
	"ReconfigureTLS":    MongoDBOpsRequestTypeReconfigureTLS,
	"Reprovision":       MongoDBOpsRequestTypeReprovision,
}

// ParseMongoDBOpsRequestType attempts to convert a string to a MongoDBOpsRequestType.
func ParseMongoDBOpsRequestType(name string) (MongoDBOpsRequestType, error) {
	if x, ok := _MongoDBOpsRequestTypeValue[name]; ok {
		return x, nil
	}
	return MongoDBOpsRequestType(""), fmt.Errorf("%s is %w", name, ErrInvalidMongoDBOpsRequestType)
}

// MustParseMongoDBOpsRequestType converts a string to a MongoDBOpsRequestType, and panics if is not valid.
func MustParseMongoDBOpsRequestType(name string) MongoDBOpsRequestType {
	val, err := ParseMongoDBOpsRequestType(name)
	if err != nil {
		panic(err)
	}
	return val
}
