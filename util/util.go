package util

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
)

var BG = context.Background()

// Hard-exits with a log statement on error.
// Only for cases like CLI startup and scripting
func Fatal(errs ...error) {
	for _, err := range errs {
		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}
}

// JSON map convenience func
func DecodeStringMap(raw []byte) (map[string]any, error) {
	var result map[string]any
	err := json.Unmarshal(raw, &result)

	return result, err
}

// Pretty-print for debugging
func Json(x interface{}) ([]byte, error) {
	return json.MarshalIndent(x, "", "  ")
}

// Stringify JSON pretty-print
func Debug(x interface{}) string {
	y, err := Json(x)
	if err != nil {
		return err.Error()
	}
	return string(y)
}

// convenience alias
type MAP = map[string]any
