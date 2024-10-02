// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package structs

import (
	"time"
)

// ActionTask represents a ActionTask
type ActionTask struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	HeadBranch   string `json:"head_branch"`
	HeadSHA      string `json:"head_sha"`
	RunNumber    int64  `json:"run_number"`
	Event        string `json:"event"`
	DisplayTitle string `json:"display_title"`
	Status       string `json:"status"`
	WorkflowID   string `json:"workflow_id"`
	URL          string `json:"url"`
	// swagger:strfmt date-time
	CreatedAt time.Time `json:"created_at"`
	// swagger:strfmt date-time
	UpdatedAt time.Time `json:"updated_at"`
	// swagger:strfmt date-time
	RunStartedAt time.Time `json:"run_started_at"`
}

// ActionTaskResponse returns a ActionTask
type ActionTaskResponse struct {
	Entries    []*ActionTask `json:"workflow_runs"`
	TotalCount int64         `json:"total_count"`
}

type ActionRun struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	RepoID        int64  `json:"repo_id"`
	OwnerID       int64  `json:"owner_id"`
	WorkflowID    string `json:"workflow_id"`
	TriggerUserID int64  `json:"trigger_user_id"`
	Ref           string `json:"ref"`
	CommitSHA     string `json:"commit_sha"`
	Status        string `json:"status"`
	Event         string `json:"event"`
	Version       int64  `json:"version"`
	// swagger:strfmt date-time
	StartedAt time.Time `json:"started_at"`
	// swagger:strfmt date-time
	StoppedAt time.Time `json:"stopped_at"`
	// swagger:strfmt date-time
	CreatedAt time.Time `json:"created_at"`
	// swagger:strfmt date-time
	UpdatedAt time.Time `json:"updated_at"`
}

type ActionRunResponse struct {
	Entries    []*ActionRun `json:"actions_runs"`
	TotalCount int64        `json:"total_count"`
}

type ActionTaskStep struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	// swagger:strfmt date-time
	StartedAt time.Time `json:"started_at"`
	// swagger:strfmt date-time
	StoppedAt time.Time `json:"stopped_at"`
	// swagger:strfmt date-time
	CreatedAt time.Time `json:"created_at"`
	// swagger:strfmt date-time
	UpdatedAt time.Time `json:"updated_at"`
}

type ActionTaskStepResponse struct {
	Entries    []*ActionTaskStep `json:"actions_tasks_steps"`
	TotalCount int64             `json:"total_count"`
}

type ActionRunJob struct {
	ID                int64             `json:"id"`
	RunID             int64             `json:"run_id"`
	RepoID            int64             `json:"repo_id"`
	OwnerID           int64             `json:"owner_id"`
	CommitSHA         string            `json:"commit_sha"`
	IsForkPullRequest bool              `json:"is_fork_pull_request"`
	Name              string            `json:"name"`
	Attempt           int64             `json:"attempt"`
	WorkflowPayload   []byte            `json:"workflow_payload"`
	JobID             string            `json:"job_id"`
	Needs             []string          `json:"needs"`
	RunsOn            []string          `json:"runs_on"`
	TaskID            int64             `json:"task_id"`
	Steps             []*ActionTaskStep `json:"steps"`
	Status            string            `json:"status"`
	// swagger:strfmt date-time
	StartedAt time.Time `json:"started_at"`
	// swagger:strfmt date-time
	StoppedAt time.Time `json:"stopped_at"`
	// swagger:strfmt date-time
	CreatedAt time.Time `json:"created_at"`
	// swagger:strfmt date-time
	UpdatedAt time.Time `json:"updated_at"`
}

type ActionRunJobResponse struct {
	Entries    []*ActionRunJob `json:"actions_run_jobs"`
	TotalCount int64           `json:"total_count"`
}
