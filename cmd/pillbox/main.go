package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fteem/pbox"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("pillbox", "Simple medication reminders")

	// Installation
	install   = app.Command("install", "Install the service")
	uninstall = app.Command("uninstall", "Uninstall the service")

	// Reminders
	reminder          = app.Command("reminders", "Reminders operations")
	reminderMorning   = reminder.Flag("morning", "Run in the morning").Short('m').Bool()
	reminderAfternoon = reminder.Flag("afternoon", "Run in the morning").Short('a').Bool()
	reminderEvening   = reminder.Flag("evening", "Run in the morning").Short('e').Bool()

	reminderAdd              = reminder.Command("add", "Add reminder")
	reminderAddMedication    = reminderAdd.Arg("medication", "Medication name").Required().String()
	reminderRemove           = reminder.Command("remove", "Remove reminder")
	reminderRemoveMedication = reminderRemove.Arg("medication", "Medication name").Required().String()

	reminderList = reminder.Command("list", "See all reminders")

	// Status
	status = app.Command("status", "Daemon status")
)

func main() {
	store := pillbox.NewStore("pillbox.db")
	if err := store.Open(); err != nil {
		fmt.Println(err)
	}
	defer store.Close()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case install.FullCommand():
		if err := pillbox.Install(agentConfig); err != nil {
			fmt.Println("Failed to install:", err)
			return
		}

		if err := pillbox.Load(agentConfig); err != nil {
			fmt.Println("Failed to load agent:", err)
			return
		}

		fmt.Printf("Service \"%s\" installed.\n", agentConfig.DisplayName)
	case uninstall.FullCommand():
		if err := pillbox.Unload(agentConfig); err != nil {
			fmt.Println("Failed to unload agent:", err)
			return
		}

		if err := pillbox.Uninstall(agentConfig); err != nil {
			fmt.Println("Failed to uninstall:", err)
			return
		}
		fmt.Printf("Service \"%s\" uninstalled.\n", agentConfig.DisplayName)
	case reminderAdd.FullCommand():
		if !(*reminderMorning || *reminderAfternoon || *reminderEvening) {
			app.Fatalf("Any (or all) of the --morning, --afternoon or --evening flags have to be present.")
		}

		reminder := pillbox.Reminder{
			Body:      *reminderAddMedication,
			Morning:   *reminderMorning,
			Afternoon: *reminderAfternoon,
			Evening:   *reminderEvening,
		}
		store.SaveReminder(reminder)
	case reminderList.FullCommand():
		reminders, err := store.FetchReminders()
		if err != nil {
			panic(err)
		}
		reporter := NewReporter(reminders)
		fmt.Println(reporter.ToTable())
	case reminderRemove.FullCommand():
		var removed pillbox.Reminder
		reminders, err := store.FetchReminders()
		if err != nil {
			panic(err)
		}

		for _, reminder := range reminders {
			if reminder.Body == *reminderRemoveMedication {
				removed = reminder
			}
		}
		err = store.DeleteReminder(removed)
		if err != nil {
			panic(err)
		}
	case status.FullCommand():
		out, err := exec.Command("sh", "-c", "launchctl list | grep Pillbox").CombinedOutput()
		if err != nil {
			fmt.Println("Process not found, exiting.")
			os.Exit(1)
		}
		parts := strings.Fields(string(out))
		if parts[1] == "0" {
			fmt.Println("Running.")
		} else {
			fmt.Println("Process active, with errors.")
		}
	}
}
