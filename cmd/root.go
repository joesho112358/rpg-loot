package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/joesho112358/rpg-loot/internal/common"
	"github.com/spf13/cobra"
	"os"
)

var cr int
var hoard bool
var system string

var tableFile string

var rootCmd = &cobra.Command{
	Use:   "rpg-loot",
	Short: "A CLI tool for generating RPG loot in Go!",
	Run: func(cmd *cobra.Command, args []string) {
		// we want gold and items
		var gold int
		var items []common.Item
		// we need a random number generator
		g := common.NewGenerator(0)
		// we need to know where to look
		var tables []common.LootTable
		var defaultFile string
		var message string

		// set defaults based on system
		switch system {
		case "dnd":
			defaultFile = "data/dnd_loot_tables.json"
			message = "D&D"
		case "pathfinder":
			defaultFile = "data/pathfinder_loot_tables.json"
			message = "Pathfinder"
		default:
			fmt.Println("Error: Invalid system. Use 'dnd' or 'pathfinder'.")
			return
		}

        fileToLoad := tableFile
        if fileToLoad == "" {
            fileToLoad = defaultFile
        }

        // load and validate JSON file
        data, err := os.ReadFile(fileToLoad)
        if err != nil {
            fmt.Printf("Error reading %s: %v\n", fileToLoad, err)
            os.Exit(1)
        }
        if err := json.Unmarshal(data, &tables); err != nil {
            fmt.Printf("Error parsing JSON in %s: %v\n", fileToLoad, err)
            os.Exit(1)
        }
        if len(tables) == 0 {
            fmt.Printf("Error: %s contains no tables\n", fileToLoad)
            os.Exit(1)
        }
        for i, tbl := range tables {
            if err := validateLootTable(tbl, i); err != nil {
                fmt.Printf("Validation error in %s: %v\n", fileToLoad, err)
                os.Exit(1)
            }
        }

		table := getTableForCR(tables, cr)

		if hoard {
			gold, items = g.GenerateHoard(table)
			message += " Treasure Hoard:"
		} else {
			gold, items = g.GenerateIndividualLoot(table)
			message += " Individual Loot:"
		}

		fmt.Println(message)
		fmt.Printf("Gold: %d gp\n", gold)
		if len(items) > 0 {
			fmt.Println("Items:")
			// ignoring the index here
			for _, item := range items {
				fmt.Printf("- %s (%s, %d gp): %s\n", item.Name, item.Type, item.Value, item.Description)
			}
		} else {
			fmt.Println("No items found.")
		}
	},
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a JSON loot table file",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(tableFile)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", tableFile, err)
			return
		}

		var tables []common.LootTable
		if err := json.Unmarshal(data, &tables); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return
		}

		if len(tables) == 0 {
			fmt.Println("Error: JSON file contains no tables")
			return
		}
		for i, table := range tables {
			if err := validateLootTable(table, i); err != nil {
				fmt.Printf("Validation error: %v\n", err)
				return
			}
		}

		fmt.Printf("JSON file %s is valid!\n", tableFile)
	},
}

func validateLootTable(table common.LootTable, index int) error {
	if table.CR < 0 {
		return fmt.Errorf("table %d: CR (%d) must be non-negative", index, table.CR)
	}
	if table.GoldRange[0] < 0 || table.GoldRange[1] < table.GoldRange[0] {
		return fmt.Errorf("table %d: invalid GoldRange [%d, %d]", index, table.GoldRange[0], table.GoldRange[1])
	}
	if table.ItemChance < 0.0 || table.ItemChance > 1.0 {
		return fmt.Errorf("table %d: ItemChance (%f) must be between 0.0 and 1.0", index, table.ItemChance)
	}
	if len(table.Items) == 0 {
		return fmt.Errorf("table %d: Items list cannot be empty", index)
	}
	for i, item := range table.Items {
		if item.Name == "" {
			return fmt.Errorf("table %d, item %d: Name cannot be empty", index, i)
		}
		if item.Type == "" {
			return fmt.Errorf("table %d, item %d: Type cannot be empty", index, i)
		}
		if item.Value < 0 {
			return fmt.Errorf("table %d, item %d: Value (%d) must be non-negative", index, i, item.Value)
		}
		if item.Description == "" {
			return fmt.Errorf("table %d, item %d: Description cannot be empty", index, i)
		}
	}
	return nil
}

func getTableForCR(tables []common.LootTable, cr int) common.LootTable {
	if len(tables) == 0 {
		fmt.Println("Error: No valid tables available")
		os.Exit(1)
	}
	for _, table := range tables {
		if cr <= table.CR {
			return table
		}
	}
	return tables[len(tables)-1]
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// flags for root command
	rootCmd.Flags().IntVarP(&cr, "cr", "c", 0, "Challenge Rating for loot")
	rootCmd.Flags().BoolVarP(&hoard, "hoard", "r", false, "Generate treasure for a hoard of enemies")
	rootCmd.Flags().StringVarP(&system, "system", "s", "dnd", "RPG system (dnd or pathfinder)")
	rootCmd.Flags().StringVarP(&tableFile, "table-file", "t", "", "Path to custom JSON loot table file")

	// add validate command with file flag
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringVarP(&tableFile, "file", "f", "", "Path to JSON loot table file")
	validateCmd.MarkFlagRequired("file")
}
