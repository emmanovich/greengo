package main

func DemoGarden() Garden {

	g := Garden{}

	g.Add(Plant{
		Name: "Monstera",
		Interval: 5,
		LastWater: 7,
	})

	g.Add(Plant{
		Name: "Aloe Vera",
		Interval: 10,
		LastWater: 1,
	})

	g.Add(Plant{
		Name: "Orchid",
		Interval: 4,
		LastWater: 6,
	})

	return g
}
