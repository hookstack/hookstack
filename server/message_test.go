package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/frain-dev/convoy"
	"github.com/frain-dev/convoy/config"
	"github.com/frain-dev/convoy/mocks"
	"github.com/frain-dev/convoy/server/models"
	"github.com/go-chi/chi/v5"
	pager "github.com/gobeam/mongo-go-pagination"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
)

func Test_ensureNewMessage(t *testing.T) {

	var app *applicationHandler

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orgRepo := mocks.NewMockOrganisationRepository(ctrl)
	appRepo := mocks.NewMockApplicationRepository(ctrl)
	msgRepo := mocks.NewMockMessageRepository(ctrl)

	orgID := "1234567890"
	appId := "12345"
	msgId := "1122333444456"

	app = newApplicationHandler(msgRepo, appRepo, orgRepo)

	message := &convoy.Message{
		UID:   msgId,
		AppID: appId,
	}

	type args struct {
		message *convoy.Message
	}

	tests := []struct {
		name       string
		method     string
		statusCode int
		args       args
		body       *strings.Reader
		dbFn       func(msgRepo *mocks.MockMessageRepository, appRepo *mocks.MockApplicationRepository, orgRepo *mocks.MockOrganisationRepository)
	}{
		{
			name:       "invalid message - no event type",
			method:     http.MethodPost,
			statusCode: http.StatusBadRequest,
			body:       strings.NewReader(`{ "data": {}}`),
			args: args{

				message: message,
			},
			dbFn: func(msgRepo *mocks.MockMessageRepository, appRepo *mocks.MockApplicationRepository, orgRepo *mocks.MockOrganisationRepository) {
				appRepo.EXPECT().
					FindApplicationByID(gomock.Any(), gomock.Any()).Times(0).
					Return(&convoy.Application{
						UID:       appId,
						OrgID:     orgID,
						Title:     "Valid application",
						Endpoints: []convoy.Endpoint{},
					}, nil)
				msgRepo.EXPECT().
					CreateMessage(gomock.Any(), gomock.Any()).Times(0).
					Return(nil)

			},
		},
		{
			name:       "valid message",
			method:     http.MethodPost,
			statusCode: http.StatusCreated,
			body:       strings.NewReader(`{"event_type": "test.event", "data": { "Hello": "World", "Test": "Data" }}`),
			args: args{
				message: message,
			},
			dbFn: func(msgRepo *mocks.MockMessageRepository, appRepo *mocks.MockApplicationRepository, orgRepo *mocks.MockOrganisationRepository) {
				appRepo.EXPECT().
					FindApplicationByID(gomock.Any(), gomock.Any()).Times(1).
					Return(&convoy.Application{
						UID:   appId,
						OrgID: orgID,
						Title: "Valid application",
						Endpoints: []convoy.Endpoint{
							{
								TargetURL: "http://localhost",
							},
						},
					}, nil)
				msgRepo.EXPECT().
					CreateMessage(gomock.Any(), gomock.Any()).Times(1).
					Return(nil)

			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			err := config.LoadFromFile("./testdata/TestRequireAuth_None/convoy.json")
			if err != nil {
				t.Error("Failed to load config file")
			}

			request := httptest.NewRequest(tc.method, fmt.Sprintf("/v1/apps/%s/events", tc.args.message.AppID), tc.body)
			responseRecorder := httptest.NewRecorder()
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("appID", tc.args.message.AppID)

			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

			if tc.dbFn != nil {
				tc.dbFn(msgRepo, appRepo, orgRepo)
			}

			ensureNewMessage(appRepo, msgRepo)(http.HandlerFunc(app.CreateAppMessage)).
				ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				logrus.Error(tc.args.message, responseRecorder.Body)
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}
		})
	}
}

func Test_fetchAllMessages(t *testing.T) {
	var app *applicationHandler

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orgRepo := mocks.NewMockOrganisationRepository(ctrl)
	appRepo := mocks.NewMockApplicationRepository(ctrl)
	msgRepo := mocks.NewMockMessageRepository(ctrl)

	appId := "12345"
	msgId := "1122333444456"

	app = newApplicationHandler(msgRepo, appRepo, orgRepo)

	message := &convoy.Message{
		UID:   msgId,
		AppID: appId,
	}

	type args struct {
		message *convoy.Message
	}

	tests := []struct {
		name       string
		method     string
		statusCode int
		args       args
		body       *strings.Reader
		dbFn       func(msgRepo *mocks.MockMessageRepository, appRepo *mocks.MockApplicationRepository, orgRepo *mocks.MockOrganisationRepository)
	}{
		{
			name:       "valid messages",
			method:     http.MethodGet,
			statusCode: http.StatusOK,
			body:       nil,
			args: args{
				message: message,
			},
			dbFn: func(msgRepo *mocks.MockMessageRepository, appRepo *mocks.MockApplicationRepository, orgRepo *mocks.MockOrganisationRepository) {
				msgRepo.EXPECT().
					LoadMessagesPaged(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).
					Return([]convoy.Message{
						*message,
					},
						pager.PaginationData{},
						nil)

			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, fmt.Sprintf("/v1/apps/%s/events", tc.args.message.AppID), nil)
			responseRecorder := httptest.NewRecorder()
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("appID", tc.args.message.AppID)

			pageable := models.Pageable{
				Page:    1,
				PerPage: 10,
			}
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))
			request = request.WithContext(context.WithValue(request.Context(), pageableCtx, pageable))

			if tc.dbFn != nil {
				tc.dbFn(msgRepo, appRepo, orgRepo)
			}

			fetchAllMessages(msgRepo)(http.HandlerFunc(app.GetAppMessagesPaged)).
				ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				logrus.Error(tc.args.message, responseRecorder.Body)
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}
		})
	}
}

func Test_resendMessage(t *testing.T) {
	var app *applicationHandler

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orgRepo := mocks.NewMockOrganisationRepository(ctrl)
	appRepo := mocks.NewMockApplicationRepository(ctrl)
	msgRepo := mocks.NewMockMessageRepository(ctrl)

	appId := "12345"
	msgId := "1122333444456"

	app = newApplicationHandler(msgRepo, appRepo, orgRepo)

	type args struct {
		message *convoy.Message
	}

	tests := []struct {
		name       string
		method     string
		statusCode int
		args       args
		body       *strings.Reader
		dbFn       func(msgRepo *mocks.MockMessageRepository, appRepo *mocks.MockApplicationRepository, orgRepo *mocks.MockOrganisationRepository)
	}{
		{
			name:       "invalid event to resend - already successful",
			method:     http.MethodPut,
			statusCode: http.StatusBadRequest,
			body:       nil,
			args: args{
				message: &convoy.Message{
					UID:    msgId,
					AppID:  appId,
					Status: convoy.SuccessMessageStatus,
				},
			},
			dbFn: func(msgRepo *mocks.MockMessageRepository, appRepo *mocks.MockApplicationRepository, orgRepo *mocks.MockOrganisationRepository) {
				msgRepo.EXPECT().
					UpdateStatusOfMessages(gomock.Any(), gomock.Any(), gomock.Any()).Times(0).
					Return(nil)

			},
		},
		{
			name:       "invalid event to resend - not failed",
			method:     http.MethodPut,
			statusCode: http.StatusBadRequest,
			body:       nil,
			args: args{
				message: &convoy.Message{
					UID:    msgId,
					AppID:  appId,
					Status: convoy.ProcessingMessageStatus,
				},
			},
			dbFn: func(msgRepo *mocks.MockMessageRepository, appRepo *mocks.MockApplicationRepository, orgRepo *mocks.MockOrganisationRepository) {
				msgRepo.EXPECT().
					UpdateStatusOfMessages(gomock.Any(), gomock.Any(), gomock.Any()).Times(0).
					Return(nil)

			},
		},
		{
			name:       "valid event to resend - previously failed",
			method:     http.MethodPut,
			statusCode: http.StatusOK,
			body:       nil,
			args: args{
				message: &convoy.Message{
					UID:    msgId,
					AppID:  appId,
					Status: convoy.FailureMessageStatus,
				},
			},
			dbFn: func(msgRepo *mocks.MockMessageRepository, appRepo *mocks.MockApplicationRepository, orgRepo *mocks.MockOrganisationRepository) {
				msgRepo.EXPECT().
					UpdateStatusOfMessages(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).
					Return(nil)

			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, fmt.Sprintf("/v1/apps/events/%s/resend", tc.args.message.UID), nil)
			responseRecorder := httptest.NewRecorder()
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("appID", tc.args.message.AppID)

			request = request.WithContext(context.WithValue(request.Context(), msgCtx, tc.args.message))

			if tc.dbFn != nil {
				tc.dbFn(msgRepo, appRepo, orgRepo)
			}

			resendMessage(msgRepo)(http.HandlerFunc(app.ResendAppMessage)).
				ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				logrus.Error(tc.args.message, responseRecorder.Body)
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			verifyMatch(t, *responseRecorder)
		})
	}
}
