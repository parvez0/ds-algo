package main

import (
	"fmt"
)

type Translator interface{
	Add(word string, langauge string, traslation string, transLang string)
	Find(word string, source string, dstLang string)
}

type node struct {
	value string
	translations map[string]*node
}

type translator struct{
	database map[string]*node
}

func(t *Translator) Add(word, lang, trans, transLang string) {
	var (
		datanode *node
		dstdatanode *node
		e bool
	) 
	srcKey, dstKey := getKey(lang, word), getKey(transLang, trans)
	if datanode, e = t.database[key]; !e {
		datanode = &node{value: word, translator: make(map[string]*node)}
		t.database[key] = datanode
	}
	if dstdatanode, e = t.database[dstKey]; !e {
		dstdatanode = &node{value: trans, translator: make(map[string]*node)}
		t.database[dstKey] = dstdatanode
	}
	datanode.traslations[srcKey], dstdatanode.translations[dstKey] = dstdatanode, datanode
}

func(t *Translator) Find(word, source, dst string) string {
	datanode := t.database[getKey(source, word)]
}

func getKey(a, b string) string {
	return fmt.Sprintf("%s_%s", a, b)
}