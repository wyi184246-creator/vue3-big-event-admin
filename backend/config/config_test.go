package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	fmt.Println(RootPath)
	load, err := Load(RootPath + "config/configs.yml")
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1", load.Mysql.Host)
}
