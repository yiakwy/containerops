/*
Copyright 2016 - 2017 Huawei Technologies Co., Ltd. All rights reserved.

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

package module

const (
	// Flow Run Model
	CliRun      = "CliRun"
	DaemonRun   = "DaemonRun"
	DaemonStart = "DaemonStart"

	// Stage Type
	StartStage  = "start"
	EndStage    = "end"
	NormalStage = "normal"
	PauseStage  = "pause"

	// Action Type
	Sequencing = "sequence"
	Parallel   = "parallel"

	// Result Type
	Cancel  = "cancel"
	Pending = "pending"
	Running = "running"
	Failure = "failure"
	Success = "success"
)

// Flow is DevOps orchestration flow struct.
type Flow struct {
	Model   string   `json:"-" yaml:"-"`
	URI     string   `json:"uri" yaml:"uri"`
	Number  int64    `json:",omitempty" yaml:",omitempty"`
	Title   string   `json:"title" yaml:"title"`
	Version int64    `json:"version" yaml:"version"`
	Tag     string   `json:"tag" yaml:"tag"`
	Timeout int64    `json:"timeout" yaml:"timeout"`
	Status  string   `json:"status,omitempty" yaml:"status,omitempty"`
	Logs    []string `json:"logs,omitempty" yaml:"logs,omitempty"`
	Stages  []Stage  `json:"stages,omitempty" yaml:"stages,omitempty"`
}

// Stage is
type Stage struct {
	T          string   `json:"type" yaml:"type"`
	Name       string   `json:"name" yaml:"name"`
	Title      string   `json:"title" yaml:"title"`
	Sequencing string   `json:"sequencing,omitempty" yaml:"sequencing,omitempty"`
	Status     string   `json:"status,omitempty" yaml:"status,omitempty"`
	Logs       []string `json:"logs,omitempty" yaml:"logs,omitempty"`
	Actions    []Action `json:"actions,omitempty" yaml:"actions,omitempty"`
}

// Action is
type Action struct {
	Name   string   `json:"name" yaml:"name"`
	Title  string   `json:"title" yaml:"title"`
	Status string   `json:"status,omitempty" yaml:"status,omitempty"`
	Jobs   []Job    `json:"jobs,omitempty" yaml:"jobs,omitempty"`
	Logs   []string `json:"logs,omitempty" yaml:"logs,omitempty"`
}

// Job is
type Job struct {
	T             string              `json:"type" yaml:"type"`
	Kubectl       string              `json:"kubectl" yaml:"kubectl"`
	Endpoint      string              `json:"endpoint" yaml:"endpoint"`
	Timeout       string              `json:"timeout" yaml:"timeout"`
	Status        string              `json:"status,omitempty" yaml:"status,omitempty"`
	Resources     Resource            `json:"resources" yaml:"resources"`
	Logs          []string            `json:"logs,omitempty" yaml:"logs,omitempty"`
	Environments  []map[string]string `json:"environments" yaml:"environments"`
	Outputs       []map[string]string `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	Subscriptions []map[string]string `json:"subscriptions,omitempty" yaml:"subscriptions,omitempty"`
}

// Resources is
type Resource struct {
	CPU    string `json:"cpu" yaml:"cpu"`
	Memory string `json:"memory" yaml:"memory"`
}
