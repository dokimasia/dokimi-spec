// Copyright Dokimasia B.V. 2026
// SPDX-License-Identifier: Apache-2.0

package spec_test

import (
	"slices"
	"testing"

	"go.dokimi.dev/assert"
	"go.dokimi.dev/assert/expect"

	"github.com/dokimasia/dokimi-spec/spec"
)

func TestClaim(t *testing.T) {
	t.Parallel()

	t.Run("Claims", func(t *testing.T) {
		t.Parallel()

		t.Run("returns the claims of claims.json in the order of the file", func(t *testing.T) {
			t.Parallel()

			assert.Equal(
				t,
				keysOf(spec.Claims(), func(c spec.Claim) spec.ClaimID { return c.ID }),
				idsOf(t, "claims", "claims"),
				"the claims keep the order of claims.json",
			)
		})

		t.Run("returns claims that the schema accepts", func(t *testing.T) {
			t.Parallel()

			validates(t, "claims")
		})

		t.Run("returns claims whose law, probe, machineries and defect the spec defines", func(t *testing.T) {
			t.Parallel()

			for _, c := range spec.Claims() {
				id := string(c.ID)
				_, ok := spec.LawOf(c.Law)
				expect.True(t, ok, "the spec has the law of claim "+id)
				p, ok := spec.ProbeOf(c.Probe)
				expect.True(t, ok, "the spec has the probe of claim "+id)
				expect.True(t, len(c.Machineries) > 0, "claim "+id+" needs a machinery")
				for _, m := range c.Machineries {
					_, found := spec.MachineryOf(m)
					expect.True(t, found, "the spec has the machinery "+string(m)+" of claim "+id)
				}
				_, ok = spec.DefectFormOf(c.Defect.Form)
				expect.True(t, ok, "the spec has the defect form of claim "+id)
				expect.True(
					t,
					slices.Contains(p.Reasons, c.Defect.Reason),
					"the probe of claim "+id+" reports the reason code of its defect",
				)
			}
		})
	})

	t.Run("ClaimOf", func(t *testing.T) {
		t.Parallel()

		t.Run("returns the claim with the identifier", func(t *testing.T) {
			t.Parallel()

			got, ok := spec.ClaimOf(spec.ReadAfterWrite)

			assert.True(t, ok, "the spec has the claim")
			assert.Equal(t, got, spec.Claim{
				ID:          spec.ReadAfterWrite,
				Law:         spec.ReadReturnsTheWriteBeforeIt,
				Machineries: []spec.MachineryID{spec.NothingBeyondACaller},
				Statement:   "{callable} returns what {write} stored before it",
				Probe:       spec.WriteThenRead,
				Defect:      spec.Defect{Form: spec.ReadsAnotherValue, Reason: spec.UnobservedWrite},
				Since:       "1.0.0",
			}, "the claim has the fields of the runtime contract")
		})

		t.Run("returns false for an identifier without a claim", func(t *testing.T) {
			t.Parallel()

			_, ok := spec.ClaimOf("no-such-claim")

			assert.False(t, ok, "the spec has no such claim")
		})
	})
}
