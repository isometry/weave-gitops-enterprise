package test

import (
	contxt "context"
	"testing"
	"time"

	cloudeventsnats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"

	"github.com/nats-io/nats-server/v2/server"
	natsserver "github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weaveworks/wks/pkg/messaging/handlers"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func RunServer() *server.Server {
	opts := natsserver.DefaultTestOptions
	opts.Port = -1 // Allocate a port dynamically
	return natsserver.RunServer(&opts)
}

func TestAgent(t *testing.T) {
	t.Run("Send event to NATS", func(t *testing.T) {
		ctx, cancel := contxt.WithCancel(contxt.Background())
		defer cancel()

		s := RunServer()
		defer s.Shutdown()

		// Set up publisher
		sender, err := cloudeventsnats.NewSender(s.ClientURL(), "test.subject", cloudeventsnats.NatsOptions(
			nats.Name("sender"),
		))
		require.NoError(t, err)
		defer sender.Close(ctx)
		publisher, err := cloudevents.NewClient(sender)
		require.NoError(t, err)
		notifier, err := handlers.NewEventNotifier("test", publisher)
		require.NoError(t, err)

		// Set up subscriber
		consumer, err := cloudeventsnats.NewConsumer(s.ClientURL(), "test.subject", cloudeventsnats.NatsOptions(
			nats.Name("consumer"),
		))
		require.NoError(t, err)
		subscriber, err := cloudevents.NewClient(consumer)
		require.NoError(t, err)

		events := make(chan cloudevents.Event)

		expected := &v1.Event{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "some-event",
				Namespace: "wkp-ns",
			},
		}

		go func() {
			if err := subscriber.StartReceiver(ctx, func(event cloudevents.Event) {
				events <- event
				// Shut down subscriber after receiving one event
				cancel()
			}); err != nil {
				t.Logf("Failed to start NATS subscriber: %v.", err)
			}
		}()

		// Wait enough time for subscriber to subscribe
		time.Sleep(50 * time.Millisecond)
		err = notifier.Notify("added", expected)
		require.NoError(t, err)

		var actual *v1.Event
		select {
		case e := <-events:
			err := e.DataAs(&actual)
			require.NoError(t, err)
		case <-time.After(1 * time.Second):
			t.Logf("Time out waiting for event to arrive")
		}

		assert.Equal(t, expected, actual)
	})

}