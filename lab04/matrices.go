package main

import ("fmt"
		"sync")


func main() {
  a := make_matrix(5, 6)
  b := make_matrix(6, 6)
  a = fill(a)
  b = fillWithDiagonal(b, 2)
  printMatrix(a)
  fmt.Printf("\n\tx\n\n")
  printMatrix(b)
  fmt.Printf("\n\t=\n\n")
  c := matrix_multiply_parallel(a, b)
  printMatrix(c)
  fmt.Printf("\n\n\n\n\n\n")

  a = make_matrix(4, 5)
  b = make_matrix(5, 3)
  a = fill(a)
  b = fillWithDiagonal(b, 3)
  printMatrix(a)
  fmt.Printf("\n\tx\n\n")
  printMatrix(b)
  fmt.Printf("\n\t=\n\n")
  c = matrix_multiply_parallel(a, b)
  printMatrix(c)
  fmt.Printf("\n\n\n\n\n\n")

  a = make_matrix(4, 4)
  b = make_matrix(4, 4)
  a = fill(a)
  b = fillWithDiagonal(b, 4)
  printMatrix(a)
  fmt.Printf("\n\tx\n\n")
  printMatrix(b)
  fmt.Printf("\n\t=\n\n")
  c = matrix_multiply_parallel(a, b)
  printMatrix(c)
  fmt.Printf("\n\n\n\n\n\n")

  a = make_matrix(4, 3)
  b = make_matrix(2, 4)
  a = fill(a)
  b = fillWithDiagonal(b, 1)
  printMatrix(a)
  fmt.Printf("\n\tx\n\n")
  printMatrix(b)
  fmt.Printf("\n\t=\n\n")
  c = matrix_multiply_parallel(a, b)
  printMatrix(c)
}



func make_matrix(n int, m int) [][]int {
	matrix := make([][]int, n);
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	return matrix
}


func printMatrix(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func fill(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j] = i*m+j
		}
	}
	return matrix
}

func fillWithDiagonal(matrix [][]int, c int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == j {
				matrix[i][j] = c
			} else {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func matrix_multiply(a [][]int, b [][]int) [][]int{
	if len(a[0]) == len(b){
		n := len(a)
		m := len(a[0])
		p := len(b[0])
		c := make_matrix(n, p)
		for i := 0; i < n; i++ {
			for j := 0; j < p; j++ {
				c[i][j] = 0
				for k := 0; k < m; k++ {
					c[i][j] += a[i][k]* b[k][j]
				}
			}
		}
		return c
	} else {
		return make_matrix(1, 1)
	}
}

func matrix_multiply_parallel(a [][]int, b [][]int) [][]int{
	if len(a[0]) == len(b){
		n := len(a)
		m := len(a[0])
		p := len(b[0])
		c := make_matrix(n, p)
		var wait_group sync.WaitGroup
		for i := 0; i < n; i++ {
			for j := 0; j < p; j++ {
				wait_group.Add(1)
				go func(i int, j int) {
    				c[i][j] = inner_mult(a, b, i, j, m)
    				wait_group.Done()
				}(i, j)
			}
		}
		wait_group.Wait()
		return c
	} else {
		return make_matrix(1, 1)
	}
}

func inner_mult(a [][]int, b [][]int, i int, j int, m int) int{

    sum := 0;

    for k:=0 ; k < m ; k++{
        sum += a[i][k] * b[k][j];
    }
    return sum;
}