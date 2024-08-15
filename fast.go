package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"time"
)

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
		1000000000,
		"The maximum amount of roll session groups before the program gives up. (A session group has 4 roll sessions)",
	)

	// If you don't want milestones to be shown, set this flag to false
	doesShowMilestonesPtr := flag.Bool("showMilestones", true, "Whether to show each milestone")

	flag.Parse()

	paralysisProcsNeeded := *paralysisProcsNeededPtr
	maxRollSessions := *maxRollSessionsPtr
	doesShowMilestones := *doesShowMilestonesPtr
	rollsPerSession := *rollsPerSessionPtr

	// One in 4 chance
	paralysisChanceDenominator := 4

	mostParalysisInASession := -1
	achievingSessionGroup := 0
	achievingSession := 0
	var achievingTime time.Duration

	currentSessionGroup := 0
	currentSession := 0

	resultsArray := make([]int, paralysisChanceDenominator)

	// THIS IS THE RELEVANT BIT OF CODE:

	// Timer setup:
	startTime := time.Now()
	fmt.Printf("Execution started %v\n", startTime)

	for (mostParalysisInASession < paralysisProcsNeeded) && (currentSession < maxRollSessions) {
		currentRoll := 0
		currentSessionGroup++

		// Roll 4 sets
		for currentRoll < rollsPerSession {
			// Generate a number 0 to 3
			number := rand.IntN(paralysisChanceDenominator)
			// Add 1 to the array index of the result.
			resultsArray[number]++
			// Attempt goes up by 1
			currentRoll++
		}

		// Get the biggest number in the results array
		// We are essentially checking for paralysis, 4 roll sessions at a time, one for each number index.
		for i, sessionParalysis := range resultsArray {
			// Check if that index value is big
			if sessionParalysis > mostParalysisInASession {
				// If it is, save it
				mostParalysisInASession = sessionParalysis
				achievingSessionGroup = currentSessionGroup
				achievingSession = currentSession
				achievingTime = time.Since(startTime)

				// Announce it. This is minimal amount of rolls.
				if doesShowMilestones {
					fmt.Printf("NEW GOAL! %v\n - Roll Session Group #%v\n - Roll Session #%v\n - Time: %v\n\n", mostParalysisInASession, achievingSessionGroup, achievingSession, achievingTime)
				}
			}
			// Add 1 to the amount of sessions ran
			currentSession++
			// Clear that index
			resultsArray[i] = 0
		}
	}

	endTime := time.Now()
	executionTime := time.Since(startTime)

	fmt.Printf("Final greatest attempt:\n - %v paralysis rolls\n - Roll Session Group #%v\n - Roll Session #%v\n - Time: %v\n\n", mostParalysisInASession, achievingSessionGroup, achievingSession, achievingTime)
	fmt.Printf("All %v simulations took %v", maxRollSessions, executionTime)
	fmt.Printf("Execution ended %v", endTime)

}
