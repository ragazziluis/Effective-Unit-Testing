package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDatabase(t *testing.T) {
	db, err := ConnectDatabase()
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
