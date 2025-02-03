package game

import (
	"fmt"
	"strconv"
)

type Dote struct {
	X int
	Y int
}

type Figure struct {
	Letter 	rune
}

type Field struct {
	World [3][3]rune
	Figures map[Dote]Figure
}

func (f *Field) AddMove(movecode string) error {
	letter := movecode[0]
	x, err := strconv.Atoi(string(movecode[1]))
	if err != nil{
		return err
	}

	y, err := strconv.Atoi(string(movecode[2]))
	if err != nil{
		return err
	}

	f.World[x][y] = rune(letter)
	return nil
}

func (f *Field) Display(){
	for i := 0; i < 3; i++ {
		for j := 0;j < 3; j++ {
			fmt.Print(f.World[i][j])
		}
		fmt.Print("\n")
	}
}