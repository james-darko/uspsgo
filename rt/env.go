package rt

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"strings"
)

func MustGetEnv(name string) string {
	v := os.Getenv(name)
	Assertf(v != "", "%s env var is required", name)
	return v
}

func MustGetEnvJson(name string, v any) {
	vStr := MustGetEnv(name)
	switch vStr[0] {
	case '{', '[':
		err := json.Unmarshal([]byte(vStr), v)
		Assertf(err == nil, "could not unmarshal required config %s: %v", name, err)
	default:
		r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(vStr))
		err := json.NewDecoder(r).Decode(v)
		Assertf(err == nil, "could not unmarshal required config %s: %v", name, err)
	}
}
