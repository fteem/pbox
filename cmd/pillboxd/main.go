package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fteem/pbox"
	"github.com/martinlindhe/notify"
)

var (
	morning = pillbox.Checkpoint{
		Hour:   10,
		Minute: 00,
	}
	afternoon = pillbox.Checkpoint{
		Hour:   16,
		Minute: 00,
	}
	evening = pillbox.Checkpoint{
		Hour:   20,
		Minute: 00,
	}
)

func main() {
	notify.Notify("💊 Pillbox", "", "Hey there! 👋 I'll stick around and remind you when it's time to take your meds!", "assets/information.png")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Exiting...")
		os.Exit(1)
	}()

	for range time.Tick(30 * time.Second) {
		store := pillbox.NewStore("pillbox.db")
		if err := store.Open(); err != nil {
			fmt.Println(err)
		}

		reminders, err := store.FetchReminders()
		if err != nil {
			panic(err)
		}

		for _, reminder := range reminders {
			if reminder.Morning && morning.Equal(time.Now()) ||
				reminder.Afternoon && afternoon.Equal(time.Now()) ||
				reminder.Evening && evening.Equal(time.Now()) {

				alert(reminder.Body)
			}

		}
		store.Close()
	}
}

func alert(body string) {
	notify.Alert("💊 Pillbox", "", fmt.Sprintf("Time to get a %s", body), "assets/information.png")
}
