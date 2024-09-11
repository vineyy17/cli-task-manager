package db

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"github.com/mergestat/timediff"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key     int       `json:"key"`
	Value   string    `json:"value"`
	Created time.Time `json:"created"`
	Done    bool      `json:"done"`
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		created := time.Now()

		taskStruct := Task{
			Key:     id,
			Value:   task,
			Created: created,
			Done:    false,
		}

		// Encode the Task struct into bytes
		encoded, err := encodeTask(taskStruct)
		if err != nil {
			return err
		}

		return b.Put(key, encoded)
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task Task
			task.Key = btoi(k)

			// Decode the stored byte slice into the Task struct
			err := decodeTask(v, &task)
			if err != nil {
				return err
			}

			tasks = append(tasks, task)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func MarkTaskDone(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		v := b.Get(itob(key))
		if v == nil {
			return nil
		}

		var task Task
		// Decode the existing task
		err := decodeTask(v, &task)
		if err != nil {
			return err
		}

		// Mark as done
		task.Done = true

		// Re-encode the updated task
		encoded, err := encodeTask(task)
		if err != nil {
			return err
		}

		return b.Put(itob(key), encoded)
	})
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

// encodeTask encodes the Task struct into a byte slice using JSON
func encodeTask(task Task) ([]byte, error) {
	return json.Marshal(task) 
}

// decodeTask decodes the byte slice into a Task struct using JSON
func decodeTask(data []byte, task *Task) error {
	return json.Unmarshal(data, task) 
}

// FormatCreatedTime formats the created time using timediff
func FormatCreatedTime(t time.Time) string {
	return timediff.TimeDiff(t)
}
