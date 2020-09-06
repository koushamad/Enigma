package reflector

import (
	"github.com/koushamad/Enigma/domain/entity"
	"github.com/koushamad/Enigma/infra/repository/file/reflector"
	"log"
	"math/rand"
	"time"
)

type Reflect struct {
	entity     entity.Reflect
	repository reflector.Reflector
}

func (r Reflect) Generate() Reflect {
	ch := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '2', '3', '4', '5', '6', '7', '='}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ch), func(i, j int) { ch[i], ch[j] = ch[j], ch[i] })

	r.entity = entity.Reflect{Chars: ch}

	return r
}

func (r Reflect) Save() Reflect {
	if err := r.repository.Add(&r.entity); err != nil {
		log.Fatal(err)
	}

	return r
}
