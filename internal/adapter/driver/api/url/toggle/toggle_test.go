package toggle

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
	mock "url-shortener/mocks"
)

var _ = Describe("ToggleHandler", Ordered, func() {
	var (
		ctrl        *gomock.Controller
		mockUseCase *mock.MockURLShortenerUseCase
		handler     *Handler
		e           *echo.Echo
		ctx         echo.Context
		rec         *httptest.ResponseRecorder
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
		shortURL = "t2xx0dWNg"
	})

	When("the param is not present", func() {
		It("should return an error", func() {
			requestBody := `{"enabled": true}`
			req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/%s", shortURL), strings.NewReader(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			ctx = e.NewContext(req, rec)

			err := handler.HandleRequest(ctx)
			Expect(err).ToNot(BeNil())
		})
	})

	When("the body is invalid", func() {
		It("should return an error", func() {
			req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/%s", shortURL), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			ctx = e.NewContext(req, rec)

			err := handler.HandleRequest(ctx)
			Expect(err).ToNot(BeNil())
		})
	})

	When("the use case fails", func() {
		It("should return an error", func() {
			requestBody := `{"enabled": true}`
			req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/%s", shortURL), strings.NewReader(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			ctx = e.NewContext(req, rec)
			ctx.SetParamNames("shortID")
			ctx.SetParamValues(shortURL)

			mockUseCase.EXPECT().
				ToggleURLStatus(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(errors.New("use case error"))

			err := handler.HandleRequest(ctx)
			Expect(err).ToNot(BeNil())
		})
	})

	When("when the request is valid", func() {
		It("should return 200", func() {
			requestBody := `{"enabled": true}`
			req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/%s", shortURL), strings.NewReader(requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			ctx = e.NewContext(req, rec)
			ctx.SetParamNames("shortID")
			ctx.SetParamValues(shortURL)

			mockUseCase.EXPECT().
				ToggleURLStatus(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil)

			err := handler.HandleRequest(ctx)
			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))
		})
	})
})
