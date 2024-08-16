package main

import (
	"flag"
	"fmt"
	"hash/maphash"
	"sync"
	"time"
)

const (
	bitShifts = 31
)

type rollSessionResult struct {
	paralysisProcAmount uint64
	achieveTime         time.Time
	achieveSession      int
}

func Rand64() uint64 {
	return new(maphash.Hash).Sum64()
}

func simulateSessionGroup(
	rollsPerSession int,
	ch chan<- rollSessionResult,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	currentRoll := 0
	var paralysisProcInt uint64 = 0

outerLoop:
	for {
		randomNumber := Rand64() // rand.Int64()
		// Roll 64 sets
		for i := 0; i < bitShifts; i++ {
			paralysisProcInt += 1 << ((randomNumber >> (i << 1) & 0b11) << 0b100)

			currentRoll++
			if currentRoll > rollsPerSession {
				break outerLoop
			}
		}
	}

	// Local minimum
	var result rollSessionResult

	result.paralysisProcAmount = 0
	achieveSession := 0

	// Get the biggest number in the results array
	// We are essentially checking for paralysis, 4 roll sessions at a time, one for each number index.
	for i := 0; i < 3; i++ {
		paralysisProcs := (paralysisProcInt >> (i << 4)) & 0xFFFF

		if (paralysisProcs) > result.paralysisProcAmount {
			result.paralysisProcAmount = paralysisProcs
			result.achieveTime = time.Now()
			result.achieveSession = achieveSession
		}
	}

	ch <- result
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
		1000000,
		"The maximum amount of roll session groups before the program gives up. (A session group has 4 roll sessions)",
	)

	// If you don't want milestones to be shown, set this flag to false
	doesShowMilestonesPtr := flag.Bool(
		"showMilestones",
		false,
		"Whether to show each milestone",
	)

	flag.Parse()

	paralysisProcsNeeded := uint64(*paralysisProcsNeededPtr)
	maxRollSessions := *maxRollSessionsPtr
	doesShowMilestones := *doesShowMilestonesPtr
	rollsPerSession := *rollsPerSessionPtr

	currentSessionGroup := 0
	currentSession := 0

	// THIS IS THE RELEVANT BIT OF CODE:

	// Timer setup:
	startTime := time.Now()
	fmt.Printf("Execution started %v\n", startTime)

	var finalResult rollSessionResult

	maxGoroutines := 1023

	for (finalResult.paralysisProcAmount < paralysisProcsNeeded) && (currentSession < maxRollSessions) {
		responseChannel := make(chan rollSessionResult)
		var waitGroup sync.WaitGroup

		for i := 1; i <= maxGoroutines; i++ {
			waitGroup.Add(1)
			go simulateSessionGroup(rollsPerSession, responseChannel, &waitGroup)
		}

		go func() {
			waitGroup.Wait()
			close(responseChannel)
		}()

		for res := range responseChannel {
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
	}

	endTime := time.Now()
	executionTime := time.Since(startTime)
	fmt.Println("Final Greatest Attempt:")
	showResult(finalResult, startTime)
	fmt.Printf("All %v simulations took %v\n", maxRollSessions, executionTime)
	fmt.Printf("Execution ended %v", endTime)

}
