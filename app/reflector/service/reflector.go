package service

import (
	"github.com/koushamad/Enigma/domain/entity"
	"github.com/koushamad/Enigma/infra/repository/file"
	"math/rand"
	"time"
)

type Reflect struct {
	entity     entity.Reflect
	repository file.Reflector
}

func (r Reflect) Generate() {
	r.repository.Erase()
	for i := 0; i < 100; i++ {
		r.repository.Push(r.getReflector())
	}

	if err := r.repository.Save(); err != nil {
		panic(err)
	}
}

func (r Reflect) getReflector() entity.Reflect {
	ch := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '2', '3', '4', '5', '6', '7', '='}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ch), func(i, j int) { ch[i], ch[j] = ch[j], ch[i] })

	return entity.Reflect{Chars: ch}
}
