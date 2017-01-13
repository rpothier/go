package exercises

import "fmt"


type bbPlayer struct {
	name string
	atBats int
	hits int
}

func div(a int,b int) float64 {
	return float64(a)/ float64(b)

}

func (bbp *bbPlayer) average() float64 {
	//bbp.name = "Test"
	return div(bbp.hits, bbp.atBats)
}

func Methods() {
	fmt.Println("Methods")

	players := []bbPlayer {
		{"David_Ortiz",8640,2772},
		{"Carl_Yastrzemski",11988,3419},
		{"Ted_Williams",7706,2654},
	}

	//players = append(players, bbPlayer {
	//	name: "Jim Rice",
	//	atBats: 8225,
	//	hits: 2452,
	//})
	//var p *bbPlayer

	for _,pl := range players {
		fmt.Printf("the average for %-16s is .%3.0f\n", pl.name, pl.average()*1000)
		//fmt.Printf("the average for %-16s is %1.3v\n", (&pl).name, (&pl).average())
		//p = &pl
		//fmt.Printf("the average for %-16s is %1.3v\n", p.name, p.average())
	}
	p2 := bbPlayer{"Jim Rice", 8225, 2452}
	fmt.Printf("the average for %-16s is .%3.0f\n", p2.name, p2.average()*1000)
}