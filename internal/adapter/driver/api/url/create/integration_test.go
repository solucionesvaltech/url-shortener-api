//go:build integration
// +build integration

package create

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"url-shortener/internal/adapter/driven/dynamo"
	"url-shortener/internal/adapter/driven/dynamo/url"
	"url-shortener/internal/adapter/driven/redis"
	"url-shortener/internal/core/usecase"
	"url-shortener/pkg/config/configyaml"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("URL Shortener Integration Test", Ordered, func() {
	var (
		handler  *Handler
		useCase  *usecase.URLUseCase
		e        *echo.Echo
		recorder *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		conf, _ := configyaml.NewConfigYaml()
		dbClient, _ := dynamo.NewDynamoDBClient(conf)
		redis := redis.NewRedisClient(conf)
		repo, _ := url.NewURLRepository(dbClient, conf)

		useCase = usecase.NewURLUseCase(repo, redis, conf)
		handler = NewHandler(useCase)

		e = echo.New()
		recorder = httptest.NewRecorder()
	})

	When("creating a new short URL", func() {
		It("should return a 201 status and store the URL", func() {
			reqBody, _ := json.Marshal(map[string]string{"url": "https://example.com"})
			req := httptest.NewRequest(http.MethodPost, "/urls", bytes.NewBuffer(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			ctx := e.NewContext(req, recorder)

			err := handler.HandleRequest(ctx)
			Expect(err).To(BeNil())
			Expect(recorder.Code).To(Equal(http.StatusCreated))

			var response string
			err = json.Unmarshal(recorder.Body.Bytes(), &response)
			Expect(err).To(BeNil())
			parts := strings.Split(response, "/")
			shortID := parts[len(parts)-1]
			Expect(shortID).NotTo(BeEmpty())

			savedURL, err := useCase.DetailURL(context.Background(), shortID)
			Expect(err).To(BeNil())
			Expect(savedURL.Original).To(Equal("https://example.com"))
		})
	})
})
