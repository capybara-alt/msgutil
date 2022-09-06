package msgutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var appMsgs map[string]interface{}
const (
	notFoundError = "Target error not found"
	targetKeyNull = "Keys cannot be empty"
	fileNotFound	= "Message resource file not found"
)

// Load json file then map to variable appMsgs
// This function must be called once
func Init(jsonpath string) error {
	jsonfiles, err := filepath.Glob(jsonpath)
	if err != nil {
		return err
	}

	appMsgs = make(map[string]interface{})
	for _, jsonfile := range jsonfiles {
		if filepath.Ext(jsonfile) != ".json" {
			continue
		}

		file, err := os.Open(jsonfile)
		if err != nil {
			return err
		}
		defer file.Close()
	
		b, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
	
		childAppMsg := make(map[string]interface{})
		if err = json.Unmarshal(b, &childAppMsg); err != nil {
			return err
		}

		filename := filepath.Base(jsonfile)
		filename = strings.Replace(filename, ".json", "", -1)
		appMsgs[filename] = childAppMsg
	}

	if len(appMsgs) < 1 {
		return errors.New(fileNotFound)
	}


	return nil
}

// Get message string
// If keys are empty or message isn't found from resource, return error
func GetMassage(keys ...string) (string, error) {
	if len(keys) < 1 {
		return "", errors.New(targetKeyNull)
	}

	tempAppMsgs := appMsgs
	for _, key := range keys {
		if reflect.ValueOf(tempAppMsgs[key]).Kind() == reflect.String {
			return tempAppMsgs[key].(string), nil
		} else if reflect.ValueOf(tempAppMsgs[key]).Kind() == reflect.Map {
			tempAppMsgs = tempAppMsgs[key].(map[string]interface{})
		} else {
			break
		}
	}

	return "", errors.New(notFoundError)
}


// Get message string then set arg values
func GetMessageArgs(keys []string, args ...string) (string, error) {
	msgStr, err := GetMassage(keys...)
	if err != nil {
		return "", err
	}

	for index, arg := range args {
		msgStr = strings.Replace(msgStr, fmt.Sprintf("{{%d}}", index), arg, -1)
	}

	return msgStr, nil
}
