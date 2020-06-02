package util

func String(unknownVal interface{}) string {
	var stringVal, ok = unknownVal.(string)

	if ok {
		return stringVal
	}
	return ""
}

func Int(unknownVal interface{}) int {
	intVal, ok := unknownVal.(int)

	if ok {
		return intVal
	}

	return 0
}
