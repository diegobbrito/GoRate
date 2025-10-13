# GoRate – Real-time Currency & Crypto Quotes API

GoRate é uma API escrita em **Go** que busca cotações de moedas e criptomoedas em **múltiplas fontes simultaneamente**, utilizando **concorrência com goroutines** e **channels**. O objetivo é demonstrar performance, clean architecture e boas práticas em Go.

---

## 🚀 Funcionalidades

- Consulta de cotações em **fontes múltiplas** (simultaneamente)
- **Concorrência** com goroutines e contextos com timeout
- Retorno com média e número de fontes válidas
- **API REST** simples e eficiente
- Pronto para deploy com Docker

---

## 📦 Estrutura do Projeto

```
go-rate/
├── cmd/
│   └── server/main.go         # ponto de entrada do servidor
├── internal/
│   ├── api/                   # controladores HTTP
│   ├── domain/                # modelos e entidades
│   ├── service/               # lógica de negócio
│   └── infra/                 # integrações externas
├── pkg/logger/                # logging centralizado
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

---

## 🧠 Endpoints

### `GET /api/v1/rates?symbol=usd`

**Resposta:**
```json
{
  "symbol": "usd",
  "average": 5.32,
  "sources": 2
}
```

---

## 🧰 Executando localmente

### 1️⃣ Clonar o repositório
```bash
git clone https://github.com/diegobrito/GoRate.git
cd go-rate
```

### 2️⃣ Rodar o projeto
```bash
go run main.go
```

Acesse em: [http://localhost:8080/api/v1/rates?symbol=usd](http://localhost:8080/api/v1/rates?symbol=usd)

---

## 🐳 Docker

### Build da imagem
```bash
docker build -t go-rate .
```

### Executar o container
```bash
docker run -p 8080:8080 go-rate
```

---

## 🧩 Tecnologias

- **Go 1.23+**
- **HTTP standard library**
- **Goroutines & channels**
- **Context com timeout**
- **Docker**

---

## 📜 Licença

MIT License © 2025 [Diego Brito](https://github.com/diegobrito)

---
