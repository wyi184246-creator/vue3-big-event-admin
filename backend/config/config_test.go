package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	fmt.Println(RootPath)
	load := GetConfig()
	assert.Equal(t, "127.0.0.1", load.Mysql.Host)
}
