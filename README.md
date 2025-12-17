# MGTOONG

Versão em Go (Golang) do MGTOON.

MGTOONG é uma biblioteca para trabalhar com Token-Oriented Object Notation no Golang.

## Installation

`go get github.com/mugomes/mgtoong`

## Example

```
toon := &mgtoong.MGTOONG{}

// Criando Usuários
toon.Create("users", []string{"id", "nome", "active"}, "id")

// Abrindo arquivo toon caso exista
toon.LoadFile("users.toon", "id")

// Adicionando registros
toon.Add(map[string]string{
	"id":     "1",
	"nome":   "Ana",
	"active": "0",
})
toon.Add(map[string]string{
	"id":     "2",
	"nome":   "Maria",
	"active": "0",
})
toon.Add(map[string]string{
	"id":     "3",
	"nome":   "João",
	"active": "0",
})

// Lendo todos registros
rows := toon.ReadAll()

for _, row := range rows {
	fmt.Println(row["id"], row["nome"])
}

// Lendo os valores das colunas da ID 1 (chave primária)
valor := toon.ReadOne("1")
fmt.Println(valor["nome"])

// Atualizando registro
toon.Update("1", map[string]string{
	"active": "1",
})

// Excluindo registro
toon.Delete("2")

// Formato TOON
fmt.Println(toon.ToString())

// Salvar arquivo em formato toon
toon.SaveFile("users.toon")
```

## Support

- GitHub Sponsors: https://github.com/sponsors/mugomes/
- More Options: https://www.mugomes.com.br/p/apoie.html

## License

The MGTOONG is provided under:

[SPDX-License-Identifier: MIT](https://github.com/mugomes/mgtoong/blob/main/LICENSE)

All contributions to the MGTOONG are subject to this license.
