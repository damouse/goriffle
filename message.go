package riffle

// Message is a generic container for a WAMP message.
type Message interface {
	messageType() messageType
}

var (
	abortUnexpectedMsg = &abort{
		Details: map[string]interface{}{},
		Reason:  "turnpike.error.unexpected_message_type",
	}
	abortNoAuthHandler = &abort{
		Details: map[string]interface{}{},
		Reason:  "turnpike.error.no_handler_for_authmethod",
	}
	abortAuthFailure = &abort{
		Details: map[string]interface{}{},
		Reason:  "turnpike.error.authentication_failure",
	}
	goodbyeSession = &goodbye{
		Details: map[string]interface{}{},
		Reason:  ErrCloseRealm,
	}
)

type messageType int

func (mt messageType) New() Message {
	switch mt {
	case HELLO:
		return new(hello)
	case WELCOME:
		return new(welcome)
	case ABORT:
		return new(abort)
	case CHALLENGE:
		return new(challenge)
	case AUTHENTICATE:
		return new(authenticate)
	case GOODBYE:
		return new(goodbye)
	case HEARTBEAT:
		return new(heartbeat)
	case ERROR:
		return new(errorMessage)

	case PUBLISH:
		return new(publish)
	case PUBLISHED:
		return new(published)

	case SUBSCRIBE:
		return new(subscribe)
	case SUBSCRIBED:
		return new(subscribed)
	case UNSUBSCRIBE:
		return new(unsubscribe)
	case UNSUBSCRIBED:
		return new(unsubscribed)
	case EVENT:
		return new(event)

	case CALL:
		return new(call)
	case CANCEL:
		return new(cancel)
	case RESULT:
		return new(result)

	case REGISTER:
		return new(register)
	case REGISTERED:
		return new(registered)
	case UNREGISTER:
		return new(unregister)
	case UNREGISTERED:
		return new(unregistered)
	case INVOCATION:
		return new(invocation)
	case INTERRUPT:
		return new(interrupt)
	case YIELD:
		return new(yield)
	default:
		// TODO: allow custom message types?
		return nil
	}
}

func (mt messageType) String() string {
	switch mt {
	case HELLO:
		return "HELLO"
	case WELCOME:
		return "WELCOME"
	case ABORT:
		return "ABORT"
	case CHALLENGE:
		return "CHALLENGE"
	case AUTHENTICATE:
		return "AUTHENTICATE"
	case GOODBYE:
		return "GOODBYE"
	case HEARTBEAT:
		return "HEARTBEAT"
	case ERROR:
		return "ERROR"

	case PUBLISH:
		return "PUBLISH"
	case PUBLISHED:
		return "PUBLISHED"

	case SUBSCRIBE:
		return "SUBSCRIBE"
	case SUBSCRIBED:
		return "SUBSCRIBED"
	case UNSUBSCRIBE:
		return "UNSUBSCRIBE"
	case UNSUBSCRIBED:
		return "UNSUBSCRIBED"
	case EVENT:
		return "EVENT"

	case CALL:
		return "CALL"
	case CANCEL:
		return "CANCEL"
	case RESULT:
		return "RESULT"

	case REGISTER:
		return "REGISTER"
	case REGISTERED:
		return "REGISTERED"
	case UNREGISTER:
		return "UNREGISTER"
	case UNREGISTERED:
		return "UNREGISTERED"
	case INVOCATION:
		return "INVOCATION"
	case INTERRUPT:
		return "INTERRUPT"
	case YIELD:
		return "YIELD"
	default:
		// TODO: allow custom message types?
		panic("Invalid message type")
	}
}

const (
	HELLO        messageType = 1
	WELCOME      messageType = 2
	ABORT        messageType = 3
	CHALLENGE    messageType = 4
	AUTHENTICATE messageType = 5
	GOODBYE      messageType = 6
	HEARTBEAT    messageType = 7
	ERROR        messageType = 8

	PUBLISH   messageType = 16 //	Tx 	Rx
	PUBLISHED messageType = 17 //	Rx 	Tx

	SUBSCRIBE    messageType = 32 //	Rx 	Tx
	SUBSCRIBED   messageType = 33 //	Tx 	Rx
	UNSUBSCRIBE  messageType = 34 //	Rx 	Tx
	UNSUBSCRIBED messageType = 35 //	Tx 	Rx
	EVENT        messageType = 36 //	Tx 	Rx

	CALL   messageType = 48 //	Tx 	Rx
	CANCEL messageType = 49 //	Tx 	Rx
	RESULT messageType = 50 //	Rx 	Tx

	REGISTER     messageType = 64 //	Rx 	Tx
	REGISTERED   messageType = 65 //	Tx 	Rx
	UNREGISTER   messageType = 66 //	Rx 	Tx
	UNREGISTERED messageType = 67 //	Tx 	Rx
	INVOCATION   messageType = 68 //	Tx 	Rx
	INTERRUPT    messageType = 69 //	Tx 	Rx
	YIELD        messageType = 70 //	Rx 	Tx
)

// [HELLO, Realm|uri, Details|dict]
type hello struct {
	Realm   string
	Details map[string]interface{}
}

func (msg *hello) messageType() messageType {
	return HELLO
}

// [WELCOME, Session|id, Details|dict]
type welcome struct {
	Id      uint
	Details map[string]interface{}
}

func (msg *welcome) messageType() messageType {
	return WELCOME
}

// [ABORT, Details|dict, Reason|uri]
type abort struct {
	Details map[string]interface{}
	Reason  string
}

func (msg *abort) messageType() messageType {
	return ABORT
}

// [CHALLENGE, AuthMethod|string, Extra|dict]
type challenge struct {
	AuthMethod string
	Extra      map[string]interface{}
}

func (msg *challenge) messageType() messageType {
	return CHALLENGE
}

// [AUTHENTICATE, Signature|string, Extra|dict]
type authenticate struct {
	Signature string
	Extra     map[string]interface{}
}

func (msg *authenticate) messageType() messageType {
	return AUTHENTICATE
}

// [GOODBYE, Details|dict, Reason|uri]
type goodbye struct {
	Details map[string]interface{}
	Reason  string
}

func (msg *goodbye) messageType() messageType {
	return GOODBYE
}

// [HEARTBEAT, IncomingSeq|integer, OutgoingSeq|integer
// [HEARTBEAT, IncomingSeq|integer, OutgoingSeq|integer, Discard|string]
type heartbeat struct {
	IncomingSeq uint
	OutgoingSeq uint
	Discard     string
}

func (msg *heartbeat) messageType() messageType {
	return HEARTBEAT
}

// [ERROR, REQUEST.Type|int, REQUEST.Request|id, Details|dict, Error|uri]
// [ERROR, REQUEST.Type|int, REQUEST.Request|id, Details|dict, Error|uri, Arguments|list]
// [ERROR, REQUEST.Type|int, REQUEST.Request|id, Details|dict, Error|uri, Arguments|list, ArgumentsKw|dict]
type errorMessage struct {
	Type        messageType
	Request     uint
	Details     map[string]interface{}
	Error       string
	Arguments   []interface{}          `wamp:"omitempty"`
	ArgumentsKw map[string]interface{} `wamp:"omitempty"`
}

func (msg *errorMessage) messageType() messageType {
	return ERROR
}

// [PUBLISH, Request|id, Options|dict, Domain|uri]
// [PUBLISH, Request|id, Options|dict, Domain|uri, Arguments|list]
// [PUBLISH, Request|id, Options|dict, Domain|uri, Arguments|list, ArgumentsKw|dict]
type publish struct {
	Request     uint
	Options     map[string]interface{}
	Domain      string
	Arguments   []interface{}          `wamp:"omitempty"`
	ArgumentsKw map[string]interface{} `wamp:"omitempty"`
}

func (msg *publish) messageType() messageType {
	return PUBLISH
}

// [PUBLISHED, PUBLISH.Request|id, Publication|id]
type published struct {
	Request     uint
	Publication uint
}

func (msg *published) messageType() messageType {
	return PUBLISHED
}

// [SUBSCRIBE, Request|id, Options|dict, Domain|uri]
type subscribe struct {
	Request uint
	Options map[string]interface{}
	Domain  string
}

func (msg *subscribe) messageType() messageType {
	return SUBSCRIBE
}

// [SUBSCRIBED, SUBSCRIBE.Request|id, Subscription|id]
type subscribed struct {
	Request      uint
	Subscription uint
}

func (msg *subscribed) messageType() messageType {
	return SUBSCRIBED
}

// [UNSUBSCRIBE, Request|id, SUBSCRIBED.Subscription|id]
type unsubscribe struct {
	Request      uint
	Subscription uint
}

func (msg *unsubscribe) messageType() messageType {
	return UNSUBSCRIBE
}

// [UNSUBSCRIBED, UNSUBSCRIBE.Request|id]
type unsubscribed struct {
	Request uint
}

func (msg *unsubscribed) messageType() messageType {
	return UNSUBSCRIBED
}

// [EVENT, SUBSCRIBED.Subscription|id, PUBLISHED.Publication|id, Details|dict]
// [EVENT, SUBSCRIBED.Subscription|id, PUBLISHED.Publication|id, Details|dict, PUBLISH.Arguments|list]
// [EVENT, SUBSCRIBED.Subscription|id, PUBLISHED.Publication|id, Details|dict, PUBLISH.Arguments|list,
//     PUBLISH.ArgumentsKw|dict]
type event struct {
	Subscription uint
	Publication  uint
	Details      map[string]interface{}
	Arguments    []interface{}          `wamp:"omitempty"`
	ArgumentsKw  map[string]interface{} `wamp:"omitempty"`
}

func (msg *event) messageType() messageType {
	return EVENT
}

// CallResult represents the result of a CALL.
type callResult struct {
	Args   []interface{}
	Kwargs map[string]interface{}
	Err    string
}

// [CALL, Request|id, Options|dict, Domain|uri]
// [CALL, Request|id, Options|dict, Domain|uri, Arguments|list]
// [CALL, Request|id, Options|dict, Domain|uri, Arguments|list, ArgumentsKw|dict]
type call struct {
	Request     uint
	Options     map[string]interface{}
	Domain      string
	Arguments   []interface{}          `wamp:"omitempty"`
	ArgumentsKw map[string]interface{} `wamp:"omitempty"`
}

func (msg *call) messageType() messageType {
	return CALL
}

// [RESULT, CALL.Request|id, Details|dict]
// [RESULT, CALL.Request|id, Details|dict, YIELD.Arguments|list]
// [RESULT, CALL.Request|id, Details|dict, YIELD.Arguments|list, YIELD.ArgumentsKw|dict]
type result struct {
	Request     uint
	Details     map[string]interface{}
	Arguments   []interface{}          `wamp:"omitempty"`
	ArgumentsKw map[string]interface{} `wamp:"omitempty"`
}

func (msg *result) messageType() messageType {
	return RESULT
}

// [REGISTER, Request|id, Options|dict, Domain|uri]
type register struct {
	Request uint
	Options map[string]interface{}
	Domain  string
}

func (msg *register) messageType() messageType {
	return REGISTER
}

// [REGISTERED, REGISTER.Request|id, Registration|id]
type registered struct {
	Request      uint
	Registration uint
}

func (msg *registered) messageType() messageType {
	return REGISTERED
}

// [UNREGISTER, Request|id, REGISTERED.Registration|id]
type unregister struct {
	Request      uint
	Registration uint
}

func (msg *unregister) messageType() messageType {
	return UNREGISTER
}

// [UNREGISTERED, UNREGISTER.Request|id]
type unregistered struct {
	Request uint
}

func (msg *unregistered) messageType() messageType {
	return UNREGISTERED
}

// [INVOCATION, Request|id, REGISTERED.Registration|id, Details|dict]
// [INVOCATION, Request|id, REGISTERED.Registration|id, Details|dict, CALL.Arguments|list]
// [INVOCATION, Request|id, REGISTERED.Registration|id, Details|dict, CALL.Arguments|list, CALL.ArgumentsKw|dict]
type invocation struct {
	Request      uint
	Registration uint
	Details      map[string]interface{}
	Arguments    []interface{}          `wamp:"omitempty"`
	ArgumentsKw  map[string]interface{} `wamp:"omitempty"`
}

func (msg *invocation) messageType() messageType {
	return INVOCATION
}

// [YIELD, INVOCATION.Request|id, Options|dict]
// [YIELD, INVOCATION.Request|id, Options|dict, Arguments|list]
// [YIELD, INVOCATION.Request|id, Options|dict, Arguments|list, ArgumentsKw|dict]
type yield struct {
	Request     uint
	Options     map[string]interface{}
	Arguments   []interface{}          `wamp:"omitempty"`
	ArgumentsKw map[string]interface{} `wamp:"omitempty"`
}

func (msg *yield) messageType() messageType {
	return YIELD
}

// [CANCEL, CALL.Request|id, Options|dict]
type cancel struct {
	Request uint
	Options map[string]interface{}
}

func (msg *cancel) messageType() messageType {
	return CANCEL
}

// [INTERRUPT, INVOCATION.Request|id, Options|dict]
type interrupt struct {
	Request uint
	Options map[string]interface{}
}

func (msg *interrupt) messageType() messageType {
	return INTERRUPT
}

////////////////////////////////////////
/*
 Begin a whole mess of code we really don't want to get into
 and which pretty much guarantees we'll have to make substantial changes to
 Riffle code: the messages don't have a standardized way of returning their
 TO identity!

 Really, really need this, Short of modifying and standardizing the WAMP changes
 this is unlikely to happen without node monkey-patching. So here we go.
*/
////////////////////////////////////////

type NoDestinationError string

func (e NoDestinationError) Error() string {
	return "cannot determine destination from: " + string(e)
}

// Given a message, return the intended endpoint
func destination(m *Message) (string, error) {
	msg := *m

	switch msg := msg.(type) {

	case *publish:
		return msg.Domain, nil
	case *subscribe:
		return msg.Domain, nil

	// Dealer messages
	case *register:
		return msg.Domain, nil
	case *call:
		return msg.Domain, nil

	default:
		//log.Println("Unhandled message:", msg.messageType())
		return "", NoDestinationError(msg.messageType())
	}
}

// Given a message, return the request uint
func requestID(m *Message) uint {
	switch msg := (*m).(type) {
	case *publish:
		return msg.Request
	case *subscribe:
		return msg.Request
	case *register:
		return msg.Request
	case *call:
		return msg.Request
	}

	return uint(0)
}
