package repository

import (
	"database/sql"

	"github.com/bruna-barbosa-guimaraes/imersao-fullstack-fullcycle-golang/internal/events/domain"
)

type myssqlEventRepository struct {
	db *sql.DB
}

// func NewMysqlEventRepository(db *sql.DB) (domain.EventRepository, error) {
// 	return &NewMysqlEventRepository{db: db}, nil
// }

func (r *myssqlEventRepository) CreateSpot(spot *domain.Spot) error {
	query := `
		INSERT INTO spots (id, event_, name, status, ticket_id)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query, spot.ID, spot.EventID, spot.Name, spot.Status, spot.TicketID)
	return err
}

func (r *myssqlEventRepository) ReserveSpot(spotID, ticketID string) error {
	query := `
		UPDATE spots
		SET status = ?, ticket_id = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(query, domain.SpotStatusSold, ticketID, spotID)
	return err
}

func (r *myssqlEventRepository) CreateTicket(ticket *domain.Ticket) error {
	query := `
		INSERT INTO tickets (id, event_id, spot_id, ticket_type, price)
		VALEUS (?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query, ticket.ID, ticket.EventID, ticket.Spot.ID, ticket.TicketType, ticket.Price)
	return err
}

func (r *myssqlEventRepository) FindEventByID(eventID string) (*domain.Event, error) {
	query := `
		SELECT id, name, location, organization, rating, date, image_url, capacity, price, partner_id
		FROM events
		WHERE id = ?
	`

	rows, err := r.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var event *domain.Event
	err = rows.Scan(
		&event.ID,
		&event.Name,
		&event.Location,
		&event.Organization,
		&event.Rating,
		&event.Date,
		&event.ImageURL,
		&event.Capacity,
		&event.Price,
		&event.PartnerID,
	)

	if err != nil {
		return nil, err
	}

	return event, nil
}
