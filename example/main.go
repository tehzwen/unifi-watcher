package main

import (
	"fmt"
	"log"

	unifi_watcher "github.com/tehzwen/unifi-watcher"
)

func main() {
	w := unifi_watcher.NewUnifiWatcher()
	if err := w.Watch(func(e unifi_watcher.UnifiEvent) {
		// I just want to print the event as it comes in
		if e.SmartDetectTypes != nil {
			fmt.Printf("Type %s, SmartDetectTypes: %s\n", e.Type, *e.SmartDetectTypes)
		} else {
			fmt.Println(e)
		}
	}); err != nil {
		log.Fatal(err)
	}
}
