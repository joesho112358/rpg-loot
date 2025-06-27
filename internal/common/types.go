package common

import "math/rand"

// single loot item
type Item struct {
    Name        string
    Type        string // like "gold", "gem", "art", "magic"
    Value       int    // value in gold pieces
    Description string
}

type LootTable struct {
    CR          int
    GoldRange   [2]int // min and max gold
    Items       []Item
    ItemChance  float64 // probability for roll
}

type Generator struct {
    rand *rand.Rand
}
