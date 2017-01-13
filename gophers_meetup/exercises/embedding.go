package exercises

import "fmt"

type stats interface {
	getStats()
	getName()
}

type player struct {
	name string
}

func (p player) getName() {
	fmt.Printf("the stats for %-16s is", p.name)
	//return bbp.p.name
	//return p.name
}
func (p player) getStats() {
	fmt.Printf("No stats")
}

type baseBallPlayer struct {
	//player   //embedded
	p player //not embedded
	atBats int
	hits int
}

func (bbp baseBallPlayer) getStats() {
	fmt.Printf(" .%3.0f\n", div(bbp.hits, bbp.atBats)*1000)
}

func (bbp baseBallPlayer) getName() {
	//return bbp.p.name
	fmt.Printf("the baseball stats for %-16s is", bbp.p.name)
	//fmt.Printf("the baseball stats for %-16s is", bbp.p.name)
}

func Embedding() {
	fmt.Println("\n\n\nEmbedding")
/*
	players := []baseBallPlayer {
		{player: player{ name:"David_Ortiz"},atBats:8640,hits:2772},
		{player: player{ name:"Carl_Yastrzemsli"},atBats:11988,hits:3419},
		{player: player{ name:"Ted_Williams"},atBats:7706,hits:2654},
	}

*/
	players := []baseBallPlayer {
		{p: player{ name:"David_Ortiz"},atBats:8640,hits:2772},
		{p: player{ name:"Carl_Yastrzemsli"},atBats:11988,hits:3419},
		{p: player{ name:"Ted_Williams"},atBats:7706,hits:2654},
	}

	for _,pl := range players {
		//printStats(pl)

		//fmt.Printf("the stats for %-16s is\n", pl.name)
		//fmt.Printf("the stats for %-16s is ", pl.p.name)
		pl.getName()
		pl.getStats()
	}

}

func printStats(pl stats){

	pl.getName()
	//pl.p.getName() // can't do here
	pl.getStats()
}