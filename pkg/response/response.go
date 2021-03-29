package response

import "github.com/beego/beego/v2/server/web/context"

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Code    int         `json:"-"`
}

func (r *Response) ServeJSON(ctx *context.Context) {
	ctx.Output.Status = r.Code
	ctx.Output.JSON(r, false, false)
}
