package Main

import (
	"log"
	"net/http"
	"sync"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Controller"
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/ConfigSetup"
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/FieldName"
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Model"
	Router "github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Routes"
	"github.com/gorilla/mux"
)

func main() {

	ConfigSetup.InitateSetup(FieldName.Dev)

	wg := sync.WaitGroup{}

	DbProc := Model.DBProcessor{}
	RedisProc := Model.RedisProcessor{}

	go Model.NewDBProcessor(&wg, &DbProc)
	wg.Add(1)

	go Model.NewRedisInstance(&wg, &RedisProc)
	wg.Add(1)

	wg.Wait()

	ctrl := Controller.NewController(DbProc, RedisProc)

	muxRouter := mux.NewRouter()

	CustomRouter := Router.NewCustomRouter(muxRouter, ctrl)

	CustomRouter.UpdateCustomRouter()

	http.Handle("/", muxRouter)

	err := http.ListenAndServe(":8080", muxRouter)

	if err != nil {
		log.Fatal(err)
	}

}
