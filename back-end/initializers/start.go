package initializers

func StartDatabase(path string) {
	LoadEnvVariables(path)
	ConnectToDB()
	SyncDatabases()
}
