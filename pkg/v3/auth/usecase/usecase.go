package usecase

import (
	"avyaas/internal/domain/interfaces"

	"github.com/go-redis/redis"
)

/*
usecase represents the authentication usecase, which contains the necessary components for handling
authentication-related business logic. It includes an authentication repository for data access and
a Redis client for token management and storage.
*/
type usecase struct {
	repository  interfaces.AuthRepository
	accountRepo interfaces.AccountRepository
	courseRepo  interfaces.CourseRepository
	redisClient *redis.Client
}

/*
New initializes and returns a new instance of the authentication usecase. It takes an authentication
repository and a Redis client as parameters. The usecase is responsible for handling business logic
related to authentication.
*/
func New(repo interfaces.AuthRepository, accountRepo interfaces.AccountRepository, courseRepo interfaces.CourseRepository, redisClient *redis.Client) *usecase {
	return &usecase{
		repository:  repo,
		accountRepo: accountRepo,
		courseRepo:  courseRepo,
		redisClient: redisClient,
	}
}
