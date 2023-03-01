package initializers

func StartDatabase() {
	LoadEnvVariables()
	ConnectToDB()
	SyncDatabases()
}
