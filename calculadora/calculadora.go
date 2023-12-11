package mycalculator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operacion := scanner.Text()

	fmt.Println(operacion)

	operador := "*"

	valores := strings.Split(operacion, operador)
	fmt.Println(valores)

	fmt.Println(valores[0] + valores[1])

	operador1, err1 := strconv.Atoi(valores[0])
	operador2, err2 := strconv.Atoi(valores[1])

	if err1 != nil {
		log.Fatal("No se pudo convertir el numero 1", err1)

	}
	if err2 != nil {
		log.Fatal("No se pudo convertir el numero 2", err2)
	}

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		break
	case "-":
		fmt.Println(operador1 - operador2)
		break
	case "*":
		fmt.Println(operador1 * operador2)
		break
	case "/":
		fmt.Println(operador1 / operador2)
		break
	default:
		fmt.Println("No se encuentra la operacion")
		break
	}

}
