package create

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
	mock "url-shortener/mocks"
)

var _ = Describe("CreateHandler", Ordered, func() {
	var (
		ctrl        *gomock.Controller
		mockUseCase *mock.MockURLShortenerUseCase
		handler     *Handler
		e           *echo.Echo
		ctx         echo.Context
		rec         *httptest.ResponseRecorder
		originalURL string
		shortURL    string
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockUseCase = mock.NewMockURLShortenerUseCase(ctrl)
		handler = NewHandler(mockUseCase)
		e = echo.New()
		rec = httptest.NewRecorder()
	})

	BeforeAll(func() {
		ctrl = gomock.NewController(GinkgoT())
		originalURL = "https://example.com"
		shortURL = "https://shortener/t2xx0dWNg"
	})

	When("the use case fails", func() {
		It("should return an error", func() {
			requestBody := `{"url": "https://example.com"}`
			req := httptest.NewRequest(http.MethodPost, "/urls", strings.NewReader(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			ctx = e.NewContext(req, rec)

			mockUseCase.EXPECT().
				CreateShortURL(gomock.Any(), originalURL).
				Return("", errors.New("use case error"))

			err := handler.HandleRequest(ctx)
			Expect(err).ToNot(BeNil())
		})
	})

	When("when the request is valid", func() {
		It("should return 201 and the new short URL", func() {
			requestBody := `{"url": "https://example.com"}`
			req := httptest.NewRequest(http.MethodPost, "/urls", strings.NewReader(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			ctx = e.NewContext(req, rec)

			mockUseCase.EXPECT().
				CreateShortURL(gomock.Any(), originalURL).
				Return(shortURL, nil)

			err := handler.HandleRequest(ctx)
			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusCreated))
			Expect(rec.Body.String()).To(ContainSubstring(shortURL))
		})
	})
})
