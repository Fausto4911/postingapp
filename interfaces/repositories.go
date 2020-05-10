package interfaces

import "fmt"

type DbHandler interface {
	Execute(statement string)
	Query(statement string) Row
}

type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

type DbRepo struct {
	dbHandlers map[string]DbHandler
	dbHandler  DbHandler
}

type DbUserRepo DbRepo
type DbCustomerRepo DbRepo
type DbOrderRepo DbRepo
type DbItemRepo DbRepo

func NewDbUserRepo(dbHandlers map[string]DbHandler) *DbUserRepo {
	dbUserRepo := new(DbUserRepo)
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers["DbUserRepo"]
	return dbUserRepo
}

func (repo *DbUserRepo) Store(user usecases.User) {
	isAdmin := "no"
	if user.IsAdmin {
		isAdmin = "yes"
	}
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO users (id, customer_id, is_admin)
                                        VALUES ('%d', '%d', '%v')`,
		user.Id, user.Customer.Id, isAdmin))
	customerRepo := NewDbCustomerRepo(repo.dbHandlers)
	customerRepo.Store(user.Customer)
}
