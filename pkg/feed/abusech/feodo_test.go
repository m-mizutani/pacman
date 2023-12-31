package abusech_test

import (
	"context"
	"testing"

	"github.com/m-mizutani/drone/pkg/domain/model"
	"github.com/m-mizutani/drone/pkg/feed/abusech"
	"github.com/m-mizutani/drone/pkg/infra/bq"
	"github.com/m-mizutani/gt"
)

func TestFeodoIntegration(t *testing.T) {
	mock := bq.NewMock()
	ctx := context.Background()

	// first time
	gt.NoError(t, abusech.NewFeodo().Import(ctx, mock))

	gt.A(t, mock.InsertedData).Longer(0)
	firstLen := len(mock.InsertedData)
	gt.A(t, mock.Records).Length(1).At(0, func(t testing.TB, rec *model.ImportLog) {
		gt.Equal(t, rec.TableName, "abusech_feodo")
	})

	// second time
	gt.NoError(t, abusech.NewFeodo().Import(ctx, mock))
	gt.A(t, mock.InsertedData).Length(firstLen)
	gt.A(t, mock.Records).Length(2)
}
