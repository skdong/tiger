package bytedance

func simplifyPath(path string) string {
	paths := []string{}
	name := ""
	for i := 0; i < len(path); i++ {
		if path[i] != '/' {
			name += string(path[i])
		}
		if path[i] == '/' || i == len(path)-1 {
			if len(name) > 0 {
				if name[0] == '.' && len(name) < 3 {
					if len(name) == 2 && len(paths) > 0 {
						paths = paths[:len(paths)-1]
					}
					name = ""
				} else {
					if len(name) > 0 {
						paths = append(paths, name)
						name = ""
					}
				}
			}
		}
	}
	absPath := "/"
	for _, name := range paths {
		if len(absPath) > 1 {
			absPath += "/"
		}
		absPath += name
	}
	return absPath
}
