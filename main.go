package main

import (
    "github.com/joesho112358/rpg-loot/cmd"
    "log"
)

func main() {
	// note to me: err is in scope only in this if statement
	// (not that it matters here), but good to know!
    if err := cmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
