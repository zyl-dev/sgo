package consts

import "github.com/pkg/errors"

var (
	DBNotFoundError = errors.New("Database not found")
	CheckSignFailed = errors.New("Check sign failed")
)
