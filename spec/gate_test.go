// Copyright Dokimasia B.V. 2026
// SPDX-License-Identifier: Apache-2.0

package spec_test

import (
	"testing"

	"go.dokimi.dev/assert"

	"github.com/dokimasia/dokimi-spec/spec"
)

func TestGate(t *testing.T) {
	t.Parallel()

	t.Run("Gates", func(t *testing.T) {
		t.Parallel()

		t.Run("returns the gates of gates.json in the order of the file", func(t *testing.T) {
			t.Parallel()

			assert.Equal(
				t,
				keysOf(spec.Gates(), func(g spec.Gate) spec.GateID { return g.ID }),
				idsOf(t, "gates", "gates"),
				"the gates keep the order of gates.json",
			)
		})

		t.Run("returns gates that the schema accepts", func(t *testing.T) {
			t.Parallel()

			validates(t, "gates")
		})
	})

	t.Run("GateOf", func(t *testing.T) {
		t.Parallel()

		t.Run("returns the gate with the identifier", func(t *testing.T) {
			t.Parallel()

			got, ok := spec.GateOf(spec.ReturnsBesideError)

			assert.True(t, ok, "the spec has the gate")
			assert.Equal(t, got, spec.Gate{ID: spec.ReturnsBesideError, Since: "1.0.0"},
				"the gate has the version that added it")
		})

		t.Run("returns false for an identifier without a gate", func(t *testing.T) {
			t.Parallel()

			_, ok := spec.GateOf("no-such-gate")

			assert.False(t, ok, "the spec has no such gate")
		})
	})
}
