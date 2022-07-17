package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	eludedGuards := rand.Intn(100)
	var isHeistOn bool = true
	fmt.Println(eludedGuards)
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	openedVault := rand.Intn(100)
	if openedVault >= 70 && isHeistOn == true {
		fmt.Println("Grab and GO!")
	} else if isHeistOn == true {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	leftSafely := rand.Intn(5)
	if isHeistOn == true {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Print("Looks like you tripped an alarm... run?")
		case 1:
			isHeistOn = false
			fmt.Print("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Print("When did they start raising dogs in vaults??")
		case 3:
			isHeistOn = false
			fmt.Print("Looks like this fingerprint scanner won’t accept any fingerprint…")
		default:
			fmt.Print("Start the getaway car!")
		}
	}
	if isHeistOn == true {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("$", amtStolen, "not bad!")
	}
}
