package routes

import (
	"github.com/coderkamlesh/portfolio_backend/config" // Import config
	"github.com/coderkamlesh/portfolio_backend/internal/http/handler"
	"github.com/coderkamlesh/portfolio_backend/internal/http/middleware"
	"github.com/coderkamlesh/portfolio_backend/internal/repository"
	"github.com/coderkamlesh/portfolio_backend/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	cld := config.SetupCloudinary()

	// 2. Repositories
	userRepo := repository.NewUserRepository()
	projectRepo := repository.NewProjectRepository()
	skillRepo := repository.NewSkillRepository()
	expRepo := repository.NewExperienceRepository()

	// 3. Services (Ab yahan CLD pass karna hai)
	authService := service.NewAuthService(userRepo, cld)          // <-- Changed
	projectService := service.NewProjectService(projectRepo, cld) // <-- Changed
	skillService := service.NewSkillService(skillRepo)            // <-- Changed
	expService := service.NewExperienceService(expRepo)           // <-- Changed

	// 4. Handlers
	authHandler := handler.NewAuthHandler(authService)
	projectHandler := handler.NewProjectHandler(projectService)
	skillHandler := handler.NewSkillHandler(skillService)
	expHandler := handler.NewExperienceHandler(expService)
	// New Upload Handler
	uploadHandler := handler.NewUploadHandler(cld)

	api := r.Group("/api")
	{
		// Public
		api.POST("/login", authHandler.Login)
		api.GET("/hero", authHandler.GetHeroInfo)
		api.GET("/projects", projectHandler.GetAll)
		api.GET("/skills", skillHandler.GetAll)
		api.GET("/experiences", expHandler.GetAll)
		api.POST("/setup", authHandler.SetupAdmin)

		// Admin Protected
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware())
		{
			// 1. Upload API (Generic for Hero, Projects, Skills)
			admin.POST("/upload", uploadHandler.UploadFile)

			// 2. Update Hero Info
			admin.PUT("/hero", authHandler.UpdateHero)

			// 3. Projects CRUD
			admin.POST("/projects", projectHandler.Create)
			admin.PUT("/projects/:id", projectHandler.Update)
			admin.DELETE("/projects/:id", projectHandler.Delete)

			// 4. Skills CRUD
			admin.POST("/skills", skillHandler.Create)
			admin.PUT("/skills/:id", skillHandler.Update)
			admin.DELETE("/skills/:id", skillHandler.Delete)

			// 5. Experience CRUD
			admin.POST("/experiences", expHandler.Create)
			admin.PUT("/experiences/:id", expHandler.Update)
			admin.DELETE("/experiences/:id", expHandler.Delete)

			admin.PATCH("/hero/avatar", authHandler.UpdateAvatar)

			// 2. Project Image
			admin.PATCH("/projects/image/:id", projectHandler.UpdateImage)
		}
	}
}
