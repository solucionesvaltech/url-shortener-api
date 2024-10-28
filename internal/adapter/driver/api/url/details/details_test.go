package details

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"url-shortener/internal/core/domain"
	mock "url-shortener/mocks"
)

var _ = Describe("DetailsHandler", Ordered, func() {
	var (
		ctrl        *gomock.Controller
		mockUseCase *mock.MockURLShortenerUseCase
		handler     *Handler
		e           *echo.Echo
		ctx         echo.Context
		rec         *httptest.ResponseRecorder
		shortURL    string
		url         domain.URL
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
		url = domain.URL{
			Short:   shortURL,
			Enabled: true,
		}
	})

	When("the param is not present", func() {
		It("should return an error", func() {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/urls/%s", shortURL), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			ctx = e.NewContext(req, rec)

			err := handler.HandleRequest(ctx)
			Expect(err).ToNot(BeNil())
		})
	})

	When("the use case fails", func() {
		It("should return an error", func() {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/urls/%s", shortURL), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			ctx = e.NewContext(req, rec)
			ctx.SetParamNames("shortID")
			ctx.SetParamValues(shortURL)

			mockUseCase.EXPECT().
				DetailURL(gomock.Any(), gomock.Any()).
				Return(nil, errors.New("use case error"))

			err := handler.HandleRequest(ctx)
			Expect(err).ToNot(BeNil())
		})
	})

	When("when the request is valid", func() {
		It("should return 200", func() {
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/urls/%s", shortURL), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			ctx = e.NewContext(req, rec)
			ctx.SetParamNames("shortID")
			ctx.SetParamValues(shortURL)

			mockUseCase.EXPECT().
				DetailURL(gomock.Any(), shortURL).
				Return(&url, nil)

			err := handler.HandleRequest(ctx)
			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(rec.Body.String()).To(ContainSubstring(shortURL))
		})
	})
})
