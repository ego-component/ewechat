package ewechat

import (
	"github.com/ego-component/ewechat/cache"
	"github.com/ego-component/ewechat/context"
	"github.com/ego-component/ewechat/oauth"
	"github.com/gotomicro/ego/core/elog"
)

type Component struct {
	config *config
	ctx    *context.Context
	client cache.Cache
	logger *elog.Component
}

func newComponent(cfg *config, ctx *context.Context, client cache.Cache, logger *elog.Component) *Component {
	return &Component{
		config: cfg,
		ctx:    ctx,
		client: client,
		logger: logger,
	}
}

// GetOauth 获取小程序的实例
func (c *Component) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(c.ctx)
}
