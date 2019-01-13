package gene

import (
	"fmt"
	"math/rand"
	"sync"
)

var m sync.Mutex
var count = 0

var firstNames = []strings{
	"Mom",
	"Bob",
	"Daibutsu",
	"Niku",
	"Shitty",
	"Fuckedup",
}

// Email generates a uniq email address every time it is called.
// It is intended to be used for creating new user accounts wo/ worrying about
// an email address already being used.
// Note: This doesn't check the DB to verify that the email is not taken,
// so if you are generating email addresses another way it is possible to have collisions.
func Email() string {
	m.Lock()
	defer m.Unlock()
	name := firstNames[rand.Intn(len(firstNames))]
	ret := fmt.Sprintf("%s-%d@example.com", name, count)
	count++
	return ret
}
