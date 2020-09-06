package reflector

import (
	"bufio"
	"errors"
	"github.com/koushamad/Enigma/domain/entity"
	"log"
	"os"
)

var (
	reflects []entity.Reflect
)

type Reflector struct{}

func init() {
	if f, err := os.Open("var/reflectors"); err == nil {
		defer f.Close()
		scaner := bufio.NewScanner(f)
		for scaner.Scan() {
			reflects = append(reflects, entity.Reflect{Chars: scaner.Bytes()})
		}
		if err := scaner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func (Reflector) GetAll() []entity.Reflect {
	return reflects
}

func (Reflector) GetByIndex(index int) (entity.Reflect, error) {
	if index < len(reflects) {
		return reflects[index], nil
	}

	return entity.Reflect{}, errors.New("can not found reflector")
}

func (Reflector) Add(r *entity.Reflect) error {
	if f, err := os.OpenFile("var/reflectors", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		defer f.Close()
		if _, err = f.Write(r.Chars); err != nil {
			return err
		} else {
			if _, err := f.WriteString("\n"); err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}
