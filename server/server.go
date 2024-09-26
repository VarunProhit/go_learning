package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	// "os"
"time"
	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Collection     *mongo.Collection
	Port   string
	Router *mux.Router
}

func (a *App) Initialize(){
	// err := godotenv.Load()
    // if err != nil {
	// 	log.Println(err)
    //     log.Fatalf("Error loading .env file")
    // }
//    uri := os.Getenv("MONGODB_URI")
uri := "mongodb+srv://20163:<pswd>@cluster0.tlpouqz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
   fmt.Println(uri)
  
  clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Ensure disconnection on function exit
    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
            log.Fatalf("Failed to disconnect from MongoDB: %v", err)
        }
    }()

    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatalf("Ping to MongoDB failed: %v", err)
    }

    fmt.Println("Connected to MongoDB!")

  a.Collection = client.Database("todo_list").Collection("vp")
  a.Router = mux.NewRouter()
  a.initializeRoutes()

}
func (a *App) initializeRoutes(){
	a.Router.HandleFunc("/task",a.GetRequest).Methods("GET")
}
func (a *App) Run(){
	fmt.Println("server on", a.Port)
	log.Fatal(http.ListenAndServe(a.Port,a.Router))
}

func (a *App)GetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get")
    // listTasks(a.Collection)
//     uri := "mongodb+srv://20163:<pswd>@cluster0.tlpouqz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
//    fmt.Println(uri)
  
//   clientOptions := options.Client().ApplyURI(uri)
//     client, err := mongo.Connect(context.TODO(), clientOptions)
//     if err != nil {
//         log.Fatalf("Failed to connect to MongoDB: %v", err)
//     }

//     // Ensure disconnection on function exit
//     defer func() {
//         if err = client.Disconnect(context.TODO()); err != nil {
//             log.Fatalf("Failed to disconnect from MongoDB: %v", err)
//         }
//     }()

//     // Check the connection
//     err = client.Ping(context.TODO(), nil)
//     if err != nil {
//         log.Fatalf("Ping to MongoDB failed: %v", err)
//     }

//     fmt.Println("Connected to MongoDB!")

//   coll := client.Database("todo_list").Collection("vp")
    // fmt.Println(coll)
    coll := a.Collection
    fmt.Println(a.Collection)
    todos := []todoModel{}
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()


    // Fetch all documents
    cursor, err := coll.Find(ctx, bson.D{})
    if err != nil {
       log.Println(err)
    }
    fmt.Println("dsss")
    defer cursor.Close(ctx)
    if err := cursor.All(ctx, &todos); err != nil {
     log.Println(err)
 }
 fmt.Println("todos",todos)
 todoList := []todo{}
 for _,t := range todos{
     todoList = append(todoList, todo{
         ID: t.ID.String(),
         Title: t.Title,
         Completed: t.Completed,
         CreatedAt: t.CreatedAt,
     })
 }
 for _,t :=range todoList{
     fmt.Println(t)
 }
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "post")
}
