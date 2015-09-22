package actions_test

import (
  "path/filepath"
  "encoding/gob"
  "bytes"
  "github.com/orfjackal/gospec/src/gospec"
  . "github.com/orfjackal/gospec/src/gospec"
  "github.com/mik3cap/haunts/base"
  "github.com/mik3cap/haunts/game"
  "github.com/mik3cap/haunts/game/actions"
)

var datadir string

func init() {
  datadir, _ = filepath.Abs("../../data_test")
  base.SetDatadir(datadir)
}

func ActionSpec(c gospec.Context) {
  game.RegisterActions()
  c.Specify("Actions are loaded properly.", func() {
    basic := game.MakeAction("Basic Test")
    _, ok := basic.(*actions.BasicAttack)
    c.Expect(ok, Equals, true)
  })

  c.Specify("Actions can be gobbed without loss of type.", func() {
    buf := bytes.NewBuffer(nil)
    enc := gob.NewEncoder(buf)

    var as []game.Action
    as = append(as, game.MakeAction("Move Test"))
    as = append(as, game.MakeAction("Basic Test"))

    err := enc.Encode(as)
    c.Assume(err, Equals, nil)

    dec := gob.NewDecoder(buf)
    var as2 []game.Action
    err = dec.Decode(&as2)
    c.Assume(err, Equals, nil)

    _, ok := as2[0].(*actions.Move)
    c.Expect(ok, Equals, true)

    _, ok = as2[1].(*actions.BasicAttack)
    c.Expect(ok, Equals, true)

  })
}
