// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/console.sol";

contract ReservationLedger {

    // Enum para controlar o estado de uma reserva de forma clara.
    enum Status { Created, Completed, Cancelled }

    // Struct para armazenar os detalhes da reserva.
    struct Reservation {
        bytes32 externalId;
        string carId;
        string postoId;
        uint256 timestamp;
        uint256 chargeAmount; // Valor da recarga para o pagamento.
        Status status;
    }

    // Mapeamento de um ID externo (gerado pelo nosso servidor Go) para os detalhes da Reserva.
    mapping(bytes32 => Reservation) public reservations;
    // Array para armazenar todos os IDs de reserva para um histórico iterável.
    bytes32[] public reservationHistory;

    // Eventos para logar atividades importantes na blockchain.
    event ReservationMade(bytes32 indexed externalId, string carId, string postoId);
    event RechargeCompleted(bytes32 indexed externalId);
    event ReservationCancelled(bytes32 indexed externalId);
    event PaymentProcessed(bytes32 indexed externalId, uint256 amount);

    /**
     * @dev Registra uma nova reserva no livro-razão usando um ID gerado externamente.
     */
    function makeReservation(bytes32 _externalId, string memory _carId, string memory _postoId) public {
        require(reservations[_externalId].timestamp == 0, "Reserva com este ID ja existe.");

        reservations[_externalId] = Reservation({
            externalId: _externalId,
            carId: _carId,
            postoId: _postoId,
            timestamp: block.timestamp,
            chargeAmount: 0,
            status: Status.Created
        });

        reservationHistory.push(_externalId);
        emit ReservationMade(_externalId, _carId, _postoId);
    }

    /**
     * @dev Marca uma recarga como concluída.
     */
    function completeRecharge(bytes32 _externalId) public {
        Reservation storage r = reservations[_externalId];
        require(r.timestamp != 0, "Reserva nao encontrada.");
        require(r.status == Status.Created, "A reserva nao esta em estado valido para ser completada.");
        
        r.status = Status.Completed;
        emit RechargeCompleted(_externalId);
    }

    /**
     * @dev Registra o pagamento de uma recarga concluída.
     */
    function processPayment(bytes32 _externalId, uint256 _amount) public {
        Reservation storage r = reservations[_externalId];
        require(r.timestamp != 0, "Reserva nao encontrada.");
        require(r.status == Status.Completed, "A recarga deve ser completada antes do pagamento.");

        r.chargeAmount = _amount;
        emit PaymentProcessed(_externalId, _amount);
    }

    /**
     * @dev Cancela uma reserva.
     */
    function cancelReservation(bytes32 _externalId) public {
        Reservation storage r = reservations[_externalId];
        require(r.timestamp != 0, "Reserva nao encontrada.");
        require(r.status == Status.Created, "Apenas reservas ativas podem ser canceladas.");

        r.status = Status.Cancelled;
        emit ReservationCancelled(_externalId);
    }

    /**
     * @dev Retorna os detalhes de uma reserva específica.
     * Ótimo para o endpoint de histórico.
     */
    function getReservation(bytes32 _externalId) public view returns (Reservation memory) {
        return reservations[_externalId];
    }
}
