package handler

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.BlogUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, blogUsecase interfaces.BlogUsecase) {
	handler := &handler{
		usecase: blogUsecase,
	}

	publicBlogHandler := publicApp.Group("/blog/")
	privateBlogHandler := privateApp.Group("/blog/")

	privateBlogHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateBlog()))
	publicBlogHandler.Get("list/", handler.ListBlog())
	privateBlogHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateBlog()))
	privateBlogHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteBlog()))
	publicBlogHandler.Get("details/:id/", handler.DetailsOfBlog())
	privateBlogHandler.Post("like-unlike/:id/", middleware.RolesAndPermissionMiddleware(handler.BlogLikeUnlike()))

	privateBlogHandler.Post("create-comment/:id/", middleware.RolesAndPermissionMiddleware(handler.CreateBlogComment()))
	publicBlogHandler.Get("list-comment/", handler.ListComments())
	privateBlogHandler.Patch("update-comment/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateBlogComment()))
	privateBlogHandler.Delete("delete-comment/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteBlogComment()))
}
