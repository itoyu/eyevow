package eyevow

import (
	"net/url"
	"os"
)

type env struct {
	Env     string
	Secret  string
	URL     *url.URL
	Bucket  string
	Twitter twitterEnv
}

type twitterEnv struct {
	ConsumerKey    string
	ConsumerSecret string
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
		Env:    envOr("ENV", "development"),
		Secret: envOr("SECRET", "deadbeef"),
		URL:    u,
		Bucket: envOr("BUCKET", "var/eyevow/"),
		Twitter: twitterEnv{
			ConsumerKey:    os.Getenv("TWITTER_CONSUMER_KEY"),
			ConsumerSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		},
	}
}
