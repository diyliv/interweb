package repository

import (
	"database/sql"
	"time"

	"go.uber.org/zap"

	"github.com/diyliv/interweb/internal/errs"
	"github.com/diyliv/interweb/internal/models"
)

type repository struct {
	psql   *sql.DB
	logger *zap.Logger
}

func NewRepository(psql *sql.DB, logger *zap.Logger) *repository {
	return &repository{
		psql:   psql,
		logger: logger,
	}
}

func (r *repository) AddUser(user models.User) error {
	_, err := r.psql.Query("INSERT INTO user_info(user_telegram_id, user_first_request, user_requests_count) VALUES($1, $2, $3)",
		user.UserTelegramId, time.Now().Local(), 1)
	if err != nil {
		r.logger.Error("Error while inserting data: " + err.Error())
		return err
	}
	return nil
}

func (r *repository) UserExists(userId int64) bool {
	row := r.psql.QueryRow("SELECT user_requests_count FROM user_info WHERE user_telegram_id = $1", userId)
	if err := row.Scan(); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return true
	}
	return false
}

func (r *repository) UpdateUser(userId int64) error {
	_, err := r.psql.Exec("UPDATE user_info SET user_requests_count = user_requests_count + 1 WHERE user_telegram_id = $1", userId)
	if err != nil {
		r.logger.Error("Error while updating user info: " + err.Error())
		return err
	}
	return nil
}

func (r *repository) GetUserInfo(userId int64) (*models.User, error) {
	row := r.psql.QueryRow("SELECT user_first_request, user_requests_count FROM user_info WHERE user_telegram_id = $1", userId)
	if row.Err() != nil {
		r.logger.Error("Error while querying: " + row.Err().Error())
		return nil, row.Err()
	}

	var resp models.User
	if err := row.Scan(&resp.UserFirstRequest, &resp.UserRequestsCount); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrNotFound
		}
		r.logger.Error("Error while scanning values: " + err.Error())
		return nil, err
	}
	return &resp, nil
}
