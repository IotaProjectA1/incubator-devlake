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

package parser

import (
	"strings"

	"github.com/apache/incubator-devlake/core/errors"
)

type GitExtractorTaskData struct {
	Options *GitExtractorOptions
	GitRepo RepoCollector
}

type GitExtractorOptions struct {
	RepoId          string `json:"repoId" mapstructure:"repoId"`
	Name            string `json:"name" mapstructure:"name"`
	Url             string `json:"url" mapstructure:"url"`
	User            string `json:"user" mapstructure:"user"`
	Password        string `json:"password" mapstructure:"password"`
	PrivateKey      string `json:"privateKey" mapstructure:"privateKey"`
	Passphrase      string `json:"passphrase" mapstructure:"passphrase"`
	Proxy           string `json:"proxy" mapstructure:"proxy"`
	UseGoGit        bool   `json:"useGoGit" mapstructure:"useGoGit"`
	SkipCommitStat  *bool  `json:"skipCommitStat" mapstructure:"skipCommitStat" comment:"skip all commit stat including added/deleted lines and commit files as well"`
	SkipCommitFiles *bool  `json:"skipCommitFiles" mapstructure:"skipCommitFiles"`
}

func (o GitExtractorOptions) Valid() errors.Error {
	if o.RepoId == "" {
		return errors.BadInput.New("empty repoId")
	}
	if o.Url == "" {
		return errors.BadInput.New("empty url")
	}
	url := strings.TrimPrefix(o.Url, "ssh://")
	if !(strings.HasPrefix(o.Url, "http") || strings.HasPrefix(url, "git@") || strings.HasPrefix(o.Url, "/")) {
		return errors.BadInput.New("wrong url")
	}
	return nil
}
