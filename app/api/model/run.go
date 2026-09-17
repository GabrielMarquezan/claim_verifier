package model

import "time"

type RunStatus string

const (
	RunPending    RunStatus = "PENDING"
	RunRunning    RunStatus = "RUNNING"
	RunSuccessful RunStatus = "SUCCESSFUL"
	RunFailed     RunStatus = "FAILED"
)

type Run struct {
	ID         int64      `db:"id" json:"id"`
	ClaimID    int64      `db:"claim_id" json:"claim_id"`
	Status     RunStatus  `db:"status" json:"status"`
	ModelID    string     `db:"model_id" json:"model_id"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	FinishedAt *time.Time `db:"finished_at" json:"finished_at,omitempty"`
}
