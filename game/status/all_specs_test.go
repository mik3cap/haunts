package status_test

import (
  "testing"
  "github.com/orfjackal/gospec/src/gospec"
  "github.com/mik3cap/haunts/game/status"
)

func TestAllSpecs(t *testing.T) {
  status.RegisterAllConditions()
  r := gospec.NewRunner()
  r.AddSpec(ConditionsSpec)
  gospec.MainGoTest(r, t)
}
