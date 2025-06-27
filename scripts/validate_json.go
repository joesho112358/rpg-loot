package main

import (
    "encoding/json"
    "fmt"
    "os"
    "github.com/joesho112358/rpg-loot/internal/common"
)

// validateLootTable checks if a LootTable meets basic rules
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

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Error: No file path provided. Usage: go run validate_json.go <file_path>")
        os.Exit(1)
    }
    filePath := os.Args[1]

	data, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Printf("Error reading %s: %v\n", filePath, err)
        os.Exit(1)
    }

    var tables []common.LootTable
    if err := json.Unmarshal(data, &tables); err != nil {
        fmt.Printf("Error parsing JSON: %v\n", err)
        os.Exit(1)
    }

    if len(tables) == 0 {
        fmt.Println("Error: JSON file contains no tables")
        os.Exit(1)
    }
    for i, table := range tables {
        if err := validateLootTable(table, i); err != nil {
            fmt.Printf("Validation error: %v\n", err)
            os.Exit(1)
        }
    }

    fmt.Printf("JSON file %s is valid!\n", filePath)
}
