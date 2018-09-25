package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Any interface{}

type Monad func(error) (Any, error)

func Return(v Any) Monad {
	return func(s error) (Any, error) {
		return v, s
	}
}

func Bind(m Monad, f func(Any) Monad) Monad {
	return func(s error) (Any, error) {
		a, b := m(s)
		if b != nil {
			return nil, b
		}
		return f(a)(b)
	}
}

func ReadFile(filename Any) Monad {
	log.Println("reading the file")

	filenameString := filename.(string)
	return func(error) (Any, error) {
		return ioutil.ReadFile(filenameString)
	}
}

func ToJSON(v Any) Monad {
	log.Println("unmarshalling the file to json")

	vBytes := v.([]byte)
	return func(s error) (Any, error) {
		type actor struct {
			Name   string `json:"name"`
			Age    uint   `json:"age"`
			Height uint   `json:"height"`
		}
		a := actor{}
		err := json.Unmarshal(vBytes, &a)
		return a, err
	}
}

func main() {
	monad := Return("actor.json")

	monad = Bind(monad, ReadFile)
	monad = Bind(monad, ToJSON)

	actorJSON, err := monad(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(actorJSON)
}
