// Copyright Dokimasia B.V. 2026
// SPDX-License-Identifier: Apache-2.0

package spec_test

import (
	"testing"

	"go.dokimi.dev/assert"

	"github.com/dokimasia/dokimi-spec/spec"
)

func TestStance(t *testing.T) {
	t.Parallel()

	t.Run("Stances", func(t *testing.T) {
		t.Parallel()

		t.Run("returns the stances of stances.json in the order of the file", func(t *testing.T) {
			t.Parallel()

			assert.Equal(
				t,
				keysOf(spec.Stances(), func(s spec.Stance) spec.StanceID { return s.ID }),
				idsOf(t, "stances", "stances"),
				"the stances keep the order of stances.json",
			)
		})

		t.Run("returns stances that the schema accepts", func(t *testing.T) {
			t.Parallel()

			validates(t, "stances")
		})
	})

	t.Run("StanceOf", func(t *testing.T) {
		t.Parallel()

		t.Run("returns the stance with the identifier", func(t *testing.T) {
			t.Parallel()

			got, ok := spec.StanceOf(spec.Needs)

			assert.True(t, ok, "the spec has the stance")
			assert.Equal(
				t,
				got.Closes,
				"The param machineries on the declaration, or a runtime that supplies the machinery.",
				"the stance has what closes it",
			)
		})

		t.Run("returns false for an identifier without a stance", func(t *testing.T) {
			t.Parallel()

			_, ok := spec.StanceOf("no-such-stance")

			assert.False(t, ok, "the spec has no such stance")
		})
	})
}
