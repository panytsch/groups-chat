package rpc

import (
	"github.com/panytsch/groups-chat/app"
	"github.com/panytsch/groups-chat/messaging/client"
	"github.com/panytsch/groups-chat/messaging/events"
	"github.com/panytsch/groups-chat/models"
	"github.com/panytsch/groups-chat/protocol"
)

type typing struct{}

func newTyping() *typing {
	return &typing{}
}

func (t *typing) Start(c *client.Client, p *protocol.RpcTypingStart) {
	withGroup(p.GroupID, c.UserID, func(group *models.Group) {
		e := events.NewEvent(protocol.EVENT_TYPING_START, app.Timestamp())
		e.Body = protocol.EventTypingStart{
			GroupID: group.ID,
			UserID:  c.UserID,
		}
		e.SendToUsersWithoutMe(c.SessionID, group.UserIDs)
	})
}

func (t *typing) End(c *client.Client, p *protocol.RpcTypingEnd) {
	withGroup(p.GroupID, c.UserID, func(group *models.Group) {
		e := events.NewEvent(protocol.EVENT_TYPING_END, app.Timestamp())
		e.Body = protocol.EventTypingEnd{
			GroupID: group.ID,
			UserID:  c.UserID,
		}
		e.SendToUsersWithoutMe(c.SessionID, group.UserIDs)
	})
}
