package main

import "fmt"

func main() {
	// fmt.Println("Please visit http://10.186.62.13:12345/")
	// // array.array_demo()
	// // string_demo()

	// // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// // 	s := fmt.Sprintf("你好，世界！ -- time %s", time.Now().String())

	// // 	fmt.Fprintf(w, "%v\n", s)

	// // 	log.Printf("%v\n", s)
	// // })

	// // if err := http.ListenAndServe(":12345", nil); err != nil {
	// // 	log.Fatal("ListenAndServer: ", err)
	// // }

	// // fmt.Print("Print", "is", "different")

	// fmt.Printf("%d", time.Now().UnixNano())

	// rand.Seed(time.Now().UnixNano())

	// isHeistOn := true

	// eludedGuards := rand.Intn(100)

	// if eludedGuards >= 50 {
	// 	fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	// } else {
	// 	isHeistOn = false
	// 	fmt.Println("Plan a better disguise next time?")
	// }

	// openedVault := rand.Intn(100)

	// if isHeistOn && openedVault >= 70 {
	// 	fmt.Println("Grab and GO!")
	// } else if isHeistOn {
	// 	isHeistOn = false
	// 	fmt.Println("not opened !")
	// }

	// if isHeistOn {
	// 	switch leftSafely := rand.Intn(5); leftSafely {
	// 	case 0:
	// 		isHeistOn = false
	// 		fmt.Println("Turns out vault doors don't open from the inside...", 0)
	// 	case 1:
	// 		isHeistOn = false
	// 		fmt.Println("Turns out vault doors don't open from the inside...", 1)
	// 	case 2:
	// 		isHeistOn = false
	// 		fmt.Println("Turns out vault doors don't open from the inside...", 2)
	// 	case 3:
	// 		isHeistOn = false
	// 		fmt.Println("Turns out vault doors don't open from the inside...", 3)

	// 	default:
	// 		fmt.Println("Start the getaway car!")
	// 	}

	// }

	// if isHeistOn {
	// 	amtStolen := 10000 + rand.Intn(1000000)

	// 	fmt.Println(amtStolen)

	// }

	// fmt.Println(isHeistOn)

	fuel := 1000000

	fuel = flyToPlanet("金星", fuel)

	fuelGauge(fuel)

	fuel = flyToPlanet("火星", fuel)

	fuelGauge(fuel)

	fuel = flyToPlanet("土星", fuel)

	fuelGauge(fuel)

}

func fuelGauge(fuel int) {
	fmt.Printf("当前燃料剩余量：%d\n", fuel)
}

func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "金星":
		fuel = 300000
	case "火星":
		fuel = 500000
	case "土星":
		fuel = 700000
	default:
		fuel = 0
	}

	return fuel
}

func greetPlanet(planet string) {
	fmt.Printf("当前所处位置: %s\n", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int

	fuelRemaining = fuel

	fuelCost = calculateFuel(planet)

	if fuelRemaining > fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}
