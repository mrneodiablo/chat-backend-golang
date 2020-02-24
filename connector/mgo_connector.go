package connector

import "hope-pet-chat-backend/models/db"

const (
	CHAT_READMESSAGE_COLLECTION = "chat_read"
	CHAT_HEARTBEAT_COLLECTION = "chat_heartbeat"
	CHAT_UNREADMESSAGE_COLLECTION = "chat_unread"

)

func SessionConnectCollection(COLLECTION string) *db.Collection {
	return db.NewCollectionSession(COLLECTION)
}
