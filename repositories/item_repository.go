package repositories

import (
	"simple-backend/config"
	"simple-backend/models"

	"gorm.io/gorm"
)

func GetAllItems() ([]models.Item, error) {
	var items []models.Item
	result := config.DB.Find(&items)
	return items, result.Error
}

func CreateItem(item *models.Item) error {
	result := config.DB.Create(item)
	return result.Error
}

// Полное обновление элемента (PUT)
func UpdateItem(id uint, updatedItem *models.Item) error {
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return err
	}
	// Полностью обновляем элемент, заменяя все поля
	return config.DB.Model(&item).Updates(updatedItem).Error
}

// Частичное обновление элемента (PATCH)
func PartialUpdateItem(id uint, updatedFields map[string]interface{}) error {
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return err
	}
	// Обновляем только указанные поля
	return config.DB.Model(&item).Updates(updatedFields).Error
}
func DeleteItem(id uint) error {
	var item models.Item
	result := config.DB.Delete(&item, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
