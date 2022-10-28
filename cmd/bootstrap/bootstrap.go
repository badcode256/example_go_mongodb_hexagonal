package bootstrap

import (
	"context"
	//"database/sql"

	"log"

	"github.com/badcode256/example_go_mongodb_hexagonal/internal/infra/database/mongoDb/user"
	"github.com/badcode256/example_go_mongodb_hexagonal/internal/infra/server"
	"github.com/badcode256/example_go_mongodb_hexagonal/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//_ "github.com/denisenkom/go-mssqldb"
)

//ZcxyWxJm5OBzu336
//mongodb+srv://user_business:<password>@cluster0.dvbglbv.mongodb.net/?retryWrites=true&w=majority
//mongodb+srv://user_business:<password>@cluster0.dvbglbv.mongodb.net/test
//mongodb+srv://user_business:ZcxyWxJm5OBzu336@cluster0.dvbglbv.mongodb.net/dbBusiness?retryWrites=true\&w=majority

var uriDb = "mongodb+srv://user_business:ZcxyWxJm5OBzu336@cluster0.dvbglbv.mongodb.net/dbBusiness?retryWrites=true&w=majority"

func Start() error {

	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uriDb))

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	/*mySqlURI := fmt.Sprintf("server=%s;user id=%s;password=%s;port=1433;database=Business;", os.Getenv("SQL_SERVER"), os.Getenv("SQL_USER"), os.Getenv("SQL_PASSWORD"))
	db, err := sql.Open("sqlserver", mySqlURI)

	if err != nil {
		return err
	}*/

	userRepository := user.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	server := server.New(context.Background(), "localhost", 3000, userService)

	return server.Run()
}
