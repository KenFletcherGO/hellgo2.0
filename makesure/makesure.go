package makesure

func Check(s string, f map[string]string) string {
	for k, v := range f {
		if k == s {
			return v
		} else if len(s) > 2 || len(s) < 2 {
			return "yo"
		} else {
			return "yo"
		}
	}
	return "yo"
}
