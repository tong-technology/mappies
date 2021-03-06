package mappies

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func Contains(stringSlice []string, stringToSearch string) bool {
	for _, stringElement := range stringSlice {
		if stringElement == stringToSearch {
			return true
		}
	}
	return false
}

func ContainsKey(thisMap interface{}, key string) bool {
	keys := reflect.ValueOf(thisMap).MapKeys()
	for _, v := range keys {
		if v.Interface().(string) == key {
			return true
		}
	}
	return false
}

func MergeMaps(map1 map[string]string, map2 map[string]string) map[string]string {
	for key, value := range map2 {
		map1[key] = value
	}
	return map1
}

func DetectMimeType(file io.Reader) (string, error) {
	buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration
	if _, err := file.Read(buff); err != nil {
		fmt.Println(err) // do something with that error
		return "", err
	}
	return http.DetectContentType(buff), nil
}