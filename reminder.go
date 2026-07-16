package main

func DuePlants(plants []Plant) []Plant {

	var due []Plant

	for _, plant := range plants {

		if NeedsWater(plant) {

			due = append(due, plant)

		}

	}

	return due
}
