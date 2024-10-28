package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"url-shortener/internal/core/domain"
	mock "url-shortener/mocks"
	"url-shortener/pkg/config"
	"url-shortener/pkg/log"
)

var _ = Describe("ResolveURL", Ordered, func() {
	var (
		ctrl        *gomock.Controller
		ctx         context.Context
		mockRepo    *mock.MockURLRepository
		mockCache   *mock.MockURLCache
		urlUC       *URLUseCase
		shortURL    string
		originalURL string
	)

	BeforeAll(func() {
		ctrl = gomock.NewController(GinkgoT())
		log.InitLogger()
	})

	BeforeEach(func() {
		mockRepo = mock.NewMockURLRepository(ctrl)
		mockCache = mock.NewMockURLCache(ctrl)
		urlUC = NewURLUseCase(mockRepo, mockCache, &config.Config{})
		ctx = context.TODO()
		shortURL = "t2xx0dWNg"
		originalURL = "https://example.com"
	})

	When("getting URL", func() {
		It("should return an error if getting URL from cache and repository fails", func() {
			mockCache.EXPECT().Get(ctx, shortURL).Return("", errors.New("cache error"))
			mockRepo.EXPECT().Find(shortURL).Return(nil, errors.New("repo error"))

			url, err := urlUC.ResolveURL(ctx, shortURL)

			Expect(url).To(BeEmpty())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("repo error"))
		})

		It("should return the URL if the cache fails", func() {
			mockCache.EXPECT().Get(ctx, shortURL).Return("", errors.New("cache error"))
			mockRepo.EXPECT().Find(shortURL).Return(&domain.URL{
				Short:    shortURL,
				Original: originalURL,
				Enabled:  true,
			}, nil)
			mockCache.EXPECT().Set(ctx, shortURL, originalURL).Return(errors.New("cache error"))

			url, err := urlUC.ResolveURL(ctx, shortURL)

			Expect(url).To(Equal(originalURL))
			Expect(err).To(BeNil())
		})

		It("should return an error if the URL is disabled", func() {
			mockCache.EXPECT().Get(ctx, shortURL).Return("", errors.New("cache error"))
			mockRepo.EXPECT().Find(shortURL).Return(&domain.URL{
				Short:    shortURL,
				Original: originalURL,
				Enabled:  false,
			}, nil)

			url, err := urlUC.ResolveURL(ctx, shortURL)

			Expect(url).To(BeEmpty())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("original url not found or disabled for id: %s", shortURL)))
		})

		It("should get the cache URL", func() {
			mockCache.EXPECT().Get(ctx, shortURL).Return(originalURL, nil)

			url, err := urlUC.ResolveURL(ctx, shortURL)

			Expect(url).To(Equal(originalURL))
			Expect(err).To(BeNil())
		})
	})
})
