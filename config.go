package main

import (
	"fmt"
	"log"
)

var transact TransactionLogger

func initializeTransactionLog() error {
	var err error

	transact, err = NewPostgresTransactionLogger(
		PostgresDBParams{
			host:     "localhost",
			dbName:   "kvs",
			user:     "admin",
			password: "1234",
		})
	if err != nil {
		return fmt.Errorf("failed to create transaction logger: %w", err)
	}

	events, errors := transact.ReadEvents()
	count, ok, e := 0, true, Event{}

	for ok && err == nil {
		select {
		case err, ok = <-errors:

		case e, ok = <-events:
			switch e.EventType {
			case EventDelete: // Got a DELETE event!
				err = Delete(e.Key)
				count++
			case EventPut: // Got a PUT event!
				err = Put(e.Key, e.Value)
				count++
			}
		}
	}

	log.Printf("%d events replayed\n", count)

	transact.Run()

	go func() {
		for err := range transact.Err() {
			log.Print(err)
		}
	}()

	return err
}
