package tools

import (
	"strconv"
	"strings"
)

type QueryString string

func (q QueryString) SplitBy(split string) []string {
	stringList := strings.Split(string(q), split)
	return stringList
}

func (q QueryString) GetInteger(split string) []uint {
	var (
		ids = make([]uint, 0)
	)
	stringList := q.SplitBy(split)
	for _, s := range stringList {
		parseUint, _ := strconv.ParseUint(s, 64, 20)
		ids = append(ids, uint(parseUint))
	}
	return ids
}
