# Makefile para o projeto de Recarga de Carros Elétricos

# Constrói as imagens para todos os serviços.
build:
	@echo "Construindo todas as imagens..."
	docker-compose build

# Sobe os serviços de backend (servidores, broker, blockchain).
up-servers:
	@echo "Iniciando os servidores, o broker MQTT e a blockchain Anvil..."
	docker-compose up -d mosquitto anvil servidor-feiradesantana servidor-ilheus servidor-salvador

# NOVO COMANDO: Implanta o contrato inteligente no contêiner Anvil que já está rodando.
deploy-contract:
	@echo "Implantando o contrato ReservationLedger na rede Anvil..."
	# CORREÇÃO: Adiciona a flag --workdir para garantir que o forge seja executado no contexto correto.
	docker-compose exec --workdir /app/blockchain anvil forge script /app/blockchain/script/Deploy.s.sol --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast

# Roda o carro de forma interativa.
run-car:
	@echo "Iniciando o contêiner do carro para interação..."
	docker-compose exec carro /car_app

# Mostra os logs dos serviços de backend.
logs:
	@echo "Mostrando logs dos servidores..."
	docker-compose logs -f mosquitto anvil servidor-feiradesantana servidor-ilheus servidor-salvador

# Para todos os contêineres e remove a rede.
down:
	@echo "Parando todos os contêineres..."
	docker-compose down

