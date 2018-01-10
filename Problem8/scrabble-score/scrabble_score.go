package scrabble

import (
	"strings"
)

//array holding each letter's scrabble score
var tileValues = [26]int{
	1, 3, 3, 2, //Tile values for letters a,b,c,d
	1, 4, 2, 4, //Tile values for letters e,f,g,h
	1, 8, 5, 1, //Tile values for letters i,j,k,l
	3, 1, 1, 3, //Tile values for letters m,n,o,p
	10, 1, 1, 1, //Tile values for letters q,r,s,t
	1, 4, 4, 8, //Tile values for letters u,v,w,x
	4, 10, //Tile values for letters y,z
}

//Score computes the score for a series of tiles
func Score(tiles string) int {
	total := 0

	//convert all tile to lower case ascii
	s := strings.ToLower(tiles)

	//look up the tiles value in array and add it to the running total
	for _, v := range s {
		total += tileValues[v-0x61] //0x61 is the hex value of char 'a'
	}

	return total
}
