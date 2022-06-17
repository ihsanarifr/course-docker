package user

import (
	"course/internal/domain"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(dbConn *gorm.DB) *UserService {
	return &UserService{
		db: dbConn,
	}
}

func (us UserService) Register(c *gin.Context) {
	var user domain.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}
	if user.Name == "" {
		c.JSON(400, gin.H{
			"message": "nama wajib di isi",
		})
		return
	}
	if user.Email == "" {
		c.JSON(400, gin.H{
			"message": "email wajib di isi",
		})
		return
	}
	if user.Password == "" {
		c.JSON(400, gin.H{
			"message": "password wajib di isi",
		})
		return
	}
	if len(user.Password) < 6 {
		c.JSON(400, gin.H{
			"message": "panjang password minimal 6 karakter",
		})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	err = us.db.Create(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "gagal registrasi. harap coba kembali",
		})
		return
	}
	token, _ := generateJWT(user.ID)
	c.JSON(201, gin.H{
		"token": token,
	})
}

func (us UserService) Login(c *gin.Context) {
	var userRequest domain.User
	err := c.ShouldBind(&userRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}
	if userRequest.Email == "" || userRequest.Password == "" {
		c.JSON(400, gin.H{
			"message": "email/password salah",
		})
		return
	}

	var user domain.User
	err = us.db.Where("email = ?", userRequest.Email).Take(&user).Error
	if err != nil || user.ID == 0 {
		c.JSON(400, gin.H{
			"message": "email/password salah",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "email/password salah",
		})
		return
	}
	token, _ := generateJWT(user.ID)
	c.JSON(200, gin.H{
		"token": token,
	})
}

var signature = []byte("Sup3rSecretK31")

func generateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "edspert",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signature)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (us UserService) DecriptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		}
		return signature, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}

	if !parsedToken.Valid {
		return data, errors.New("token invalid")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
