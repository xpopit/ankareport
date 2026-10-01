package handlers

import (
	"backend/models"
	"backend/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetReports(c *fiber.Ctx) error {
	reports, err := repositories.GetReports()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(reports)
}

func GetReport(c *fiber.Ctx) error {
	id := c.Params("id")
	report, err := repositories.GetReport(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Report not found"})
	}
	return c.JSON(report)
}

func CreateReport(c *fiber.Ctx) error {
	var body struct {
		Report     models.Report           `json:"report"`
		Definition models.ReportDefinition `json:"definition"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if body.Report.TenantID == "" {
		body.Report.TenantID = "default"
	}

	err := repositories.CreateReport(&body.Report, &body.Definition)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(body.Report)
}

func UpdateReport(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		Report     models.Report           `json:"report"`
		Definition models.ReportDefinition `json:"definition"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := repositories.UpdateReport(id, &body.Report, &body.Definition)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success"})
}

func DeleteReport(c *fiber.Ctx) error {
	id := c.Params("id")
	err := repositories.DeleteReport(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "deleted"})
}

func CloneReport(c *fiber.Ctx) error {
	id := c.Params("id")
	newReport, err := repositories.CloneReport(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(newReport)
}
