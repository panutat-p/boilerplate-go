package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"boilerplate-go/internal/mock"
	"boilerplate-go/internal/model"
)

type UseCaseTestSuite struct {
	suite.Suite
	ctrl         *gomock.Controller
	mockStore    *mock.MockIStore
	mockExternal *mock.MockIExternal
	useCase      *UseCase
}

func (suite *UseCaseTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockStore = mock.NewMockIStore(suite.ctrl)
	suite.mockExternal = mock.NewMockIExternal(suite.ctrl)
	suite.useCase = NewUseCase(nil, suite.mockStore, suite.mockExternal)
}

func (suite *UseCaseTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *UseCaseTestSuite) TestGetFruits() {
	ctx := context.Background()

	suite.Run("success", func() {
		want := []model.Fruit{
			{Name: "Apple"},
			{Name: "Banana"},
		}
		suite.mockStore.
			EXPECT().
			ReadFruitFile(gomock.Any()).
			Return(want, nil)

		got, err := suite.useCase.GetFruits(ctx)
		suite.NoError(err)
		suite.Equal(want, got)
	})

	suite.Run("empty fruit", func() {
		suite.mockStore.
			EXPECT().
			ReadFruitFile(gomock.Any()).
			Return(nil, nil)

		got, err := suite.useCase.GetFruits(ctx)
		suite.Error(err)
		suite.Len(got, 0)
	})

	suite.Run("cannot read file", func() {
		want := []model.Fruit{
			{Name: "Apple"},
			{Name: "Banana"},
		}
		suite.mockStore.
			EXPECT().
			ReadFruitFile(gomock.Any()).
			Return(want, errors.New("unit-test"))

		_, err := suite.useCase.GetFruits(ctx)
		suite.Error(err)
	})
}

func TestUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UseCaseTestSuite))
}
