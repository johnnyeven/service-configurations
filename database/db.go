package database

import (
	"github.com/johnnyeven/libtools/sqlx"
)

var DBConfiguration = sqlx.NewDatabase("configuration")
