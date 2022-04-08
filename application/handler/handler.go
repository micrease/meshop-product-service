package handler

import (
	"context"
	mctx "github.com/micrease/micrease-core/context"
	"meshop-product-service/datasource"
)

type RpcHandler struct {
	ctx *mctx.Context
}

func (this *RpcHandler) NewContext() {
	this.ctx = new(mctx.Context)
	this.ctx.Orm = datasource.GetDB()
}

func (this *RpcHandler) NewWithContext(ctx context.Context) {
	if this.ctx == nil {
		this.ctx = new(mctx.Context)
		this.ctx.Orm = datasource.GetDB()
	}
	this.ctx.Ctx = ctx
}
