package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"zargram"
	"zargram/configs"
	"zargram/db"
	"zargram/pkg/handlers"
	"zargram/pkg/repository"
	"zargram/pkg/service"
)

//1. CRUD для пользователей /
//2. CRUD для постов  /
//3. CRUD для сторисов  /
//4. CRUD для комментариев  /
//5. Подписаться на пользователя  /
//6. Отписаться от пользователя   /
//7. CRUD для подписчиков       /
//8. Архивировать пост         /
//9. Ставить реакции на посты
//10. Показ количества реакции
//11. Прикрепить аватарку     /
//12. Сохранить пост другого пользователя
//13. Поиск пользователя    /
import (
	"context"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	//reading from yaml
	if err := InitConfigs(); err != nil {
		log.Fatalf("error while reading config file. error is %v", err.Error())
	}

	var cfg configs.DatabaseConnConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Couldn't unmarshal the config into struct. error is %v", err.Error())
	}

	cfg.Password = os.Getenv("DB_PASSWORD")

	conn := repository.GetDBConnection(cfg)

	db.Init(conn)

	//---------- Dependency injection-----------
	repos := repository.NewRepository(conn)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)
	//--------------------------------------------

	server := new(zargram.Server)
	go func() {
		if err := server.Run(os.Getenv("HTTP_PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error while running http.server. Error is %s", err.Error())
		}
	}()

	fmt.Printf("Server is listening to port: %s\n", os.Getenv("PORT"))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	fmt.Println("server is shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error while shutting server down. Error: %s", err.Error())
	}

	repository.Close(conn)
}

func InitConfigs() error {
	viper.AddConfigPath("configs") //адрес директории
	viper.SetConfigName("config")  //имя файла
	viper.SetConfigType("yml")
	return viper.ReadInConfig() //считывает config и сохраняет данные во внутренний объект viper
}
