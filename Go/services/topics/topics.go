package topics

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	. "pyro.com/codefunn/models"
)

// GetAllTopics Fetch all topics from DB and return
func GetAllTopics(db *gorm.DB) []Topic {
	var topics []Topic
	db.Find(&topics)
	return topics
}

func GetTopicById(db *gorm.DB, id string) Topic {
	var topic Topic
	db.Find(&topic, "ID=?", id)
	return topic
}

func AddTopic(db *gorm.DB, ctx *fiber.Ctx) (*Topic, error) {
	topic := new(Topic)

	err := ctx.BodyParser(topic)
	if err != nil {
		log.Fatal("AddTopic: BodyParser Error: ", err)
		return topic, err
	}
	db.Create(&topic)

	return topic, nil
}

func UpdateTopic(){

}

func DeleteTopic(){

}
