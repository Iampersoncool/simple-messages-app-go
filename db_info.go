package main

import "os"

func GetDbInfoFromEnv(defConnUrl string) string {
	url := os.Getenv("DB_URL")

	if url == "" {
		return defConnUrl
	}

	return url
}
