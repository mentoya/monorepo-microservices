package httphandler

import "github.com/irahardianto/monorepo-mocroservices/bookings/model"

type (
	// For Get - /bookings
	BookingsResource struct {
		Data []model.Booking `json:"data"`
	}
	// For Post/Put - /bookings
	BookingResource struct {
		Data model.Booking `json:"data"`
	}
)
