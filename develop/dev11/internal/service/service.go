package service

import (
	"errors"
	"time"
)

type Event struct {
	ID     int
	UserID int
	Date   time.Time
}

func Create(userID int, date time.Time) (*Event, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}
	event := &Event{
		ID:     1,
		UserID: userID,
		Date:   date,
	}
	return event, nil
}

func Update(eventID, userID int, date time.Time) (*Event, error) {
	if eventID <= 0 {
		return nil, errors.New("invalid event ID")
	}
	event := &Event{
		ID:     eventID,
		UserID: userID,
		Date:   date,
	}
	return event, nil
}

func Delete(eventID int) error {
	if eventID <= 0 {
		return errors.New("invalid event ID")
	}
	return nil
}

func GetDay(date time.Time) ([]*Event, error) {
	events := []*Event{}
	return events, nil
}

func GetWeek(date time.Time) ([]*Event, error) {
	events := []*Event{}
	return events, nil
}

func GetMonth(date time.Time) ([]*Event, error) {
	events := []*Event{}
	return events, nil
}
