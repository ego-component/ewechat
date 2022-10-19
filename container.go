package ewechat

import (
	"sync"

	"github.com/ego-component/eredis"
	"github.com/ego-component/ewechat/cache"
	"github.com/ego-component/ewechat/context"
	"github.com/gotomicro/ego/client/ehttp"
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/core/elog"
)

type Container struct {
	config *config
	name   string
	ctx    *context.Context
	client cache.Cache
	logger *elog.Component
}

func DefaultContainer() *Container {
	return &Container{
		config: DefaultConfig(),
		logger: elog.EgoLogger.With(elog.FieldComponent(ModName)),
	}
}

// Load ...
func Load(key string) *Container {
	container := DefaultContainer()
	if err := econf.UnmarshalKey(key, container.config); err != nil {
		container.logger.Panic("parse wechat config panic",
			elog.FieldErr(err), elog.FieldKey(key), elog.FieldValueAny(container.config))
	}
	container.logger = container.logger.With(elog.FieldComponentName(key))
	return container
}

// Build 构建配置
func (con *Container) Build(options ...Option) *Component {
	cfg := con.config
	ctx := new(context.Context)
	ctx.AppID = cfg.AppID
	ctx.AppSecret = cfg.AppSecret
	ctx.Token = cfg.Token
	ctx.EncodingAESKey = cfg.EncodingAESKey
	ctx.PayMchID = cfg.PayMchID
	ctx.PayKey = cfg.PayKey
	ctx.PayNotifyURL = cfg.PayNotifyURL
	con.ctx = ctx
	ctx.SetAccessTokenLock(new(sync.RWMutex))
	ctx.SetJsAPITicketLock(new(sync.RWMutex))
	ctx.RestyClient = ehttp.DefaultContainer().Build(
		ehttp.WithDebug(cfg.Debug),
		ehttp.WithRawDebug(cfg.RawDebug),
		ehttp.WithReadTimeout(cfg.ReadTimeout),
		ehttp.WithSlowLogThreshold(cfg.SlowLogThreshold),
		ehttp.WithEnableAccessInterceptor(cfg.EnableAccessInterceptor),
		ehttp.WithEnableAccessInterceptorRes(cfg.EnableAccessInterceptorRes),
	)

	for _, option := range options {
		option(con)
	}
	ctx.Cache = con.client
	return newComponent(cfg, ctx, con.client, con.logger)
}

func WithRedis(client *eredis.Component) Option {
	return func(c *Container) {
		c.client = client
	}
}
