package maps

// MergeStringMaps will merge all keys/values (non-recursivelly) of two map[string]string
func MergeStringMaps(a map[string]string, b map[string]string) map[string]string {
	c := make(map[string]string)

	for k, v := range a {
		c[k] = v
	}

	for k, v := range b {
		c[k] = v
	}

	return c
}