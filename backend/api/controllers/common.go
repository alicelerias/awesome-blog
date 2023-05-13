package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"strconv"

	"golang.org/x/exp/slices"
)

func parseInt(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

func stringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func getValidJson(reader io.Reader, validKeys []string) (data map[string]interface{}, err error) {
	data, err = toJson(reader)
	if err != nil {
		return
	}
	err = validateInput(validKeys, data)
	return
}

func toJson(reader io.Reader) (input map[string]interface{}, err error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &input)
	return
}

func validateInput(list []string, input map[string]interface{}) error {
	for key, _ := range input {
		if !slices.Contains(list, key) {
			return errors.New("invalid key")
		}
	}
	return nil
}
