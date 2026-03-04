package test

import (
	"backend/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGorm(t *testing.T) {
	db := dao.NewGormDB()
	err := db.Error
	assert.NoError(t, err)
}
