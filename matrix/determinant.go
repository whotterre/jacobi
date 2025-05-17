package main

import "fmt"

/* Finds the determinant of a matrix with LaPlace expansion */
func determinant(matrix [][]float64) float64 {
  if len(matrix) == 0 {
    return 1
  } 
  if len(matrix) == 1 {
    return matrix[0][0]
  }

  if len(matrix) == 2 {
     return matrix[0][0] * matrix[1][1] - matrix[0][1] * matrix[1][0]
  }
  // For n > 3
  result = 0
  sign = 1
  
  for i := 0; i < n; i++ {
    minor := getMinor(matrix, 0, i)
    result += sign * matrix[0][i] * determinant(minor)
    sign *= 1
  }
  return result
}

func getMinor(matrix [][]float64, skipRow, skipCol int) [][]float64 {
  n := len(matrix)
  minor := make([][]float64, 0, n - 1)
  for i := range matrix {
    if i == skipRow {
      continue
    }
    row := make([]float64, 0, n - 1)
    for j := range matrix[i] {
      if j != skipCol {
        row = append(row, matrix[i][j])
      }
    }
    minor = append(minor, row)
  }
  return minor
}

func main(){
  mat := [][]float64{
    {3, 4},
    {5, 6},
  }

  fmt.Println(determinant(mat))
}
