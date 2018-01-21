package main

import (
	"github.com/apcera/termtables"
	"github.com/fteem/pbox"
)

type Reporter struct {
	Reminders []pillbox.Reminder
	Table     *termtables.Table
}

func NewReporter(reminders []pillbox.Reminder) *Reporter {
	table := termtables.CreateTable()
	table.AddHeaders("Name", "Morning", "Afternoon", "Evening")
	r := Reporter{
		Reminders: reminders,
		Table:     table,
	}

	return &r
}

func (r *Reporter) ToTable() string {
	for _, reminder := range r.Reminders {
		r.Table.AddRow(
			reminder.Body,
			r.boolToWord(reminder.Morning),
			r.boolToWord(reminder.Afternoon),
			r.boolToWord(reminder.Evening),
		)
	}
	return r.Table.Render()
}

func (r *Reporter) boolToWord(b bool) string {
	if b {
		return "Yes"
	} else {
		return "No"
	}
}

func (r *Reporter) boolToIcon(b bool) string {
	if b {
		return "✅"
	} else {
		return "❌"
	}
}
