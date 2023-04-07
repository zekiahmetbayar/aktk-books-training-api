package books

import (
	"errors"
	"go-training/app/entities"
	"go-training/internal/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func Create(c *fiber.Ctx) error {
	item := entities.Book{}
	if err := c.BodyParser(&item); err != nil {
		return err
	}

	err := database.Connection().Model(&entities.Book{}).Create(&item).Error
	if err != nil {
		return err
	}

	return c.JSON(item)
}

func Index(c *fiber.Ctx) error {
	items := []entities.Book{}
	err := database.Connection().Model(&entities.Book{}).Preload(clause.Associations).Find(&items).Error
	if err != nil {
		return err
	}

	return c.JSON(items)
}

func Show(c *fiber.Ctx) error {
	if len(c.Params("id")) <= 0 {
		return errors.New("please set id as parameter")
	}

	item := entities.Book{}
	err := database.Connection().Model(&item).Preload(clause.Associations).Where("id = ?", c.Params("id")).First(&item).Error
	if err != nil {
		return err
	}

	return c.JSON(item)
}

func Update(c *fiber.Ctx) error {
	if len(c.Params("id")) <= 0 {
		return errors.New("please set id as parameter")
	}

	request := entities.Book{}
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	item := entities.Book{}
	err := database.Connection().Model(&item).Where("id = ?", c.Params("id")).First(&item).Error
	if err != nil {
		return err
	}

	err = database.Connection().Model(&item).Updates(&request).Error
	if err != nil {
		return err
	}

	return c.JSON(request)
}

func Delete(c *fiber.Ctx) error {
	if len(c.Params("id")) <= 0 {
		return errors.New("please set id as parameter")
	}

	int_value, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = database.Connection().Model(&entities.Book{}).Delete(&entities.Book{
		ID: int_value,
	}).Error
	if err != nil {
		return err
	}

	return c.JSON("Item deleted successfully.")
}
