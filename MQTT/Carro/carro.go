package main

import (
	consts "MQTT/utils/Constantes"
	topics "MQTT/utils/Topicos"
	clientemqtt "MQTT/utils/mqttLib/ClienteMQTT"
	router "MQTT/utils/mqttLib/Router"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// MqttMessage representa uma mensagem MQTT recebida para ser enviada pelo canal
type MqttMessage struct {
	Topic   string
	Payload []byte
}

// Canais globais para comunicação entre goroutines
var (
	incomingMqttChan = make(chan MqttMessage, 100) // Canal para mensagens MQTT recebidas, com buffer
	userInputChan    = make(chan string)           // Canal para entrada do usuário
	quitChan         = make(chan os.Signal, 1)     // Canal para sinal de encerramento
	promptChan       = make(chan Prompt)
)

type Prompt struct {
	Pergunta   string
	RespostaCh chan string
}

type Carro struct {
	ID                string                 `json:"id"`
	Bateria           float64                `json:"bateria"`
	Clientemqtt       clientemqtt.MQTTClient `json:"-"`
	X                 float64                `json:"x"`
	Y                 float64                `json:"y"`
	CapacidadeBateria float64                `json:"capacidadebateria"`
	Consumobateria    float64                `json:"consumobateria"`
	CidadeAtual       string                 `json:"cidadeatual"`
}

func (c *Carro) SolicitarReserva(rotas map[string][]consts.Parada, cidadeDestino string, serverID string) {

	rotasIndexadas := []string{}
	for nome, paradas := range rotas {
		fmt.Printf("\n[%d] %s:\n", len(rotasIndexadas), nome)
		for i, parada := range paradas {
			fmt.Printf("  \t [%d] %s (ID: %s)\n", i+1, parada.NomePosto, parada.IDPosto)
			fmt.Printf("      \t Localização: (X: %.2f, Y: %.2f)\n", parada.X, parada.Y)
		}
		rotasIndexadas = append(rotasIndexadas, nome)
	}

	input := perguntarUsuario("Digite o número da rota desejada: ")
	escolha, _ := strconv.Atoi(input)

	if escolha < 0 || escolha >= len(rotasIndexadas) {
		fmt.Println("❌ Escolha inválida.")
		return
	}

	nomeRotaEscolhida := rotasIndexadas[escolha]
	fmt.Printf("Você escolheu a rota: %s\n", nomeRotaEscolhida)
	paradasEscolhidas := rotas[nomeRotaEscolhida]

	reserva := consts.Reserva{
		Carro: consts.Carro{
			ID:                c.ID,
			Bateria:           c.Bateria,
			X:                 c.X,
			Y:                 c.Y,
			CapacidadeBateria: c.CapacidadeBateria,
			Consumobateria:    c.Consumobateria,
		},
		Paradas: paradasEscolhidas,
	}

	ConteudoJSON, err := json.Marshal(reserva)
	if err != nil {
		log.Printf("[ERRO] Falha ao serializar mensagem de reserva: %v\n", err)
		return
	}

	topic := topics.CarroRequestReserva(c.ID, serverID, cidadeDestino)
	log.Println("[CARRO] Publicando solicitação de reserva no tópico: ", topic)
	c.publicarAoServidor(ConteudoJSON, topic)
}

func (c *Carro) CancelarReserva() {
	topic := topics.CarroRequestCancel(c.ID)
	log.Println("[CARRO] Publicando cancelamento de reserva no tópico: ", topic)
	msg := map[string]string{
		"IDCarro": c.ID,
		"Msg":     "Cancelar Reserva",
	}
	msgJSON, _ := json.Marshal(msg)
	c.Clientemqtt.Publish(topic, msgJSON)
}

func (c *Carro) FinalizarRecarga() {
	topic := topics.CarroSendsRechargeFinish(c.ID)
	log.Println("[CARRO] Publicando finalização de recarga no tópico: ", topic)
	msg := map[string]string{
		"IDCarro": c.ID,
		"Msg":     "Finalizar Recarga",
	}
	msgJSON, _ := json.Marshal(msg)
	c.Clientemqtt.Publish(topic, msgJSON)
}

func (c *Carro) publicarAoServidor(conteudoJSON []byte, topico string) {
	if conteudoJSON == nil {
		log.Println("[CARRO] Não foi possível publicar: conteúdo JSON é nulo.")
		return
	}
	log.Printf("[CARRO] Publicando no tópico: %s com payload: %s\n", topico, string(conteudoJSON))
	c.Clientemqtt.Publish(topico, conteudoJSON)
}

func (c *Carro) solicitarRota(cidadeInicial string, cidadeDestino string) {
	log.Println("[CARRO] Função solicitarRota foi chamada")
	topic := topics.CarroRequestRotas(c.ID, cidadeDestino)
	log.Printf("[CARRO] Topico para solicitação de rota: %s", topic)

	trajeto := consts.Trajeto{
		CarroMQTT: consts.Carro{
			ID:                c.ID,
			Bateria:           c.Bateria,
			X:                 c.X,
			Y:                 c.Y,
			CapacidadeBateria: c.CapacidadeBateria,
			Consumobateria:    c.Consumobateria,
		},
		Inicio:  cidadeInicial,
		Destino: cidadeDestino,
	}
	ConteudoJSON, err := json.Marshal(trajeto)
	if err != nil {
		log.Printf("[ERRO] Falha ao serializar trajeto para rota: %v\n", err)
		return
	}
	c.publicarAoServidor(ConteudoJSON, topic)
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func (c *Carro) PorcentagemBateria() float64 {
	return (c.Bateria / c.CapacidadeBateria) * 100
}

func desserializarMensagem(mensagem []byte) consts.Mensagem {
	var msg consts.Mensagem
	if err := json.Unmarshal(mensagem, &msg); err != nil {
		fmt.Printf("[ERRO] Erro ao decodificar mensagem: %v\n", err)
		return consts.Mensagem{}
	}
	return msg
}

func (c *Carro) exibirMenu() {
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("            🚀 MENU PRINCIPAL 🚀")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("  🆔 Carro ID: %s \n", c.ID)
	fmt.Printf("  🔋 Bateria: %.2f%%\n", c.PorcentagemBateria())
	fmt.Println("  1️⃣  | Solicitar Nova Rota")
	fmt.Println("  2️⃣  | Cancelar Rota Atual")
	fmt.Println("  3️⃣  | Finalizar Recarga")
	fmt.Println("  4️⃣  | Encerrar Conexão")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func setupMqttHandlers(router *router.Router, carID string) {
	router.Register(topics.ServerResponseToCar(carID), func(payload []byte) {
		incomingMqttChan <- MqttMessage{Topic: topics.ServerResponseToCar(carID), Payload: payload}
	})

	router.Register(topics.ServerResponteRoutes(carID, "+"), func(payload []byte) {
		incomingMqttChan <- MqttMessage{Topic: topics.ServerResponteRoutes(carID, "+"), Payload: payload}
	})

	router.Register(topics.ServerReserveStatus("+", carID), func(payload []byte) {
		incomingMqttChan <- MqttMessage{Topic: topics.ServerReserveStatus("+", carID), Payload: payload}
	})
}

func processIncomingMqttMessages(car *Carro) {
	for msg := range incomingMqttChan {
		log.Printf("[Processador MQTT] Recebeu mensagem do tópico: %s\n", msg.Topic)

		if strings.HasPrefix(msg.Topic, topics.ServerResponseToCar(car.ID)) {
			fmt.Printf(">> [Resposta Servidor] %s\n", string(msg.Payload))

		} else if strings.HasPrefix(msg.Topic, topics.ServerResponteRoutes(car.ID, "")) {
			// CORREÇÃO APLICADA AQUI
			msgServer := desserializarMensagem(msg.Payload)
			fmt.Println(">> Rotas Recebidas do IP:", msgServer.ID)

			paradasMap := make(map[string][]consts.Parada)

			// 1. Acessa o objeto 'paradas' que está aninhado dentro do conteúdo.
			rotasPayload, ok := msgServer.Conteudo["paradas"].(map[string]interface{})
			if !ok {
				log.Println("Erro: O payload de rotas recebido nao tem o formato esperado.")
				continue
			}

			// 2. Agora, itera sobre o mapa de rotas real (ex: "Rota1", "Rota2").
			for nomeRota, listaParadasInterface := range rotasPayload {
				// Converte a lista de paradas (que é um []interface{}) para JSON bytes.
				bytes, err := json.Marshal(listaParadasInterface)
				if err != nil {
					log.Printf("Erro ao serializar paradas para a rota '%s': %v\n", nomeRota, err)
					continue
				}

				// Converte os JSON bytes de volta para o tipo correto []consts.Parada.
				var paradas []consts.Parada
				if err := json.Unmarshal(bytes, &paradas); err != nil {
					log.Printf("Erro ao desserializar paradas para a rota '%s': %v\n", nomeRota, err)
					continue
				}
				paradasMap[nomeRota] = paradas
			}

			if len(paradasMap) > 0 {
				car.SolicitarReserva(paradasMap, msgServer.Origem, msgServer.ID)
			} else {
				fmt.Println(">> Nenhuma rota viavel encontrada. Por favor, tente outro destino.")
			}

		} else if strings.HasPrefix(msg.Topic, topics.ServerReserveStatus("+", car.ID)) {
			msgServer := desserializarMensagem(msg.Payload)
			fmt.Println(">> Status de Reserva recebido do IP:", msgServer.ID)

			var status struct {
				Status string `json:"status"`
			}
			json.Unmarshal(msg.Payload, &status)

			if status.Status == "OK" {
				fmt.Println("✅ Reserva confirmada com sucesso em todos os postos!")
			} else if status.Status == "ERRO" {
				fmt.Println("❌ Falha ao realizar a reserva. Um dos postos pode estar ocupado. Tente outra rota.")
			}
		} else {
			log.Printf("[Processador MQTT] Tópico desconhecido ou não tratado: %s\n", msg.Topic)
		}
	}
}

func (c *Carro) AssinarRespostaServidor() {
	c.Clientemqtt.Subscribe(topics.ServerResponseToCar(c.ID))
	c.Clientemqtt.Subscribe(topics.ServerResponteRoutes(c.ID, "+"))
	c.Clientemqtt.Subscribe(topics.ServerReserveStatus("+", c.ID))
}

func (c *Carro) selecionarCidade() string {
	var cidades []string
	for _, cidade := range consts.CidadesArray {
		if cidade != c.CidadeAtual {
			cidades = append(cidades, cidade)
		}
	}

	fmt.Println("Cidades disponíveis para rota:")
	for i, cidade := range cidades {
		fmt.Printf("  %d - %s\n", i+1, cidade)
	}

	input := perguntarUsuario("Digite a opção para cidade de destino: ")
	escolha, err := strconv.Atoi(input)
	if err != nil || escolha < 1 || escolha > len(cidades) {
		fmt.Println("Opção inválida. Tente novamente.")
		return ""
	}
	return cidades[escolha-1]
}

func readUserInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case prompt := <-promptChan:
			fmt.Print(prompt.Pergunta)
			input, err := reader.ReadString('\n')
			if err != nil {
				prompt.RespostaCh <- ""
			} else {
				prompt.RespostaCh <- strings.TrimSpace(input)
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func perguntarUsuario(pergunta string) string {
	respCh := make(chan string)
	promptChan <- Prompt{
		Pergunta:   pergunta,
		RespostaCh: respCh,
	}
	return <-respCh
}

func main() {
	log.Println("[CARRO] Inicializando aplicação...")
	ip, _ := getLocalIP()
	routerCarro := router.NewRouter()
	mqttClient := *clientemqtt.NewClient(consts.Broker, routerCarro, topics.CarroDesconectado(ip), ip)

	conn := mqttClient.Connect()
	if conn.Wait() && conn.Error() != nil {
		log.Fatalf("[CARRO] Erro ao conectar ao broker: %v", conn.Error())
	}
	log.Println("[CARRO] Conectado ao broker MQTT.")

	rand.Seed(time.Now().UnixNano())
	randomX := rand.Float64()*(355.0-60.0) + 60.0
	randomY := rand.Float64()*(270.0-50.0) + 50.0
	cidadeInicial := consts.CidadeAtualDoCarro(randomX, randomY)
	log.Printf("Cidade inicial [%s]: (%.2f, %.2f)\n", cidadeInicial, randomX, randomY)

	carro := Carro{
		ID:                ip,
		Bateria:           60.0,
		Clientemqtt:       mqttClient,
		X:                 randomX,
		Y:                 randomY,
		CapacidadeBateria: 60.0,
		Consumobateria:    0.20,
		CidadeAtual:       cidadeInicial,
	}

	carro.AssinarRespostaServidor()
	setupMqttHandlers(routerCarro, carro.ID)

	go processIncomingMqttMessages(&carro)
	go readUserInput()

	// Loop principal da aplicação
	for {
		carro.exibirMenu()
		opcao := perguntarUsuario("Digite a opção desejada: ")
		switch opcao {
		case "1":
			cidadeDestino := carro.selecionarCidade()
			if cidadeDestino != "" {
				carro.solicitarRota(carro.CidadeAtual, cidadeDestino)
			}
		case "2":
			carro.CancelarReserva()
			fmt.Println("[CARRO] Solicitação de cancelamento enviada.")
		case "3":
			carro.FinalizarRecarga()
			fmt.Println("[CARRO] Informação de recarga finalizada enviada.")
		case "4":
			fmt.Println("Encerrando conexão...")
			return // Sai do loop e encerra a função main
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
