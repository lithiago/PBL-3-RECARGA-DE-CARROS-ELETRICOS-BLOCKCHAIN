package api

import (
	consts "MQTT/utils/Constantes"
	storage "MQTT/utils/storage"
	"encoding/hex" // Pacote para converter bytes em string hexadecimal
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"MQTT/Servidor/blockchain"
	"MQTT/Servidor/contracts" // Importa para usar a struct do contrato

	"github.com/gin-gonic/gin"
)

// ReservationResponse é uma struct limpa para a resposta da API,
// facilitando a conversão dos dados da blockchain para JSON.
type ReservationResponse struct {
	ExternalId   string `json:"externalId"`
	CarId        string `json:"carId"`
	PostoId      string `json:"postoId"`
	Timestamp    string `json:"timestamp"`
	ChargeAmount string `json:"chargeAmount"`
	Status       string `json:"status"`
}

var postosMutex sync.Mutex

func ServerAPICommunication(arquivoPontos string) {

	r := gin.Default()

	// ... (TODAS AS SUAS ROTAS EXISTENTES: /postos, /2pc/prepare, etc. FICAM AQUI SEM ALTERAÇÕES) ...
	r.GET("/postos", func(c *gin.Context) {
		postos, err := storage.GetPostosFromJSON(arquivoPontos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(postos) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Nenhum posto encontrado"})
			return
		}
		c.JSON(http.StatusOK, postos)
	})
	r.GET("/postos/disponiveis", func(c *gin.Context) {
		postos, err := storage.GetPostosDisponiveis(arquivoPontos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(postos) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "nenhum posto disponível encontrado"})
			return
		}
		c.JSON(http.StatusOK, postos)
	})
	r.PATCH("/postos/:id/adicionar", func(c *gin.Context) {
		id := c.Param("id")
		var carro consts.Carro
		if err := c.ShouldBindJSON(&carro); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}
		postos, err := storage.GetPostosFromJSON(arquivoPontos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var postoAtualizado *consts.Posto
		for _, p := range postos {
			if p.Id == id {
				if len(p.Fila) > 0 {
					c.JSON(http.StatusConflict, gin.H{"error": "Já existe um carro na fila"})
					return
				}
				p.Fila = append(p.Fila, carro)
				postoAtualizado = p
				break
			}
		}
		if postoAtualizado == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Posto não encontrado"})
			return
		}
		if err := storage.AtualizarArquivo(arquivoPontos, postos); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o arquivo JSON"})
			return
		}
		c.JSON(http.StatusOK, postoAtualizado)
	})
	r.PATCH("/postos/:id/remover", func(c *gin.Context) {
		id := c.Param("id")
		var carro consts.Carro
		if err := c.ShouldBindJSON(&carro); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}
		postos, err := storage.GetPostosFromJSON(arquivoPontos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var postoAtualizado *consts.Posto
		for _, p := range postos {
			if p.Id == id {
				for i, c := range p.Fila {
					if c.ID == carro.ID {
						p.Fila = append(p.Fila[:i], p.Fila[i+1:]...)
						postoAtualizado = p
						break
					}
				}
				break
			}
		}
		if postoAtualizado == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Posto ou carro não encontrado"})
			return
		}
		if err := storage.AtualizarArquivo(arquivoPontos, postos); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o arquivo JSON"})
			return
		}
		c.JSON(http.StatusOK, postoAtualizado)
	})
	r.POST("/2pc/prepare", func(c *gin.Context) {
		var req struct {
			PostoID string       `json:"posto_id"`
			Carro   consts.Carro `json:"carro"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "abort", "error": "Dados inválidos"})
			return
		}
		postosMutex.Lock()
		defer postosMutex.Unlock()
		postos, err := storage.GetPostosFromJSON(arquivoPontos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": "abort", "error": "Erro ao ler postos"})
			return
		}
		for _, p := range postos {
			if p.Id == req.PostoID {
				if len(p.Fila) > 0 || p.Pendente != nil {
					c.JSON(http.StatusOK, gin.H{"result": "abort"})
					return
				}
				p.Pendente = &req.Carro
				storage.AtualizarArquivo(arquivoPontos, postos)
				c.JSON(http.StatusOK, gin.H{"result": "ok"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"result": "abort", "error": "Posto não encontrado"})
	})
	r.POST("/2pc/commit", func(c *gin.Context) {
		var req struct {
			PostoID string       `json:"posto_id"`
			Carro   consts.Carro `json:"carro"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "abort", "error": "Dados inválidos"})
			return
		}
		postosMutex.Lock()
		defer postosMutex.Unlock()
		postos, err := storage.GetPostosFromJSON(arquivoPontos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": "abort", "error": "Erro ao ler postos"})
			return
		}
		var postoAtualizado *consts.Posto
		for _, p := range postos {
			if p.Id == req.PostoID {
				if p.Pendente != nil && p.Pendente.ID == req.Carro.ID {
					p.Fila = append(p.Fila, req.Carro)
					p.Pendente = nil
					postoAtualizado = p
				}
				break
			}
		}
		if postoAtualizado == nil {
			c.JSON(http.StatusNotFound, gin.H{"result": "abort", "error": "Posto não encontrado"})
			return
		}
		if err := storage.AtualizarArquivo(arquivoPontos, postos); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": "abort", "error": "Erro ao atualizar o arquivo JSON"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "committed"})
	})
	r.POST("/2pc/abort", func(c *gin.Context) {
		var req struct {
			PostoID string       `json:"posto_id"`
			Carro   consts.Carro `json:"carro"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "abort", "error": "Dados inválidos"})
			return
		}
		postosMutex.Lock()
		defer postosMutex.Unlock()
		postos, err := storage.GetPostosFromJSON(arquivoPontos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": "abort", "error": "Erro ao ler postos"})
			return
		}
		for _, p := range postos {
			if p.Id == req.PostoID && p.Pendente != nil && p.Pendente.ID == req.Carro.ID {
				p.Pendente = nil
				storage.AtualizarArquivo(arquivoPontos, postos)
				break
			}
		}
		c.JSON(http.StatusOK, gin.H{"result": "aborted"})
	})
	r.POST("/2pc/release", func(c *gin.Context) {
		var req struct {
			PostoID string       `json:"posto_id"`
			Carro   consts.Carro `json:"carro"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Dados inválidos"})
			return
		}
		postos, err := storage.GetPostosFromJSON(arquivoPontos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Erro ao ler postos"})
			return
		}
		postoEncontrado := false
		for _, p := range postos {
			if p.Id == req.PostoID {
				newFila := []consts.Carro{}
				removed := false
				for _, fCarro := range p.Fila {
					if fCarro.ID == req.Carro.ID {
						removed = true
					} else {
						newFila = append(newFila, fCarro)
					}
				}
				p.Fila = newFila
				postoEncontrado = true
				if removed {
					log.Printf("[API - RELEASE] Carro %s removido do posto %s.\n", req.Carro.ID, req.PostoID)
				} else {
					log.Printf("[API - RELEASE] Carro %s nao encontrado no posto %s.\n", req.Carro.ID, req.PostoID)
				}
				break
			}
		}
		if !postoEncontrado {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Posto nao encontrado"})
			return
		}
		if err := storage.AtualizarArquivo(arquivoPontos, postos); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Erro ao atualizar o arquivo JSON"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Posto liberado."})
	})

	// ALTERADO: A rota de histórico agora está totalmente implementada.
	r.GET("/history/:carId", func(c *gin.Context) {
		targetCarId := c.Param("carId")

		// 1. Chama a função do pacote blockchain para buscar todo o histórico.
		fullHistory, err := blockchain.GetFullHistoryOnChain()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar historico da blockchain", "details": err.Error()})
			return
		}

		var carHistory []ReservationResponse
		// 2. Itera sobre o histórico e filtra pelo carId desejado.
		for _, reservation := range fullHistory {
			if reservation.CarId == targetCarId {
				// 3. Mapeia os dados para a struct de resposta limpa.
				carHistory = append(carHistory, mapToResponse(reservation))
			}
		}

		// 4. Retorna o histórico filtrado como JSON.
		c.JSON(http.StatusOK, carHistory)
	})

	porta := os.Getenv("PORTA")
	if porta == "" {
		log.Fatalf("[SERVIDOR] Variavel de ambiente PORTA nao definida")
	}
	r.Run(":" + porta)
}

// mapToResponse é uma função auxiliar para converter o tipo do contrato em um tipo de resposta amigável.
func mapToResponse(reservation contracts.ReservationLedgerReservation) ReservationResponse {
	// Converte o status (uint8) para uma string legível.
	var status string
	switch reservation.Status {
	case 0:
		status = "Created"
	case 1:
		status = "Completed"
	case 2:
		status = "Cancelled"
	default:
		status = "Unknown"
	}

	return ReservationResponse{
		ExternalId:   "0x" + hex.EncodeToString(reservation.ExternalId[:]),
		CarId:        reservation.CarId,
		PostoId:      reservation.PostoId,
		Timestamp:    time.Unix(reservation.Timestamp.Int64(), 0).Format(time.RFC3339),
		ChargeAmount: reservation.ChargeAmount.String(), // Converte o *big.Int para string.
		Status:       status,
	}
}

func ObterPostosDeOutroServidor(url string) ([]*consts.Posto, error) {
	resp, err := http.Get(url + "/postos/disponiveis")
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisicao para %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("requisicao falhou com status %d", resp.StatusCode)
	}

	var postos []consts.Posto
	if err := json.NewDecoder(resp.Body).Decode(&postos); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %v", err)
	}

	var postosPointers []*consts.Posto
	for i := range postos {
		postosPointers = append(postosPointers, &postos[i])
	}
	return postosPointers, nil
}

func TwoPhaseCommit(participantes []consts.Participante2PC, carro consts.Carro) error {
	payloadTemplate := `{"posto_id":"%s","carro":%s}`

	okCount := 0
	for _, p := range participantes {
		carroJSON, _ := json.Marshal(carro)
		payload := fmt.Sprintf(payloadTemplate, p.PostoID, string(carroJSON))
		resp, err := http.Post(p.URL+"/2pc/prepare", "application/json", strings.NewReader(payload))
		if err != nil {
			log.Printf("[2PC] Erro ao enviar prepare para %s: %v", p.URL, err)
			break
		}
		defer resp.Body.Close()
		var res map[string]string
		json.NewDecoder(resp.Body).Decode(&res)
		if res["result"] == "ok" {
			okCount++
		} else {
			break
		}
	}

	if okCount == len(participantes) {
		for _, p := range participantes {
			carroJSON, _ := json.Marshal(carro)
			payload := fmt.Sprintf(payloadTemplate, p.PostoID, string(carroJSON))
			http.Post(p.URL+"/2pc/commit", "application/json", strings.NewReader(payload))
		}
		log.Println("[2PC] Commit enviado para todos os participantes")
		return nil
	}

	for _, p := range participantes {
		carroJSON, _ := json.Marshal(carro)
		payload := fmt.Sprintf(payloadTemplate, p.PostoID, string(carroJSON))
		http.Post(p.URL+"/2pc/abort", "application/json", strings.NewReader(payload))
	}
	log.Println("[2PC] Abort enviado para todos os participantes")
	return fmt.Errorf("2PC abortado por algum participante")
}
