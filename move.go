package main

type Move struct {
	Dmg int    // Damage (points)
	Lvl rune   // Attack level
	Exe int    // Execution (frames)
	Act int    // Active (frames)
	Tot int    // Total duration (frames)
	Grd int    // Frame difference after Guard
	NH  int    // Frame difference after Normal Hit
	CH  int    // Frame difference after Counter Hit
	Esc string // Evade Direction
	Sob int    // Number of drink points sobered from Shun Di
	Inp string // Input string that corresponds to move
}

func Shoda() *Move {
	shoda := &Move{
		Dmg: 10,
		Lvl: 'H',
		Exe: 12,
		Act: 2,
		Tot: 27,
		Grd: 2,
		NH:  5,
		CH:  8,
		Esc: "any",
		Sob: 0,
		Inp: "p",
	}
	return shoda
}

func Elbow() *Move {
	elbow := &Move{
		Dmg: 17,
		Lvl: 'M',
		Exe: 14,
		Act: 2,
		Tot: 36,
		Grd: -4,
		NH:  0,
		CH:  7,
		Esc: "any",
		Sob: 0,
		Inp: "6p",
	}
	return elbow
}

func LowKick() *Move {
	lowKick := &Move{
		Dmg: 15,
		Lvl: 'L',
		Exe: 16,
		Act: 2,
		Tot: 51,
		Grd: -15,
		NH:  -6,
		CH:  0,
		Esc: "any",
		Sob: 0,
		Inp: "2k",
	}
	return lowKick
}

func Throw() *Move {
	throw := &Move{
		Dmg: 50,
		Lvl: 'H',
		Exe: 10,
		Act: 1,
		Tot: 0,
		Grd: 0,
		NH:  0,
		CH:  0,
		Esc: "",
		Sob: 0,
		Inp: "6p+g",
	}
	return throw
}
