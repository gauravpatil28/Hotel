package dbrepo

import (
	"errors"
	"time"

	"github.com/gauravpatil28/booking/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID == 2 {
		return 0, errors.New("some error")
	}
	return 1, nil
}

func (m *testDBRepo) InsertRoomRestriciton(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("Error")
	}
	return nil
}

func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {

	return false, nil
}

func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {

	var room models.Room

	if id > 2 {
		return room, errors.New("Can't find room")
	}

	return room, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) GetUserById(id int) (models.User, error) {

	var u models.User

	return u, nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {

	return 0, "", nil
}

func (m *testDBRepo) AllReservation() ([]models.Reservation, error) {

	var reservation []models.Reservation

	return reservation, nil
}

func (m *testDBRepo) AllNewReservation() ([]models.Reservation, error) {
	var reservation []models.Reservation

	return reservation, nil
}

func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {

	var reservation models.Reservation
	return reservation, nil
}

func (m *testDBRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {

	return nil

}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

func (m *testDBRepo) GetRestrictionsForRoomByDate(roomId int, start, end time.Time)([]models.RoomRestriction, error) {
	var restrictions []models.RoomRestriction

	return restrictions, nil
}
