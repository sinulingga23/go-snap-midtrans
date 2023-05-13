package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/sinulingga23/go-snap-midtrans/constants"
	"github.com/sinulingga23/go-snap-midtrans/payload"
)

type snapService struct {
	serverKey string
	host      string
}

func NewSnapService() snapService {

	snapService := snapService{}

	if os.Getenv("ENV") == "production" {
		snapService.serverKey = os.Getenv("API_SERVER_KEY_MIDTRANS_PRODUCTION")
		snapService.host = os.Getenv("HOST_MIDTRANS_PRODUCTION")
	}

	if os.Getenv("ENV") == "development" {
		snapService.serverKey = os.Getenv("API_SERVER_KEY_MIDTRANS_SANDBOX")
		snapService.host = os.Getenv("HOST_MIDTRANS_SANDBOX")
	}

	return snapService
}

func (s *snapService) AcquireToken(requestAcquireToken payload.AcquireTokenSnapRequest) (payload.AcquireTokenSnapResponse, error) {

	serviceName := "snap_service:acquire_token"

	bytesMarshal, errMarshal := json.Marshal(requestAcquireToken)

	if errMarshal != nil {
		log.Printf("%s: errMarshal: %v", serviceName, errMarshal)
		return payload.AcquireTokenSnapResponse{}, errMarshal
	}

	bytesPayload := bytes.NewReader(bytesMarshal)

	request, errNewRequest := http.NewRequest(
		http.MethodPost,
		s.host,
		bytesPayload)
	if errNewRequest != nil {
		log.Printf("%s: errNewRequest: %v", serviceName, errNewRequest)
		return payload.AcquireTokenSnapResponse{}, errNewRequest
	}
	defer func() {
		if errClose := request.Body.Close(); errClose != nil {
			log.Printf("%s: errClose: %v", serviceName, errClose)
		}
	}()

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(s.serverKey, ":")

	response, errDo := http.DefaultClient.Do(request)
	if errDo != nil {
		log.Printf("%s: errDo: %v", serviceName, errDo)
		return payload.AcquireTokenSnapResponse{}, errDo
	}
	defer func() {
		if errClose := response.Body.Close(); errClose != nil {
			log.Printf("%s: errClose: %v", serviceName, errClose)
		}
	}()

	bytesResponse, errReadAll := io.ReadAll(response.Body)
	if errReadAll != nil {
		log.Printf("%s: errReadAll: %v", serviceName, errReadAll)
		return payload.AcquireTokenSnapResponse{}, errReadAll
	}

	if response.StatusCode == constants.RC_SUCCESS_CREATE_SNAP_TOKEN {

		acquireTokenSnapResponse := payload.AcquireTokenSnapResponse{}
		if errUnmarshal := json.Unmarshal(bytesResponse, &acquireTokenSnapResponse); errUnmarshal != nil {
			log.Printf("%s: errUnmarshal: %v", serviceName, errUnmarshal)
			return payload.AcquireTokenSnapResponse{}, errUnmarshal
		}

		return acquireTokenSnapResponse, nil
	}

	if response.StatusCode == constants.RC_FAILED_CREATE_SNAP_TOKEN {
		log.Printf("%s: Failed Create Snap Token: %v, StatusCode: %v", serviceName, string(bytesResponse), response.StatusCode)
		return payload.AcquireTokenSnapResponse{}, errors.New(constants.RD_FAILED_CREATE_SNAP_TOKEN)
	}

	log.Printf("%s: Error Create Snap Token: %v, StatusCode: %v", serviceName, string(bytesResponse), response.StatusCode)
	return payload.AcquireTokenSnapResponse{}, errors.New(constants.RD_ERROR_CREATE_SNAP_TOKEN)
}
