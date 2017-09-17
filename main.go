package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"encoding/json"
)

type App struct{
	Client *redis.Client
}
func (a *App) init(){
	a.Client =  redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main(){

	a := App{}
	a.init()

	fmt.Printf("The response of the db is %s\n",a.Client.Ping())

	conn := &Connection{Uid:"Test", Number:1}
	fmt.Printf("The Uid is %s and the Number is %d\n",conn.Uid,conn.Number)

	serialized,err := json.Marshal(conn)

	if err != nil{
		fmt.Printf("Fuck there is an error")
	}

	fmt.Print(serialized)
	serializedString := string(serialized)
	fmt.Print(serializedString)

	fmt.Printf("The serialized string is %s\n ",serializedString)

	if err != nil{
		panic(err)
	}

	a.Client.Set("connection",serializedString,12121212121212)
	value,err :=a.Client.Get("connection").Result()

	conn = &Connection{}

	json.Unmarshal([]byte(value),conn)

	if err==nil{
		fmt.Printf("The result of the querry is %s \n",value)
		fmt.Printf("The Uid is %s and the Number is %d\n",conn.Uid,conn.Number)
	}
}

type Connection struct{
	Uid    string
	Number int
}

