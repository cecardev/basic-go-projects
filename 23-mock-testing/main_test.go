package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "1",
						Name: "Cesar",
						Age:  23,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  23,
					DNI:  "1",
					Name: "Cesar",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ftEmployee, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error getting the employee")
		}
		if ftEmployee.Age != test.expectedEmployee.Age {
			t.Errorf("Error, expected %d got %d", ftEmployee.Age, test.expectedEmployee.Age)
		}

	}

	GetPersonByDNI = originalGetPersonByByDNI
	GetEmployeeById = originalGetEmployeeById

}
