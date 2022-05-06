package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/gitnyasha/go-hekani-backend/db/mock"
	db "github.com/gitnyasha/go-hekani-backend/db/sqlc"
	"github.com/gitnyasha/go-hekani-backend/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetQuestionAPI(t *testing.T) {
	question := createRandomQuestion()

	testCases := []struct {
		name          string
		questionID    int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "OK",
			questionID: question.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetQuestion(gomock.Any(), gomock.Eq(question.ID)).
					Times(1).
					Return(question, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchQuestion(t, recorder.Body, question)
			},
		},
		{
			name:       "NotFound",
			questionID: question.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetQuestion(gomock.Any(), gomock.Eq(question.ID)).
					Times(1).
					Return(db.Question{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:       "InternalError",
			questionID: question.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetQuestion(gomock.Any(), gomock.Eq(question.ID)).
					Times(1).
					Return(db.Question{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:       "InvalidID",
			questionID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetQuestion(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			store := mockdb.NewMockStore(controller)
			tc.buildStubs(store)
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/questions/%d", tc.questionID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func createRandomQuestion() db.Question {
	return db.Question{
		ID:                 util.RandomInt(1, 100),
		Title:              util.RandomString(),
		UserID:             3,
		QuestionCategoryID: 3,
	}
}

func requireBodyMatchQuestion(t *testing.T, body *bytes.Buffer, question db.Question) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotQuestion db.Question
	err = json.Unmarshal(data, &gotQuestion)
	require.NoError(t, err)
	require.Equal(t, question, gotQuestion)
}
