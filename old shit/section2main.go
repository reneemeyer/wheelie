package main

import (
	"errors"
	"log"
)
import "fmt"
import "time"
import "github.com/reneemyer/wheelie/helpers"

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	birthDate   time.Time
}

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

type Animal interface {
	// Says list of functions that any Animal var must have
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func PrintInfo(a Animal) {
	fmt.Println("says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "rarrrrr"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(10)
	intChan <- randomNumber
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	result, err := divide(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result of division is", result)
	////dog := Dog{
	////	Name:  "samson",
	////	Breed: "Gernamn",
	////}
	////gorilla := Gorilla{
	////	Name:          "Jock",
	////	Color:         "black",
	////	NumberOfTeeth: 38,
	////}
	////PrintInfo(&dog)
	////PrintInfo(&gorilla)
	////var newVar3 helpers.SomeType
	////newVar3.TypeName = "megatron type"
	////log.Println(newVar3, "star")
	//
	////intChan := make(chan int)
	////defer close(intChan)
	////
	////go CalculateValue(intChan)
	////
	////num := <-intChan
	////log.Println(num, "radnommmmm")
	//
	////animals := []string{"dog", "fish", "hors"}
	//
	////animals := make(map[string]string)
	////animals["dog"] = "cooper"
	////var firstLine = "once upon a midnight dreary"
	//
	////var users []User
	////users = append(users, User{
	////	FirstName:   "Renee",
	////	LastName:    "Meyr",
	////	PhoneNumber: "51222222222",
	////	Age:         52,
	////})
	////users = append(users, User{
	////	FirstName:   "Mary",
	////	LastName:    "Brown",
	////	PhoneNumber: "51222222222",
	////	Age:         18,
	////})
	////for _, l := range users {
	////	log.Println(l.FirstName)
	////}
	////for animalType, animal := range animals {
	////	log.Println(animal, "insideeee", animalType)
	////}
	//myJson := `
	//	[
	//		{
	//			"first_name": "Clark",
	//			"last_name": "Kent",
	//			"hair_color": "brown",
	//			"has_dog": true
	//		},
	//		{
	//			"first_name": "Bruce",
	//			"last_name": "Wayne",
	//			"hair_color": "black",
	//			"has_dog": false
	//		}
	//	]
	//`
	//// json to a struct
	//var unmarshalled []Person
	//err := json.Unmarshal([]byte(myJson), &unmarshalled)
	//if err != nil {
	//	log.Println("Error unmarshalling", err)
	//}
	//log.Printf("unmarshalled: %v", unmarshalled)
	//// struct tp json
	//var mySlice []Person
	//var m1 Person
	//m1.FirstName = "wally"
	//m1.LastName = "west"
	//m1.HairColor = "red"
	//m1.HasDog = false
	//var m2 Person
	//m1.FirstName = "dianna"
	//m1.LastName = "prince"
	//m1.HairColor = "yellow"
	//m1.HasDog = false
	//mySlice = append(mySlice, m1)
	//mySlice = append(mySlice, m2)
	//
	//newJson, err := json.MarshalIndent(mySlice, "", "    ")
	//if err != nil {
	//	log.Println("error marshalling")
	//}
	//fmt.Println(string(newJson))

}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cant divide by 0")
	}
	result = x / y
	return result, nil
}

func saySomething() (string, string) {
	return "something", "else"
}

func saySomethingElse(s3 string) (string, string) {
	log.Println("recieved is", s)
	return s3, "else"
}
func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
