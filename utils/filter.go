package utils

func Filter(items []string, filter string) []string {
	var result []string
	for _, item := range items {
		if item != filter {
			result = append(result, item)
		}
	}
	return result
}