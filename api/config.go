package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// function to get the Twilio sid
func envACCOUNTSID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatalf("Error while loading .env")
	}
	//use the built in os package to get the variable
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

// function to load auth token
func envAuthToken() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatalf("Error while loading .env")
	}
	return os.Getenv("TWILIO_AUTHTOKEN")
}

// load service id.....
func envSERVICEID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatalf("Error while loading .env")
	}
	return os.Getenv("TWILIO_SERVICE_ID")
}
