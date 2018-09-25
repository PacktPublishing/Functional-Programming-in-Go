package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedConn struct {
	mock.Mock
}

func (m *MockedConn) isAuthenticated(username string) bool {
	args := m.Called(username)
	return args.Bool(0)

}

func TestAPI(t *testing.T) {
	conn := new(MockedConn)
	api := API{
		conn: conn,
	}

	conn.On("isAuthenticated", "forest.whitaker").Return(false)

	assert.Equal(t, api.isAuthenticated("forest.whitaker"), true)
}
