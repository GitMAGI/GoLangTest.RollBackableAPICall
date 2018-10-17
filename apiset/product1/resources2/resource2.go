package resources2

func Test(args ...interface{}) (interface{}, error) {
	var input = args[0].([]interface{})

	var b = input[0].(string)
	var d = input[1].(float64)

	return test(b, d), nil
}

func test(b string, d float64) map[string]interface{} {
	var result = make(map[string]interface{})

	result["b"] = b
	result["d"] = d

	return result
}

func untest(b string, d float64) map[string]interface{} {
	var result = make(map[string]interface{})

	result["b"] = b
	result["d"] = d

	return result
}
