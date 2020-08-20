package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

type mapa struct {
	guzanito gusano
	numero   int  // dato para llenar alimento
	activo   bool // true hay gusano // False no hay gusano
}

type gusano struct {
	tamaño     int
	id         int
	horizontal bool
	x          int
	y          int
}

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func iniciarlizar_gusano() gusano {
	m := gusano{
		tamaño:     3,
		id:         0,
		horizontal: true,
	}
	return m
}
func crear_map(mp [][]mapa, cantidad_gusano int, comida int) {

	for i := 0; i < cantidad_gusano; i++ {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(len(mp) - 2)
		y := rand.Intn(len(mp) - 2)
		if mp[x][y].activo == false { // si la casilla esta desocupada
			if mp[x+1][y].activo == false && mp[x+2][y].activo == false && x < 5 { // si las siguientes dos casillas en direccion abajo estan desocupadas llenar
				for j := 0; j < 3; j++ { //llene de forma vertical
					mp[x+j][y].guzanito = iniciarlizar_gusano()
					mp[x+j][y].guzanito.id = +i
					mp[x+j][y].activo = true
				}
				mp[x][y].guzanito.x = x
				mp[x][y].guzanito.y = y
			} else if mp[x][y+1].activo == false && mp[x][y+2].activo == false { //llene de forma horizontal
				for j := 0; j < 3; j++ {
					mp[x][y+j].guzanito = iniciarlizar_gusano()
					mp[x][y+j].guzanito.id = +i
					mp[x][y+j].activo = true
				}
				mp[x][y].guzanito.x = x
				mp[x][y].guzanito.y = y
			}
		} else if mp[x][y].activo == true { // si esta ocupada la casilla con un gusano
			for { // randomear hasta encontrar una casilla vacia
				s1 := rand.Intn(len(mp) - 2)
				s2 := rand.Intn(len(mp) - 2)
				if mp[s1][s2+1].activo == false && mp[s1][s2+2].activo == false && mp[s1][s2].activo == false {
					x = s1
					y = s2
					break
				}
			}
			for j := 0; j < 3; j++ { //llene de forma vertical
				mp[x][y+j].guzanito = iniciarlizar_gusano()
				mp[x][y+j].guzanito.id = +i
				mp[x][y+j].activo = true

			}
			mp[x][y].guzanito.x = x
			mp[x][y].guzanito.y = y
		} else {
			print("n")
		}
	}
	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp); j++ {
			if mp[j][i].activo == false { // si es 0 llenar con numeros random con limite entregado por consola
				s2 := rand.Intn(comida)
				mp[i][j].numero = s2
			} else {
				print("\nencontre casilla ocupada en = ", j, i)
			}
		}
	}
}

func color(dato string, numero int) string { //funcion para pasar cambiar de color
	print("codigo facilito")
}

func imprimir(mp [][]mapa) {

	for i := 0; i < len(mp); i++ {
		print("\n")
		for j := 0; j < len(mp); j++ {
			if mp[i][j].activo == true {
				print(ColorBlue, " ", mp[i][j].guzanito.id, " ")
				print(ColorReset)
			} else {
				print(" ", mp[i][j].numero, " ")
			}
		}
	}
}
func main() {

	gusanos, _ := strconv.Atoi(os.Args[1])
	x, _ := strconv.Atoi(os.Args[2])
	y, _ := strconv.Atoi(os.Args[3])
	comida, _ := strconv.Atoi(os.Args[4])

	print(gusanos, x, y, comida)

	mp := make([][]mapa, x)
	for i := 0; i < x; i++ {
		mp[i] = make([]mapa, y)
	}
	crear_map(mp, gusanos, comida)
	print("\n")
	imprimir(mp)
}
