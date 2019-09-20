package models

import (
	"go-stunning-garbanzo/utils"
)

// Card é a estrutura basica para os objetos de cartas
type Card struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	ManaCost *Mana             `json:"mana_cost"`
	Image    *Image            `json:"image"`
	Type     string            `json:"type"`
	Spells   map[string]string `json:"spells"`
	Detail   string            `json:"detail"`
}

// Mana é a estrutura com o conjunto de manas de um Card
type Mana struct {
	Any   int8 `json:"any"`
	Black int8 `json:"black"`
	Blue  int8 `json:"blue"`
	Green int8 `json:"green"`
	Red   int8 `json:"red"`
	White int8 `json:"white"`
}

// Image é a estrutura com os detalhes da imagem do Card
type Image struct {
	Path   string `json:"path"`
	Author string `json:"author"`
}

// NewCard retorna um novo card
func NewCard() *Card {
	return &Card{
		ID:       utils.GenerateULID(),
		ManaCost: new(Mana),
		Image:    new(Image),
		Spells:   make(map[string]string),
	}
}

/*
CreateCard cria um novo card no banco

@param newCard *Card - Card que deverá ser criado
@return id string - ID do card criado
@return err error - Erro que foi gerado ao tentar criar card
*/
func CreateCard(newCard *Card) (id string, err error) {
	return NewCard().ID, nil
}

/*
GetCard busca por um card

@param idCard string - ID do card que está buscando
@return card *Card - Objeto do card que foi buscado
@return err error - Erro que foi gerado ao tentar encontrar o card
*/
func GetCard(idCard string) (card *Card, err error) {
	return NewCard(), nil
}

/*
GetCards busca por todos os cards

@param idCard string - ID do card que está buscando
@return card *Card - Objeto do card que foi buscado
@return err error - Erro que foi gerado ao tentar encontrar os cards
*/
func GetCards() (cards []*Card, err error) {
	for i := 0; i < 10000; i++ {
		cards = append(cards, NewCard())
	}
	return
}

/*
UpdateCard busca por todos os cards

@param card *Card - Card que será atualizado
@return id string - ID do objeto do card que foi atualizado
@return err error - Erro que foi gerado ao tentar atualizar o card
*/
func UpdateCard(card *Card) (id string, err error) {
	return card.ID, nil
}

/*
DeleteCard busca por todos os cards

@param cardID string - ID do card que será apagado
@return id string - ID do objeto do card que foi apagado
@return err error - Erro que foi gerado ao tentar atualizar o card
*/
func DeleteCard(cardID string) (id string, err error) {
	return cardID, nil
}
