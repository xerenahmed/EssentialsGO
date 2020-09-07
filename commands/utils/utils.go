package utils

func SliceUnique(slice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range slice {
		if _, value := keys[entry]; !value && entry != "" {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}