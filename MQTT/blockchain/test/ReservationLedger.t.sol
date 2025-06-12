// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/ReservationLedger.sol";

contract ReservationLedgerTest is Test {
    
    ReservationLedger public ledger;

    // Variáveis que podemos reutilizar nos testes
    bytes32 private constant MOCK_ID_1 = keccak256("reserva-1");
    string private constant CAR_ID_1 = "carro-123";
    string private constant POSTO_ID_1 = "posto-abc";

    function setUp() public {
        ledger = new ReservationLedger();
    }

    // --- Testes para makeReservation ---

    function test_FazerReservaComSucesso() public {
        // Prepara o teste para esperar o evento com os parâmetros corretos
        vm.expectEmit(true, true, true, false);
        emit ReservationLedger.ReservationMade(MOCK_ID_1, CAR_ID_1, POSTO_ID_1);

        // Executa a função
        ledger.makeReservation(MOCK_ID_1, CAR_ID_1, POSTO_ID_1);

        // Verifica os dados armazenados usando a nova função getReservation
        ReservationLedger.Reservation memory r = ledger.getReservation(MOCK_ID_1);

        assertEq(r.externalId, MOCK_ID_1, "O ID externo nao foi armazenado corretamente.");
        assertEq(r.carId, CAR_ID_1, "O ID do carro nao foi armazenado corretamente.");
        assertEq(r.postoId, POSTO_ID_1, "O ID do posto nao foi armazenado corretamente.");
        assertEq(uint(r.status), uint(ReservationLedger.Status.Created), "O status inicial deve ser Created.");
    }

    function test_FalhaAoFazerReservaDuplicada() public {
        ledger.makeReservation(MOCK_ID_1, CAR_ID_1, POSTO_ID_1);

        // Espera a falha ao tentar criar a mesma reserva novamente
        vm.expectRevert("Reserva com este ID ja existe.");
        ledger.makeReservation(MOCK_ID_1, "carro-diferente", "posto-diferente");
    }

    // --- Testes para completeRecharge ---

    function test_CompletarRecargaComSucesso() public {
        ledger.makeReservation(MOCK_ID_1, CAR_ID_1, POSTO_ID_1);

        vm.expectEmit(true, false, false, false);
        emit ReservationLedger.RechargeCompleted(MOCK_ID_1);

        ledger.completeRecharge(MOCK_ID_1);

        ReservationLedger.Reservation memory r = ledger.getReservation(MOCK_ID_1);
        assertEq(uint(r.status), uint(ReservationLedger.Status.Completed), "O status da reserva deveria ser Completed.");
    }

    function test_FalhaAoCompletarRecargaInexistente() public {
        vm.expectRevert("Reserva nao encontrada.");
        ledger.completeRecharge(MOCK_ID_1);
    }
    
    // --- Testes para processPayment ---

    function test_ProcessarPagamentoComSucesso() public {
        uint256 paymentAmount = 100 * 1e18; // Simula um pagamento de 100 "tokens"
        ledger.makeReservation(MOCK_ID_1, CAR_ID_1, POSTO_ID_1);
        ledger.completeRecharge(MOCK_ID_1); // A recarga precisa ser completada primeiro

        vm.expectEmit(true, false, false, false);
        emit ReservationLedger.PaymentProcessed(MOCK_ID_1, paymentAmount);

        ledger.processPayment(MOCK_ID_1, paymentAmount);

        ReservationLedger.Reservation memory r = ledger.getReservation(MOCK_ID_1);
        assertEq(r.chargeAmount, paymentAmount, "O valor do pagamento nao foi registrado corretamente.");
    }

    function test_FalhaAoProcessarPagamentoAntesDeCompletar() public {
        uint256 paymentAmount = 100 * 1e18;
        ledger.makeReservation(MOCK_ID_1, CAR_ID_1, POSTO_ID_1);

        vm.expectRevert("A recarga deve ser completada antes do pagamento.");
        ledger.processPayment(MOCK_ID_1, paymentAmount);
    }
    
    // --- Testes para cancelReservation ---

    function test_CancelarReservaComSucesso() public {
        ledger.makeReservation(MOCK_ID_1, CAR_ID_1, POSTO_ID_1);

        vm.expectEmit(true, false, false, false);
        emit ReservationLedger.ReservationCancelled(MOCK_ID_1);

        ledger.cancelReservation(MOCK_ID_1);

        ReservationLedger.Reservation memory r = ledger.getReservation(MOCK_ID_1);
        assertEq(uint(r.status), uint(ReservationLedger.Status.Cancelled), "O status da reserva deveria ser Cancelled.");
    }

    function test_FalhaAoCancelarReservaConcluida() public {
        ledger.makeReservation(MOCK_ID_1, CAR_ID_1, POSTO_ID_1);
        ledger.completeRecharge(MOCK_ID_1);

        vm.expectRevert("Apenas reservas ativas podem ser canceladas.");
        ledger.cancelReservation(MOCK_ID_1);
    }
}
