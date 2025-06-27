package common

import (
	"math/rand"
	"time"
)

// NewGenerator initializes a random number generator and returns the address
func NewGenerator(seed int64) *Generator {
	if seed == 0 {
		// seed with current time as default
        seed = time.Now().UnixNano()
    }
    return &Generator{
        rand: rand.New(
            rand.NewSource(seed),
        ),
    }
}

// GenerateIndividualLoot acts on a Generator and generates loot for a single 
// encounter using the `table` parameter
func (g *Generator) GenerateIndividualLoot(table LootTable) (int, []Item) {
	// roll for gold amount
    gold := g.rand.Intn(table.GoldRange[1]-table.GoldRange[0]) + table.GoldRange[0]
    
    var items []Item

	// roll to see if we get an item
    if g.rand.Float64() < table.ItemChance {
		// roll to see which item we get
        itemIndex := g.rand.Intn(len(table.Items))
        items = append(items, table.Items[itemIndex])
    }
    
    return gold, items
}

// GenerateHoard acts on a Generator and generates loot for a hoard 
// encounter using the `table` parameter
func (g *Generator) GenerateHoard(table LootTable) (int, []Item) {
	// roll for gold amount
	gold := g.rand.Intn(table.GoldRange[1]*2-table.GoldRange[0]) + table.GoldRange[0]
    
    var items []Item
	// we get an item! ... at random!
    itemIndex := g.rand.Intn(len(table.Items))
    items = append(items, table.Items[itemIndex])
    
	// roll to see how many more items we might get! (0-2 items)
    numItems := g.rand.Intn(3)
    for range numItems {
		// roll to see if we get an item
        if g.rand.Float64() < table.ItemChance {
            // roll to see which item
            itemIndex := g.rand.Intn(len(table.Items))
            items = append(items, table.Items[itemIndex])
        }
    }
    
    return gold, items
}
