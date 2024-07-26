package db

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/vadskev/go_final_project/internal/lib/logger"
	"go.uber.org/zap"
)

func (r *Repository) Delete(id string) error {
	const op = "storage.db.Delete"
	query := sq.
		Delete(tableName).
		Where(sq.Eq{"id": id})

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
		return fmt.Errorf("задача не найдена")
	}
	return nil
}
