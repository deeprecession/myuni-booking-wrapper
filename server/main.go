package main

import (
	"inno-booking-wrapper/pkg/bookings"
	"inno-booking-wrapper/pkg/myuniversity"
	"time"

	"go.uber.org/zap"
)

func main() {
	log := zap.Must(zap.NewDevelopment())

	email := ""
	password := ""

	log.Sugar().Infow("Getting my.university cookies", "email", email)

	myUniversityCookies, err := myuniversity.GetMyUniversityCookies(email, password, log)
	if err != nil {
		log.Sugar().Fatalf("failed to get uni cookies: %v", err)
	}

	log.Sugar().Infow("Got my.univesity cookies!")

	start := time.Now()
	end := time.Now().Add(time.Hour * 24)

	_, err = bookings.GetRooms(myUniversityCookies, start, end)
	if err != nil {
		log.Sugar().Fatalf("failed to get rooms with bookings: %v", err)
	}

	log.Sugar().Infow("Got bookings for every room!")
}
