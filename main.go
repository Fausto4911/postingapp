package main

import (
	"fmt"
	"log"
	"postingapp/model"
	"postingapp/repository/postgres"
)

func main() {
	fmt.Println("======== Running main ==========")
	e := model.Estudiante{
		Name: "Alejandro",
		Age:  45,
		Active: true,
	}
	
	err := postgres.EstudianteCreate(e)
	if  err != nil {
		log.Fatal()
	}

	fmt.Println(" Object Estuiante created !")
	
}
