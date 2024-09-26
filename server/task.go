package server

import(
//    "bufio"
	"context"
	// "encoding/hex"
	// "fmt"
	// "os"
	// "strconv"
	"time"

	// "encoding/json"
	"log"

	// "github.com/joho/godotenv"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)
type(
	todoModel struct{
		ID primitive.ObjectID `bson:"_id,omitEmpty"`
		Title string `bson:"title"`
		Completed bool `bson:"completed"`
		CreatedAt time.Time `bson:"createdAt"`
		UpdatedAt time.Time `bson:"updatedAt"`
	
	}
	todo struct{
		ID string `json:"id"`
		Title string `json:"title"`
		Completed bool `json:"completed"`
		CreatedAt time.Time `json:"createdAt`
		UpdatedAt time.Time `json:"updatedAt"`
	}
	)

	func listTasks(coll *mongo.Collection){
		// fmt.Println(coll)

	   todos := []todoModel{}
	   ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	   defer cancel()
	
	//    // Fetch all documents
	   cursor, err := coll.Find(ctx, bson.D{})
	   if err != nil {
		  log.Println("e1",err)
	   }
	   defer cursor.Close(ctx)
	   if err := cursor.All(ctx, &todos); err != nil {
		log.Println("e2",err)
	}
	// todoList := []todo{}
	// for _,t := range todos{
	// 	todoList = append(todoList, todo{
	// 		ID: t.ID.String(),
	// 		Title: t.Title,
	// 		Completed: t.Completed,
	// 		CreatedAt: t.CreatedAt,
	// 	})
	// }
	// for _,t :=range todoList{
	// 	fmt.Println(t)
	// }
	}
	

	