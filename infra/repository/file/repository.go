package file

import (
	"bufio"
	"errors"
	"github.com/koushamad/Enigma/domain/entity"
	"log"
	"os"
)

const (
	FILE_PATH = "var/reflectors"
)

var (
	reflects []entity.Reflect
)

type Reflector struct{}

func init() {
	if f, err := os.Open(FILE_PATH); err == nil {
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

func (Reflector) Erase() {
	reflects = reflects[0:0]
}

func (Reflector) GetByIndex(index int) (entity.Reflect, error) {
	if index < len(reflects) {
		return reflects[index], nil
	}

	return entity.Reflect{}, errors.New("can not found reflector")
}

func (r Reflector) Push(reflect entity.Reflect) {
	reflects = append(reflects, reflect)
}

func (r Reflector) Pop() entity.Reflect {
	var reflect entity.Reflect
	reflect, reflects = reflects[len(reflects)-1], reflects[:len(reflects)-1]

	return reflect
}

func (r Reflector) Save() error {

	if err := r.removeFileIfExist(); err != nil {
		return err
	}

	for _, ref := range r.GetAll() {
		if err := r.push(&ref); err != nil {
			return err
		}
	}
	return nil
}

func (Reflector) push(r *entity.Reflect) error {
	if f, err := os.OpenFile(FILE_PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
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

func (r Reflector) removeFileIfExist() error {
	if r.fileIsExist() {
		if err := os.Remove(FILE_PATH); err != nil {
			return err
		}
	}

	return nil
}

func (r Reflector) fileIsExist() bool {
	info, err := os.Stat(FILE_PATH)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
