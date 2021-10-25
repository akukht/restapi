package model

import (
	"context"

	_ "github.com/lib/pq"
)

const (
	deleteEvent    = `DELETE FROM events WHERE id = $1`
	checkUserEvent = `SELECT event_id FROM users_events WHERE event_id = $1 AND user_id = $2`
)

func (q *Queries) DeleteEvent(ctx context.Context, eventID int, userID int) error {
	row := q.db.QueryRowContext(ctx, checkUserEvent, eventID, userID)
	var checkEvent int
	err := row.Scan(&checkEvent)

	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, deleteEvent, eventID)

	return err
}
