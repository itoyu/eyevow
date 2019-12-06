package eyevow

import (
	"net/url"
	"os"
)

type env struct {
	Secret string
	URL    *url.URL
	Bucket string
}

func envOr(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func parseEnv() *env {
	s := envOr("URL", "http://localhost:1256")
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	return &env{
		Secret: envOr("SECRET", "deadbeef"),
		URL:    u,
		Bucket: envOr("BUCKET", "var/eyevow/"),
	}
}
