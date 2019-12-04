package eyevow

import "os"

type env struct {
	Secret string
}

func envOr(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func parseEnv() *env {
	return &env{
		Secret: envOr("SECRET", "deadbeef"),
	}
}
