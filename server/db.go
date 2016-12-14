package server

type DB map[string]ResumeObject

var (
	// localDB is the local storage used for storing resume details
	localDB = make(DB)
)
