package db

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/vadskev/go_final_project/internal/lib/logger"
	"github.com/vadskev/go_final_project/internal/models/task"
	"go.uber.org/zap"
)

func (r *Repository) Update(task *task.Task) error {
	const op = "storage.db.Update"
	query := sq.
		Update(tableName).
		Set("date", task.Date).
		Set("title", task.Title).
		Set("comment", task.Comment).
		Set("repeat", task.Repeat).
		Where(sq.Eq{"id": task.ID})

	sql, args, err := query.ToSql()
	if err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
		return err
	}

	res, err := r.DB().Exec(sql, args...)
	if err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
		return err
	}

	if affected == 0 {
		logger.Error(op, zap.Any("error:", errors.New("task not found")))
		return fmt.Errorf("task not found")
	}
	return nil
}
