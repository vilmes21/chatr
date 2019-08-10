package main

func (store *dbStore) CreateSentence(s *Sentence) error {
	_, err := store.db.Query("INSERT INTO sentence(chat_speaker_id, content, time) VALUES ($1,$2, $3)", s.ChatSpeakerId, s.Content, s.Time)
	return err
}

func (store *dbStore) GetSentences(chatId int) ([]*Sentence, error) {
	//todo: use chatId to fetch chatSpeakerIds
	chatSpeakerIds := []int{1,2,3}
	rows, err := store.db.Query("SELECT * from sentence where chat_speaker_id in $1", chatSpeakerIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
 
	sentences := []*Sentence{}
	for rows.Next() {
		s := &Sentence{}
		if err := rows.Scan(&s.ChatSpeakerId, &s.Content, &s.Time); err != nil {
			return nil, err
		}
		
		sentences = append(sentences, s)
	}
	return sentences, nil
}