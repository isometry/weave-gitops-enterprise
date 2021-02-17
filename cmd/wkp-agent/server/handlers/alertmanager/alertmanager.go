package alertmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"github.com/prometheus/alertmanager/api/v2/client"
	"github.com/prometheus/alertmanager/notify/webhook"
	log "github.com/sirupsen/logrus"
	"github.com/weaveworks/wks/common/messaging/payload"
)

const (
	eventType   = "PrometheusAlerts"
	contentType = "application/json"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) PushHandler(req *http.Request) (event.Event, error) {
	var msg webhook.Message

	if req.Body == nil {
		return event.Event{}, errors.New("empty payload")
	}

	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	err := decoder.Decode(&msg)
	if err != nil {
		return event.Event{}, err
	}

	ce := cloudevents.NewEvent()
	ce.SetID(uuid.New().String())
	ce.SetType(eventType)
	ce.SetTime(time.Now())
	ce.SetSource(msg.ExternalURL)
	if err := ce.SetData(contentType, msg); err != nil {
		log.Errorf("Unable to set event as data: %v.", err)
		return event.Event{}, err
	}

	return ce, nil
}

func NewWebhookHandler(fn func(event.Event)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ce, err := NewParser().PushHandler(r)
		if err != nil {
			log.Fatalf("Failed to parse alert: %v", err)
		}
		fn(ce)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(""))
	}
}

func GetAlerts(token string, alertmanagerAddress string) (*payload.PrometheusAlerts, error) {
	u, err := url.Parse(alertmanagerAddress)
	if err != nil {
		return nil, err
	}
	cfg := client.DefaultTransportConfig().WithHost(u.Host).WithBasePath(u.Path)
	c := client.NewHTTPClientWithConfig(nil, cfg)
	resp, err := c.Alert.GetAlerts(nil)
	if err != nil {
		return nil, err
	}

	pa := &payload.PrometheusAlerts{
		Token:  token,
		Alerts: resp.Payload,
	}

	return pa, nil
}

func ToCloudEvent(source string, alerts *payload.PrometheusAlerts) (event.Event, error) {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.New().String())
	ce.SetType(eventType)
	ce.SetTime(time.Now())
	ce.SetSource(source)
	if err := ce.SetData(contentType, alerts); err != nil {
		log.Errorf("Unable to set event as data: %v.", err)
		return event.Event{}, err
	}
	return ce, nil
}

func GetAlertsAsEvent(token string, alertmanagerAddress string) (event.Event, error) {
	alerts, err := GetAlerts(token, alertmanagerAddress)
	if err != nil {
		return event.Event{}, err
	}
	return ToCloudEvent(alertmanagerAddress, alerts)
}
