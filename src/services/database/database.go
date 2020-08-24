package database

import (
	"github.com/kyani-inc/kms/v2"
)

var (
	SBI kms.Database
)

func Setup() {
	SetupSBI()
}
