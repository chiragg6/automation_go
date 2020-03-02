package server

// "github.com/Code_collection/routes"
// "github.com/Code_collection/routes"

import (
	"github.com/fix_automater/storage"
)

func StartServer() {
	storage.Initialize()
	// router := routes.Routes()

	// err := http.ListenAndServe(os.Getenv("HOST_ADDRESSES")+","+os.Getenv("PORT"), router)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
