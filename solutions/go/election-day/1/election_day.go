package electionday

import ( 
	"fmt"
)
// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var counter = initialVotes
    return &counter
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    if (counter == nil){
        return 0;
    }
	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	if (counter != nil){
    	*counter += increment
    }
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var result = ElectionResult {
        Name: candidateName,
        Votes: votes,
    }
    return &result
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
    if (result == nil){
        panic("Null pointer passed.")
    }
	return fmt.Sprintf("%s (%d)", (*result).Name, (*result).Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	_, exists := results[candidate]
    if (exists){
        results[candidate] -= 1
    }
}
