/*
 Copyright 2018 Padduck, LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 	http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package environments

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pufferpanel/apufferi/common"
	"github.com/pufferpanel/pufferd/cache"
	"github.com/pufferpanel/pufferd/utils"
)

var mapping map[string]EnvironmentFactory

func LoadModules() {
	mapping = make(map[string]EnvironmentFactory)

	mapping["standard"] = StandardFactory{}

	loadAdditionalModules(mapping)
}

func Create(environmentType, folder, id string, environmentSection map[string]interface{}) (Environment, error) {
	factory := mapping[environmentType]

	if factory == nil {
		return nil, errors.New(fmt.Sprintf("undefined environment: %s", environmentType))
	}

	serverRoot := common.JoinPath(folder, id)
	rootDirectory := common.GetStringOrDefault(environmentSection, "root", serverRoot)
	cache := cache.CreateCache()
	wsManager := utils.CreateWSManager()

	env := factory.Create(folder, id, environmentSection, rootDirectory, cache, wsManager)

	return env, nil
}

func GetSupportedEnvironments() []string {
	result := make([]string, len(mapping))
	i := 0
	for _, v := range mapping {
		result[i] = v.Key()
		i++
	}

	return result
}