package processor

import (
	"database/sql"
	"log"
	"time"
)

type processor struct {
	db      *sql.DB
	workers map[string]*Worker
}

const (
	getAllUsersQuery = `
		SELECT id FROM test_user`
)

func NewProcessor(db *sql.DB) *processor {
	return &processor{
		db:      db,
		workers: make(map[string]*Worker),
	}
}

func (p *processor) Run() {
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			p.update()
		}
	}
}

func (p *processor) update() {
	rows, err := p.db.Query(getAllUsersQuery)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return
		default:
			log.Fatal(err)
		}
	}
	defer rows.Close()

	for rows.Next() {
		var userID string
		err = rows.Scan(&userID)
		if err != nil {
			log.Fatal(err)
		}

		_, exists := p.workers[userID]
		if !exists {
			p.workers[userID] = NewWorker(userID, p.db)
			go p.workers[userID].Work()
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
