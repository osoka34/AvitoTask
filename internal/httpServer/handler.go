package httpServer

import (
	userHttp "AvitoTask/internal/account/delivery/http"
	userRepository "AvitoTask/internal/account/repository"
	userUsecase "AvitoTask/internal/account/usecase"
	"AvitoTask/internal/s_constant"
	segmentHttp "AvitoTask/internal/segment/delivery/http"
	segmentRepository "AvitoTask/internal/segment/repository"
	segmentUsecase "AvitoTask/internal/segment/usecase"
	"AvitoTask/pkg/storage"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	serverLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
	"os"
	"time"

	usersInSegHttp "AvitoTask/internal/users_in_segm/delivery/http"
	usersInSegRepository "AvitoTask/internal/users_in_segm/repository"
	usersInSegUsecase "AvitoTask/internal/users_in_segm/usecase"

	statHttp "AvitoTask/internal/statistics/delivery/http"
	statRepository "AvitoTask/internal/statistics/repository"
	statUsecase "AvitoTask/internal/statistics/usecase"
)

func (s *Server) MapHandlers(ctx context.Context, app *fiber.App, logger *zap.SugaredLogger) error {

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
	statRepo := statRepository.NewPostgresRepository(db, poolDb)

	// -------------------------------------------------------------------------------------

	statUC := statUsecase.NewStatUsecase(statRepo, logger, s.cfg)
	segmentUC := segmentUsecase.NewSegmentUsecase(segmentRepo, logger, s.cfg)
	userUC := userUsecase.NewUserUsecase(userRepo, logger, s.cfg)
	usersInSegUC := usersInSegUsecase.NewUsersInSegUsecase(usersInSegRepo, logger, s.cfg, statUC)

	// -------------------------------------------------------------------------------------

	segmentHandlers := segmentHttp.NewSegmentHandler(segmentUC, s.cfg, logger)
	userHandlers := userHttp.NewUserHandler(userUC, s.cfg, logger)
	usersInSegHandlers := usersInSegHttp.NewUsersInSegHandler(usersInSegUC, s.cfg, logger)
	statHandlers := statHttp.NewStatHandler(statUC, s.cfg, logger)

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
	statHttp.MapStatRoutes(app, statHandlers)

	// -------------------------------------------------------------------------------------

	go func(ctx context.Context) {

		ctx, cancel := context.WithCancel(ctx)

		for {
			select {
			case <-ctx.Done():
				logger.Error("error in goroutine with deleting by ttl")
				return
			default:
				toDelete, err := usersInSegUC.CountUsersWithExpiredTtl()
				if err != nil {
					cancel()
					continue
				}
				if toDelete {
					err = usersInSegUC.DeleteByTtl()
					if err != nil {
						cancel()
					}
				}
				time.Sleep(time.Minute * s_constant.DelayForDeleteByTtl)
			}
		}
	}(ctx)

	return nil
}
