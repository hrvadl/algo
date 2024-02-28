package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hrvadl/algo/internal/matrix"
)

func GetMatrix(rows, columns int) (matrix.Matrix, error) {
	if rows < 0 || columns < 0 {
		return matrix.Matrix{}, errors.New("can't create a matrix with the zero size")
	}

	m := matrix.Matrix{
		Rows: make([]matrix.Row, rows),
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nType the elements in the row  with the space between each:\n")
	for i := 0; i < rows && scanner.Scan(); i++ {
		tokens := strings.Fields(scanner.Text())

		if len(tokens) != columns {
			return matrix.Matrix{}, fmt.Errorf("invalid amount of tokens, expected: %v", columns)
		}

		for _, token := range tokens {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return matrix.Matrix{}, err
			}

			m.Rows[i] = append(m.Rows[i], num)
		}
	}

	return m, nil
}
