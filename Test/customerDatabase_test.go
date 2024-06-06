
package test

import (
	"testing"

	"assesment.com/client"
)

func TestCreateEmployee(t *testing.T) {
	emp := client.CustomerDetails{ID: 1, Name: "John Doe", Position: "Developer", Salary: 60000}
	if err := client.InsertCustomer(emp); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

}

func TestGetEmployeeByID(t *testing.T) {
	emp := client.CustomerDetails{ID: 1, Name: "John Doe", Position: "Developer", Salary: 60000}
	client.InsertCustomer(emp)
	if _, err := client.GetCustomerByID(1); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestUpdateEmployee(t *testing.T) {
	emp := client.CustomerDetails{ID: 1, Name: "John Doe", Position: "Developer", Salary: 60000}
	client.InsertCustomer(emp)
	emp.Name = "Jane Doe"
	if err := client.UpdateCustomer(emp); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if emp, err := client.GetCustomerByID(1); err != nil || emp.Name != "Jane Doe" {
		t.Errorf("expected updated employee, got %v, error %v", emp, err)
	}
}

func TestDeleteEmployee(t *testing.T) {

	emp := client.CustomerDetails{ID: 1, Name: "John Doe", Position: "Developer", Salary: 60000}
	client.InsertCustomer(emp)
	if err := client.DeleteCustomerByID(1); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

}
