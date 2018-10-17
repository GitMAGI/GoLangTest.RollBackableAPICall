package resources3

func Test(args ...interface{}) (interface{}, error) {
	var input = args[0].([]interface{})

	var b = input[0].(string)

	return test(b), nil
}

func test(b string) map[string]interface{} {
	var result = make(map[string]interface{})

	result["b"] = b

	return result
}

func untest(b string) map[string]interface{} {
	var result = make(map[string]interface{})

	result["b"] = b

	return result
}
