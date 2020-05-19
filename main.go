package main

import (
	"context"
	"fmt"
	"log"
	"postingapp/model"
	"postingapp/repository/postgres"
	"time"
)

var ctx context.Context

func main() {
	fmt.Println("*** Running main ***")
	ctx = context.Background()
	repo := postgres.IEstudianteRepository{
		Ctx: ctx,
	}

	e, err := repo.GetById(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("======== GetById ==========")
	fmt.Println(e)

	list, err := repo.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("======== GetAll ==========")
	fmt.Println(list)

	m := model.Estudiante{
		ID:        1,
		Name:      "PEPE",
		Age:       99,
		UpdatedAt: time.Now(),
	}
	fmt.Println("======== Update ==========")
	err = repo.Update(m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("======== Store ==========")
	es := model.Estudiante {
		Name: "Emmanuel",
		Active: true,
		Age: 25,
	}

	err = repo.Store(es)

	if err != nil {
		log.Fatal(err)
	}
	//
	//err := postgres.EstudianteCreate(e)
	//if  err != nil {
	//	log.Fatal()
	//}

	//fmt.Println(" Object Estuiante created !")

}
