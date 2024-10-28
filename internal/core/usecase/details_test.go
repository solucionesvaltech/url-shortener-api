package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"url-shortener/internal/core/domain"
	mock "url-shortener/mocks"
	"url-shortener/pkg/config"
)

var _ = Describe("DetailURL", Ordered, func() {
	var (
		ctrl        *gomock.Controller
		ctx         context.Context
		mockRepo    *mock.MockURLRepository
		urlUC       *URLUseCase
		shortURL    string
		originalURL string
	)

	BeforeAll(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	BeforeEach(func() {
		mockRepo = mock.NewMockURLRepository(ctrl)
		urlUC = NewURLUseCase(mockRepo, nil, &config.Config{})
		ctx = context.TODO()
		shortURL = "t2xx0dWNg"
		originalURL = "https://example.com"
	})

	When("getting URL details", func() {
		It("should return an error if getting URL from repository fails", func() {
			mockRepo.EXPECT().Find(shortURL).Return(nil, errors.New("repo error"))

			url, err := urlUC.DetailURL(ctx, shortURL)

			Expect(url).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("repo error"))
		})

		It("should get URL", func() {
			mockRepo.EXPECT().Find(gomock.Any()).Return(&domain.URL{
				Short:    shortURL,
				Original: originalURL,
				Enabled:  true,
			}, nil)

			url, err := urlUC.DetailURL(ctx, shortURL)

			Expect(err).To(BeNil())
			Expect(url).NotTo(BeNil())
		})
	})
})
