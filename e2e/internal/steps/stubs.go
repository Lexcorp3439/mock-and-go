package steps

import (
	"e2e/internal/testutils"
	"github.com/google/uuid"
)

type StubSteps struct {
}

func NewStubSteps() *StubSteps {
	return &StubSteps{}
}

func (s StubSteps) GetStubUsersData() (string, string, int32) {
	return uuid.NewString(), testutils.GetRandomPhone(), int32(testutils.GetRandomInt(50))
}
