package main

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func throw(expression string) ([]int, int, error) {
	rand.Seed(time.Now().UnixNano())
	var matcher = regexp.MustCompile(`^(\d+)d(\d+)?$`)
	var matches = matcher.FindStringSubmatch(expression)
	if matches == nil {
		return make([]int, 0), 0, errors.New("Invalid format!")
	}
	var quantity, _ = strconv.Atoi(matches[1])
	var edges, _ = strconv.Atoi(matches[2])
	var results = make([]int, quantity)
	var sum = 0

	for i := range results {
		results[i] = rand.Intn(edges) + 1
		sum += results[i]
	}

	return results, sum, nil
}
