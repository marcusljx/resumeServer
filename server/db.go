package server

type DB map[string][]byte

var (
	// localDB is the local storage used for storing resume details
	localDB DB
)
