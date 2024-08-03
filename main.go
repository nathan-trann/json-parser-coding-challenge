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

		key, err := parseKey(jsonStr, &i)
		if err != nil {
			return nil, err
		}
		fmt.Println(key)
		// res[key] := parseValue(jsonStr, &i)

		if i == len(jsonStr)-1 && ch != '}' {
			err := getErrorMsg("Invalid JSON. Missing `}` at the end")
			return nil, err
		}
	}
	return res, nil
}

func parseKey(jsonStr string, idx *int) (string, error) {
	var key string
	for i := *idx; i < len(jsonStr); i++ {
		if jsonStr[i] == '"' {
			for jsonStr[i] != ':' {
				key += string(jsonStr[i])
			}
		}

		if key[len(key)-1] != '"' {
			err := getErrorMsg(`Invalid JSON. Missing " at the end`)
			return "", err
		}

		*idx = i
	}

	return key, nil
}

func getErrorMsg(msg string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(msg, args...))
}
