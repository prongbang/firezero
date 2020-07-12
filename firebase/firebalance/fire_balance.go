package firebalance

import (
	"context"
	"log"

	"endtry.com/travelio/config/firebase/fireutil"
	firemodel "endtry.com/travelio/grpc/firebase/model"
)

type FireBalance interface {
	GetCloud() *firemodel.FirebaseApp
}

type fireBalance struct {
	FireConfig *firemodel.FirebaseConfig
}

func NewFireBalance(fireConfig *firemodel.FirebaseConfig) FireBalance {
	return &fireBalance{
		FireConfig: fireConfig,
	}
}

func (f *fireBalance) GetCloud() *firemodel.FirebaseApp {
	fireCloud := fireutil.InitializeAppWithServiceAccount(f.FireConfig.ServiceAccountPath, f.FireConfig.DatabaseURL)
	client, err := fireCloud.Database(context.Background())
	if err != nil {
		log.Println(err)
	}
	firebaseApp := firemodel.FirebaseApp{
		App: fireCloud,
		DB:  client,
	}
	return &firebaseApp
}

func Initial() {

	// firecloud1 := fireutil.InitializeAppWithServiceAccount(os.Getenv("1CLOUD_SERVICE_ACCOUNT"), os.Getenv("1CLOUD_DATABASE_URL"))
	// client, err := firecloud1.Database(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var dbSize fireutil.FireSize
	// if err := client.NewRef("size").Get(context.Background(), &dbSize); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("dbSize: ", dbSize)

	// maxSize, err := strconv.ParseInt(os.Getenv("DATABASE_MAX_SIZE"), 10, 64)
	// if dbSize.Length > maxSize {
	// 	accessToken := fireutil.GetAccessToken(firecloud1, os.Getenv("1CLOUD_SERVICE_ACCOUNT"))
	// 	size := fireutil.GetDatabaseSize(os.Getenv("1APP_NAME"), accessToken)
	// 	log.Println("size:", size, " Kb")
	// } else {

	// }

	// uDataSource := datasource.NewUserRemoteDtaSource(client, context.Background())
	// uRepo := repository.NewUserRepository(uDataSource)
	// uUseCase := domain.NewUserUseCase(uRepo)
	// uPresenter := presenter.NewUserPresenter(uUseCase)

	// data := model.User{
	// 	ID:    "123456789",
	// 	First: "Hello",
	// 	Last:  "World",
	// 	Cloud: config.FIRE_CLOUD01,
	// }
	// err = uPresenter.Add(&data)
	// if err != nil {
	// 	log.Println(err)
	// }
}
