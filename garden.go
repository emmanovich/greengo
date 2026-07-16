package main

type Garden struct {
	Plants []Plant
}

func (g *Garden) Add(p Plant) {
	g.Plants = append(g.Plants, p)
}
