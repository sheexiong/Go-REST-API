package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sheexiong/Go-REST-API/bin/models"
	"github.com/sheexiong/Go-REST-API/bin/store"
)

func GetProperties(c *gin.Context) {
	// Connection to the database
	db := store.InitDb()
	// Close connection database
	defer db.Close()

	var properties []models.Property
	// SELECT * FROM properties
	db.Find(&properties)

	// Display JSON result
	c.JSON(200, properties)
}

func GetProperty(c *gin.Context) {
	// Connection to the database
	db := store.InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var property models.Property
	// SELECT * FROM properties WHERE id = 1;
	db.First(&property, id)

	if property.ID != 0 {
		// Display JSON result
		c.JSON(200, property)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Property not found"})
	}
}

func GetCountries(c *gin.Context) {
	// Connection to the database
	db := store.InitDb()
	// Close connection database
	defer db.Close()

	var countries []models.Country
	// SELECT * FROM countries
	db.Find(&countries)

	// Display JSON result
	c.JSON(200, countries)
}

func GetCountry(c *gin.Context) {
	// Connection to the database
	db := store.InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var country models.Country
	// SELECT * FROM countries WHERE id = 1;
	db.First(&country, id)

	if country.ID != 0 {
		// Display JSON result
		c.JSON(200, country)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Country not found"})
	}
}

func PostProperty(c *gin.Context) {
	db := store.InitDb()
	defer db.Close()

	var property models.Property
	var country models.Country
	c.Bind(&property)

	if property.PostType == "" && property.PropertyType == "" && property.Addr == "" && property.Bathrooms == 0 {
		// Display error
		c.JSON(422, gin.H{"error": "Some fields are empty or wrong"})
	} else {
		db.Where("country = ?", property.Country).First(&country)
		if country.ID == 0 {
			c.JSON(422, gin.H{"error": "Invalid Country"})
		} else {
			db.Create(&property)
			c.JSON(201, gin.H{"success": property})
		}
	}
}

func PostCountry(c *gin.Context) {
	db := store.InitDb()
	defer db.Close()

	var country models.Country
	c.Bind(&country)

	if country.Country == "" {
		// Display error
		c.JSON(422, gin.H{"error": "Some fields are empty or wrong"})
	} else {
		result := db.Create(&country)
		if result.RowsAffected == 0 {
			c.JSON(422, gin.H{"error": "The country already exist"})
		} else {
			c.JSON(201, gin.H{"success": country})
		}
	}
}

func UpdateProperty(c *gin.Context) {
	// Connection to the database
	db := store.InitDb()
	// Close connection database
	defer db.Close()

	// Get id
	id := c.Params.ByName("id")
	var property models.Property
	db.First(&property, id)

	if property.PostType != "" && property.PropertyType != "" && property.Addr != "" && property.Bathrooms != 0 {

		if property.ID != 0 {
			var newProperty models.Property
			c.Bind(&newProperty)

			var country models.Country
			db.Where("country = ?", newProperty.Country).First(&country)
			if country.ID == 0 {
				c.JSON(422, gin.H{"error": "Invalid Country"})
			} else {
				result := models.Property{
					ID:           property.ID,
					PostType:     newProperty.PostType,
					Price:        newProperty.Price,
					Desc:         newProperty.Desc,
					PropertyType: newProperty.PropertyType,
					UnitSize:     newProperty.UnitSize,
					Addr:         newProperty.Addr,
					Bedrooms:     newProperty.Bedrooms,
					Bathrooms:    newProperty.Bathrooms,
					Country:      newProperty.Country,
				}

				db.Save(&result)
				// Display modified data in JSON message "success"
				c.JSON(200, gin.H{"success": result})
			}

		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Property not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func DeleteProperty(c *gin.Context) {
	// Connection to the database
	db := store.InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var property models.Property
	db.First(&property, id)

	if property.ID != 0 {
		db.Delete(&property)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Property #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Property not found"})
	}
}
