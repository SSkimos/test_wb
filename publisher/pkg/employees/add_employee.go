package employees

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"publisher/pkg/common/models"
)

type AddEmployeeRequestBody struct {
	EmployeeId  int    `json:"employee_id"`
	FirstName   string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Position    string `json:"position"`
	Department  string `json:"department"`
	HireDate    string `json:"hire_date"`
	Salary      int    `json:"salary"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phone_number"`
	Address     string `json:"address"`
}

func (h handler) AddEmployee(c *gin.Context) {
	body := AddEmployeeRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var employee models.Employee

	employee.EmployeeId = body.EmployeeId
	employee.FirstName = body.FirstName
	employee.Last_name = body.Last_name
	employee.Position = body.Position
	employee.Department = body.Department
	employee.HireDate = body.HireDate
	employee.Salary = body.Salary
	employee.Email = body.Email
	employee.PhoneNumber = body.PhoneNumber
	employee.Address = body.Address

	jsonData, err := json.Marshal(employee)
	subj, msg := "new employee", jsonData
	if err != nil {

	}

	h.con.Publish(subj, msg)

	h.con.Flush()

	if err := h.con.LastError(); err != nil {
		log.Panicf("Last error: %v", err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}

	c.JSON(http.StatusOK, &employee)
}