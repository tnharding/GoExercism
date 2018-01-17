//Package grains ...
package grains

var board = make([]uint64, 64)

func init() {
	var startValue uint64 = 1

	for i := 0; i < 64; i++ {
		board[i] = startValue
		startValue *= 2
	}
}

func Square(input int) (uint64, error) {
	return board[input-1], nil
}

func Total() uint64 {

	return 0
}
