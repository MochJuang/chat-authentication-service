package postgresql

import (
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func (r *userRepository) Begin() (string, error) {
	tx := r.db.Begin()
	sessionId := uuid.New().String()

	r.tx[sessionId] = tx

	return sessionId, tx.Error
}

func (r *userRepository) Rollback(sessionId string) error {
	tx, exist := r.tx[sessionId]
	if !exist {
		return errors.New("transaction session not found")
	}

	log.Info("Rollback transaction")
	delete(r.tx, sessionId)
	return tx.Rollback().Error
}

func (r *userRepository) Commit(sessionId string) error {
	tx, exist := r.tx[sessionId]
	if !exist {
		return errors.New("transaction session not found")
	}

	log.Info("Commit transaction")
	delete(r.tx, sessionId)
	return tx.Commit().Error
}
