package storage

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

type Storage struct {
	session *gocql.Session
}

func NewStorage(hosts ...string) (*Storage, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = "chat"
	cluster.NumConns = 2
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	log.Printf("storage is connected")
	return &Storage{session: session}, nil
}

func (s *Storage) SaveMessage(channelId gocql.UUID, userID gocql.UUID, content string, timestamp time.Time) error {
	messageId := gocql.TimeUUID()
	query := "INSERT INTO chat.messages (channel_id, id, user_id, content, created_at) VALUES (?, ?, ?, ?, ?)"
	return s.session.Query(query, channelId, messageId, userID, content, timestamp).Exec()
}

func (s *Storage) GetMessages(channelId string, limit int) ([]map[string]interface{}, error) {
	query := "SELECT timestamp, user_id, message FROM chat.messages WHERE channel_name = ? LIMIT ?"
	iter := s.session.Query(query, channelId, limit).Iter()
	var messages []map[string]interface{}
	for {
		row := make(map[string]interface{})
		if !iter.MapScan(row) {
			break
		}
		messages = append(messages, row)
	}
	return messages, iter.Close()
}

func (s *Storage) Close() {
	s.session.Close()
}
