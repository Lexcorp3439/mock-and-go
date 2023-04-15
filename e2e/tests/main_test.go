package tests

import (
	"testing"

	simple "e2e/tests/identification"
	complex "e2e/tests/identification_complex"
	inservice "e2e/tests/identification_inservice"
	wiremock "e2e/tests/identification_wiremock"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

func TestIdentificationService(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(simple.Suite))
}

func TestIdentificationServiceWithInServiceMock(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(inservice.Suite))
}

func TestIdentificationServiceWithWireMock(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(wiremock.Suite))
}

func TestIdentificationServiceWithComplexMock(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(complex.Suite))
}
