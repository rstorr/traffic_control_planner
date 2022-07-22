package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

type ControlType int

const (
	Roundabout ControlType = iota
	StopSign
	TrafficLight
)

const rBoutBonusValue = 0.1

// TODO break main.go into multiple files for easy reading.

func main() {
	app := cli.NewApp()
	app.Name = "traffic_control_planner"
	app.Usage = "determine optimal traffic control method based on Cars-Per-Minute."
	app.Action = func(c *cli.Context) error {

		var r1, r2, r3, r4 int
		var input string
		var err error

		// TODO: fix for loop code duplication
		for {
			fmt.Println("Enter Road 1 (North) Cars Per Minute number:")
			_, err := fmt.Scanf("%s", &input)
			if err != nil {
				return err
			}

			r1, err = strconv.Atoi(input)
			if err == nil {
				break
			}
		}

		for {
			fmt.Println("Enter Road 2 (East) Cars Per Minute number:")
			_, err := fmt.Scanf("%s", &input)
			if err != nil {
				return err
			}

			r2, err = strconv.Atoi(input)
			if err == nil {
				break
			}
		}

		for {
			fmt.Println("Enter Road 3 (South) Cars Per Minute number:")
			_, err := fmt.Scanf("%s", &input)
			if err != nil {
				return err
			}

			r3, err = strconv.Atoi(input)
			if err == nil {
				break
			}
		}

		for {
			fmt.Println("Enter Road 4 (West) Cars Per Minute number:")
			_, err := fmt.Scanf("%s", &input)
			if err != nil {
				return err
			}

			r4, err = strconv.Atoi(input)
			if err == nil {
				break
			}
		}

		roundAboutBonus := roundAboutBonus(r1, r2, r3, r4)

		totalCPM := r1 + r2 + r3 + r4

		trafficControl, err := determineTrafficControl(totalCPM, roundAboutBonus)
		if err != nil {
			return err
		}

		controlName, err := getControlName(trafficControl)
		if err != nil {
			return err
		}

		fmt.Printf("The most suitable traffic control system based on the provided CPM values: %s \n", controlName)

		// TODO print efficiency of chosen system and other systems.

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// roundAboutBonus is determined true if one through road is twice (or more) busy than the
// other through road.
func roundAboutBonus(r1, r2, r3, r4 int) bool {

	r12 := r1 + r2
	r34 := r3 + r4

	if r12*2 <= r34 {
		return true
	}

	if r34*2 <= r12 {
		return true
	}

	return false

}

// getControlName returns the string value of the enum.
func getControlName(t ControlType) (string, error) {
	if t == Roundabout {
		return "roundabout", nil
	}
	if t == StopSign {
		return "stop signs", nil
	}
	if t == TrafficLight {
		return "traffic lights", nil
	}

	return "", errors.New("control type not found.")
}

// determineTrafficControl takes the total CPM and returns optimal control type.
func determineTrafficControl(cpm int, rBoutBonus bool) (ControlType, error) {

	switch {
	case cpm >= 20:
		return highThroughput(rBoutBonus), nil
	case cpm < 20 && cpm >= 10:
		return medThroughput(rBoutBonus), nil
	case cpm < 10:
		return lowThroughput(rBoutBonus), nil
	}

	return 0, errors.New("developer error.")
}

func highThroughput(rBoutBonus bool) ControlType {
	roundaboutEfficiency := 0.5
	stopSignEfficiency := 0.2
	trafficLightEfficiency := 0.9

	if rBoutBonus {
		roundaboutEfficiency += rBoutBonusValue
	}

	return mostEfficient(
		roundaboutEfficiency,
		stopSignEfficiency,
		trafficLightEfficiency,
	)
}

func medThroughput(rBoutBonus bool) ControlType {
	roundaboutEfficiency := 0.75
	stopSignEfficiency := 0.3
	trafficLightEfficiency := 0.75

	if rBoutBonus {
		roundaboutEfficiency += rBoutBonusValue
	}

	return mostEfficient(
		roundaboutEfficiency,
		stopSignEfficiency,
		trafficLightEfficiency,
	)
}

func lowThroughput(rBoutBonus bool) ControlType {
	roundaboutEfficiency := 0.9
	stopSignEfficiency := 0.4
	trafficLightEfficiency := 0.3

	if rBoutBonus {
		roundaboutEfficiency += rBoutBonusValue
	}

	return mostEfficient(
		roundaboutEfficiency,
		stopSignEfficiency,
		trafficLightEfficiency,
	)
}

// mostEfficient calculates most efficient controlType from supplied values.
func mostEfficient(roundabout, stopSign, trafficLight float64) ControlType {
	if roundabout >= stopSign &&
		roundabout >= trafficLight {
		return Roundabout
	}
	if stopSign >= roundabout &&
		stopSign >= trafficLight {
		return StopSign
	}

	return TrafficLight
}
