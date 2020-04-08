package server

import "errors"

type submission struct {
	Sender  string
	Message *MessageEvent
	next    string // doubly  .-''-.''-.
	prev    string // linked  '-..'-..-' bois
}

//SubmissionCache holds submissions for a SubmissionRoom
type SubmissionCache struct {
	Submissions map[string]*submission //Essentially a doubly linked list stored in a map for random access.
	Capacity    int
	Len         int
}

func newSubmissionCache(capacity int) *SubmissionCache {
	sc := SubmissionCache{
		Submissions: make(map[string]*submission, capacity),
		Capacity:    capacity,
		Len:         0,
	}
	sc.Submissions["HEAD"] = &submission{
		prev: "TAIL",
	}
	sc.Submissions["TAIL"] = &submission{
		next: "HEAD",
	}
	return &sc
}

func (sc *SubmissionCache) add(s *submission) bool {
	//If we're at capacity, the submission at the tail is rejected.
	if sc.Len == sc.Capacity {
		sc.remove(sc.Submissions["TAIL"].next)
	}

	_, alreadySubmitted := sc.Submissions[s.Sender]

	if !alreadySubmitted {
		//Update submission's prev&next
		s.next = "HEAD"
		s.prev = sc.Submissions["HEAD"].prev

		//Squishing the new submission between the HEAD elem and the most recent submission
		sc.Submissions[s.prev].next = s.Sender //Update previously most recent submission's 'next' field
		sc.Submissions["HEAD"].prev = s.Sender //Update head's prev field

		sc.Len++
	}

	//Add/update submission to/in Submissions map
	sc.Submissions[s.Sender] = s

	return alreadySubmitted
}

func (sc *SubmissionCache) remove(sender string) error {
	toDel, exists := sc.Submissions[sender]
	if !exists {
		return errors.New("could not find submission from sender: " + sender)
	}

	//Patching the submission's neighbours together before yeeting it out
	sc.Submissions[toDel.prev].next = toDel.next //Update previous elem's next field
	sc.Submissions[toDel.next].prev = toDel.prev //Update next elem's prev field

	delete(sc.Submissions, toDel.Sender) //Remove the submission from the Submissions map

	sc.Len--

	return nil
}

//getAll should return the submissions in the order of oldest first.
func (sc *SubmissionCache) getAll() []*submission {
	submissions := make([]*submission, sc.Len)

	i := 0
	key := sc.Submissions["TAIL"].next //The key of the oldest submission

	for key != "HEAD" {
		submissions[i] = sc.Submissions[key]
		key = sc.Submissions[key].next //Move onto the next submission
	}

	return submissions
}
