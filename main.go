package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	res, err := parseJSON(`{"key": "value"}`)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(res)
}

func parseJSON(jsonStr string) (map[string]interface{}, error) {
	jsonStr = strings.TrimSpace(jsonStr)

	res := make(map[string]interface{})

	for i, ch := range jsonStr {
		if i == 0 && ch != '{' {
			err := getErrorMsg("Invalid JSON. Cannot start with: %c", ch)
			return nil, err
		}

		i++

		key, err := parseKey(jsonStr, &i)
		if err != nil {
			return nil, err
		}

		value, err := parseValue(jsonStr, &i)
		if err != nil {
			return nil, err
		}

		res[key] = value

		if i == len(jsonStr)-1 && ch != '}' {
			err := getErrorMsg("Invalid JSON. Missing `}` at the end")
			return nil, err
		}
	}
	return res, nil
}

func parseKey(jsonStr string, idx *int) (string, error) {
	var key string
	for jsonStr[*idx] != ':' {
		key += string(jsonStr[*idx])
		*idx++
	}

	if key[len(key)-1] != '"' {
		err := getErrorMsg(`Invalid JSON. Missing " at the end`)
		return "", err
	}

	return key, nil
}

func parseValue(jsonStr string, idx *int) (string, error) {
	return "value", nil
}

func getErrorMsg(msg string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(msg, args...))
}
