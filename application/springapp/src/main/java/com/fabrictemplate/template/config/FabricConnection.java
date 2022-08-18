package com.fabrictemplate.template.config;


import lombok.Generated;
import org.hyperledger.fabric.gateway.Gateway;
import org.hyperledger.fabric.gateway.Wallet;
import org.hyperledger.fabric.gateway.Wallets;

import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Map;


@Generated
public class FabricConnection {

    static {
        System.setProperty("org.hyperledger.fabric.sdk.service_discovery.as_localhost", "true");
    }

    // Helper function for getting connected to the gateway
    public static Gateway connect() throws Exception {

        // Load a file system based wallet for managing identities.
        Path walletPath = Paths.get("wallet");
        Wallet wallet = Wallets.newFileSystemWallet(walletPath);
        // load a CCP
        Path networkConfigPath = Paths.get("src", "main", "resources","connection-org1.yaml");
        Gateway.Builder builder = Gateway.createBuilder();
        builder.identity(wallet, "admin").networkConfig(networkConfigPath).discovery(true);
        return builder.connect();
    }


    public static Gateway makeConnection() throws Exception {

        try {
            EnrollAdmin.main();
        } catch (Exception e) {
            System.err.println(e);
        }

        try (Gateway gateway = connect()) {
            return gateway;
        } catch (Exception e) {
            System.err.println(e);
            return null;
        }
    }
}