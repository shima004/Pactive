package http

import (
	"context"
	"log"

	"github.com/go-fed/activity/streams/vocab"
	"github.com/shima004/pactive/streams"
)

type InboxCallback interface {
	PersonCallback(c context.Context, person vocab.ActivityStreamsPerson) error
	AcceptCallback(c context.Context, accept vocab.ActivityStreamsAccept) error
	FollowCallback(c context.Context, follow vocab.ActivityStreamsFollow) (Reaction, error)
	GetJsonResolver() (*streams.JSONResolver, error)
}

type Reaction struct {
	Accept vocab.ActivityStreamsAccept
	Reject vocab.ActivityStreamsReject
}

type InboxCallbackFuncs struct {
}

func (f *InboxCallbackFuncs) GetJsonResolver() (*streams.JSONResolver, error) {
	return streams.NewJSONResolver(
		f.PersonCallback,
		f.AcceptCallback,
		f.FollowCallback,
	)
}

func (f *InboxCallbackFuncs) PersonCallback(c context.Context, person vocab.ActivityStreamsPerson) error {
	log.Println("PersonCallback")
	log.Println(person.Serialize())
	return nil
}

func (f *InboxCallbackFuncs) AcceptCallback(c context.Context, accept vocab.ActivityStreamsAccept) error {
	log.Println("AcceptCallback")
	log.Println(accept.Serialize())
	return nil
}

func (f *InboxCallbackFuncs) FollowCallback(c context.Context, follow vocab.ActivityStreamsFollow) (Reaction, error) {
	log.Println("FollowCallback")
	log.Println(follow.Serialize())

	accept := streams.NewActivityStreamsAccept()
	accept.SetActivityStreamsObject(follow.GetActivityStreamsObject())
	accept.SetActivityStreamsActor(follow.GetActivityStreamsActor())

	return Reaction{
		Accept: accept,
		Reject: nil,
	}, nil

}
