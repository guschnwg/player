package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ericdaugherty/alexa-skills-kit-golang"
)

var a = &alexa.Alexa{
	ApplicationID:       "amzn1.ask.skill.4653a02c-e0e1-46b3-ade8-2315b997c94f",
	RequestHandler:      &HelloWorld{},
	IgnoreApplicationID: true,
	IgnoreTimestamp:     true,
}

const cardTitle = "HelloWorld"

// HelloWorld handles reqeusts from the HelloWorld skill.
type HelloWorld struct{}

// Handle processes calls from Lambda
func Handle(ctx context.Context, requestEnv *alexa.RequestEnvelope) (interface{}, error) {
	return a.ProcessRequest(ctx, requestEnv)
}

// OnSessionStarted called when a new session is created.
func (h *HelloWorld) OnSessionStarted(ctx context.Context, request *alexa.Request, session *alexa.Session, ctxPtr *alexa.Context, response *alexa.Response) error {

	log.Printf("OnSessionStarted requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	return nil
}

// OnLaunch called with a reqeust is received of type LaunchRequest
func (h *HelloWorld) OnLaunch(ctx context.Context, request *alexa.Request, session *alexa.Session, ctxPtr *alexa.Context, response *alexa.Response) error {
	speechText := "Welcome to the Alexa Skills Kit, you can say hello"

	log.Printf("OnLaunch requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	response.SetSimpleCard(cardTitle, speechText)
	response.SetOutputText(speechText)
	response.SetRepromptText(speechText)

	response.ShouldSessionEnd = true

	return nil
}

// OnIntent called with a reqeust is received of type IntentRequest
func (h *HelloWorld) OnIntent(ctx context.Context, request *alexa.Request, session *alexa.Session, ctxPtr *alexa.Context, response *alexa.Response) error {

	log.Printf("OnIntent requestId=%s, sessionId=%s, intent=%s", request.RequestID, session.SessionID, request.Intent.Name)

	switch request.Intent.Name {
	case "LaunchRequest":
		log.Println("HelloWorldIntent triggered")
		speechText := "Hello World"

		response.SetSimpleCard(cardTitle, speechText)
		response.SetOutputText(speechText)

		log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
	case "SongPlayerIntent":
		response.AddAudioPlayer(
			"AudioPlayer.Play",
			"REPLACE_ALL",
			"testinho",
			"https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3",
			0,
		)
	case "AMAZON.HelpIntent":
		log.Println("AMAZON.HelpIntent triggered")
		speechText := "You can say hello to me!"

		response.SetSimpleCard("HelloWorld", speechText)
		response.SetOutputText(speechText)
		response.SetRepromptText(speechText)
	case "AMAZON.StopIntent":
		response.SetOutputText("Parar")
		response.Directives = append(response.Directives, alexa.AudioPlayerDirective{
			Type: "AudioPlayer.Stop",
		})
	case "AMAZON.PauseIntent":
		response.Directives = append(response.Directives, alexa.AudioPlayerDirective{
			Type: "AudioPlayer.Stop",
		})
		response.SetOutputText("Pausar")
	case "AMAZON.ResumeIntent":
		response.AddAudioPlayer(
			"AudioPlayer.Play",
			"REPLACE_ALL",
			"testinho",
			"https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3",
			50000,
		)
		response.SetOutputText("Continuar")
	default:
		response.SetOutputText("Sei lá o que você quer")
	}

	return nil
}

// OnSessionEnded called with a reqeust is received of type SessionEndedRequest
func (h *HelloWorld) OnSessionEnded(ctx context.Context, request *alexa.Request, session *alexa.Session, ctxPtr *alexa.Context, response *alexa.Response) error {

	log.Printf("OnSessionEnded requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	return nil
}

// Alexa ...
func Alexa(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)

	request := alexa.RequestEnvelope{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	resp, _ := a.ProcessRequest(r.Context(), &request)

	json.NewEncoder(w).Encode(resp)
}
