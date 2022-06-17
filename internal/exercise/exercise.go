package exercise

import (
	"course/internal/domain"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExerciseService struct {
	db *gorm.DB
}

func NewExerciseService(dbConn *gorm.DB) *ExerciseService {
	return &ExerciseService{
		db: dbConn,
	}
}

func (es ExerciseService) GetExerciseByID(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid exercise id",
		})
		return
	}
	var exercise domain.Exercise
	err = es.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "exercise not found",
		})
		return
	}
	c.JSON(200, exercise)
}

func (es ExerciseService) GetUserScore(c *gin.Context) {
	paramID := c.Param("id")
	exerciseID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid exercise id",
		})
		return
	}
	userID := int(c.Request.Context().Value("user_id").(float64))

	var exercise domain.Exercise
	err = es.db.Where("id = ?", exerciseID).Preload("Questions").Take(&exercise).Error
	if err != nil || len(exercise.Questions) == 0 {
		c.JSON(404, gin.H{
			"message": "exercise not found",
		})
		return
	}

	var answers []domain.Answer
	err = es.db.Where("exercise_id = ? AND user_id = ?", exerciseID, userID).Find(&answers).Error
	if err != nil || len(answers) == 0 {
		c.JSON(200, gin.H{
			"score": 0,
		})
		return
	}

	mapQA := make(map[int]domain.Answer)
	for _, answer := range answers {
		mapQA[answer.QuestionID] = answer
	}

	var score int
	for _, question := range exercise.Questions {
		if strings.EqualFold(question.CorrectAnswer, mapQA[question.ID].Answer) {
			score += question.Score
		}
	}
	c.JSON(200, gin.H{
		"score": score,
	})
}
