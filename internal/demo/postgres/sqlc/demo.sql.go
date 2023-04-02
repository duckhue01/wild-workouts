// Code generated by sqlc. DO NOT EDIT.
// source: demo.sql

package sqlc

import (
	"context"
)

const listAllDemos = `-- name: ListAllDemos :many
SELECT id, name FROM demo
`

func (q *Queries) ListAllDemos(ctx context.Context) ([]Demo, error) {
	rows, err := q.db.QueryContext(ctx, listAllDemos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Demo{}
	for rows.Next() {
		var i Demo
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
