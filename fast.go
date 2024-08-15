package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"time"
)

type rollSessionResult struct {
	paralysisProcAmount int
	achieveTime         time.Time
	achieveSession      int
}

func simulateSessionGroup(
	rollsPerSession int,
	chanceDenominator int,
	// ch chan<- rollSessionResult,
	// wg *sync.WaitGroup,
) rollSessionResult {

	// defer wg.Done()

	currentRoll := 0
	paralysisProcArray := make([]int, chanceDenominator)

	// Roll 4 sets
	for currentRoll < rollsPerSession {
		// Generate a number 0 to 3
		number := rand.IntN(chanceDenominator)
		// Add 1 to the array index of the result.
		paralysisProcArray[number]++
		// Attempt goes up by 1
		currentRoll++
	}

	// Local minimum
	var result rollSessionResult

	result.paralysisProcAmount = 0
	achieveSession := 0

	// Get the biggest number in the results array
	// We are essentially checking for paralysis, 4 roll sessions at a time, one for each number index.
	for _, sessionParalysis := range paralysisProcArray {
		achieveSession++

		// Check if that index value is big
		if sessionParalysis > result.paralysisProcAmount {
			// If it is, save it
			result.paralysisProcAmount = sessionParalysis
			result.achieveTime = time.Now()
			result.achieveSession = achieveSession
		}
	}
	// ch <- result

	return result
}

func showResult(result rollSessionResult, startingTime time.Time) {
	fmt.Printf(
		" - Paralysis procs: %v\n - Roll Session: #%v\n - Time: %v\n - Date: %v\n",
		result.paralysisProcAmount,
		result.achieveSession,
		result.achieveTime.Sub(startingTime),
		result.achieveTime,
	)

}

func main() {
	// Get the amount of non self-destructive PP we have.
	defenseCurlPP := 40
	mudSportPP := 15
	// Reduce the number by 1  because Graveler faints due to Self destruct, because it goes first.
	safeTurns := defenseCurlPP + mudSportPP - 1 // -> 54

	opponentPP := 160                                         // Obtained directly from the video, not fact checking this
	struggleTakeDownTurns := 71                               // Obtaining directly from the video, not fact checking this
	defaultTurnsToStall := opponentPP + struggleTakeDownTurns // -> 231

	defaultParalysisProcsNeeded := defaultTurnsToStall - safeTurns // -> 177

	// For execution configuration, we can set the amount of paralysis procs manually
	paralysisProcsNeededPtr := flag.Int(
		"paralysisProcs",
		defaultParalysisProcsNeeded,
		"The amount of paralysis procs needed to end the simulation in victory.",
	)

	// For execution configuration, we can set the amount of turns available manually
	rollsPerSessionPtr := flag.Int(
		"turnsToStall",
		defaultTurnsToStall,
		"The amount of turns available in the simulation in victory. This is the amount of rolls in a roll session.",
	)

	// For execution configuration, we can set the amount of roll sessions manually, too.
	maxRollSessionsPtr := flag.Int(
		"maxAttempts",
		100000,
		"The maximum amount of roll session groups before the program gives up. (A session group has 4 roll sessions)",
	)

	// If you don't want milestones to be shown, set this flag to false
	doesShowMilestonesPtr := flag.Bool(
		"showMilestones",
		false,
		"Whether to show each milestone",
	)

	flag.Parse()

	paralysisProcsNeeded := *paralysisProcsNeededPtr
	maxRollSessions := *maxRollSessionsPtr
	doesShowMilestones := *doesShowMilestonesPtr
	rollsPerSession := *rollsPerSessionPtr

	// One in 4 chance
	paralysisChanceDenominator := 4

	currentSessionGroup := 0
	currentSession := 0

	// THIS IS THE RELEVANT BIT OF CODE:

	// Timer setup:
	startTime := time.Now()
	fmt.Printf("Execution started %v\n", startTime)

	var finalResult rollSessionResult

	// responseChannel := make(chan rollSessionResult)
	// var waitGroup sync.WaitGroup

	for (finalResult.paralysisProcAmount < paralysisProcsNeeded) && (currentSession < maxRollSessions) {
		res := simulateSessionGroup(rollsPerSession, paralysisChanceDenominator)

		if res.paralysisProcAmount > finalResult.paralysisProcAmount {
			finalResult = res
			finalResult.achieveSession += currentSessionGroup * 4
			if doesShowMilestones {
				fmt.Println("New Goal!")
				showResult(finalResult, startTime)
			}
		}

		currentSessionGroup++
		currentSession += 4
	}

	endTime := time.Now()
	executionTime := time.Since(startTime)
	fmt.Println("Final Greatest Attempt:")
	showResult(finalResult, startTime)
	fmt.Printf("All %v simulations took %v\n", maxRollSessions, executionTime)
	fmt.Printf("Execution ended %v", endTime)

}
