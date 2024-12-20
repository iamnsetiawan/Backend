package model

type TicketResponse struct {
	ID         string         `json:"id"`
	EventID    uint           `json:"event_id"`
	OrderID    uint           `json:"order_id"`
	Price      float64        `json:"price"`
	Type       string         `json:"type"`
	SeatNumber string         `json:"seat_number"`
	Event      *EventResponse `json:"event,omitempty"`
	Order      *OrderResponse `json:"order,omitempty"`
}

type CreateTicketRequest struct {
	EventID uint    `json:"event_id" validate:"required"`
	Price   float64 `json:"price" validate:"required"`
	Type    string  `json:"type" validate:"required,oneof=vip regular VIP REGULAR"`
	Count   int     `json:"count" validate:"numeric,required,min=1"`
}

type UpdateTicketRequest struct {
	ID         string  `param:"id" validate:"required"`
	EventID    uint    `json:"event_id" validate:"numeric,omitempty"`
	OrderID    *uint   `json:"order_id,omitempty" validate:"numeric,omitempty"`
	Price      float64 `json:"price" validate:"numeric,omitempty"`
	Type       string  `json:"type" validate:"omitempty"`
	SeatNumber string  `json:"seat_number" validate:"omitempty"`
}

type GetTicketRequest struct {
	ID string `param:"id" validate:"required"`
}

type TicketsRequest struct {
	Page  int    `query:"page" validate:"numeric"`
	Size  int    `query:"size" validate:"numeric"`
	Sort  string `query:"sort" validate:"omitempty,oneof=id event_id eventId order_id orderID price type seat_number seatNumber"`
	Order string `query:"order" validate:"omitempty"`
}

type TicketSearchRequest struct {
	ID         string  `query:"id" validate:"omitempty"`
	EventID    uint    `query:"event_id" validate:"omitempty"`
	OrderID    uint    `query:"order_id" validate:"omitempty"`
	Price      float64 `query:"price" validate:"omitempty"`
	Type       string  `query:"type" validate:"omitempty"`
	SeatNumber string  `query:"seat_number" validate:"omitempty"`
	Page       int     `query:"page" validate:"numeric"`
	Size       int     `query:"size" validate:"numeric"`
	Sort       string  `query:"sort" validate:"omitempty,oneof=id event_id eventId order_id orderID price type seat_number seatNumber"`
	Order      string  `query:"order" validate:"omitempty"`
}

type TicketQueryOptions struct {
	ID          *string
	EventID     *uint
	OrderID     *uint
	Price       *float64
	Type        *string
	SeatNumbers []string `json:"seat_numbers,omitempty"`
	Page        int
	Size        int
	Sort        string
	Order       string
}
