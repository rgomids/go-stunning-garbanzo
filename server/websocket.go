package server

import (
	"encoding/json"
	"fmt"
	"go-stunning-garbanzo/utils"
	"log"

	"github.com/gorilla/websocket"
)

var (
	// ClientGroupsLength é o tamanho dos grupos que estão no Hub
	ClientGroupsLength = 5
)

// ClientGroup é responsável por manter
// as informações dos usuários do mesmo grupo
// para que possam fazer broadcast das mensagens
type ClientGroup struct {
	// O ID do grupo serve para identifica-lo em meio a outros grupos
	ID string
	// Lista com todas as sessões de clientes conectados nesse grupo
	ClientSessions []*ClientSession
}

// NewClientGroup retorna um novo grupo sem sessões
func NewClientGroup(groupID string) *ClientGroup {
	if groupID == "" {
		return &ClientGroup{
			ID: utils.GenerateULID(),
		}
	}
	return &ClientGroup{
		ID: groupID,
	}
}

// AddClientSession coloca uma sessão nova dentro do grupo
func (cg *ClientGroup) AddClientSession(clientSession *ClientSession) {
	cg.ClientSessions = append(cg.ClientSessions, clientSession)
}

func (cg *ClientGroup) sendMessageToGroup(message *EventMessage) {
	for _, clientSession := range cg.ClientSessions {
		clientSession.SendMessage(message)
	}
}

// ClientSession é responsável por manter as
// informações do usuário que fez a solicitação
type ClientSession struct {
	// ID serve para diferencia-lo dos outros dãã...
	ID string
	// Esse é o grupo que esse client está inserido
	Group string
	// WebsocketConnection carrega a conexão WS do cliente para que ele possa continuar se comunicando
	WebsocketConnection *websocket.Conn
	// SendResponse envia para o usuário as respostas das chamadas
	SendResponse chan []byte
	// FinishClientSession finaliza o hub de operação dele
	FinishClientSession chan bool
	// Este cara vai receber a solicitação e vai trata-la
	EventsHub *EventHub
}

// NewClientSession cria uum novo usuário
func NewClientSession() *ClientSession {
	return &ClientSession{
		ID:                  utils.GenerateULID(),
		SendResponse:        make(chan []byte),
		FinishClientSession: make(chan bool),
	}
}

// SendMessage envia uma mensagem no padrão EventMessage para o Client
func (cs *ClientSession) SendMessage(message *EventMessage) {
	msg, err := json.Marshal(message)
	if err != nil {
		log.Printf("[ERRO] SendMessage can't marshal message: %v", err)
		return
	}
	cs.SendResponse <- msg
}

// SendBroadcast envia uma mensagem no padrão EventMessage para todos os Clients do mesmo grupo
func (cs *ClientSession) SendBroadcast(message *EventMessage) {
	cs.EventsHub.ClientGroups[cs.Group].sendMessageToGroup(message)
}

// ReadFromSocket Pega as mensagens que vem do websocket
func (cs *ClientSession) ReadFromSocket() {
	eventMessageRaw := &EventMessage{}
	for {
		err := cs.WebsocketConnection.ReadJSON(eventMessageRaw)
		if err != nil {
			log.Printf("[ERRO] ReadFromSocket can't read message: %v\n", err)
		}
		cs.SendResponse <- []byte(fmt.Sprintf(`{"event": "%s_PROCESSING"}`, eventMessageRaw.Event))
		go log.Printf("[WS] New Event: %s\n", eventMessageRaw.Event)
		eventMessageRaw.Client = cs
		cs.EventsHub.Messaging <- eventMessageRaw
	}
}

// WriteToSocket Envia a mensagem para o cliente
func (cs *ClientSession) WriteToSocket() {
	defer func() {
		close(cs.SendResponse)
		close(cs.FinishClientSession)
	}()
	for {
		select {
		case message, isOk := <-cs.SendResponse:
			if !isOk {
				log.Printf("[ERRO] WriteToSocket can't get message: %+v", cs.ID)
			}
			// Nessa parte deve ser utilizado a conexão
			cs.WebsocketConnection.WriteMessage(websocket.TextMessage, message)
		// Aqui a sessão do cliente é fechada
		case <-cs.FinishClientSession:
			return
		}
	}
}

// EventMessage é o modelo de
// mensagens que serão compartilhados no webscoket
type EventMessage struct {
	// Event é o tipo de evento que está relacionado a essa chamada
	Event string `json:"event"`

	// Data é a informação que a mensagem está transportando
	Data interface{} `json:"data"`

	// Client é o usuário que fez a solicitação
	Client *ClientSession `json:"-"`
}

// EventHub é o centralizador das mensagens,
// ele é responsável por pegar as mensagens e as enviar
// para as rotas
type EventHub struct {
	// Este é o canal que vai distribuir as mensagens
	Messaging chan *EventMessage
	// Este canal finaliza o hub
	Finish chan bool
	// Essa é a lista com todos os Handlers
	Handlers *EventHandlers
	// Armazena todos os grupos de mensagem
	ClientGroups map[string]*ClientGroup
}

// AddGroup update group list with one new group
func (eh *EventHub) AddGroup(clientGroup *ClientGroup) {
	eh.ClientGroups[clientGroup.ID] = clientGroup
}

// EventHandlers carrega a lista com todas as possiveis
// chamadas e seus handlers
type EventHandlers struct {
	HandlerList map[string]func(*EventMessage)
}

// NewEventHub cria o novo EventHub com o channel já iniciado
func NewEventHub() *EventHub {
	return &EventHub{
		Messaging: make(chan *EventMessage),
		Finish:    make(chan bool),
		Handlers: &EventHandlers{
			HandlerList: make(map[string]func(*EventMessage)),
		},
		ClientGroups: make(map[string]*ClientGroup, ClientGroupsLength),
	}
}

// AddHandler adiciona um novo handler para as chamadas
func (eh *EventHub) AddHandler(event string, f func(*EventMessage)) {
	eh.Handlers.HandlerList[event] = f
}

// Run aqui é o hub onde as mensagens vão ser lidas do messaging
// e posteriormente serem enviadas
func (eh *EventHub) Run() {
	defer func() {
		close(eh.Messaging)
		close(eh.Finish)
	}()
	for {
		select {
		// Sempre que uma mensagem for recebida ela deve ser enviada aqté este lugar
		case message := <-eh.Messaging:
			// Aqui devem ficar as verificações sobre quais eventos estão sendo enviados, pois assim podemos direcionar para cada handler'
			go EventDispatcher(eh.Handlers, message)
		// Se receber algo nesse canal o hub é finalizado
		// Essa chamada deve ser feita apenas em casos especificos pois
		// assim que o hub for fechado a aplicação é encerrada
		case <-eh.Finish:
			return
		}
	}
}

// EventDispatcher é o responsável por tratar as mensagens recebidas
// pelo websocket e direcionalas ao handler correto
//
// Essa função poderia estar um um arquivo próprio dentro deste pacote, pois a mesma vai acabar ficando muito grande
func EventDispatcher(handlers *EventHandlers, message *EventMessage) {
	if f, ok := handlers.HandlerList[message.Event]; ok {
		f(message)
	} else {
		message.Client.SendMessage(&EventMessage{Event: "EVENT_NOT_FOUND"})
	}
}
