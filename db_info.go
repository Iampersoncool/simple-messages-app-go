package main

import "os"

const EnvDatabaseUrlKey = "DB_URL"

func GetDbInfoFromEnv(defConnUrl string) string {
	url := os.Getenv(EnvDatabaseUrlKey)

	if url == "" {
		return defConnUrl
	}

	return url
}
