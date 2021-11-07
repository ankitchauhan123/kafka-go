package lib


const (
	Testing= Env("testing")
	Development = Env("development")
	Staging = Env("staging")
	Production = Env("production")
	)


type Env string

