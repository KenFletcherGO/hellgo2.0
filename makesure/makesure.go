package makesure

func Check(s string, f map[string]string) string {
	var holder string
	for k, v := range f {
		if k == s {
			holder = v
		} else if len(s) > 2 || len(s) < 2 {
			holder = "yo"
		} else {
			holder = "Yo"
		}
	}
	return holder
}
