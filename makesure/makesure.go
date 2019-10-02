package makesure

func Check(s string, f map[string]string) string {
	var holder string
	holder = f[s]
	//if its not the key then the value would be
	//empty remeber this
	if holder == "" {
		holder = "Yo"
	}
	/*
		for k, v := range f {
			if len(s) == 2 && k == s {
				holder = v
			} else {
				holder = "Yo"
			}
		}
	*/
	return holder
}
