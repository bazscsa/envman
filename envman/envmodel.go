package envman

import (
	"strconv"

	log "github.com/Sirupsen/logrus"
)

const (
	// IsExpandKey ...
	IsExpandKey string = "is_expand"
	// TrueKey ...
	TrueKey string = "true"
	// FalseKey ...
	FalseKey string = "false"
)

// EnvModel ... This is the model of ENVIRONMENT in envman, for methods
type EnvModel struct {
	Key      string
	Value    string
	IsExpand bool
}

// EnvMapItem ... This is the model of ENVIRONMENT in envman, for storing in file
type EnvMapItem map[string]string

type envsYMLModel struct {
	Envs []EnvMapItem `yml:"environments"`
}

// Convert envsYMLModel to envModel array
func (envYML envsYMLModel) convertToEnvModelArray() []EnvModel {
	var envModels []EnvModel
	for _, envMapItem := range envYML.Envs {
		envModel := envMapItem.convertToEnvModel()
		envModels = append(envModels, envModel)
	}
	return envModels
}

func (eMap EnvMapItem) convertToEnvModel() EnvModel {
	var eModel EnvModel
	for key, value := range eMap {
		if key != IsExpandKey {
			eModel.Key = key
			eModel.Value = value
		}
	}
	eModel.IsExpand = ParseBool(eMap[IsExpandKey])
	return eModel
}

// ParseBool ...
func ParseBool(s string) bool {
	if s == "" {
		return true
	}

	expand, err := strconv.ParseBool(s)
	if err != nil {
		log.Errorln("[ENVMAN] - isExpand: Failed to parse input:", err)
		return true
	}
	return expand
}

// Convert envModel array to envsYMLModel
func convertToEnvsYMLModel(eModels []EnvModel) envsYMLModel {
	var envYML envsYMLModel
	var envMaps []EnvMapItem
	for _, eModel := range eModels {
		eMap := eModel.convertToEnvMap()
		envMaps = append(envMaps, eMap)
	}
	envYML.Envs = envMaps
	return envYML
}

func (eModel EnvModel) convertToEnvMap() EnvMapItem {
	eMap := make(EnvMapItem)
	if eModel.IsExpand == false {
		eMap[IsExpandKey] = FalseKey
	}
	eMap[eModel.Key] = eModel.Value
	return eMap
}