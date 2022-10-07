package util

func Pick(mapData map[string]interface{}, keys []string) map[string]interface{} {
	rlt := make(map[string]interface{})

	for _, key := range keys {
		if val, ok := mapData[key]; ok {
			rlt[key] = val
		}
	}

	return rlt
}
