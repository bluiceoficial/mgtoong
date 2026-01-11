# MGTOONG

**MGTOONG** é uma biblioteca em Go para trabalhar com TOON (Token-Oriented Object Notation), um formato simples e leve para armazenar coleções de registros com chave primária.

Projetada para ser direta e sem dependências externas, MGTOONG facilita criação, leitura, atualização e exclusão de registros em arquivos `.toon` ou em memória.

---

## ✨ Recursos

* ✅ CRUD simples (Create / Read / Update / Delete)
* 💾 Suporte a arquivo `.toon` com LoadFile / SaveFile
* 🔑 Chave primária por coleção
* 🗂️ Manipulação de coleções (tables) independentes
* 📦 Leve — usa apenas a biblioteca padrão do Go
* 🧪 API intuitiva para uso em CLIs, ferramentas e protótipos

---

## 📦 Instalação

```bash
go get github.com/mugomes/mgtoong
```

---

## 🚀 Exemplo de uso

```go
import (
    "fmt"
    "log"

    "github.com/mugomes/mgtoong"
)

func main() {
    toon := &mgtoong.MGTOONG{}

    // Criar uma coleção "users" com colunas e chave primária "id"
    toon.Create("users", []string{"id", "nome", "active"}, "id")

    // Carregar arquivo caso exista (mantém a chave primária)
    toon.LoadFile("users.toon", "id")

    // Adicionar registros
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

    // Ler todos os registros
    rows := toon.ReadAll("users")
    for _, row := range rows {
        fmt.Println(row["id"], row["nome"])
    }

    // Ler um registro pela chave primária
    valor := toon.ReadOne("users", "1")
    fmt.Println(valor["nome"])

    // Atualizar registro
    toon.Update("users", "1", map[string]string{
        "active": "1",
    })

    // Excluir registro
    toon.Delete("users", "2")

    // Representação em string no formato TOON
    fmt.Println(toon.ToString("users"))

    // Salvar em disco
    if err := toon.SaveFile("users.toon"); err != nil {
        log.Fatal(err)
    }
}
```

---

## 🧠 Como funciona

* Cada coleção tem colunas definidas e uma chave primária.
* Registros são armazenados como pares coluna→valor (strings).
* O arquivo `.toon` é uma serialização do estado interno; o projeto expõe métodos para carregar e salvar.
* O design prioriza simplicidade e previsibilidade, ideal para dados de configuração leve e protótipos.

---

## 🔍 Formato de exemplo (.toon)

Exemplo simplificado de saída gerada por `ToString("users")`:

```
TABLE users PK id
COLUMNS id nome active
ROW 1 Ana 0
ROW 2 Maria 0
END
```

---

## 👤 Autor

**Murilo Gomes Julio**

🔗 https://mugomes.github.io

📺 https://youtube.com/@mugomesoficial

---

## License

Copyright (c) 2025-2026 Murilo Gomes Julio

Licensed under the [MIT](https://github.com/mugomes/mgtoong/blob/main/LICENSE) license.

All contributions to MGTOONG are subject to this license.
