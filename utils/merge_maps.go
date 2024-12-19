package utils

func MergeMaps(map1, map2 map[string]string) map[string]string {
	errMap := make(map[string]string)

	for k, v := range map1 {
		errMap[k] = v
	}

	for k, v := range map2 {
		errMap[k] = v
	}

	return errMap
}
