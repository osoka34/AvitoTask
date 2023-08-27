package httpServer

import (
	userHttp "AvitoTask/internal/account/delivery/http"
	userRepository "AvitoTask/internal/account/repository"
	userUsecase "AvitoTask/internal/account/usecase"
	"AvitoTask/pkg/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	serverLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
	"os"

	segmentHttp "AvitoTask/internal/segment/delivery/http"
	segmentRepository "AvitoTask/internal/segment/repository"
	segmentUsecase "AvitoTask/internal/segment/usecase"

	usersInSegHttp "AvitoTask/internal/users_in_segm/delivery/http"
	usersInSegRepository "AvitoTask/internal/users_in_segm/repository"
	usersInSegUsecase "AvitoTask/internal/users_in_segm/usecase"
)

func (s *Server) MapHandlers(app *fiber.App, logger *zap.SugaredLogger) error {

	db, err := storage.InitPsqlDB(s.cfg)
	if err != nil {
		log.Info(err)
		logger.Fatalf("err is: %v", err)
	}

	// -------------------------------------------------------------------------------------

	poolDb, err := storage.InitConnectionPoolPsqlDB(s.cfg)
	if err != nil {
		logger.Fatalf("err is: %v", err)
	}

	// -------------------------------------------------------------------------------------

	segmentRepo := segmentRepository.NewPostgresRepository(db, poolDb)
	userRepo := userRepository.NewPostgresRepository(db, poolDb)
	usersInSegRepo := usersInSegRepository.NewPostgresRepository(db, poolDb)

	// -------------------------------------------------------------------------------------

	segmentUC := segmentUsecase.NewSegmentUsecase(segmentRepo, logger, s.cfg)
	userUC := userUsecase.NewUserUsecase(userRepo, logger, s.cfg)
	usersInSegUC := usersInSegUsecase.NewUsersInSegUsecase(usersInSegRepo, logger, s.cfg)

	// -------------------------------------------------------------------------------------

	segmentHandlers := segmentHttp.NewSegmentHandler(segmentUC, s.cfg, logger)
	userHandlers := userHttp.NewUserHandler(userUC, s.cfg, logger)
	usersInSegHandlers := usersInSegHttp.NewUsersInSegHandler(usersInSegUC, s.cfg, logger)

	// -------------------------------------------------------------------------------------

	app.Use(serverLogger.New())
	if _, ok := os.LookupEnv("LOCAL"); !ok {
		app.Use(recover.New())
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// -------------------------------------------------------------------------------------

	segmentHttp.MapSegmentRoutes(app, segmentHandlers)
	userHttp.MapUserRoutes(app, userHandlers)
	usersInSegHttp.MapUsersInSegRoutes(app, usersInSegHandlers)

	// -------------------------------------------------------------------------------------

	return nil
}
