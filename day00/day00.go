package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func readAndValidateNumbers() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	var numbs []int

	for scanner.Scan() {
		tmpNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		// validate tmpNumber
		if tmpNumber < -100000 || tmpNumber > 100000 {
			return nil, errors.New(fmt.Sprintf("not valid number %d", tmpNumber))
		}

		numbs = append(numbs, tmpNumber)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(numbs) == 0 {
		return nil, errors.New("need at least one number")
	}

	return numbs, nil
}

func getNumbsSum(numbs []int) (sum int) {
	sum = 0
	if numbs == nil || len(numbs) == 0 {
		return
	}

	for _, numb := range numbs {
		sum += numb
	}

	return
}

func getMean(numbsLen int, numbsSum int) float64 {
	if numbsLen == 0 {
		return 0
	}

	return float64(numbsSum) / float64(numbsLen)
}

func getSortedNumbsMedian(numbs []int) (median float64) {
	median = 0
	if numbs == nil || len(numbs) == 0 {
		return
	}

	numbsLen := len(numbs)
	if numbsLen%2 == 0 {
		median = float64(numbs[numbsLen/2-1]+numbs[numbsLen/2]) / 2
	} else {
		median = float64(numbs[numbsLen/2])
	}

	return
}

func getMinMode(numbs []int) int {
	if numbs == nil || len(numbs) == 0 {
		return 0
	}

	modeMap := make(map[int]int)
	for _, numb := range numbs {
		modeMap[numb] += 1
	}

	minMode := numbs[0]
	maxOccurred := modeMap[minMode]
	for _, numb := range numbs {
		tmpNumbMode := modeMap[numb]
		if (tmpNumbMode > maxOccurred) || (tmpNumbMode == maxOccurred && numb < minMode) {
			minMode = numb
			maxOccurred = tmpNumbMode
		}
	}

	return minMode
}

func getSD(numbs []int, mean float64) float64 {
	if numbs == nil || len(numbs) == 0 {
		return 0
	}

	var difBetweenNumbAndMeanSquaresSum float64 = 0
	for _, numb := range numbs {
		difBetweenNumbAndMeanSquare := math.Pow(mean-float64(numb), 2)
		difBetweenNumbAndMeanSquaresSum += difBetweenNumbAndMeanSquare
	}

	return math.Sqrt(difBetweenNumbAndMeanSquaresSum / float64(len(numbs)))
}

func main() {
	meanFlagIsSet := flag.Bool("mean", false, "Calculate average")
	medianFlagIsSet := flag.Bool("median", false, "Calculate median (middle number of a sorted sequence if its size is odd, and an average between two middle ones if their count is even)")
	modeFlagIsSet := flag.Bool("mode", false, "Calculate mode (number which is occurring most frequently. if there are several, the smallest one among those)")
	sdFlagIsSet := flag.Bool("sd", false, "Calculate standard deviation")
	flag.Parse()

	numbs, err := readAndValidateNumbers()
	if err != nil {
		fmt.Printf("error while reading numbers: %s\n", err.Error())
		return
	}

	// sort slice of numbers
	sort.Ints(numbs)

	// by default print all
	isDefault := false
	if !*meanFlagIsSet && !*medianFlagIsSet && !*modeFlagIsSet && !*sdFlagIsSet {
		isDefault = true
	}

	mean := getMean(len(numbs), getNumbsSum(numbs))

	if *meanFlagIsSet || isDefault {
		fmt.Printf("Mean: %.2f\n", mean)
	}

	if *medianFlagIsSet || isDefault {
		fmt.Printf("Median: %.2f\n", getSortedNumbsMedian(numbs))
	}

	if *modeFlagIsSet || isDefault {
		fmt.Printf("Mode: %d\n", getMinMode(numbs))
	}

	if *sdFlagIsSet || isDefault {
		fmt.Printf("SD: %.2f\n", getSD(numbs, mean))
	}
}
