package house

import (
  "github.com/runningwild/haunts/base"
  "github.com/runningwild/haunts/texture"
)

// RELICS ********************************************************************
func MakeRelic(name string) *Relic {
  r := Relic{ Defname: name }
  base.GetObject("relic", &r)
  return &r
}

func GetAllRelicNames() []string {
  return base.GetAllNamesInRegistry("relic")
}

func LoadAllRelicsInDir(dir string) {
  base.RemoveRegistry("relic")
  base.RegisterRegistry("relic", make(map[string]*relicDef))
  base.RegisterAllObjectsInDir("relic", dir, ".json", "json")
}

type relicDef struct {
  Name  string
  Text  string
  Image texture.Object
}

type Relic struct {
  Defname string
  *relicDef

  // The pointer is used in the editor, but also stores the position of the
  // spawn point for use when the game is actually running.
  Pointer *Furniture  `registry:"loadfrom-furniture"`
}
func (s *Relic) Furniture() *Furniture {
  if s.Pointer == nil {
    s.Pointer = MakeFurniture("SpawnRelic")
  }
  return s.Pointer
}



// CLUES *********************************************************************
func MakeClue(name string) *Clue {
  c := Clue{ Defname: name }
  base.GetObject("clue", &c)
  return &c
}

func GetAllClueNames() []string {
  return base.GetAllNamesInRegistry("clue")
}

func LoadAllCluesInDir(dir string) {
  base.RemoveRegistry("clue")
  base.RegisterRegistry("clue", make(map[string]*clueDef))
  base.RegisterAllObjectsInDir("clue", dir, ".json", "json")
}

type clueDef struct {
  Name  string
  Text  string
  Image texture.Object
}

type Clue struct {
  Defname string
  *clueDef

  // The pointer is used in the editor, but also stores the position of the
  // spawn point for use when the game is actually running.
  Pointer *Furniture  `registry:"loadfrom-furniture"`
}
func (s *Clue) Furniture() *Furniture {
  if s.Pointer == nil {
    s.Pointer = MakeFurniture("SpawnClue")
  }
  return s.Pointer
}



// EXITS *********************************************************************
func MakeExit(name string) *Exit {
  c := Exit{ Defname: name }
  base.GetObject("exit", &c)
  return &c
}

func GetAllExitNames() []string {
  return base.GetAllNamesInRegistry("exit")
}

func LoadAllExitsInDir(dir string) {
  base.RemoveRegistry("exit")
  base.RegisterRegistry("exit", make(map[string]*exitDef))
  base.RegisterAllObjectsInDir("exit", dir, ".json", "json")
}

type exitDef struct {
  Name  string
  Text  string
  Image texture.Object
}

type Exit struct {
  Defname string
  *exitDef

  // The pointer is used in the editor, but also stores the position of the
  // spawn point for use when the game is actually running.
  Pointer *Furniture  `registry:"loadfrom-furniture"`
}
func (s *Exit) Furniture() *Furniture {
  if s.Pointer == nil {
    s.Pointer = MakeFurniture("SpawnExit")
  }
  return s.Pointer
}



// EXPLORERS *****************************************************************
func MakeExplorer(name string) *Explorer {
  c := Explorer{ Defname: name }
  base.GetObject("explorer", &c)
  return &c
}

func GetAllExplorerNames() []string {
  return base.GetAllNamesInRegistry("explorer")
}

func LoadAllExplorersInDir(dir string) {
  base.RemoveRegistry("explorer")
  base.RegisterRegistry("explorer", make(map[string]*explorerDef))
  base.RegisterAllObjectsInDir("explorer", dir, ".json", "json")
}

type explorerDef struct {
  Name  string
  Text  string
  Image texture.Object
}

type Explorer struct {
  Defname string
  *explorerDef

  // The pointer is used in the editor, but also stores the position of the
  // spawn point for use when the game is actually running.
  Pointer *Furniture  `registry:"loadfrom-furniture"`
}
func (s *Explorer) Furniture() *Furniture {
  if s.Pointer == nil {
    s.Pointer = MakeFurniture("SpawnExplorer")
  }
  return s.Pointer
}



// HAUNTS ********************************************************************
func MakeHaunt(name string) *Haunt {
  c := Haunt{ Defname: name }
  base.GetObject("haunt", &c)
  return &c
}

func GetAllHauntNames() []string {
  return base.GetAllNamesInRegistry("haunt")
}

func LoadAllHauntsInDir(dir string) {
  base.RemoveRegistry("haunt")
  base.RegisterRegistry("haunt", make(map[string]*hauntDef))
  base.RegisterAllObjectsInDir("haunt", dir, ".json", "json")
}

type hauntDef struct {
  Name  string
  Text  string
  Image texture.Object
}

type Haunt struct {
  Defname string
  *hauntDef

  // The pointer is used in the editor, but also stores the position of the
  // spawn point for use when the game is actually running.
  Pointer *Furniture  `registry:"loadfrom-furniture"`
}
func (s *Haunt) Furniture() *Furniture {
  if s.Pointer == nil {
    s.Pointer = MakeFurniture("SpawnHaunt")
  }
  return s.Pointer
}



type spawnError struct {
  msg string
}
func (se *spawnError) Error() string {
  return se.msg
}

// func verifyRelicSpawns(h *HouseDef) error {
//   total := 0
//   for i := range h.Floors {
//     total += len(h.Floors[i].Relics)
//   }
//   if total < 5 {
//     return &spawnError{ "House needs at least five relic spawn points." }
//   }
//   return nil
// }

// func verifyPlayerSpawns(h *HouseDef) error {
//   total := 0
//   for i := range h.Floors {
//     total += len(h.Floors[i].Players)
//   }
//   if total < 1 {
//     return &spawnError{ "House needs at least one player spawn point." }
//   }
//   return nil
// }

// func verifyCleanseSpawns(h *HouseDef) error {
//   total := 0
//   for i := range h.Floors {
//     total += len(h.Floors[i].Cleanse)
//   }
//   if total < 3 {
//     return &spawnError{ "House needs at least cleanse spawn points." }
//   }
//   return nil
// }

// func verifyClueSpawns(h *HouseDef) error {
//   total := 0
//   for i := range h.Floors {
//     total += len(h.Floors[i].Clues)
//   }
//   if total < 10 {
//     return &spawnError{ "House needs at least ten clue spawn points." }
//   }
//   return nil
// }

// func verifyExitSpawns(h *HouseDef) error {
//   total := 0
//   for i := range h.Floors {
//     total += len(h.Floors[i].Exits)
//   }
//   if total < 1 {
//     return &spawnError{ "House needs at least one exit spawn point." }
//   }
//   return nil
// }