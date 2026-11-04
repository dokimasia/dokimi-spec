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

func TestLaw(t *testing.T) {
	t.Parallel()

	t.Run("Laws", func(t *testing.T) {
		t.Parallel()

		t.Run("returns the laws of laws.json in the order of the file", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, keysOf(spec.Laws(), func(l spec.Law) spec.LawID { return l.ID }), idsOf(t, "laws", "laws"),
				"the laws keep the order of laws.json")
		})

		t.Run("returns laws that the schema accepts", func(t *testing.T) {
			t.Parallel()

			validates(t, "laws")
		})

		t.Run("returns laws that a claim witnesses", func(t *testing.T) {
			t.Parallel()
			witnessed := make([]spec.LawID, 0, len(spec.Claims()))
			for _, c := range spec.Claims() {
				witnessed = append(witnessed, c.Law)
			}

			for _, l := range spec.Laws() {
				expect.True(t, slices.Contains(witnessed, l.ID), "a claim witnesses the law "+string(l.ID))
			}
		})
	})

	t.Run("LawOf", func(t *testing.T) {
		t.Parallel()

		t.Run("returns the law with the identifier", func(t *testing.T) {
			t.Parallel()

			got, ok := spec.LawOf(spec.NoValueBesideAnError)

			assert.True(t, ok, "the spec has the law")
			assert.Equal(t, got.Statement, "A call that fails returns the zero value of each value result.",
				"the law has its statement")
		})

		t.Run("returns false for an identifier without a law", func(t *testing.T) {
			t.Parallel()

			_, ok := spec.LawOf("no-such-law")

			assert.False(t, ok, "the spec has no such law")
		})
	})
}
