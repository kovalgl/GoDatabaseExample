package models

/*
Implements IBookManager interface
*/
type FakeBookManager struct {}

func (fbm FakeBookManager) AllBooks() ([]*Book, error) {
	bks := make([]*Book, 0)
	bks = append(bks, &Book{"978-1503261969", "Emma", "Jayne Austen", 9.440000})
	bks = append(bks, &Book{"978-1503379640", "The Prince", "Niccolò Machiavelli", 6.990000})

	return bks, nil
}