package url

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"url-shortener/internal/core/domain"
	mock "url-shortener/mocks"
	"url-shortener/pkg/config"
	"url-shortener/pkg/log"
)

var _ = Describe("DynamoDBRepository", Ordered, func() {
	var (
		ctrl     *gomock.Controller
		mockDB   *mock.MockClientInterface
		repo     *Repository
		url      domain.URL
		shortURL string
	)

	BeforeAll(func() {
		ctrl = gomock.NewController(GinkgoT())
		log.InitLogger()
		shortURL = "t2xx0dWNg"
	})

	BeforeEach(func() {
		mockDB = mock.NewMockClientInterface(ctrl)
		repo, _ = NewURLRepository(mockDB, &config.Config{
			DatabasesConfig: config.DatabasesConfig{
				DynamoDB: config.DynamoConfig{
					TableName: "urls",
				},
			},
		})
	})

	When("saving", func() {
		It("should return an error if the db fails", func() {
			mockDB.EXPECT().PutItem(gomock.Any()).Return(nil, errors.New("db error"))

			err := repo.Save(url)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("db error"))
		})

		It("should save successfully", func() {
			mockDB.EXPECT().PutItem(gomock.Any()).Return(nil, nil)

			err := repo.Save(url)

			Expect(err).To(BeNil())
		})
	})

	When("getting", func() {
		It("should return an error if the db fails", func() {
			mockDB.EXPECT().GetItem(gomock.Any()).Return(nil, errors.New("db error"))

			_, err := repo.Find(shortURL)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("db error"))
		})

		It("should return an error if the record is nil", func() {
			mockDB.EXPECT().GetItem(gomock.Any()).Return(&dynamodb.GetItemOutput{
				Item: nil,
			}, nil)

			_, err := repo.Find(shortURL)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("URL with shortID: %s not found", shortURL)))
		})

		It("should get successfully", func() {
			mockDB.EXPECT().GetItem(gomock.Any()).Return(&dynamodb.GetItemOutput{
				Item: make(map[string]*dynamodb.AttributeValue),
			}, nil)

			url, err := repo.Find(shortURL)

			Expect(url).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})

	When("updating", func() {
		It("should return an error if the db fails", func() {
			mockDB.EXPECT().UpdateItem(gomock.Any()).Return(nil, errors.New("db error"))

			err := repo.Update(url)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("db error"))
		})

		It("should update successfully", func() {
			mockDB.EXPECT().UpdateItem(gomock.Any()).Return(nil, nil)

			err := repo.Update(url)

			Expect(err).To(BeNil())
		})
	})
})
