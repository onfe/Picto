package server

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

type submission struct {
	ID      string
	Addr    string
	Message *messageEvent
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

func genSubmissionID(addr string) string {
	_, month, day := time.Now().Date()
	addrSansPort := strings.Split(addr, ":")[0]
	id := addrSansPort + "-" + strconv.Itoa(day) + "-" + month.String()
	log.Println(id)
	return id
}

func (sc *submissionCache) add(s *submission) bool {
	//If we're at capacity, the submission at the tail is rejected.
	if sc.Len == sc.Capacity {
		sc.remove(sc.Submissions["TAIL"].next)
	}

	//Populate submission's ID&next fields
	s.ID = genSubmissionID(s.Addr)

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
