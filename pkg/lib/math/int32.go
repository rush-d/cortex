/*
Copyright 2020 Cortex Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package math

func MinInt32(val int32, vals ...int32) int32 {
	min := val
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func MaxInt32(val int32, vals ...int32) int32 {
	max := val
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func IsDivisibleByInt32(num int32, divisor int32) bool {
	return num%divisor == 0
}

func FactorsInt32(num int32) []int32 {
	divisibleNumbers := []int32{}
	for divisor := int32(1); divisor <= num/2; divisor++ {
		if IsDivisibleByInt32(num, divisor) {
			divisibleNumbers = append(divisibleNumbers, divisor)
		}
	}
	return append(divisibleNumbers, num)
}
