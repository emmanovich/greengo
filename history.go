package main

type History struct {
	Records []string
}

func DemoHistory() History {

	return History{

		Records: []string{

			"Monstera watered",

			"Orchid fertilized",

			"Aloe Vera watered",

			"Monstera cleaned",

			"Orchid watered",
		},
	}
}
