package main

import (
	"flag"
	"math/rand"
	"strconv"
	"time"
)

var (
	name     = flag.String("name", "sensor", "name of the sensor")
	freq     = flag.Uint("freq", 5, "update frequency in cycle/sec")
	max      = flag.Float64("max", 5., "maximum value for generated readings")
	min      = flag.Float64("min", 1., "minimum value for generated readings")
	stepSize = flag.Float64("step", 0.1, "maximum allowable change per measurement")

	randomNumberGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	value                 = randomNumberGenerator.Float64()*(*max-*min) + *min
	nominalValue          = (*max-*min)/2 + *min
)

func main() {
	flag.Parse()
	duration, _ := time.ParseDuration(strconv.Itoa(1000/int(*freq)) + "ms")

	signal := time.Tick(duration)
	for range signal {
		calcValue()
	}
}

func calcValue() {
	var maxStep, minStep float64
	if value < nominalValue {
		maxStep = *stepSize
		minStep = -1 * *stepSize * (value - *min) / (nominalValue - *min)
	} else {
		maxStep = *stepSize * (*max - value) / (*max - nominalValue)
		minStep = -1 * *stepSize
	}

	value = randomNumberGenerator.Float64()*(maxStep-minStep) + minStep
}
