package main

import (
	"github.com/manveru/faker"

	"fmt"
	"math/rand"
	"time"
)

func main() {
	fake, _ := faker.New("en")
	fake.Rand = rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println("Name: ", fake.Name())
	fmt.Println("Country: ", fake.Country())
	fmt.Println("Email: ", fake.Email())
	fmt.Println("Lat: ", fake.Latitude())
	fmt.Println("Lng: ", fake.Longitude())
	fmt.Println("CompanyName: ", fake.CompanyName())
	fmt.Println("PostCode: ", fake.PostCode())
	fmt.Println("StreetAddress: ", fake.StreetAddress())
	fmt.Println("City: ", fake.City())
	fmt.Println("StreetName: ", fake.StreetName())
	fmt.Println("PhoneNumber: ", fake.PhoneNumber())
}
