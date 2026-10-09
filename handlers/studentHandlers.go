package handlers

import (
	"db-api/models"
	"db-api/repository"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	repo repository.StudentRepository
}

func NewStudentHandler(repo repository.StudentRepository) *StudentHandler {
	return &StudentHandler{repo: repo}
}

func (h *StudentHandler) GetAll(c *fiber.Ctx) error {
	students, err := h.repo.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(students)
}

func (h *StudentHandler) GetByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid ID"})
	}

	student, err := h.repo.GetByID(id)
	if err != nil {
		log.Printf("error: %v", err)
		return c.Status(404).JSON(fiber.Map{"error": "student not found"})
	}
	return c.Status(200).JSON(student)
}

func (h *StudentHandler) Create(c *fiber.Ctx) error {
	student := new(models.Student)
	if err := c.BodyParser(student); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "student not found"})
	}

	if err := h.repo.Create(student); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(student)
}

func (h *StudentHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalide ID"})
	}

	var student models.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalide data format"})
	}

	if err := h.repo.Update(id, &student); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(student)
}

func (h *StudentHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalide ID"})
	}

	if err := h.repo.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "the student was successfully deleted."})
}
