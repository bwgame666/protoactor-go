package actor

// The Producer type is a function that creates a new actor
type Producer func() Actor
type ProducerWithActorSystem func(system *ActorSystem) Actor

// Actor is the interface that defines the Receive method.
//
// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
type Actor interface {
	Receive(c Context)
}

// The ReceiveFunc type is an adapter to allow the use of ordinary functions as actors to process messages
type ReceiveFunc func(c Context)

// Receive calls f(c)
func (f ReceiveFunc) Receive(c Context) {
	f(c)
}

type ReceiverFunc func(c ReceiverContext, envelope *MessageEnvelope)

type SenderFunc func(c SenderContext, target *PID, envelope *MessageEnvelope)

type ContextDecoratorFunc func(ctx Context) Context
