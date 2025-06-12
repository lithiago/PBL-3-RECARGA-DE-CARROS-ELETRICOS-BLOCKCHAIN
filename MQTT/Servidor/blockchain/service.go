package blockchain

import (
	"context"
	"log"
	"math/big"
	"time" // Adicionado para a lógica de espera

	// IMPORTANTE: Substitua 'MQTT' se o nome do seu módulo no arquivo go.mod for diferente.
	"MQTT/Servidor/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

// Constantes para a conexão com a blockchain
const (
	// A URL aponta para o nome do serviço 'anvil' dentro da rede Docker.
	blockchainNodeURL = "http://anvil:8545"
	// Chave privada de uma das contas do Anvil
	deployerPrivateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	// MUITO IMPORTANTE: Substitua pelo endereço do seu contrato, se for diferente!
	contractHexAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
)

var ledgerInstance *contracts.ReservationLedger
var transactor *bind.TransactOpts

// InitBlockchainService inicializa a conexão com a blockchain, agora com lógica de repetição.
func InitBlockchainService() {
	log.Println("[BLOCKCHAIN] Inicializando servico...")

	var client *ethclient.Client
	var err error
	maxRetries := 5
	retryDelay := 10 * time.Second

	// Tenta conectar várias vezes antes de desistir.
	for i := 0; i < maxRetries; i++ {
		client, err = ethclient.Dial(blockchainNodeURL)
		if err == nil {
			// Sucesso na conexão, agora tenta obter o ChainID para confirmar
			_, err = client.ChainID(context.Background())
			if err == nil {
				log.Println("[BLOCKCHAIN] Conectado e validado com o no Ethereum com sucesso.")
				break // Sai do loop se a conexão e a validação do ChainID funcionarem
			}
		}
		log.Printf("[BLOCKCHAIN] Falha ao conectar/validar com o cliente Ethereum (tentativa %d/%d): %v. Tentando novamente em %v...\n", i+1, maxRetries, err, retryDelay)
		time.Sleep(retryDelay)
	}

	// Se ainda houver um erro após todas as tentativas, o programa para.
	if err != nil {
		log.Fatalf("Falha ao conectar ao cliente Ethereum apos %d tentativas: %v", maxRetries, err)
	}

	// Como a conexão foi bem-sucedida, podemos obter o ChainID sem erro
	chainID, _ := client.ChainID(context.Background())

	privateKey, err := crypto.HexToECDSA(deployerPrivateKey)
	if err != nil {
		log.Fatalf("Falha ao carregar chave privada: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Falha ao criar transator: %v", err)
	}
	transactor = auth

	address := common.HexToAddress(contractHexAddress)
	instance, err := contracts.NewReservationLedger(address, client)
	if err != nil {
		log.Fatalf("Falha ao instanciar o contrato ReservationLedger: %v", err)
	}
	ledgerInstance = instance

	log.Println("[BLOCKCHAIN] Servico inicializado com sucesso.")
}

// GenerateReservationID cria um novo ID único para a reserva.
func GenerateReservationID() [32]byte {
	var id [32]byte
	newUUID := uuid.New()
	copy(id[:], newUUID[:])
	return id
}

// RecordReservationOnChain registra uma nova reserva no livro-razão da blockchain.
func RecordReservationOnChain(externalID [32]byte, carID string, postoID string) error {
	log.Printf("[BLOCKCHAIN] Registrando reserva ID %x para Carro: %s\n", externalID, carID)
	if ledgerInstance == nil {
		return log.Output(1, "Servico blockchain nao inicializado")
	}

	tx, err := ledgerInstance.MakeReservation(transactor, externalID, carID, postoID)
	if err != nil {
		log.Printf("Falha ao enviar a transacao 'makeReservation': %v", err)
		return err
	}

	log.Printf("[BLOCKCHAIN] Transacao 'makeReservation' enviada! Hash: %s\n", tx.Hash().Hex())
	return nil
}

// RecordRechargeCompletedOnChain chama a função completeRecharge do contrato.
func RecordRechargeCompletedOnChain(externalID [32]byte) error {
	log.Printf("[BLOCKCHAIN] Registrando conclusao da recarga para o ID %x\n", externalID)
	if ledgerInstance == nil {
		return log.Output(1, "Servico blockchain nao inicializado")
	}

	tx, err := ledgerInstance.CompleteRecharge(transactor, externalID)
	if err != nil {
		log.Printf("Falha ao enviar a transacao 'completeRecharge': %v", err)
		return err
	}

	log.Printf("[BLOCKCHAIN] Transacao 'completeRecharge' enviada! Hash: %s\n", tx.Hash().Hex())
	return nil
}

// RecordPaymentOnChain chama a função processPayment do contrato.
func RecordPaymentOnChain(externalID [32]byte, amount *big.Int) error {
	log.Printf("[BLOCKCHAIN] Registrando pagamento de %s para o ID %x\n", amount.String(), externalID)
	if ledgerInstance == nil {
		return log.Output(1, "Servico blockchain nao inicializado")
	}

	tx, err := ledgerInstance.ProcessPayment(transactor, externalID, amount)
	if err != nil {
		log.Printf("Falha ao enviar a transacao 'processPayment': %v", err)
		return err
	}

	log.Printf("[BLOCKCHAIN] Transacao 'processPayment' enviada! Hash: %s\n", tx.Hash().Hex())
	return nil
}

// RecordReservationCancelledOnChain chama a função cancelReservation do contrato.
func RecordReservationCancelledOnChain(externalID [32]byte) error {
	log.Printf("[BLOCKCHAIN] Registrando cancelamento para o ID %x\n", externalID)
	if ledgerInstance == nil {
		return log.Output(1, "Servico blockchain nao inicializado")
	}

	tx, err := ledgerInstance.CancelReservation(transactor, externalID)
	if err != nil {
		log.Printf("Falha ao enviar a transacao 'cancelReservation': %v", err)
		return err
	}

	log.Printf("[BLOCKCHAIN] Transacao 'cancelReservation' enviada! Hash: %s\n", tx.Hash().Hex())
	return nil
}

// GetFullHistoryOnChain busca todas as reservas na blockchain.
func GetFullHistoryOnChain() ([]contracts.ReservationLedgerReservation, error) {
	if ledgerInstance == nil {
		return nil, log.Output(1, "Servico blockchain nao inicializado")
	}

	var allReservations []contracts.ReservationLedgerReservation
	for i := int64(0); ; i++ {
		id, err := ledgerInstance.ReservationHistory(nil, big.NewInt(i))
		if err != nil {
			break
		}
		reservation, err := ledgerInstance.GetReservation(nil, id)
		if err != nil {
			log.Printf("Falha ao buscar detalhes da reserva para o ID %x: %v", id, err)
			continue
		}
		allReservations = append(allReservations, reservation)
	}

	return allReservations, nil
}
