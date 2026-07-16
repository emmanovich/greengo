package main

import "fmt"

func PrintReport(
	g Garden,
	h History,
) {

	fmt.Println("Today's Plant Checklist")
	fmt.Println()

	for _, plant := range g.Plants {

		fmt.Println("✓", plant.Name)

		if NeedsWater(plant) {

			fmt.Println("Needs watering")

		} else {

			fmt.Println("Watered recently")

		}

		fmt.Println()
	}

	fmt.Println("----------------------")
	fmt.Println()

	fmt.Println("Plants:", len(g.Plants))

	fmt.Println(
		"Need watering:",
		len(DuePlants(g.Plants)),
	)

	fmt.Println(
		"History records:",
		len(h.Records),
	)
}
