package random

import (
	"github.io/MXuDong/example/pkg/constant"
	"math/rand"
	"strconv"
	"strings"
)

// Percentage(short for re) define the random value of every object.
// The format of percentage is present%int:int:value:value:value...:value,...
// The return is string
// any error input will return error for reason
func AnalyzePercentage(input string) (string, error) {
	inputs := strings.Split(input, constant.RandomItemSpiltChar)
	var percentageList []randomItem
	for _, inputItem := range inputs {
		randomObject := randomItem{}
		// get percentage
		items := strings.Split(inputItem, constant.RandomPercentageChar)
		values := ""
		if len(items) == 1 {
			randomObject.percentage = 0
			values = items[0]
		}
		if len(items) >= 2 {
			p := items[1]
			pi, err := strconv.Atoi(p)
			if err != nil {
				return "", err
			}
			randomObject.percentage = pi
			values = strings.Join(items[1:], constant.RandomPercentageChar)
		}
		// get value
		randomObject.values = strings.Split(values, constant.RandomIntervalChar)
		percentageList = append(percentageList, randomObject)
	}

	// generator rate
	total := 0

	for _, item := range percentageList {
		total += item.percentage
	}
	if total <= 0 {
		total = 1 * len(percentageList)
		for index := range percentageList {
			percentageList[index].percentage = 1
		}
	}

	// get random
	randInt := rand.Int() % total
	index := 0
	for _, item := range percentageList {
		if index+item.percentage > randInt {
			// random value
			return item.values[rand.Int()%len(item.values)], nil
		}
	}
	return "", nil
}

type randomItem struct {
	percentage int // if nil, set 0
	values     []string
}
