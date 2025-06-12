package main

import (
	api "MQTT/Servidor/API"
	"MQTT/Servidor/blockchain"
	consts "MQTT/utils/Constantes"
	rotaslib "MQTT/utils/Rotas"
	topics "MQTT/utils/Topicos"
	clientemqtt "MQTT/utils/mqttLib/ClienteMQTT"
	router "MQTT/utils/mqttLib/Router"
	storage "MQTT/utils/storage"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Servidor struct {
	IP                    string
	ID                    string
	Cidade                string
	Client                clientemqtt.MQTTClient
	carrosConectados      map[string]*ConnectedCarStatus
	carrosConectadosMutex sync.Mutex
}

// ConnectedCarStatus armazena os IDs das reservas na blockchain.
type ConnectedCarStatus struct {
	LastActivity     time.Time
	CommittedReserva *consts.Reserva
	Participantes2PC []consts.Participante2PC
	ReservationIDs   map[string][32]byte // Mapeia postoID para o ID da reserva na blockchain
}

var (
	arquivoPontos = os.Getenv("ARQUIVO_JSON")
)

var cidadeConfig = map[string]struct {
	Container string
	Porta     string
}{
	"FSA": {"feiradesantana", "8080"},
	"ILH": {"ilheus", "8081"},
	"SSA": {"salvador", "8082"},
}

func (s *Servidor) AssinarEventosDoCarro() {
	topicsToSubscribe := []string{
		topics.CarroRequestReserva("+", s.IP, s.Cidade),
		topics.CarroRequestCancel("+"),
		topics.CarroRequestRotas("+", s.Cidade),
		topics.CarroDesconectado("+"),
		topics.CarroSendsRechargeFinish("+"),
	}
	for _, topic := range topicsToSubscribe {
		s.Client.Subscribe(topic)
	}
}

func inicializarServidor() Servidor {
	ip, _ := consts.GetLocalIP()
	routerServidor := router.NewRouter()
	mqttClient := *clientemqtt.NewClient(consts.Broker, routerServidor, topics.CarroDesconectado(ip), ip)
	token := mqttClient.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalf("Erro ao conectar ao broker: %v", token.Error())
	}
	return Servidor{
		IP:               ip,
		Client:           mqttClient,
		Cidade:           os.Getenv("CIDADE"),
		carrosConectados: make(map[string]*ConnectedCarStatus),
	}
}

func (S *Servidor) regitrarHandlersMQTT() {
	routerServidor := S.Client.Router

	routerServidor.Register(topics.CarroRequestReserva("+", S.IP, S.Cidade), func(payload []byte) {
		var reserva consts.Reserva
		if err := json.Unmarshal(payload, &reserva); err != nil {
			log.Println("Erro ao decodificar mensagem de reserva:", err)
			return
		}

		serverURLs := make(map[string]string)
		for cidade, configs := range cidadeConfig {
			serverURLs[cidade] = fmt.Sprintf("http://servidor-%s:%s", configs.Container, configs.Porta)
		}

		var participantes []consts.Participante2PC
		reservationIDs := make(map[string][32]byte)
		for _, parada := range reserva.Paradas {
			if serverURL, ok := serverURLs[parada.Cidade]; ok {
				participantes = append(participantes, consts.Participante2PC{URL: serverURL, PostoID: parada.IDPosto})
				reservationIDs[parada.IDPosto] = blockchain.GenerateReservationID()
			}
		}

		S.carrosConectadosMutex.Lock()
		S.carrosConectados[reserva.Carro.ID] = &ConnectedCarStatus{
			CommittedReserva: &reserva,
			Participantes2PC: participantes,
			ReservationIDs:   reservationIDs,
			LastActivity:     time.Now(),
		}
		S.carrosConectadosMutex.Unlock()

		topic := topics.ServerReserveStatus(S.IP, reserva.Carro.ID)
		if err := api.TwoPhaseCommit(participantes, reserva.Carro); err != nil {
			log.Printf("[ERRO] Two-Phase Commit falhou: %v\n", err)
			S.Client.Publish(topic, []byte(`{"status":"ERRO"}`))
			S.carrosConectadosMutex.Lock()
			delete(S.carrosConectados, reserva.Carro.ID)
			S.carrosConectadosMutex.Unlock()
		} else {
			log.Println("[INFO] 2PC concluído. Registrando na blockchain...")
			S.Client.Publish(topic, []byte(`{"status":"OK"}`))

			for postoID, reservationID := range reservationIDs {
				blockchain.RecordReservationOnChain(reservationID, reserva.Carro.ID, postoID)
			}
		}
	})

	routerServidor.Register(topics.CarroRequestRotas("+", S.Cidade), func(payload []byte) {
		var conteudoMsg consts.Trajeto
		json.Unmarshal(payload, &conteudoMsg)
		dadosRotas := storage.LerRotas()
		rotasValidas := rotaslib.GetRotasValidas(dadosRotas.Rotas, conteudoMsg)
		mapaCompleto := make(map[string][]consts.Posto)
		paradas := make(map[string][]consts.Parada)
		for nome, rota := range rotasValidas {
			for _, cidade := range rota {
				if _, ok := mapaCompleto[cidade]; ok {
					continue
				}
				if cidade == S.Cidade {
					mapaCompleto[cidade] = storage.CarregarPostos()
				} else {
					config, _ := cidadeConfig[cidade]
					url := "http://servidor-" + config.Container + ":" + config.Porta
					postos, _ := api.ObterPostosDeOutroServidor(url)
					var postosSemPonteiro []consts.Posto
					for _, posto := range postos {
						postosSemPonteiro = append(postosSemPonteiro, *posto)
					}
					mapaCompleto[cidade] = postosSemPonteiro
				}
			}
			paradas[nome] = rotaslib.GerarRotas(conteudoMsg.CarroMQTT, rota, dadosRotas.Cidades, mapaCompleto)
		}
		msg, _ := json.Marshal(consts.Mensagem{ID: S.IP, Origem: S.Cidade, Conteudo: map[string]interface{}{"paradas": paradas}})
		topic := topics.ServerResponteRoutes(conteudoMsg.CarroMQTT.ID, S.Cidade)
		S.Client.Publish(topic, msg)
	})

	routerServidor.Register(topics.CarroDesconectado("+"), func(payload []byte) {
		var msg map[string]string
		json.Unmarshal(payload, &msg)
		if id, ok := msg["ID"]; ok {
			log.Printf("[LWT] Carro %s desconectado inesperadamente.\n", id)
			S.processCarroDisconnected(id, true)
		}
	})

	// ALTERADO: O nome do handler foi melhorado para clareza
	handlerFinishOrCancel := func(payload []byte) {
		var msg map[string]string
		json.Unmarshal(payload, &msg)
		carroID := msg["IDCarro"]
		if carroID == "" {
			return
		}

		S.carrosConectadosMutex.Lock()
		carStatus, exists := S.carrosConectados[carroID]
		if !exists {
			S.carrosConectadosMutex.Unlock()
			return
		}
		S.carrosConectadosMutex.Unlock()

		// ALTERADO: A lógica agora está mais clara e completa.
		if msg["Msg"] == "Finalizar Recarga" {
			log.Printf("[MANUAL] Carro %s finalizou recarga. Registrando conclusão e pagamento na blockchain...\n", carroID)

			// Itera sobre todas as reservas feitas para este carro.
			for _, reservationID := range carStatus.ReservationIDs {
				// Cria uma goroutine para cada reserva para não bloquear o fluxo principal.
				go func(id [32]byte) {
					// 1. Marca a recarga como completa na blockchain.
					err := blockchain.RecordRechargeCompletedOnChain(id)
					if err != nil {
						log.Printf("Erro ao completar recarga para o ID %x: %v", id, err)
						return // Se falhar, não tenta pagar.
					}

					// 2. Simula o cálculo do pagamento e o registra na blockchain.
					// Para uma simulação, vamos usar um valor fixo (ex: 10.5 tokens).
					// O valor é em Wei (1 Ether = 10^18 Wei), então 10.5 Ether seria 105 * 10^17
					amount := new(big.Int)
					amount.SetString("10500000000000000000", 10) // Exemplo: 10.5

					err = blockchain.RecordPaymentOnChain(id, amount)
					if err != nil {
						log.Printf("Erro ao processar pagamento para o ID %x: %v", id, err)
					}
				}(reservationID)
			}
			// Após processar, limpa o estado do carro.
			S.processCarroDisconnected(carroID, false) // false = não cancelar na blockchain

		} else { // Se for qualquer outra coisa (cancelamento manual)
			log.Printf("[MANUAL] Carro %s cancelou. Iniciando limpeza...\n", carroID)
			S.processCarroDisconnected(carroID, true) // true = cancelar na blockchain
		}
	}
	routerServidor.Register(topics.CarroRequestCancel("+"), handlerFinishOrCancel)
	routerServidor.Register(topics.CarroSendsRechargeFinish("+"), handlerFinishOrCancel)
}

func (s *Servidor) processCarroDisconnected(carroID string, shouldCancelOnBlockchain bool) {
	s.carrosConectadosMutex.Lock()
	carStatus, exists := s.carrosConectados[carroID]
	if !exists {
		s.carrosConectadosMutex.Unlock()
		return
	}
	delete(s.carrosConectados, carroID)
	s.carrosConectadosMutex.Unlock()

	if carStatus.CommittedReserva != nil {
		for _, p := range carStatus.Participantes2PC {
			releasePayload, _ := json.Marshal(map[string]interface{}{"posto_id": p.PostoID, "carro": carStatus.CommittedReserva.Carro})
			http.Post(p.URL+"/2pc/release", "application/json", strings.NewReader(string(releasePayload)))
		}
	}

	if shouldCancelOnBlockchain && carStatus.ReservationIDs != nil {
		log.Println("Registrando cancelamentos na blockchain para o carro", carroID)
		for _, reservationID := range carStatus.ReservationIDs {
			go blockchain.RecordReservationCancelledOnChain(reservationID)
		}
	}
}

func main() {
	log.Println("[SERVIDOR] Inicializando...")
	blockchain.InitBlockchainService()
	server := inicializarServidor()
	log.Println("[SERVIDOR] IP:", server.IP)
	server.regitrarHandlersMQTT()
	server.AssinarEventosDoCarro()
	go api.ServerAPICommunication(arquivoPontos)
	log.Println("[SERVIDOR] Iniciando comunicação MQTT...")
	select {}
}
