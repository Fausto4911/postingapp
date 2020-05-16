package main

import (
	"fmt"
	"postingapp/repository/postgres"
)

func main() {
	fmt.Println("======== Running main ==========")
	//e := model.Estudiante{
	//	Name: "Fausto",
	//	Active: true,
	//}
	//
	//err := postgres.EstudianteCreate(e)
	//if  err != nil {
	//	log.Fatal()
	//}

	//fmt.Println(" Object Estuiante created !")

	es, err := postgres.EstudianteGetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(es)
	
}
