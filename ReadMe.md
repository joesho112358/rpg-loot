# RPG Loot

This is a tool I built to use and learn a bit of Golang along the way.

How to use:
```
./rpg-loot --cr 5 --system dnd
./rpg-loot --cr 5 --system pathfinder

./rpg-loot --cr 5 --system dnd --hoard
./rpg-loot --cr 5 --system pathfinder --hoard
```

### Data file structure
```
[
    {
        "CR": >0,
        "GoldRange": [ > 0, < first entry ],
        "ItemChance": a probability (0 <= p <= 1),
        "Items": [
            {
                "name": "non-empty",
                "type": "non-empty",
                "value": > 0,
                "description": "non-empty"
            },
        ]
    }
]
```

`./rpg-loot validate --file data/dnd_loot_tables.json` to validate your json file
or there is a script to validate  in the scripts dir!
