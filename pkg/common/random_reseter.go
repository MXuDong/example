package common

import (
	"math"
	"math/rand"
)

// RandomChecker packing the random, it provide safe random values.
// It will auto reset the present of target value.
type RandomChecker struct {
	seed         int64             // the seed of random
	randInstance *rand.Rand        // the rand instance
	v            map[string]uint64 // the value and present, the value is unchangeable
	vr           map[string]uint64 // the real value present of v
	vri          []string          // the index of vr(v real value)
	vra          []uint64          // the boom of vr(v real value)
	vSum         uint64            // dynamic sum of present
}

// reset value of r
func (r *RandomChecker) reset() *RandomChecker {
	// auto value will auto change present of values
	// copy the value
	// reset all value of r
	r.vr = map[string]uint64{}
	r.vri = []string{}
	r.vra = []uint64{}

	// copy value
	for key, value := range r.v {
		r.vr[key] = value          // reset the value to key
		r.vri = append(r.vri, key) // append the key to r
	}

	// prune the present
	prunePresent(r.vr)

	// update r.vra

	return r
}

func (r *RandomChecker) updateValue(targetIndex int) {
	value := r.vri[targetIndex] // get origin value

	otherSum := r.vSum - r.vra[targetIndex]
	onePresent := r.vSum / uint64(len(r.vri))

	// target less one present
	if r.vr[value] <= onePresent {
		r.vr[value] = 0
	} else {
		r.vr[value] -= onePresent
	}

	// update other value
	for keyItem, valueItem := range r.vr {
		if keyItem != value {
			r.vr[keyItem] += valueItem / otherSum * onePresent
		}
	}

	r.updateIndexArray()
}

func (r *RandomChecker) NextValue() string {
	randomValue := r.getNextUint64()
	index := r.findIndex(randomValue)
	if index >= 0 {
		r.updateValue(index)
		return r.vri[index]
	}
	return ""
}

// find the index of input value
func (r RandomChecker) findIndex(value uint64) int {
	value = value % r.vSum
	// todo update search method
	for index, valueI := range r.vra {
		if value < valueI {
			return index
		}
	}
	return -1
}

func (r *RandomChecker) updateIndexArray() {
	var sum uint64 = 0
	r.vra = []uint64{}
	for _, key := range r.vri {
		sum += r.vr[key]
		r.vra = append(r.vra, sum)
	}
	r.vSum = sum
}

// PutValue will put value to r.v
// If target object already exits, it will cover the present to target
// Note that every put action will reset dynamic present
// Warn: if all the value present's sum is great than math.MaxUint64, will reset all value << 1
func (r *RandomChecker) PutValue(target string, present uint64) {
	r.v[target] = present
	r.reset()
}

// ResetSeed will reset seed of RandomChecker
func (r *RandomChecker) ResetSeed(seed int64) {
	r.seed = seed
	r.randInstance = rand.New(rand.NewSource(seed))
}

// getNextUint64 will return random uint64
func (r *RandomChecker) getNextUint64() uint64 {
	return r.randInstance.Uint64()
}

// RandomCheckerGenerator will return the random checker
func RandomCheckerGenerator(seed int64) *RandomChecker {
	r := RandomChecker{
		seed: seed,
	}
	return r.reset()
}

// DefaultRandomChecker will set seed to 0
func DefaultRandomChecker() *RandomChecker {
	return RandomCheckerGenerator(0)
}

// the all the value present's sum is great than math.MaxUint64 will reset to >> 1
func prunePresent(values map[string]uint64) {
	var sum uint64
	flag := true

	// if sum > math.MaxUint64, set all value to >> 1
	for flag {
		flag = false
		sum = 0
		for _, value := range values {
			if math.MaxUint8 <= sum+value {
				//do reset
				flag = true
				break
			} else {
				sum += value
			}
		}
		if flag {
			// do reset
			for key, value := range values {
				values[key] = value >> 1
			}
		}
	}
}
