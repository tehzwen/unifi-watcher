package unifiwatcher

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DEFAULT_CONN_STR = "user=unifi-protect dbname=unifi-protect sslmode=disable port=5433"
	GET_EVENTS_QUERY = `
	SELECT id, type, "createdAt", "cameraId", "smartDetectTypes" FROM events 
	WHERE "createdAt" > $1
	ORDER BY "createdAt" DESC;`
)

type UnifiEvent struct {
	Id               string    `db:"id"`
	Type             string    `db:"type"`
	CreatedAt        time.Time `db:"createdAt"`
	CameraId         string    `db:"cameraId"`
	SmartDetectTypes string    `db:"smartDetectTypes" json:"smartDetectTypes"`
}

type UnifiWatcher struct {
	NumEvents int // total number of processed events

	connectionString string        // the connection string used for the db
	pollFrequency    time.Duration // how often the watcher polls the unifi db
	pollQuery        string        // the query used to poll the database
}

func (w *UnifiWatcher) Watch(eventFunc func(e UnifiEvent)) error {
	currentDate := time.Now()
	db, err := sqlx.Open("postgres", w.connectionString)
	if err != nil {
		return err
	}
	fmt.Println("Connected to database, listening for changes")

	currentIndex := 0
	for {
		result, err := db.QueryxContext(context.Background(), GET_EVENTS_QUERY, currentDate)
		if err != nil {
			return err
		}

		for result.Next() {
			var e UnifiEvent
			err = result.StructScan(&e)
			if err != nil {
				return err
			}
			if currentIndex == 0 {
				currentDate = e.CreatedAt
			}
			// now mak
			eventFunc(e)
			w.NumEvents++
			currentIndex++
		}
		result.Close()
		currentIndex = 0

		time.Sleep(w.pollFrequency)
	}
}

func WithPollFrequency(freq time.Duration) func(*UnifiWatcher) {
	return func(uw *UnifiWatcher) {
		uw.pollFrequency = freq
	}
}

func WithCustomConnString(connS string) func(*UnifiWatcher) {
	return func(uw *UnifiWatcher) {
		uw.connectionString = connS
	}
}

func WithCustomQuery(q string) func(*UnifiWatcher) {
	return func(uw *UnifiWatcher) {
		uw.pollQuery = q
	}
}

func NewUnifiWatcher(opts ...func(*UnifiWatcher)) *UnifiWatcher {
	// set defaults
	uw := &UnifiWatcher{
		pollFrequency:    time.Second * 2,
		connectionString: DEFAULT_CONN_STR,
		pollQuery:        GET_EVENTS_QUERY,
	}

	for _, o := range opts {
		o(uw)
	}
	return uw
}
