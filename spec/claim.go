// Copyright Dokimasia B.V. 2026
// SPDX-License-Identifier: Apache-2.0

package spec

// ClaimID identifies a claim. It is the segment of the identity of each
// check of the claim, such as zero-on-error.
type ClaimID string

// Claim is a statement about a subject, with the law that its check
// witnesses, the probe that runs the check and the planted defect on which
// the check fails.
type Claim struct {
	ID ClaimID `json:"id"`

	// Law is the law that the check witnesses.
	Law LawID `json:"law"`

	// Statement is the sentence of the check. The generator replaces
	// {declaration}, {callable} and the name of each partner in braces.
	Statement string `json:"statement"`

	// Probe is the probe that the check calls.
	Probe ProbeID `json:"probe"`

	// Defect is the planted defect of the check.
	Defect Defect `json:"defect"`

	Since string `json:"since"`

	// Machineries are what the run supplies before the check exercises its
	// claim. The first is the deepest.
	Machineries []MachineryID `json:"machineries"`
}

func (c *Claim) key() ClaimID { return c.ID }

// Defect is the planted defect of a claim.
type Defect struct {
	// Form is the defect form that each template tree writes.
	Form DefectFormID `json:"form"`

	// Reason is the reason code with which the check must fail on the
	// defect.
	Reason ReasonID `json:"reason"`
}

var claims = section[ClaimID, Claim]("claims.json", "claims")

// Claims returns every claim in the order of claims.json, which is the order
// of the checks of a suite. The slice is shared with every other call, so a
// caller does not modify it.
func Claims() []Claim { return claims().entries }

// ClaimOf returns the claim with the identifier id, and false when the spec
// has no such claim.
func ClaimOf(id ClaimID) (Claim, bool) { return claims().of(id) }
