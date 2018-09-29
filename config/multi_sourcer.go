package config

// NewMultiSourcer creates a sourcer that reads form each sourcer
// sequentially until a suitable value is found. A value found in
// a sourcer earlier in the list will override any later values.
func NewMultiSourcer(sourcers ...Sourcer) Sourcer {
	return func(envTag string) (string, bool) {
		for _, sourcer := range sourcers {
			if val, ok := sourcer(envTag); ok {
				return val, ok
			}
		}

		return "", false
	}
}