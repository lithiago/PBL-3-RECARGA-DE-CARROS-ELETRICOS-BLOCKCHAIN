# 🚗 🔗 Recarga de Carros Elétricos com Blockchain

## Componentes

- Lucas de Andrade Pereira Mendes
- Thiago Ramon Santos de Jesus
- Felipe Pinto Silva

**Professor:** Elinaldo Santos de Gois Junior

---

## 📑 Descrição

Este projeto implementa um sistema de recarga de carros elétricos totalmente integrado a uma blockchain, que atua como um livro-razão distribuído para registrar todas as transações de reserva, recarga e pagamento. Cada operação realizada entre usuários e empresas de recarga é registrada de forma imutável e transparente em um contrato inteligente, garantindo segurança, auditabilidade e confiança entre todas as partes envolvidas.

No funcionamento do sistema, sempre que um usuário realiza uma reserva de ponto de recarga, inicia uma recarga ou efetua um pagamento, essas ações são enviadas para um contrato inteligente desenvolvido em Solidity e implantado em uma blockchain privada. O contrato inteligente valida e armazena cada transação, tornando impossível a alteração ou exclusão dos registros após sua inclusão. Isso garante que todo o histórico de operações — desde a reserva até o pagamento — fique permanentemente disponível para consulta e auditoria.

A blockchain implementada elimina a necessidade de confiança em uma autoridade central, permitindo que múltiplas empresas de recarga e usuários interajam de forma descentralizada. Todas as partes podem verificar, a qualquer momento, o status e o histórico das transações, promovendo transparência e prevenindo fraudes ou disputas. Dessa forma, o sistema assegura a integridade dos dados e a confiabilidade das operações em todo o ecossistema de recarga.

---

## 🎯 Objetivo

O objetivo central deste projeto é aprimorar o sistema de recarga de carros elétricos desenvolvido anteriormente, integrando uma solução de blockchain para registrar, de maneira segura e transparente, todas as transações realizadas entre usuários e empresas de recarga. A blockchain atua como um livro-razão descentralizado, eliminando a necessidade de intermediários e prevenindo fraudes, disputas e inconsistências nos dados.

Com essa abordagem, o sistema passa a oferecer:

- **Transparência:** Todas as operações ficam publicamente registradas na blockchain, permitindo auditoria por qualquer parte interessada.
- **Segurança:** A imutabilidade dos registros impede alterações ou exclusões não autorizadas, protegendo contra fraudes.
- **Descentralização:** Nenhuma entidade central controla os dados, aumentando a confiabilidade e a resiliência do sistema.
- **Auditabilidade:** Todo o histórico de reservas, recargas e pagamentos pode ser facilmente verificado, promovendo confiança entre usuários e empresas.

A solução utiliza frameworks e ferramentas modernas para garantir a implementação eficiente e segura da blockchain, atendendo às restrições de não utilizar soluções centralizadas e de demonstrar, por meio de documentação e relatórios, a robustez e a segurança do sistema proposto.

---

## ⚙️ Estrutura dos Principais Diretórios

- `Carro/` — Código do cliente simulando o carro elétrico
- `Servidor/` — Código dos servidores dos postos de recarga
- `blockchain/` — Contrato inteligente, scripts de deploy e testes
- `utils/` — Utilitários, constantes, tópicos MQTT, rotas e dados de cidades

---

## 💻 Tecnologias e Recursos Utilizados

- **Go (Golang):** Backend dos serviços (Carro, Servidor, integração MQTT e Blockchain)
- **Solidity:** Contrato inteligente para registro das operações na blockchain
- **Foundry:** Ferramenta para desenvolvimento, teste e deploy de contratos Ethereum
- **Docker & Docker Compose:** Orquestração dos serviços (servidores, carros, broker MQTT, blockchain)
- **Mosquitto:** Broker MQTT para comunicação entre carros e servidores
- **Anvil:** Nó local Ethereum para testes e deploy do contrato inteligente
- **Gin:** Framework web para APIs REST dos servidores
- **MQTT:** Protocolo de comunicação leve entre carros e servidores
  
---

## 🚀 Como Executar o Projeto

### 1. Clonar o repositório

```sh
git clone <url-do-repositorio>
cd MQTT
```

### 2. Instalar o Foundry

O Foundry é necessário para compilar e implantar o contrato inteligente.

```sh
curl -L https://foundry.paradigm.xyz | bash
foundryup
```

> Consulte a [documentação oficial do Foundry](https://book.getfoundry.sh/) para mais detalhes.

### 3. Preparar o ambiente

Certifique-se de ter o Docker e o Docker Compose instalados.

### 4. Subir todos os serviços

```sh
docker-compose up --build -d
```

### 5. Implantar o contrato inteligente na blockchain local

```sh
make deploy-contract
```

### 6. Rodar o carro (interface interativa)

```sh
make run-car
```

### 7. Visualizar logs dos serviços

```sh
make logs
```

### 8. Finalizar e parar todos os serviços

```sh
make down
```

---

## 📌 Observações

- O contrato inteligente é implantado automaticamente na rede local Anvil.
- O sistema utiliza tópicos MQTT para comunicação entre carros e servidores.
- Todas as operações críticas (reserva, recarga, pagamento, cancelamento) são registradas na blockchain.
- Para modificar ou testar o contrato, utilize os comandos do Foundry em [`blockchain/README.md`](blockchain/README.md).

---

## 📚 Referências

- [Foundry Book](https://book.getfoundry.sh/)
- [Documentação do projeto blockchain](blockchain/README.md)
