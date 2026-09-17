package model

import "time"

type ClaimVerdict string

const (
	ClaimProcessing   ClaimVerdict = "PROCESSING"
	ClaimRefuted      ClaimVerdict = "REFUTED"
	ClaimInconclusive ClaimVerdict = "INCONCLUSIVE"
	ClaimSupported    ClaimVerdict = "SUPPORTED"
)

type Claim struct {
	ID        int64        `db:"id" json:"id"`
	Content   string       `db:"content" json:"content"`
	Verdict   ClaimVerdict `db:"verdict" json:"verdict"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
}
