package ejabber

import (
	"fmt"
	"log"
	"os"

	"gosrc.io/xmpp"
	"gosrc.io/xmpp/stanza"
)

type EjabberClient struct {
	Client     *xmpp.Client
	StreamMngr *xmpp.StreamManager
}

func (ec *EjabberClient) ConnectToIMserver(jid, passwd string) error {
	config := xmpp.Config{
		TransportConfiguration: xmpp.TransportConfiguration{
			Address: "babyeah.online:5222",
		},
		Jid:          jid,
		Credential:   xmpp.Password(passwd),
		StreamLogger: os.Stdout,
		Insecure:     false,
	}

	router := xmpp.NewRouter()
	router.HandleFunc("message", echoMessage)

	client, err := xmpp.NewClient(&config, router, func(e error) { log.Fatal(e.Error()) })
	if err != nil {
		return fmt.Errorf("error creating client: %v", err)
	}
	ec.Client = client

	cm := xmpp.NewStreamManager(client, nil)
	ec.StreamMngr = cm

	return nil
}

func echoMessage(s xmpp.Sender, p stanza.Packet) {
	msg, ok := p.(stanza.Message)
	if !ok {
		_, _ = fmt.Fprintf(os.Stdout, "Ignoring packet: %T\n", p)
		return
	}

	_, _ = fmt.Fprintf(os.Stdout, "Body = %s - from = %s\n", msg.Body, msg.From)
	reply := stanza.Message{Attrs: stanza.Attrs{To: msg.From}, Body: msg.Body}
	_ = s.Send(reply)
}
