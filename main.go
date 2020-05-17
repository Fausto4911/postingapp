package main

import (
	"context"
	"fmt"
	"log"
	"postingapp/repository/postgres"
)
var ctx context.Context
func main() {
	fmt.Println("======== Running main ==========")
    ctx = context.Background()
	repo := postgres.IEstudianteRepository{
		Ctx: ctx,
	}

	e, err := repo.GetById(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
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

}
