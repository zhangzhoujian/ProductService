// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostAPIV1ProductsHandlerFunc turns a function with the right signature into a post API v1 products handler
type PostAPIV1ProductsHandlerFunc func(PostAPIV1ProductsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAPIV1ProductsHandlerFunc) Handle(params PostAPIV1ProductsParams) middleware.Responder {
	return fn(params)
}

// PostAPIV1ProductsHandler interface for that can handle valid post API v1 products params
type PostAPIV1ProductsHandler interface {
	Handle(PostAPIV1ProductsParams) middleware.Responder
}

// NewPostAPIV1Products creates a new http.Handler for the post API v1 products operation
func NewPostAPIV1Products(ctx *middleware.Context, handler PostAPIV1ProductsHandler) *PostAPIV1Products {
	return &PostAPIV1Products{Context: ctx, Handler: handler}
}

/*PostAPIV1Products swagger:route POST /api/v1/products postApiV1Products

Create

*/
type PostAPIV1Products struct {
	Context *middleware.Context
	Handler PostAPIV1ProductsHandler
}

func (o *PostAPIV1Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAPIV1ProductsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
