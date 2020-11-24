/*
This file will be used to simulate different seating algorithums for an inline
grid, meaning that the seats are paraell to each other both in the x and y
directions, meaning your head is aligned with the person in front of you (as)
 opposed to staggered
*/

package main

import (
//    "bufio"
    "fmt"
//    "os"
    "math"
)

const AVIL = 0 //a seat that is empty and far enough away to be socally distant
const EMPY = 1 //a seat that must be empty to mainatin safe distance
const OCCU = 2 //a seat that has a human in it

const SAFEDIST = 2 //2 meters distance required for socail distanceing

const SEATWIDTH = 0.5842 //space 1 seat takes up 23' converted to meters
const ROWSPACE = 0.9144 //space between rows of seats 36' converted to meters
//source for seat dimentions (page 3 paragraph 2) -> http://www.renewhempstead.com/wp-content/uploads/2012/08/hem_movie-theater-feasibility-brief-2012.pdf

const ROWSEATS = 15 //seats in a row
const ROWNUM = 20 //rows in the theader
// this leads to a total of 300 seats (medium to large theader)
//based off of this forum from 2005 http://www.film-tech.com/ubb/f5/t001340.html

var tdr [ROWNUM][ROWSEATS] int //matrix representing a theater with 20 rows of 15 seats

type seat struct {
  seat int //the seat in the row (x cord)
	row int //the row the seat is in (y cord)
}

func main() {

  fmt.Println("Pogram Started\n")

	//reader := bufio.NewReader(os.Stdin)

  initTheader()

	var full = false

	for full == false {
		party := 0

		fmt.Print("Whats the size of the party?: ")
		fmt.Scanln(&party)

    placeParty(party)
    printTheader()
	}
}

func placeParty(size int) bool {
	room := false
	consAvil := 0 //consecuitive avalible seats
	startSeat := 0 //seat that the party will start at
	//startRow := 0 //row that the party will start at

	for c := 0 ; c < ROWNUM ; c++ { //loops for each row
		for r := 0 ; r <  ROWSEATS ; r++ { //loops for each colum
			if(tdr[c][r] == 0){ //when a seat is avilible
				if(startSeat == -1){ //if its the first avalible seat in a row
					startSeat = r //current seat becomes start seat
				}
				consAvil++
				if(consAvil >= size){ //if there are enough seats
					room = true; //there is enough room for the party
					for p := 0 ; p < size ; p++{ //loops through each person in the party
						tdr[c][startSeat] = 2 //sets seat to occupied
            blockOff(seat{seat: startSeat , row: c}) //will block off seats that need to be empty
						startSeat++
					}
          return room
				}
			}
			if(tdr[c][r] != 0){
				startSeat = -1 //indicates a new start seat is needed
				consAvil = 0
			}
		}
	}
	return room
}

func blockOff(s seat) {
  //fmt.Print("seat ")
  //fmt.Print(s.seat)
  //fmt.Print(" ")
  //fmt.Println(s.row)
  for c := 0 ; c < ROWNUM ; c++ { //loops for each row
    for r := 0 ; r <  ROWSEATS ; r++ { //loops for each colum
      if(tdr[c][r] == 0){
        if(getSeatDist(s , seat{seat: r , row: c}) < SAFEDIST){
          tdr[c][r] = 1
          //fmt.Println("block off called")
        }
      }
    }
  }
}

func getSeatDist(s1 seat , s2 seat) float64 {
	xdist := math.Abs(float64(s1.seat) - float64(s2.seat)) * SEATWIDTH
	ydist := math.Abs(float64(s1.row) - float64(s2.row)) * ROWSPACE
	dist := (xdist * xdist) + (ydist * ydist)
	return dist
}

func printTheader () {
  fmt.Println()
  for c := 0 ; c < ROWNUM ; c++ { //loops for each row
    for r := 0 ; r <  ROWSEATS ; r++ { //loops for each colum
      fmt.Print(tdr[c][r])
      fmt.Print(" ")
    }
    fmt.Println()
  }
  fmt.Println()
}

func initTheader () {
  for c := 0 ; c < ROWNUM ; c++ { //loops for each row
    for r := 0 ; r <  ROWSEATS ; r++ { //loops for each colum
      tdr[c][r] = 0
    }
  }
}
