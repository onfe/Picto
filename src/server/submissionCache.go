package server

import (
	"errors"
	"log"
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
	Message *eventWrapper
	State   string
	next    *submission // doubly  .-''-.''-.
	prev    *submission // linked  '-..'-..-' bois
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
		ID:   "HEAD",
		next: nil,
		prev: nil,
	}
	sc.Submissions["TAIL"] = &submission{
		ID:   "TAIL",
		next: sc.Submissions["HEAD"],
		prev: nil,
	}
	sc.Submissions["HEAD"].prev = sc.Submissions["TAIL"]
	return &sc
}

func (sc *submissionCache) genSubmissionID(addr, state string) string {
	_, month, day := time.Now().Date()
	//addrSansPort := strings.Split(addr, ":")[0]
	//by default, submission IDs are submitted.
	id := addr + "-" + strconv.Itoa(day) + "-" + month.String() + "-" + state

	//If the state is just submitted, only one submission may exist, so we don't need to add an iterator.
	if state == submitted {
		return id
	}

	i := 0
	iteratedID := id + "-" + strconv.Itoa(i)

	for _, idTaken := sc.Submissions[iteratedID]; idTaken; {
		i++
		iteratedID = id + "-" + strconv.Itoa(i)
		_, idTaken = sc.Submissions[iteratedID]
	}

	return iteratedID
}

func (sc *submissionCache) add(s *submission) bool {
	//If we're at capacity, the submission at the tail is rejected.
	if sc.Len == sc.Capacity {
		sc.remove(sc.Submissions["TAIL"].next.ID)
	}

	//Populate submission's ID*state fields
	s.ID = sc.genSubmissionID(s.Addr, submitted)
	s.State = submitted

	_, alreadySubmitted := sc.Submissions[s.ID]

	if !alreadySubmitted {
		//If we're not overwriting a submission, it's squished between HEAD and HEAD.prev
		s.next = sc.Submissions["HEAD"]
		s.prev = sc.Submissions["HEAD"].prev

		//Squishing the new submission between the HEAD elem and the most recent submission
		sc.Submissions["HEAD"].prev.next = s //Update previously most recent submission's 'next' field
		sc.Submissions["HEAD"].prev = s      //Update head's prev field

		sc.Len++
	} else {
		//If we're overwriting a submission, we just need to update pointers
		s.next = sc.Submissions[s.ID].next
		s.prev = sc.Submissions[s.ID].prev

		s.prev.next = s
		s.next.prev = s
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
	toDel.prev.next = toDel.next //Update previous elem's next field
	toDel.next.prev = toDel.prev //Update next elem's prev field

	delete(sc.Submissions, toDel.ID) //Remove the submission from the Submissions map

	sc.Len--

	return nil
}

func (sc *submissionCache) setState(ID, newState string) (string, error) {
	submission, exists := sc.Submissions[ID]
	if !exists {
		return "", errors.New("could not find submission with ID: " + ID)
	}

	for _, state := range []string{submitted, published, held} {
		if state == newState {
			log.Println("Updating submission ID " + submission.ID + " of state " + submission.State)
			//Delete the old submission
			delete(sc.Submissions, submission.ID)

			//Update the state and ID of the submission
			submission.State = newState
			submission.ID = sc.genSubmissionID(submission.Addr, newState)

			//Put it back into the submissions map
			sc.Submissions[submission.ID] = submission

			log.Println("Updated submission ID " + submission.ID + " to " + newState + " state")
			log.Println(sc.getChainString())

			return submission.ID, nil
		}
	}

	return "", errors.New("unrecognised submission state")
}

//getAll should return the submissions in the order of oldest first.
func (sc *submissionCache) getAll() []*submission {
	submissions := make([]*submission, sc.Len)

	i := 0
	submission := sc.Submissions["TAIL"].next //The key of the oldest submission

	for submission.ID != "HEAD" {
		submissions[i] = submission
		submission = submission.next //Move onto the next submission
		i++
	}

	return submissions
}

func (sc *submissionCache) getChainString() string {
	chainString := "TAIL"

	submission := sc.Submissions["TAIL"]

	for submission.ID != "HEAD" {
		submission = submission.next
		chainString += " <-> " + submission.ID
	}

	return chainString
}
