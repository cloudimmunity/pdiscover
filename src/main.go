package main

import (
	"log"
    "pdiscover"
)

func main() {
	watcher, err := pdiscover.NewAllWatcher(pdiscover.PROC_EVENT_ALL)
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        for {
            select {
            case ev := <-watcher.Fork:
                log.Println("fork event:", ev)
            case ev := <-watcher.Exec:
                log.Println("exec event:", ev)
                log.Printf("exec event: e => %#v p => %#v\n",ev,pdiscover.GetProcInfo(ev.Pid))
            case ev := <-watcher.Exit:
                log.Println("exit event:", ev)
            case err := <-watcher.Error:
                log.Println("error:", err)
            }
        }
    }()
    
    err = watcher.WatchAll()
    if err != nil {
        log.Fatal(err)
    }

    done := make(chan bool)

    <- done
    //...
    watcher.Close()
}
