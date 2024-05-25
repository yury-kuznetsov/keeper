package data

import (
	"errors"
	mock_data "gophkeeper/internal/client/mock/data"
	"gophkeeper/internal/client/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	idLocal := uuid.New()
	idRemove := uuid.New()
	idLocalAndRemove := uuid.New()

	r := mock_data.NewMockRepository(ctrl)
	c := mock_data.NewMockClient(ctrl)

	dataLocal := []model.Data{
		{
			ID:       idLocal,
			Category: model.CategoryCredentials,
			Data:     []byte("LOGIN_AND_PASSWORD"),
			Version:  0,
		},
		{
			ID:       idLocalAndRemove,
			Category: model.CategoryText,
			Data:     []byte("TEXT"),
			Version:  0,
		},
	}

	dataRemote := []model.DataVersionRemote{
		{
			ID:      idRemove,
			Version: 1,
		},
		{
			ID:      idLocalAndRemove,
			Version: 1,
		},
	}

	r.EXPECT().FindAll().Return(nil, errors.New("LOCAL_FOUNDING_ERROR"))
	r.EXPECT().FindAll().Return(dataLocal, nil).Times(2)

	c.EXPECT().Status().Return(nil, errors.New("REMOTE_FOUNDING_ERROR"))
	c.EXPECT().Status().Return(dataRemote, nil)

	svc := NewService(r, c)

	tests := []struct {
		name  string
		error string
	}{
		{
			name:  "Local founding error",
			error: "LOCAL_FOUNDING_ERROR",
		},
		{
			name:  "Remote founding error",
			error: "REMOTE_FOUNDING_ERROR",
		},
		{
			name:  "Update data",
			error: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			versions, err := svc.Status()
			if tt.error == "" {
				assert.Nil(t, err)
				assert.Len(t, versions, 3)

				assert.Equal(t, versions[0].ID, idLocal)
				assert.Equal(t, versions[0].VersionLocal, 0)
				assert.Equal(t, versions[0].VersionRemote, 0)

				assert.Equal(t, versions[1].ID, idLocalAndRemove)
				assert.Equal(t, versions[1].VersionLocal, 0)
				assert.Equal(t, versions[1].VersionRemote, 1)

				assert.Equal(t, versions[2].ID, idRemove)
				assert.Equal(t, versions[2].VersionLocal, 0)
				assert.Equal(t, versions[2].VersionRemote, 1)
			} else {
				assert.Equal(t, err.Error(), tt.error)
			}
		})
	}
}
