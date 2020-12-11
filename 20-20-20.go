package main

import "github.com/gen2brain/beeep"
import "time"

func main() {
	// beeep.Notify("Schau aus dem Fenster", "Sonst werden deine Augen schlecht", "assets/icon.png")
	// beeep.A32lert("Title", "Message body", "assets/icon.png")
	freq := beeep.DefaultFreq
	dur := beeep.DefaultDuration
    ticker := time.NewTicker(20 * time.Minute)
    done := make(chan bool)
    go func() {
        for {
            select {
            case <-done:
                return
            case <-ticker.C:
				beeep.Beep(freq, dur)
            }
        }
	}()
	
	time.Sleep(10 * time.Hour)
	ticker.Stop()
	done <- true 
	beeep.Beep(freq, dur)
	time.Sleep(1* time.Second)
	beeep.Beep(freq, dur)
}
