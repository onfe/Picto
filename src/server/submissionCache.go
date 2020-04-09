package server

import (
	"errors"
	"strconv"
	"time"
)

const (
	submitted string = "submitted"
	published string = "published"
	held      string = "held"
)

type submission struct {
	ID      string
	Addr    string
	Message *messageEvent
	State   string
	next    string // doubly  .-''-.''-.
	prev    string // linked  '-..'-..-' bois
}

//SubmissionCache holds submissions for a SubmissionRoom
type submissionCache struct {
	Submissions map[string]*submission //Essentially a doubly linked list stored in a map for random access.
	Capacity    int
	Len         int
}

func newSubmissionCache(capacity int) *submissionCache {
	sc := submissionCache{
		Submissions: make(map[string]*submission, capacity),
		Capacity:    capacity,
		Len:         0,
	}
	sc.Submissions["HEAD"] = &submission{
		next: "END",
		prev: "TAIL",
	}
	sc.Submissions["TAIL"] = &submission{
		next: "HEAD",
		prev: "START",
	}
	return &sc
}

func genSubmissionID(addr, state string) string {
	_, month, day := time.Now().Date()
	//addrSansPort := strings.Split(addr, ":")[0]
	//by default, submission IDs are submitted.
	id := addr + "-" + strconv.Itoa(day) + "-" + month.String() + "-" + state
	return id
}

func (sc *submissionCache) add(s *submission) bool {
	//If we're at capacity, the submission at the tail is rejected.
	if sc.Len == sc.Capacity {
		sc.remove(sc.Submissions["TAIL"].next)
	}

	//Populate submission's ID*state fields
	s.ID = genSubmissionID(s.Addr, submitted)
	s.State = submitted

	_, alreadySubmitted := sc.Submissions[s.ID]

	if !alreadySubmitted {
		//If we're not overwriting a submission, it's squished between HEAD and HEAD.prev
		s.next = "HEAD"
		s.prev = sc.Submissions["HEAD"].prev

		//Squishing the new submission between the HEAD elem and the most recent submission
		sc.Submissions[sc.Submissions["HEAD"].prev].next = s.ID //Update previously most recent submission's 'next' field
		sc.Submissions["HEAD"].prev = s.ID                      //Update head's prev field

		sc.Len++
	} else {

		//If we're overwriting a submission, we just need to update s.prev and s.next
		s.next = sc.Submissions[s.ID].next
		s.prev = sc.Submissions[s.ID].prev
	}

	//Add/update submission to/in Submissions map
	sc.Submissions[s.ID] = s

	return alreadySubmitted
}

func (sc *submissionCache) remove(ID string) error {
	toDel, exists := sc.Submissions[ID]
	if !exists {
		return errors.New("could not find submission with ID: " + ID)
	}

	//Patching the submission's neighbours together before yeeting it out
	sc.Submissions[toDel.prev].next = toDel.next //Update previous elem's next field
	sc.Submissions[toDel.next].prev = toDel.prev //Update next elem's prev field

	delete(sc.Submissions, toDel.ID) //Remove the submission from the Submissions map

	sc.Len--

	return nil
}

func (sc *submissionCache) setState(ID, newState string) (string, error) {
	toChange, exists := sc.Submissions[ID]
	if !exists {
		return "", errors.New("could not find submission with ID: " + ID)
	}

	for _, state := range []string{submitted, published, held} {
		if state == newState {
			//Delete the old submission
			delete(sc.Submissions, toChange.ID)

			//Update the state and ID of the submission
			toChange.State = state
			toChange.ID = genSubmissionID(toChange.Addr, toChange.State)

			//Update the neighbour's references to the submission
			sc.Submissions[toChange.prev].next = toChange.ID
			sc.Submissions[toChange.next].prev = toChange.ID

			//Put it back into the submissions map
			sc.Submissions[toChange.ID] = toChange

			return toChange.ID, nil
		}
	}

	return "", errors.New("unrecognised submission state")
}

//getAll should return the submissions in the order of oldest first.
func (sc *submissionCache) getAll() []*submission {
	submissions := make([]*submission, sc.Len)

	i := 0
	key := sc.Submissions["TAIL"].next //The key of the oldest submission

	for key != "HEAD" {
		submissions[i] = sc.Submissions[key]
		key = sc.Submissions[key].next //Move onto the next submission
		i++
	}

	return submissions
}

func (sc *submissionCache) getChainString() string {
	chainString := "TAIL"

	key := "TAIL"

	for key != "HEAD" {
		key = sc.Submissions[key].next
		chainString += " <-> " + key
	}

	return chainString
}
