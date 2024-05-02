package storage

import (
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
		// TODO: If you can't get this connection to work, you can temporarily disable it & try to connect with cqlsh first
		return nil, err
	}
	return &Storage{session: session}, nil
}

func (s *Storage) SaveMessage(roomName string, userID string, message string, timestamp time.Time) error {
	query := "INSERT INTO chat (room_name, timestamp, user_id, message) VALUES (?, ?, ?, ?)"
	return s.session.Query(query, roomName, timestamp, userID, message).Exec()
}

func (s *Storage) GetMessages(roomName string, limit int) ([]map[string]interface{}, error) {
	query := "SELECT timestamp, user_id, message FROM chat WHERE room_name = ? LIMIT ?"
	iter := s.session.Query(query, roomName, limit).Iter()
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
