package main

import (
	"fmt"
	"sort"
	"time"
)

const dateFormat = "2006-01-02"

type Booking struct {
	CheckInDate  *time.Time
	CheckOutDate *time.Time
}

type Event struct {
	Date          *time.Time
	OccupancyDiff int
}

func parseDateString(dateString string) *time.Time {
	result, _ := time.Parse(dateFormat, dateString)
	return &result
}

func readInput() (int, []*Booking) {
	totalBookings := 0
	fmt.Scanf("%d\n", &totalBookings)

	checkInDateInput := "2006-01-02"
	checkoutDateInput := "2006-01-02"

	bookings := make([]*Booking, totalBookings)
	for i := 0; i < totalBookings; i++ {
		booking := &Booking{}
		bookings[i] = booking

		fmt.Scanf("%s %s\n", &checkInDateInput, &checkoutDateInput)
		booking.CheckInDate = parseDateString(checkInDateInput)
		booking.CheckOutDate = parseDateString(checkoutDateInput)
	}

	return totalBookings, bookings
}

func sortBookingsByCheckInDate(bookings []*Booking) {
	sort.SliceStable(bookings, func(i int, j int) bool {
		return bookings[i].CheckInDate.Before(*bookings[j].CheckInDate)
	})
}

func sortBookingsByCheckOutDate(bookings []*Booking) {
	sort.SliceStable(bookings, func(i int, j int) bool {
		return bookings[i].CheckOutDate.Before(*bookings[j].CheckOutDate)
	})
}

func sortEventsByDateWithLowerOccupancyDiffFirst(events []*Event) {
	sort.SliceStable(events, func(i int, j int) bool {
		if events[i].Date.Before(*events[j].Date) {
			return true
		}
		if events[i].Date.After(*events[j].Date) {
			return false
		}
		return events[i].OccupancyDiff < events[j].OccupancyDiff
	})
}

func findMaxOccupancyDate(totalBookings int, bookings []*Booking) *time.Time {
	events := make([]*Event, totalBookings*2)

	sortBookingsByCheckInDate(bookings)
	for i, booking := range bookings {
		events[i] = &Event{
			Date:          booking.CheckInDate,
			OccupancyDiff: 1,
		}
	}

	sortBookingsByCheckOutDate(bookings)
	for i, booking := range bookings {
		events[totalBookings+i] = &Event{
			Date:          booking.CheckOutDate,
			OccupancyDiff: -1,
		}
	}

	sortEventsByDateWithLowerOccupancyDiffFirst(events)
	currentOccupancy := 0
	maxOccupancy := 0
	var maxOccupancyDate *time.Time
	for _, event := range events {
		currentOccupancy += event.OccupancyDiff
		if maxOccupancy < currentOccupancy {
			maxOccupancy = currentOccupancy
			maxOccupancyDate = event.Date
		}
	}

	return maxOccupancyDate
}

func main() {
	totalBookings, bookings := readInput()
	maxOccupancyDate := findMaxOccupancyDate(totalBookings, bookings)

	fmt.Printf(maxOccupancyDate.Format(dateFormat))
}
