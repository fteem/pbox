package pillbox

import (
	"encoding/binary"
	"encoding/json"
	"os/user"
	"time"

	"github.com/boltdb/bolt"
)

const (
	remindersBucket = "reminders"
)

type Store struct {
	db   *bolt.DB
	path string
}

func NewStore(path string) *Store {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return &Store{
		path: user.HomeDir + "/." + path,
	}
}

func (s *Store) Open() error {
	// Open database connection
	db, err := bolt.Open(s.path, 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return err
	}

	// Assign connection handler to Store
	s.db = db

	// Initialize needed buckets (if non-existent)
	if err := s.db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(remindersBucket))
		return nil
	}); err != nil {
		s.Close()
		return err
	}

	return nil
}

func (s *Store) Close() error {
	if s.db != nil {
		s.db.Close()
	}

	return nil
}

func (s *Store) SaveReminder(reminder Reminder) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(remindersBucket))

		id, _ := b.NextSequence()
		reminder.ID = int(id)
		jsonBlob, err := json.Marshal(reminder)
		if err != nil {
			return err
		}

		return b.Put(itob(reminder.ID), jsonBlob)
	})

}

func (s *Store) FetchReminders() ([]Reminder, error) {
	var reminders []Reminder
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(remindersBucket))

		c := b.Cursor()

		var reminder Reminder

		for k, v := c.First(); k != nil; k, v = c.Next() {
			err := json.Unmarshal(v, &reminder)
			if err != nil {
				return err
			}
			reminders = append(reminders, reminder)
		}
		return nil
	})

	if err != nil {
		return []Reminder{}, err
	}

	return reminders, nil
}

func (s *Store) DeleteReminder(reminder Reminder) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(remindersBucket))
		if err := b.Delete(itob(reminder.ID)); err != nil {
			return err
		}
		return nil
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
