package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/RachaelLuo/kex/pkg/apis/config"
	"github.com/RachaelLuo/kex/pkg/server/handlers/clusters"
	"github.com/RachaelLuo/kex/pkg/server/middleware/auth"
	"github.com/RachaelLuo/kex/pkg/server/middleware/monitor/prom"
	"github.com/RachaelLuo/kex/pkg/zone/clientset"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	cfg    *config.Config
	engine *gin.Engine
	client clientset.Interface
}

func Run(cfg *config.Config, client clientset.Interface) error {
	ctx, cancel := context.WithCancel(context.Background())

	s := &Server{
		ctx:    ctx,
		cancel: cancel,
		cfg:    cfg,
		client: client,
		engine: gin.Default(),
	}

	s.InstallHandlers()

	return s.engine.Run(":" + strconv.Itoa(s.cfg.Port))
}

func (s *Server) InstallHandlers() {
	if !s.cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		pprof.Register(s.engine)
	}

	// install healthz
	s.engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "luo luo happy everyday")
	})

	s.engine.Use(prom.PromMiddleware(nil), gin.Recovery())
	s.engine.GET("/metrics", prom.PromHandler(promhttp.Handler()))

	authorized := s.engine.Group("/", auth.MultiAuth(gin.Accounts{
		s.cfg.BasicAuthUser: s.cfg.BasicAuthPassword,
	}))

	clusters.InstallHandlers(authorized, s.cfg.NameSpace, s.cfg.ClusterInfos, s.cfg.LocalClusterInfos, s.client)
}
