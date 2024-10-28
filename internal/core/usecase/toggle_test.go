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

var _ = Describe("ToggleURLStatus", Ordered, func() {
	var (
		ctrl        *gomock.Controller
		ctx         context.Context
		mockRepo    *mock.MockURLRepository
		mockCache   *mock.MockURLCache
		urlUC       *URLUseCase
		shortURL    string
		originalURL string
		url         domain.URL
	)

	BeforeAll(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	BeforeEach(func() {
		mockRepo = mock.NewMockURLRepository(ctrl)
		mockCache = mock.NewMockURLCache(ctrl)
		urlUC = NewURLUseCase(mockRepo, mockCache, &config.Config{})
		ctx = context.TODO()
		shortURL = "t2xx0dWNg"
		originalURL = "https://example.com"
		url = domain.URL{
			Short:    shortURL,
			Original: originalURL,
			Enabled:  true,
		}
	})

	When("toggling the URL", func() {
		It("should return an error if getting URL from repository fails", func() {
			mockRepo.EXPECT().Find(shortURL).Return(nil, errors.New("repo error"))

			err := urlUC.ToggleURLStatus(ctx, shortURL, true)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("repo error"))
		})

		It("should return and error if update fails", func() {
			mockRepo.EXPECT().Find(shortURL).Return(&url, nil)
			mockRepo.EXPECT().Update(url).Return(errors.New("repo error"))

			err := urlUC.ToggleURLStatus(ctx, shortURL, true)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("repo error"))
		})

		It("should return an error if cache fails", func() {
			mockRepo.EXPECT().Find(shortURL).Return(&url, nil)
			mockRepo.EXPECT().Update(url).Return(nil)
			mockCache.EXPECT().Clean(ctx, shortURL).Return(errors.New("cache error"))

			err := urlUC.ToggleURLStatus(ctx, shortURL, true)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("cache error"))
		})

		It("should update the URL and clean the cache", func() {
			mockRepo.EXPECT().Find(shortURL).Return(&url, nil)
			mockRepo.EXPECT().Update(url).Return(nil)
			mockCache.EXPECT().Clean(ctx, shortURL).Return(nil)

			err := urlUC.ToggleURLStatus(ctx, shortURL, true)

			Expect(err).To(BeNil())
		})
	})
})
