package public

import (
	"context"
	"fmt"
	"github.com/google/logger"
	"heisenbug/complex/internal/pkg/model"
	"heisenbug/complex/internal/pkg/repository"
	"heisenbug/complex/internal/pkg/store"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	TemplateRepo  repository.Template
	PublicService *Implementation
)

func TestMain(m *testing.M) {
	db, err := store.ConnectToPostgres()
	if err != nil {
		logger.Fatalf("main: cannot create Store: %v", err)
	}

	storage := store.NewStorage(db)
	TemplateRepo = repository.NewTemplate(storage)

	PublicService = NewPublicService(TemplateRepo)

	time.Sleep(1 * time.Second)

	exitVal := m.Run()
	defer db.Close()
	os.Exit(exitVal)
}

func prepareTemplates(ctx context.Context, t *testing.T, count int, namePrefix *string) []string {
	t.Helper()

	templates := make([]string, count)

	var prefix string

	if namePrefix != nil {
		prefix = *namePrefix
	}

	for i := range templates {
		templates[i] = fmt.Sprintf("%s%s", prefix, uuid.NewString())
		err := TemplateRepo.CreateTemplate(ctx, &model.Template{
			Name:      templates[i],
			Status:    0,
			CreatedAt: time.Now().Add(-24 * time.Hour * time.Duration(i)),
			UpdatedAt: time.Now().Add(-1 * time.Hour * time.Duration((i%2)*i)),
		})
		require.NoError(t, err, "prepare db error")
	}

	return templates
}
