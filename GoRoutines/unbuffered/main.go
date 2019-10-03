package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	Tennis()
	// RunningRace()
}

//========================================================================================
/*                                                                                      *
 *                                      RUNNING RACE                                    *
 *                                                                                      */
//========================================================================================

func RunningRace() {

	// Create an unbuffered channel.
	baton := make(chan int)

	// Three palyers are running hence three go routines
	wg.Add(3)

	// First runner to his mark.
	go Runner(baton)

	// Start the race.
	baton <- 1

	// Wait for the race to finish.
	wg.Wait()
}

// Runner simulates a person running in the relay race.
func Runner(baton chan int) {
	var newRunner int

	// Wait to receive the baton.
	runner := <-baton

	// Start running around the track.
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// New runner to the line.
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// Running around the track.
	time.Sleep(100 * time.Millisecond)

	// Is the race over.
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// Exchange the baton for the next runner.
	fmt.Printf(
		"Runner %d Exchange With Runner %d\n",
		runner,
		newRunner,
	)

	baton <- newRunner
	wg.Done()
}

//========================================================================================
/*                                                                                      *
 *                                      TENNIS GAME                                     *
 *                                                                                      */
//========================================================================================

func Tennis() {

	// Create an unbuffered channel.
	court := make(chan int)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Launch two players.
	go player("Nadal", court)
	go player("Djokovic", court)

	// Start the set.
	court <- 1

	// Wait for the game to finish.
	wg.Wait()
}

// player simulates a person playing the game of tennis.
func player(name string, court chan int) {
	defer wg.Done()
	for {

		// Wait for the ball to be hit back to us.
		ball, ok := <-court
		if !ok {

			// If the channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}
