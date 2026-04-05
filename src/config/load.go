package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func Load(cfg any) error {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		structField := t.Field(i)

		tag := structField.Tag.Get("env")
		if tag == "" {
			continue
		}

		parts := strings.Split(tag, ",")
		key := parts[0]

		var required bool
		var def string

		for _, p := range parts[1:] {
			if p == "required" {
				required = true
			}
			if after, ok := strings.CutPrefix(p, "default="); ok {
				def = after
			}
		}

		val := os.Getenv(key)

		if val == "" {
			if def != "" {
				val = def
			} else if required {
				return fmt.Errorf("missing required env: %s", key)
			}
		}

		field.SetString(val)
	}

	return nil
}
