# 🔗 Luke Shortener

Um encurtador de URLs simples desenvolvido em **Go**, utilizando **Gin**, **SQLite**, **HTML**, **JavaScript** e **Tailwind CSS**.

## ✨ Funcionalidades

- 🔗 Encurta qualquer URL
- 💾 Armazena os links em SQLite
- ♻️ Reutiliza o mesmo link curto para URLs já cadastradas
- 🚀 Redirecionamento automático
- 🔒 Adiciona `https://` automaticamente quando necessário
- 🎲 Gera códigos aleatórios entre 5 e 10 caracteres

## 🛠️ Tecnologias

- Go
- Gin
- SQLite
- HTML
- JavaScript
- Tailwind CSS

## 📁 Estrutura

```
.
├── database/
│   └── database.go
├── handlers/
│   ├── home.go
│   ├── redirect.go
│   └── shorten.go
├── routes/
│   └── routes.go
├── utils/
│   └── random.go
├── database.db
├── index.html
├── main.go
└── go.mod
```

## 🚀 Como executar

### Clone o projeto

```bash
git clone https://github.com/lukex404/luke-shortener.git
cd luke-shortener
```

### Instale as dependências

```bash
go mod tidy
```

### Execute

```bash
go run .
```

O servidor ficará disponível em:

```
http://localhost:8080
```

## 📌 Como usar

Cole uma URL, por exemplo:

```
google.com
```

O sistema retornará algo como:

```
http://localhost:8080/aB92Kd
```

Ao acessar esse link, você será redirecionado automaticamente para a URL original.

## 📅 Melhorias futuras

- [ ] Links personalizados
- [ ] Contador de cliques
- [ ] Painel administrativo
- [ ] Exclusão de links
- [ ] Estatísticas por link
- [ ] API pública
- [ ] QR Code para links
- [ ] Docker
- [ ] Deploy automático

## 📄 Licença

Este projeto está licenciado sob a licença MIT.