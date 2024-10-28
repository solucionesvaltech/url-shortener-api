package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	mock "url-shortener/mocks"
	"url-shortener/pkg/config"
)

var _ = Describe("CreateShortURL", Ordered, func() {
	var (
		ctrl        *gomock.Controller
		ctx         context.Context
		mockRepo    *mock.MockURLRepository
		mockCache   *mock.MockURLCache
		urlUC       *URLUseCase
		originalURL string
	)

	BeforeAll(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		originalURL = "https://example.com"
	})

	BeforeEach(func() {
		mockRepo = mock.NewMockURLRepository(ctrl)
		mockCache = mock.NewMockURLCache(ctrl)
		urlUC = NewURLUseCase(mockRepo, mockCache, &config.Config{})
	})

	When("creating a new short URL", func() {
		It("should return an error if saving to the repository fails", func() {
			mockRepo.EXPECT().Save(gomock.Any()).Return(errors.New("repo error"))

			id, err := urlUC.CreateShortURL(ctx, originalURL)

			Expect(id).To(BeEmpty())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("repo error"))
		})

		It("should return an error if setting the cache fails", func() {
			mockRepo.EXPECT().Save(gomock.Any()).Return(nil)
			mockCache.EXPECT().Set(ctx, gomock.Any(), originalURL).Return(errors.New("cache error"))

			id, err := urlUC.CreateShortURL(ctx, originalURL)

			Expect(id).To(BeEmpty())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("cache error"))
		})

		It("should return an error if the url is invalid", func() {
			id, err := urlUC.CreateShortURL(ctx, "invalid-url")

			Expect(id).To(BeEmpty())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Message: invalid format for URL: invalid-url"))
		})

		It("should generate a valid short ID, save it to repo, and cache it", func() {
			mockRepo.EXPECT().Save(gomock.Any()).Return(nil)
			mockCache.EXPECT().Set(ctx, gomock.Any(), originalURL).Return(nil)

			id, err := urlUC.CreateShortURL(ctx, originalURL)

			Expect(err).To(BeNil())
			Expect(id).NotTo(BeEmpty())
		})
	})
})
