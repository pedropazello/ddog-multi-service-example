package controllers

import (
	"checkoutservice/commands"
	"checkoutservice/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type checkoutController struct {
}

func NewCheckoutController() *checkoutController {
	return &checkoutController{}
}

func (controller *checkoutController) Create(c *gin.Context) {
	tracer.Start(tracer.WithServiceName("checkout_service"), tracer.WithEnv("env"))
	defer tracer.Stop()

	fmt.Printf("c.Request.Header: %v\n", c.Request.Header)

	spanName := "checkout_controller.create"
	var span ddtrace.Span

	sctx, err := tracer.Extract(tracer.HTTPHeadersCarrier(c.Request.Header))
	if err != nil {
		span = tracer.StartSpan(spanName, tracer.ResourceName("checkout_service.controller"))
	} else {
		span = tracer.StartSpan(spanName, tracer.ChildOf(sctx))
	}

	ctx := tracer.ContextWithSpan(c.Request.Context(), span)
	span.SetTag("http.base_url", "http://localhost:8080")
	span.SetTag("http.method", "POST")
	span.SetTag("http.path_group", "/checkouts")

	order := models.Order{}

	if err := c.ShouldBindJSON(&order); err != nil {
		span.SetTag("response.status_code", "400")

		span.Finish(tracer.WithError(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if order.Products[0].Name == "error" {
			err = errors.New("failed to checkout")
			span.SetTag("http.response.status_code", "400")
			span.Finish(tracer.WithError(err))

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			span.SetTag("http.response.status_code", "201")
			span.Finish()

			command := commands.NewSendCheckoutCompletedMessage()
			command.SendMessage(ctx, "checkout completed")
			c.JSON(http.StatusCreated, order)
		}
	}
}
