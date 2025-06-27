#!/bin/bash
go run scripts/validate_json.go data/dnd_loot_tables.json
if [ $? -ne 0 ]; then
    echo "D&D JSON validation failed"
    exit 1
fi
go run scripts/validate_json.go data/pathfinder_loot_tables.json
if [ $? -ne 0 ]; then
    echo "Pathfinder JSON validation failed"
    exit 1
fi
echo "All JSON files valid"
