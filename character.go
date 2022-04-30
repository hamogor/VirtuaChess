package main

type Match struct {
	Player1     *Character
	Player2     *Character
	TurnCounter int
}

type Character struct {
	ID    int
	Name  string
	HP    int
	Moves map[string]*Move
}

func CreateGoh() *Character {
	goh := &Character{
		ID:    0,
		Name:  "Goh",
		HP:    220,
		Moves: make(map[string]*Move, 0),
	}
	goh.Moves["jab"] = Shoda()
	goh.Moves["elbow"] = Elbow()
	goh.Moves["lowKick"] = LowKick()
	goh.Moves["throw"] = Throw()
	return goh
}
