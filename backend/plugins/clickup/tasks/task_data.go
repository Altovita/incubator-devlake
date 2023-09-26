/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tasks

import (
	"time"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	api "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

type ClickupApiParams struct {
	TeamId  string
	SpaceId string
}

type ClickupOptions struct {
	ConnectionId     uint64   `json:"connectionId"`
	ScopeId          string   `json:"scopeId"`
	Tasks            []string `json:"tasks,omitempty"`
	CreatedDateAfter string   `json:"createdDateAfter" mapstructure:"createdDateAfter,omitempty"`
}

type ClickupTaskData struct {
	Options          *ClickupOptions
	ApiClient        *api.ApiAsyncClient
	CreatedDateAfter *time.Time
	TeamId           string
}

func DecodeAndValidateTaskOptions(options map[string]interface{}) (*ClickupOptions, errors.Error) {
	var op ClickupOptions
	if err := helper.Decode(options, &op, nil); err != nil {
		return nil, err
	}
	return &op, nil
}

func CreateRawDataSubTaskArgs(taskCtx plugin.SubTaskContext, rawTable string) (*api.RawDataSubTaskArgs, *ClickupTaskData) {
	data := taskCtx.GetData().(*ClickupTaskData)
	filteredData := *data
	filteredData.Options = &ClickupOptions{}
	*filteredData.Options = *data.Options
	var params = ClickupApiParams{
		TeamId:  data.TeamId,
		SpaceId: data.Options.ScopeId,
	}
	rawDataSubTaskArgs := &api.RawDataSubTaskArgs{
		Ctx:    taskCtx,
		Params: params,
		Table:  rawTable,
	}
	return rawDataSubTaskArgs, &filteredData
}
