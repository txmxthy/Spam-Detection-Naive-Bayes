package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

//Process file
func readFile(fname string) (data [][]int, err error) {
	// Read in file
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	///Split into lines
	lines := strings.Split(string(b), "\n")
	var dataset = make([][]int, 0)

	//For each line
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		//Split into values only
		a := strings.Split(line, "     ")
		b := make([]int, len(a)-1) //12 Features, 1 Class
		//For each value
		for i, value := range a {
			//Convert to int
			if value == "" {
				continue
			}
			//Convert (With offset for "")
			b[i-1], err = strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
		}
		dataset = append(dataset, b)
	}
	return dataset, err
}

//Separate by class label
func separateClasses(dataset [][]int) (separated map[string][][]int, err error) {
	sep := make(map[string][][]int)
	for _, v := range dataset {
		if v[len(v)-1] == 0 {
			sep["0"] = append(sep["0"], v[:len(v)-1])
		} else if v[len(v)-1] == 1 {
			sep["1"] = append(sep["1"], v[:len(v)-1])
		} else {
			return nil, err
		}

	}
	return sep, err
}

//Get Average
func mean(data []int) (avg float64, err error) {
	total := 0
	for _, v := range data {
		total += v
	}
	return float64(total) / float64(len(data)), err
}

//Get Standard Deviation
func standardDev(data []int) (stdev float64, meanavg float64, err error) {
	avg, err := mean(data)
	var sd float64
	for _, v := range data {
		// The use of Pow math function func Pow(x, y float64) float64
		sd += math.Pow(float64(v)-avg, 2)
	}
	// The use of Sqrt math function func Sqrt(x float64) float64
	return math.Sqrt(sd / float64(len(data))), avg, err
}

// Return a 12 long list summarising each column into 3 key stats
func summarise(data [][]int) (sum [][]float64, err error) {

	temp := make(map[string][]int)
	// Store every column in map
	for _, row := range data { // For each row
		for col := 0; col < len(row); col++ { // For each column
			key := strconv.Itoa(col)                // Use the col number as a key
			temp[key] = append(temp[key], row[col]) // Add the array stored at [row], [col], to the map under the key col
		}
	}

	//Do summaries
	summaries := make(map[string][]float64)
	for column, summs := range temp {
		stdevi, meanavg, _ := standardDev(summs)
		s := []float64{meanavg, stdevi, float64(len(summs))}
		summaries[column] = s
	}

	//Convert
	var list = make([][]float64, 12)
	for k, v := range summaries {
		pos, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		list[pos] = v
	}

	return list, err
}

//take separated classes in map and summarise
func summariseClasses(classes map[string][][]int) (classSumms map[string][][]float64, err error) {
	summaries := make(map[string][][]float64)
	for class, v := range classes {
		summaries[class], err = summarise(v)
	}
	return summaries, err
}

func gaussian(x float64, mean float64, stddev float64) float64 {
	exponent := math.Exp(-math.Pow(x-mean, 2) / (2.0 * math.Pow(stddev, 2)))
	return (1 / (math.Sqrt(2*math.Pi) * stddev)) * exponent
}

func classProb(summary map[string][][]float64, row []int) (probs map[string]float64, err error) {
	totalRows := len(summary)
	probabilities := make(map[string]float64)
	for classValue, classSummary := range summary {
		//Frequency
		probabilities[classValue] = summary[classValue][0][2] / float64(totalRows)
		for i := range classSummary {
			mean := classSummary[i][0]
			stdev := classSummary[i][1]
			probabilities[classValue] *= gaussian(float64(row[i]), mean, stdev)
		}

	}
	return probabilities, err
}

func main() {
	train, err := readFile(os.Args[1])
	test, err := readFile(os.Args[2])
	classes, err := separateClasses(train)
	if err != nil {
		panic(err)
	}

	summary, err := summariseClasses(classes)
	probabilities := make([]map[string]float64, 0)
	var spam = 0.0
	var notSpam = 0.0

	for _, row := range test {
		classProbability, _ := classProb(summary, row)
		probabilities = append(probabilities, classProbability)
		if classProbability["0"] < classProbability["1"] {
			spam += 1
			fmt.Printf("\tSPAM: CERTAINTY [ %.2f ]\n", classProbability["0"]*100)
		} else {
			notSpam += 1
			fmt.Printf("NOT SPAM: CERTAINTY [ %.2f ]\n", classProbability["1"]*100)

		}
	}
	fmt.Println("===== TOTAL PROBABILITIES =====")
	a := spam + notSpam
	fmt.Println("	SPAM:", spam/a)
	fmt.Println("NOT SPAM:", notSpam/a)

}
