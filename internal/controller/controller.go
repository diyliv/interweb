package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/diyliv/interweb/internal/errs"
	"github.com/diyliv/interweb/internal/interfaces"
	"github.com/diyliv/interweb/internal/models"
)

type controller struct {
	repo   interfaces.Repository
	logger *zap.Logger
}

func NewController(repo interfaces.Repository, logger *zap.Logger) *controller {
	return &controller{
		repo:   repo,
		logger: logger,
	}
}

func (c *controller) FindInfo(userId int64, category string) ([]models.APIResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.publicapis.org/entries?category=%s", category))
	if err != nil {
		c.logger.Error("Error while sending GET request to https://api.publicapis.org/entries " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	var apiData models.APIResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Error while reading response body: " + err.Error())
		return nil, err
	}
	if err := json.Unmarshal(body, &apiData); err != nil {
		c.logger.Error("Error while unmarshalling data: " + err.Error())
		return nil, err
	}

	userExists := c.repo.UserExists(userId)
	if !userExists {
		if err := c.repo.AddUser(models.User{
			UserTelegramId:    userId,
			UserFirstRequest:  time.Now().Local(),
			UserRequestsCount: 1,
		}); err != nil {
			c.logger.Error("Error while calling repository.AddUser(): " + err.Error())
			return nil, err
		}
	} else {
		if err := c.repo.UpdateUser(userId); err != nil {
			c.logger.Error("Error while calling repository.UpdateUser(): " + err.Error())
			return nil, err
		}
	}

	return []models.APIResponse{apiData}, err
}

func (c *controller) GetUserInfo(userId int64) (*models.User, error) {
	userInfo, err := c.repo.GetUserInfo(userId)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return userInfo, nil
}
