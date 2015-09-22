// Stubbed version of the sound manager - lets us test things without having
// to link in fmod.

// +build nosound
package sound

import (
  "github.com/mik3cap/glop/sprite"
  "github.com/mik3cap/haunts/base"
)

func Init()                                 {}
func MapSounds(m map[string]string)         {}
func trigger(s *sprite.Sprite, name string) {}
func PlaySound(p base.Path)                 {}
func SetBackgroundMusic(file string)        {}
