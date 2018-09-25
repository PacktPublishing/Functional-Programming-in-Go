package main

type Connection interface {
	isAuthenticated(username string) bool
}

type API struct {
	conn Connection
}

func (a API) isAuthenticated(username string) bool {
	if a.conn.isAuthenticated(username) {
		return true
	}

	return false
}
