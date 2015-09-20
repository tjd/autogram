// autogram.go

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// lowercase letters of the alphabet
const alpha = "abcdefghijklmnopqrstuvwxyz"

// array of 26 ints representing letter counts e.g. arr[0] is the number of
// as, arr[1] is the number of bs, ..., and arr[25] is the number of zs
type alphavec [26]int

// initialization vector
var zero = alphavec{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

// initialization vector
var ones = alphavec{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}

// copy one alphavec into another
func (v *alphavec) assign(other *alphavec) {
	for i := 0; i < len(v); i++ {
		v[i] = other[i]
	}
}

// add an alphavec to another
func (v *alphavec) add(other *alphavec) {
	for i := 0; i < len(v); i++ {
		v[i] += other[i]
	}
}

// returns the number of items in v equal to n
func (v *alphavec) count(n int) int {
	result := 0
	for i := 0; i < len(v); i++ {
		if v[i] == n {
			result++
		}
	}
	return result
}

// test if one alphavec is the same as another
func (v *alphavec) equals(other *alphavec) bool {
	for i := 0; i < len(v); i++ {
		if v[i] != other[i] {
			return false
		}
	}
	return true
}

// converts v to a convenient string representation
func (v *alphavec) String() string {
	var result []string
	for i := 0; i < len(v); i++ {
		s := fmt.Sprintf("%c:%v", alpha[i], v[i])
		result = append(result, s)
	}
	return "[" + strings.Join(result, " ") + "]"
}

// sets the elements of v to be randomly chosen values
func (v *alphavec) randomize() {
	for i := 0; i < len(v); i++ {
		v[i] = rand.Intn(50)
	}
}

// only counts lowercase letters, i.e. characters in alpha;
// all other characters are ignored
func vec(s string) *alphavec {
	result := zero // make a copy of the zero vector
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if 0 <= c && c < 26 {
			result[c]++
		}
	}
	return &result
}

///////////////////////////////////////////////////////////////////////////////////////

// English number names
var numbers = []string{
	"zero", "one", "two", "three", "four", "five", "six",
	"seven", "eight", "nine", "ten", "eleven", "twelve",
	"thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
	"eighteen", "nineteen",

	"twenty", "twenty-one", "twenty-two", "twenty-three", "twenty-four",
	"twenty-five", "twenty-six", "twenty-seven", "twenty-eight", "twenty-nine",

	"thirty", "thirty-one", "thirty-two", "thirty-three", "thirty-four",
	"thirty-five", "thirty-six", "thirty-seven", "thirty-eight", "thirty-nine",

	"forty", "forty-one", "forty-two", "forty-three", "forty-four",
	"forty-five", "forty-six", "forty-seven", "forty-eight", "forty-nine",

	"fifty", "fifty-one", "fifty-two", "fifty-three", "fifty-four",
	"fifty-five", "fifty-six", "fifty-seven", "fifty-eight", "fifty-nine",

	"sixty", "sixty-one", "sixty-two", "sixty-three", "sixty-four",
	"sixty-five", "sixty-six", "sixty-seven", "sixty-eight", "sixty-nine",

	"seventy", "seventy-one", "seventy-two", "seventy-three", "seventy-four",
	"seventy-five", "seventy-six", "seventy-seven", "seventy-eight", "seventy-nine",

	"eighty", "eighty-one", "eighty-two", "eighty-three", "eighty-four",
	"eighty-five", "eighty-six", "eighty-seven", "eighty-eight", "eighty-nine",

	"ninety", "ninety-one", "ninety-two", "ninety-three", "ninety-four",
	"ninety-five", "ninety-six", "ninety-seven", "ninety-eight", "ninety-nine",
}

// vector of English number names
var numberCounts = initNumberCounts()

func initNumberCounts() []alphavec {
	var result []alphavec
	for _, n := range numbers {
		result = append(result, *vec(n))
	}
	return result
}

///////////////////////////////////////////////////////////////////////////////////////

type vectrans func(*alphavec, *alphavec)

// Creates, and returns, the functions actualCounts and toString. Both of the
// returned functions take into account prefix and finalAnd.
func makeFunctions(prefix, finalAnd string) (vectrans, func(*alphavec) string) {
	extra := ones // every letter occurs at least once in the sentence
	extra.add(vec(prefix))
	extra.add(vec(finalAnd))

	actualCounts := func(v, result *alphavec) {
		result.assign(&extra)
		for i := 0; i < len(v); i++ {
			cnt := v[i]
			result.add(&numberCounts[cnt])
		}

		// add extra "s"s
		s_count := len(v) - v.count(1)
		result[18] += s_count // result[18] is "s" count
	}

	toString := func(v *alphavec) string {
		result := []string{prefix + " "}
		for i := 0; i < len(v)-1; i++ {
			num := v[i]
			result = append(result, numbers[num]+" "+string(alpha[i]))
			if num == 1 {
				result = append(result, ", ")
			} else {
				result = append(result, "s, ")
			}
		}
		if v[25] == 1 {
			result = append(result, finalAnd+" one z.")
		} else {
			result = append(result, finalAnd+" "+numbers[v[25]]+" zs.")
		}
		return strings.Join(result, "")
	}

	return actualCounts, toString
}

func findSelfRefSentence(prefix, finalAnd string) {
	fmt.Printf("findSelfRefSentence(\"%v\", \"%v\") launched ...\n", prefix, finalAnd)
	var count int64 = 1
	iter_count := 1

	actualCounts, toString := makeFunctions(prefix, finalAnd)
	v := new(alphavec)
	v.randomize()
	actual := new(alphavec)
	actualCounts(v, actual)

	cache := make(map[alphavec]int)

	for !v.equals(actual) {
		if cache[*actual] == 1 {
			v.randomize()
			actualCounts(v, actual)
			cache = make(map[alphavec]int) // empty the cache
			iter_count = 0
		} else {
			cache[*actual] = 1
			v, actual = actual, v
			actualCounts(v, actual)
			iter_count++
		}
		count++
	} // for

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("!!")
	fmt.Println("!! Success!")
	fmt.Println("!!")
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	fmt.Println(toString(v))
	fmt.Println(v)
	fmt.Println(actual)
	fmt.Printf("count = %v\n", count)
}

///////////////////////////////////////////////////////////////////////////////////////

type sentence struct {
	prefix, finalAnd string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Searching for self-referential sentences ...\n")

	inputs := []sentence{
		{"kurt's self-referential sentence has", "and"},
		{"ray's sentence has", "and"},
		{"bonnie's sentence has", "and"},
	}

	// use a WaitGroup to wait for all goroutines finish
	// see http://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/
	var wg sync.WaitGroup

	wg.Add(len(inputs)) // # of goroutines to wait for
	for _, s := range inputs {
		go func(s sentence) {
			defer wg.Done()
			findSelfRefSentence(s.prefix, s.finalAnd)
		}(s)
	}

	wg.Wait()
	fmt.Println("Program done!")
} // main
