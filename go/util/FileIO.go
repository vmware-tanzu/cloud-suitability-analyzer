/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

const JSON string = "json"
const YAML string = "yaml"
const YML string = "yml"

type FileEncoder interface {
	Encode(v interface{}) (err error)
}

type FileDecoder interface {
	Decode(v interface{}) error
}

func IsJson(filename string) bool {
	if strings.HasSuffix(strings.ToLower(filename), JSON) {
		return true
	}
	return false
}

func IsYaml(filename string) bool {
	if strings.HasSuffix(strings.ToLower(filename), YAML) ||
		strings.HasSuffix(strings.ToLower(filename), YML) {
		return true
	}
	return false
}

func WriteStructToFile(target interface{}, name string, path string, extension string, createDir bool) {

	if createDir {
		CheckAndCreateDir(path)
	}
	file := CreateFile(name, extension, path)
	defer file.Close()

	if *Verbose {
		fmt.Printf("Writing struct to file [%s]\n", file.Name())
	}

	var out []byte
	var err error
	if extension == JSON {
		out, err = json.MarshalIndent(target, "", "    ")
	} else {
		out, err = yaml.Marshal(target)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprint(file, string(out))
}

func WriteStructsToFile(target interface{}, name string, path string, extension string, createDir bool) (exportCnt int) {

	if reflect.TypeOf(target).Kind() != reflect.Slice {
		log.Panic("WriteStructsToFile() requires a target that is a slice of structs!")
	}

	slice := reflect.ValueOf(target)

	if *Verbose {
		fmt.Printf("Writing [%d] structs to file [%s]\n", slice.Len(), name)
	}
	if createDir {
		CheckAndCreateDir(path)
	}
	file := CreateFile(name, extension, path)
	defer file.Close()

	var encoder FileEncoder

	if extension == JSON {
		e := json.NewEncoder(file)
		e.SetIndent("", "    ")
		encoder = FileEncoder(e)
	} else {
		encoder = yaml.NewEncoder(file)
	}

	for i := 0; i < slice.Len(); i++ {

		err := encoder.Encode(slice.Index(i).Interface())

		if err != nil {
			fmt.Printf("Error Writing To File [%s]! Details: %v", file.Name(), err)
			return
		}

		exportCnt++
	}
	return
}

func SliceToFile(slice []string, path string, name string, ext string) {
	//Ensure path exists!
	CheckAndCreateDir(path)

	file := CreateFile(name, ext, path)
	defer file.Close()

	for _, line := range slice {
		fmt.Fprintln(file, line)
	}
}

func CheckAndCreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if !strings.Contains(path, string(os.PathSeparator)) {
			path = filepath.Join(".", path)
		}
		os.MkdirAll(path, os.ModePerm)
	}
}

func CreateFile(name string, extension string, path string) (file *os.File) {
	file, err := os.Create(fmt.Sprintf("%s/%s.%s", path, name, extension))
	checkFileError(name, err)
	return file
}

func checkFileError(name string, err error) {
	if err != nil {
		log.Printf("Failed creating/writing/reading file [%s]! Details: %v\n", name, err)
	}
}

func Exists(filePath string) (exists bool) {
	exists = true

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exists = false
	}

	return
}

func DeleteFileOrDir(path string) {
	os.Remove(path)
}

func OpenFileAtPath(path string, homepath string) (reader *os.File, err error) {

	if strings.HasPrefix(path, "~") {
		path = strings.Replace(path, "~", homepath, -1)
	}

	abspath, _ := filepath.Abs(path)

	return os.Open(abspath)
}
