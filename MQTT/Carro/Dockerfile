# Use a imagem base do Go
FROM golang:1.24

# Define o diretório de trabalho raiz da aplicação no contêiner
WORKDIR /app

# Copia os arquivos de gerenciamento de dependências primeiro para aproveitar o cache do Docker.
# Isso garante que as dependências só sejam baixadas novamente se o go.mod ou go.sum mudar.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia todo o resto do código-fonte para o diretório de trabalho.
# O contexto do build é a pasta 'MQTT/', então '.' copia tudo de lá para '/app' no contêiner.
COPY . .

# Define o diretório de trabalho para a pasta específica do carro
WORKDIR /app/Carro

# Compila a aplicação do carro.
# O Go vai encontrar as dependências em /app/utils/... porque todo o projeto foi copiado.
RUN go build -o /car_app .

# O comando para executar a aplicação quando o contêiner iniciar
CMD [ "/car_app" ]
