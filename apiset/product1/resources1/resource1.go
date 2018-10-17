package resources1

import "time"

func Test(args ...interface{}) (interface{}, error) {
	var input = args[0].([]interface{})

	var a = input[0].(int)
	var b = input[1].(string)
	var c = input[2].(time.Time)
	var d = input[3].(float64)

	return test(a, b, c, d), nil
}

func test(a int, b string, c time.Time, d float64) map[string]interface{} {
	var result = make(map[string]interface{})

	result["a"] = a
	result["b"] = b
	result["c"] = c
	result["d"] = d

	return result
}

func untest(a int, b string, c time.Time, d float64) map[string]interface{} {
	var result = make(map[string]interface{})

	result["a"] = a
	result["b"] = b
	result["c"] = c
	result["d"] = d

	return result
}
