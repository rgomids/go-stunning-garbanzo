# go-stunning-garbanzo
🌀 API com dupla implementação, sendo possível realizar chamadas via HTTP e Websocket.

## Configurar
Para configurar o projeto em sua maquina tenha a versão do `go` >= `go1.12`;
Utilize `go get` para abaixar todas as dependências;
Para inciar o server:
 - *Run*: Basta estar na pasta raiz do projeto e usar o comando: `make run`;
 - *Build*: Basta estar na pasta raiz do projeto e usar o comando: `make` ou `make build`;
 - *Build And Run*: Basta estar na pasta raiz do projeto e usar o comando: `make br`;
 
## Utilização
Recomendo a utilização do [Postman](https://www.getpostman.com/), para realizar as chamadas de testes. Existe inclusive um arquivo já configurado com todas as chamadas necessárias para esse projeto, você pode encontra-lo [aqui](https://github.com/RafaelGomides/go-stunning-garbanzo/blob/master/configurations/go-stunning-garbanzo.postman_collection.json).
Para realizar as chamadas websocket recomendo a utilização desse [site](https://www.websocket.org/echo.html). Minha ideia é criar uma interface bem simples e funcional para realizar essas chamadas, dentro do próprio projeto. Já existe um esboço disso em `views` sinta-se livre para mandar um PR com algo funcionando.
Para que a configuração do WS funcione, necessário estar com o servidor sendo executado; `make run` ou `make rb`.
Ao acessar o site coloque no input **Location**: ws://localhost:8080/ws; Em seguida, clique em **Connect**, quando fizer isso perceba se aparece **"CONNECTED"** no console ao lado direito do botão. Se aparecer, está tudo pronto.
Para fazer as requisições via Websocket utilize esse padrão:
```json
{
 "event": "",
 "data": ""
}
```
Trocando o `"event"` pelos eventos das rotas, que são:
```
ADD_CARD
GET_CARD
GET_CARDS
UPDATE_CARD
DELETE_CARD
```
Em `"data"`, você deve passar o objeto ou o ID dependendo da chamada sendo assim:
```
ADD_CARD: Object
GET_CARD: ID
GET_CARDS: null
UPDATE_CARD: Object
DELETE_CARD: ID
```
Uma requisição de criação de card seria:
```json
{
 "event": "ADD_CARD",
 "data": {
    "name": "",
    "mana_cost": {
        "any": 0,
        "black": 0,
        "blue": 0,
        "green": 0,
        "red": 0,
        "white": 0
    },
    "image": {
        "path": "",
        "author": ""
    },
    "type": "",
    "spells": {},
    "detail": ""
  }
}
```
Ou uma requisição para apagar um card:
```json
{
 "event": "DELETE_CARD",
 "data": "01DNF9X87YHG19QM8F2VN9YQZ2"
}
```

## Conceito
A ideia por trás desse projeto é implementar meus conhecimentos sobre go, e testar as conexões (WS e HTTP).
Vou fazer testes e elaborar alguns gráficos para demosntrar quais as vantagens de cada uma. Mas por enquanto seguimos com o projeto em forma de demonstração.
