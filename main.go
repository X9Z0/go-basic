package main

import (
    "fmt"
    "os"
    "strconv"

    "go-arith/cmd"
)

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Usage: go-arith <operation> <num1> <num2>")
        fmt.Println("Operations: add, subtract, multiply, divide")
        return
    }

    operation := os.Args[1]
    num1, err1 := strconv.ParseFloat(os.Args[2], 64)
    num2, err2 := strconv.ParseFloat(os.Args[3], 64)

    if err1 != nil || err2 != nil {
        fmt.Println("Error: Please provide valid numbers.")
        return
    }

    switch operation {
    case "add":
        result := cmd.Add(num1, num2)
        fmt.Printf("Result: %.2f\n", result)
    case "subtract":
        result := cmd.Subtract(num1, num2)
        fmt.Printf("Result: %.2f\n", result)
    case "multiply":
        result := cmd.Multiply(num1, num2)
        fmt.Printf("Result: %.2f\n", result)
    case "divide":
        result, err := cmd.Divide(num1, num2)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Printf("Result: %.2f\n", result)
    default:
        fmt.Println("Unknown operation:", operation)
    }
}

