package server

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	invisible string = "invisible"
	visible   string = "visible"
)

type moderatedMessage struct {
	ID       string
	SenderIP string
	Message  *eventWrapper
	State    string
	next     *moderatedMessage // doubly  .-''-.''-.
	prev     *moderatedMessage // linked  '-..'-..-' bois
}

//ModerationCache holds moderatedMessages for a ModeratedRoom
type moderationCache struct {
	Messages     map[string]*moderatedMessage //Essentially a doubly linked list stored in a map for random access.
	Capacity     int
	Len          int
	VisibleCount int
}

func newModerationCache(capacity int) *moderationCache {
	sc := moderationCache{
		Messages: make(map[string]*moderatedMessage, capacity),
		Capacity: capacity,
		Len:      0,
	}
	sc.Messages["HEAD"] = &moderatedMessage{
		ID:   "HEAD",
		next: nil,
		prev: nil,
	}
	sc.Messages["TAIL"] = &moderatedMessage{
		ID:   "TAIL",
		next: sc.Messages["HEAD"],
		prev: nil,
	}
	sc.Messages["HEAD"].prev = sc.Messages["TAIL"]
	return &sc
}

func (sc *moderationCache) genMessageID(ip string) string {
	tenthminute := time.Now().Minute() / 10
	ipSansPort := strings.Split(ip, ":")[0]
	id := ipSansPort + "-" + strconv.Itoa(tenthminute)
	return id
}

func (sc *moderationCache) add(s *moderatedMessage) bool {
	//Populate moderatedMessage's ID & state fields
	s.ID = sc.genMessageID(s.SenderIP)
	s.State = visible

	prevMessage, alreadyMessaged := sc.Messages[s.ID]

	//If the message they've already sent within this time interval was hidden, then we ignore it.
	if alreadyMessaged && prevMessage.State == invisible {
		return alreadyMessaged
	}

	//If it's a new moderatedMessage, and we're at capacity, the oldest moderatedMessage is discarded.
	if !alreadyMessaged && sc.Len == sc.Capacity {
		sc.remove(sc.Messages["TAIL"].next.ID)
	}

	if !alreadyMessaged {
		//If we're not overwriting a moderatedMessage, it's squished between HEAD and HEAD.prev
		s.next = sc.Messages["HEAD"]
		s.prev = sc.Messages["HEAD"].prev

		//Squishing the new moderatedMessage between the HEAD elem and the most recent moderatedMessage
		sc.Messages["HEAD"].prev.next = s //Update previously most recent moderatedMessage's 'next' field
		sc.Messages["HEAD"].prev = s      //Update head's prev field

		sc.Len++
		sc.VisibleCount++
	} else {
		//If we're overwriting a moderatedMessage, we just need to update pointers
		s.next = sc.Messages[s.ID].next
		s.prev = sc.Messages[s.ID].prev

		s.prev.next = s
		s.next.prev = s
	}

	//Add/update moderatedMessage to/in Messages map
	sc.Messages[s.ID] = s

	return alreadyMessaged
}

func (sc *moderationCache) remove(ID string) (*moderatedMessage, error) {
	toDel, exists := sc.Messages[ID]
	if !exists {
		return nil, errors.New("could not find moderatedMessage with ID: " + ID)
	}

	//Patching the moderatedMessage's neighbours together before yeeting it out
	toDel.prev.next = toDel.next //Update previous elem's next field
	toDel.next.prev = toDel.prev //Update next elem's prev field

	//Updating VisibleCount
	if toDel.State == visible {
		sc.VisibleCount--
	}

	delete(sc.Messages, toDel.ID) //Remove the moderatedMessage from the Messages map

	sc.Len--

	return toDel, nil
}

func (sc *moderationCache) setState(ID, newState string) (string, error) {
	moderatedMessage, exists := sc.Messages[ID]
	if !exists {
		return "", errors.New("could not find moderatedMessage with ID: " + ID)
	}

	for _, state := range []string{visible, invisible} {
		if state == newState {
			log.Println("Updating moderatedMessage ID " + moderatedMessage.ID + " of state " + moderatedMessage.State)
			//Delete the old message
			delete(sc.Messages, moderatedMessage.ID)

			//Update VisibleCount
			if moderatedMessage.State != visible && newState == visible {
				sc.VisibleCount++
			} else if moderatedMessage.State == visible && newState != visible {
				sc.VisibleCount--
			}

			//Update the state and ID of the message
			moderatedMessage.State = newState
			moderatedMessage.ID = sc.genMessageID(moderatedMessage.SenderIP)

			//Put it back into the message map
			sc.Messages[moderatedMessage.ID] = moderatedMessage

			//If VisibleCount > MaxVisibleMessages, delete the oldest visible moderatedMessage
			if sc.VisibleCount > MaxVisibleMessages {
				moderatedMessage := sc.Messages["TAIL"].next
				for moderatedMessage.State != visible {
					moderatedMessage = moderatedMessage.next
				}
				sc.remove(moderatedMessage.ID)
			}

			log.Println("Updated moderatedMessage ID " + moderatedMessage.ID + " to " + newState + " state")
			log.Println(sc.getChainString())

			return moderatedMessage.ID, nil
		}
	}

	return "", errors.New("unrecognised moderatedMessage state")
}

//getAll should return the moderatedMessage in the order of oldest first.
func (sc *moderationCache) getAll() []*moderatedMessage {
	moderatedMessages := make([]*moderatedMessage, sc.Len)

	i := 0
	moderatedMessage := sc.Messages["TAIL"].next //The key of the oldest message

	for moderatedMessage.ID != "HEAD" {
		moderatedMessages[i] = moderatedMessage
		moderatedMessage = moderatedMessage.next //Move onto the next message
		i++
	}

	return moderatedMessages
}

func (sc *moderationCache) getChainString() string {
	chainString := "TAIL"

	moderatedMessage := sc.Messages["TAIL"]

	for moderatedMessage.ID != "HEAD" {
		moderatedMessage = moderatedMessage.next
		chainString += " <-> " + moderatedMessage.ID
	}

	return chainString
}
