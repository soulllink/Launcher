package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type JSON struct {
	Name      string
	AliasName string
	Path      string
	Args      []string
	TimesRun  int
}

type items struct {
	items []JSON
}

func init() {
	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println("Error Open File: ", err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var out []JSON
	json.Unmarshal(byteValue, &out)
	defer jsonFile.Close()
	fmt.Println("Games: ", out)
	for _, v := range out {
		games.additem(v)
	}
}

func (j *items) additem(item JSON) []JSON {
	j.items = append(j.items, item)
	return j.items
}

func (j *items) find(s string) int {
	for k, v := range j.items {
		if v.Name == s || v.AliasName == s {
			return k
		}
	}
	return -1
}

func (j *items) delete(s string) []JSON {
	i := j.find(s)
	if i != -1 {
		j.items[i] = j.items[len(j.items)-1]
		j.items[len(j.items)-1] = JSON{}
		j.items = j.items[:len(j.items)-1]
	}
	return j.items
}

func (j *items) updateField(s string, vl string) {
	val := reflect.ValueOf(&j.items[0])
	(val.Elem()).FieldByName(s).SetString(vl)
}

func (j *items) save() {
	out, _ := json.MarshalIndent(j.items, "", " ")
	_ = ioutil.WriteFile("test.json", out, 0644)
}
