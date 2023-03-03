package main

import (
    "fmt"
)

func main() {
    var row, col int
    fmt.Println("Digite o número de linhas da matriz: ")
    fmt.Scan(&row)
    fmt.Println("Digite o número de colunas da matriz: ")
    fmt.Scan(&col)

    // Criação da matriz
    matrix := make([][]int, row)
    for i := 0; i < row; i++ {
        matrix[i] = make([]int, col)
    }

    // Preenchimento da matriz
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            fmt.Printf("Digite o elemento (%d,%d) da matriz: ", i, j)
            fmt.Scan(&matrix[i][j])
        }
    }

    // Cálculo da soma dos elementos da matriz
    sum := 0
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            sum += matrix[i][j]
        }
    }

    fmt.Println("A soma dos elementos da matriz é:", sum)
}
