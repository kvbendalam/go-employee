package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kvbendalam/goemployee/database"
	"github.com/kvbendalam/goemployee/models"
)

func setupRoutes(app *fiber.App) {
	app.Get("/emp", ListEmployees)
	app.Get("/project", ListProjects)
	app.Get("/manager", ListManagers)
	app.Get("/empSalaries", getEmployeeSalaries)

	app.Get("/emp/:id", GetEmployeeById)

	app.Post("/emp", CreteEmployee)
	app.Post("/manager", CreateManager)
	app.Post("/project", CreateProject)

	app.Put("/updateemployee/:id", UpdateEmployee)

	app.Delete("/deleteemployee/:id", DeleteEmployee)

}

func ListEmployees(c *fiber.Ctx) error {
	employees := []models.Employee{}
	database.DB.Db.Find(&employees)
	return c.Status(200).JSON(employees)
}

func ListProjects(c *fiber.Ctx) error {
	projects := []models.Project{}
	database.DB.Db.Find(&projects)
	return c.Status(200).JSON(projects)
}

func ListManagers(c *fiber.Ctx) error {
	managers := []models.Manager{}
	database.DB.Db.Find(&managers)
	return c.Status(200).JSON(managers)
}

func getEmployeeSalaries(c *fiber.Ctx) error {
	var results []map[string]interface{}
	database.DB.Db.Model(&models.Employee{}).Select("name, salary").Scan(&results)
	fmt.Println(results)
	return c.Status(200).JSON(results)
}

func GetEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee models.Employee

	result := database.DB.Db.Find(&employee, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(employee)
}

func CreteEmployee(c *fiber.Ctx) error {
	employee := new(models.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&employee)
	return c.Status(200).JSON(employee)
}

func CreateManager(c *fiber.Ctx) error {
	manager := new(models.Manager)

	if err := c.BodyParser(manager); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&manager)
	return c.Status(200).JSON(manager)
}

func CreateProject(c *fiber.Ctx) error {
	project := new(models.Project)

	if err := c.BodyParser(project); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&project)
	return c.Status(200).JSON(project)
}

func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee models.Employee

	if result := database.DB.Db.Find(&employee, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	database.DB.Db.Delete(&employee)

	return c.Status(200).JSON(&employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee models.Employee

	database.DB.Db.Find(&employee, "id =?", id)

	if employee.ID == " " {
		return c.Status(404).JSON(fiber.Map{"status": "No data found"})
	}

	var updateEmployee models.Employee

	err := c.BodyParser(&updateEmployee)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Internal server error"})
	}

	employee.ID = updateEmployee.ID
	employee.Name = updateEmployee.Name
	employee.Salary = updateEmployee.Salary
	employee.ProjectID = updateEmployee.ProjectID
	employee.ProjectName = updateEmployee.ProjectName
	employee.Duration = updateEmployee.Duration
	employee.Billing = updateEmployee.Billing
	employee.NoOfEmpWork = updateEmployee.NoOfEmpWork
	employee.Technologies = updateEmployee.Technologies

	database.DB.Db.Save(&employee)
	return c.JSON(fiber.Map{"data": employee})

}
