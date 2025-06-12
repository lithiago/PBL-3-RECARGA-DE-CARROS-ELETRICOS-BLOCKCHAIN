// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

// Importa a biblioteca de scripting do Foundry e o contrato que queremos implantar.
import {Script} from "forge-std/Script.sol";
import {ReservationLedger} from "../src/ReservationLedger.sol";

// O contrato de script herda de `Script` para ter acesso aos comandos de script.
contract DeployReservationLedger is Script {
    
    // A função `run` é o ponto de entrada principal do nosso script.
    function run() external returns (ReservationLedger) {
        // `vm.startBroadcast()`: Inicia uma "transmissão". Todas as chamadas de contrato
        // a partir daqui serão enviadas como transações reais para a blockchain.
        vm.startBroadcast();

        // Implanta uma nova instância do contrato `ReservationLedger`.
        // O endereço do novo contrato será retornado pela função.
        ReservationLedger ledger = new ReservationLedger();

        // `vm.stopBroadcast()`: Para a transmissão.
        vm.stopBroadcast();

        // Retorna a instância do contrato implantado.
        return ledger;
    }
}
