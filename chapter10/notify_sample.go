package main

import (
	"log"

	"gopkg.in/fsnotify.v1"
)

func notifyMain() {
	counter := 0
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(nil)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				// ANDのビット演算だこれ
				if event.Op & fsnotify.Create == fsnotify.Create {
					log.Println("created file:", event.Name)
					counter++
				} else if event.Op & fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					counter++
				} else if event.Op & fsnotify.Remove == fsnotify.Remove {
					log.Println("removed file:", event.Name)
					counter++
				} else if event.Op & fsnotify.Rename == fsnotify.Rename {
					log.Println("renamed file:", event.Name)
					counter++
				} else if event.Op & fsnotify.Chmod == fsnotify.Chmod {
					log.Println("chmod file:", event.Name)
					counter++
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
			if counter >3 {
				done<-true
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

// func main() {
// 	notifyMain()
// }