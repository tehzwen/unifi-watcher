package main

import (
	"fmt"
	"log"
	"time"

	unifi_watcher "github.com/tehzwen/unifi-watcher"
)

func main() {
	// use defaults
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

	// custom config
	cw := unifi_watcher.NewUnifiWatcher(
		unifi_watcher.WithCustomConnString("fake-conn-string"),
		unifi_watcher.WithCustomQuery("fake custom sql query"),
		unifi_watcher.WithPollFrequency(time.Second*5),
	)

	if err := cw.Watch(func(e unifi_watcher.UnifiEvent) {
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
