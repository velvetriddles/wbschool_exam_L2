package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

func parseCreate(r *http.Request) (int, time.Time, error) {
	strUserId := r.FormValue("user_id")
	strDate := r.FormValue("date")

	if strUserId == "" || strDate == "" {
		return 0, time.Time{}, errors.New("user_id and date are required")
	}

	userID, err := strconv.Atoi(strUserId)
	if err != nil {
		return 0, time.Time{}, errors.New("invalid user ID")
	}

	date, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		return 0, time.Time{}, errors.New("invalid date format")
	}

	return userID, date, nil
}

func parseUpdate(r *http.Request) (int, int, time.Time, error) {
	eventID := 0
	userID := 0
	date := time.Now()

	strEventId := r.FormValue("event_id")
	strUserId := r.FormValue("user_id")
	strDate := r.FormValue("date")

	var err error
	if strEventId != "" {
		eventID, err = strconv.Atoi(strEventId)
		if err != nil {
			return 0, 0, time.Time{}, errors.New("invalid event_id")
		}
	}

	if strUserId != "" {
		userID, err = strconv.Atoi(strUserId)
		if err != nil {
			return 0, 0, time.Time{}, errors.New("invalid user_id")
		}
	}

	if strDate != "" {
		date, err = time.Parse("2006-01-02", strDate)
		if err != nil {
			return 0, 0, time.Time{}, errors.New("invalid date")
		}
	}

	return eventID, userID, date, nil
}

func parseDelete(r *http.Request) (int, error) {
	strEventId := r.FormValue("event_id")

	eventID, err := strconv.Atoi(strEventId)
	if err != nil {
		return 0, errors.New("invalid event ID")
	}

	return eventID, nil
}

func parseForDay(r *http.Request) (time.Time, error) {
	strDate := r.FormValue("date")

	date, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		return time.Time{}, errors.New("invalid date format")
	}

	return date, nil
}

func parseForWeek(r *http.Request) (time.Time, error) {
	strDate := r.FormValue("date")

	date, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		return time.Time{}, errors.New("invalid date format")
	}

	return date, nil
}

func parseForMonth(r *http.Request) (time.Time, error) {
	strDate := r.FormValue("date")

	date, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		return time.Time{}, errors.New("invalid date format")
	}

	return date, nil
}
