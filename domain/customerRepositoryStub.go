package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Mert", City: "Trabzon", Zipcode: "61000", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Ali", City: "Trabzon", Zipcode: "61000", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1003", Name: "Veli", City: "Trabzon", Zipcode: "61000", DateofBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
