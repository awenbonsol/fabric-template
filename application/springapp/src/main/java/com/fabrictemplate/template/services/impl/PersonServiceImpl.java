package com.fabrictemplate.template.services.impl;

import com.fabrictemplate.template.config.FabricConnection;
import com.fabrictemplate.template.dto.APIResponse;
import com.fabrictemplate.template.dto.Person;
import com.fabrictemplate.template.services.PersonService;
import com.google.gson.Gson;
import lombok.Generated;
import lombok.extern.slf4j.Slf4j;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.Gateway;
import org.hyperledger.fabric.gateway.Network;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.HttpServerErrorException;

import java.nio.charset.StandardCharsets;
import java.util.HashMap;
import java.util.Map;


@Service
@Generated
@Slf4j
public class PersonServiceImpl implements PersonService {

    public ResponseEntity<Object> createPerson(Person person) {

        byte[] byteResponse = null;
        APIResponse response = null;
        Gson gson = new Gson();
        HashMap<String, Object> myMap = new HashMap<>();
        Object jsonObj = gson.toJson(person, Object.class);
        try {
            log.info("Start: Creating person.");
            Gateway gateway = FabricConnection.makeConnection();
            Network network = gateway.getNetwork("mychannel");
            Contract contract = network.getContract("basic");

            byteResponse = contract.submitTransaction("CreatePerson", String.valueOf(jsonObj));
            log.info("End: Personhave been created.");

            String result = new String(byteResponse, StandardCharsets.UTF_8);
            myMap.put("transactionId", result);

            System.out.println(result);

            if (byteResponse != null) {
                response = APIResponse.success(myMap);
            }
            gateway.close();
        }catch (Exception e) {
            log.error(e.toString());
            response = APIResponse.error(e.toString());
            return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
        }

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    public ResponseEntity<Object> updatePerson(String nationalId, String lastName) {

        byte[] byteResponse = null;
        APIResponse response = null;
        Gson gson = new Gson();
        HashMap<String, Object> myMap = new HashMap<>();
        try {
            log.info("Start: Creating person.");
            Gateway gateway = FabricConnection.makeConnection();
            Network network = gateway.getNetwork("mychannel");
            Contract contract = network.getContract("basic");

            byteResponse = contract.submitTransaction("UpdatePerson", nationalId, lastName);
            log.info("End: Person have been updated.");

            String result = new String(byteResponse, StandardCharsets.UTF_8);
            myMap.put("transactionId", result);

            if (byteResponse != null) {
                response = APIResponse.success(myMap);
            }
            gateway.close();
        }catch (Exception e) {
            log.error(e.toString());
            response = APIResponse.error(e.toString());
            return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
        }

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    public ResponseEntity<Object> getPersonById(String nationalId) {

        byte[] byteResponse;
        APIResponse response = null;
        Gson gson = new Gson();
        try {
            log.info("Start: Fetching person by National Id.");
            Gateway gateway = FabricConnection.makeConnection();
            Network network = gateway.getNetwork("mychannel");
            Contract contract = network.getContract("basic");

            byteResponse = contract.submitTransaction("GetByNationalId", nationalId);
            log.info("End: Done fetching person by National Id.");

            String result = new String(byteResponse, StandardCharsets.UTF_8);

            if (byteResponse != null) {
                response = APIResponse.success(gson.fromJson(result, Object.class));
            }
            gateway.close();
        }catch (Exception e) {
            log.error(e.toString());
            response = APIResponse.error(e.toString());
            return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
        }

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    public ResponseEntity<Object> getPersonByFirstName(String firstName) {

        byte[] byteResponse;
        APIResponse response = null;
        Gson gson = new Gson();
        try {
            log.info("Start: Fetching person by Firstname.");
            Gateway gateway = FabricConnection.makeConnection();
            Network network = gateway.getNetwork("mychannel");
            Contract contract = network.getContract("basic");

            byteResponse = contract.submitTransaction("GetByFirstName", firstName);
            log.info("End: Done fetching person by Firstname.");

            String result = new String(byteResponse, StandardCharsets.UTF_8);

            if (byteResponse != null) {
                response = APIResponse.success(gson.fromJson(result, Object.class));
            }
            gateway.close();
        }catch (Exception e) {
            log.error(e.toString());
            response = APIResponse.error(e.toString());
            return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
        }

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    public ResponseEntity<Object> getPersonByLastName(String lastName, boolean isMarried) {

        byte[] byteResponse;
        APIResponse response = null;
        Gson gson = new Gson();
        try {
            log.info("Start: Fetching person by Lastname.");
            Gateway gateway = FabricConnection.makeConnection();
            Network network = gateway.getNetwork("mychannel");
            Contract contract = network.getContract("basic");

            byteResponse = contract.submitTransaction("GetByLastName", lastName, String.valueOf(isMarried));
            log.info("End: Done fetching person by Lastname.");

            String result = new String(byteResponse, StandardCharsets.UTF_8);

            if (byteResponse != null) {
                response = APIResponse.success(gson.fromJson(result, Object.class));
            }
            gateway.close();
        }catch (Exception e) {
            log.error(e.toString());
            response = APIResponse.error(e.toString());
            return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
        }

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    public ResponseEntity<Object> getAll() {

        byte[] byteResponse;
        APIResponse response = null;
        Gson gson = new Gson();
        try {
            log.info("Start: Getting all records.");
            Gateway gateway = FabricConnection.makeConnection();
            Network network = gateway.getNetwork("mychannel");
            Contract contract = network.getContract("basic");

            byteResponse = contract.submitTransaction("GetAll");
            log.info("End: Done fetching all records.");

            String result = new String(byteResponse, StandardCharsets.UTF_8);

            if (byteResponse != null) {
                response = APIResponse.success(gson.fromJson(result, Object.class));
            }
            gateway.close();
        }catch (Exception e) {
            log.error(e.toString());
            response = APIResponse.error(e.toString());
            return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
        }

        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    public ResponseEntity<Object> deletePerson(String key) {

        byte[] byteResponse = null;
        APIResponse response = null;
        HashMap<String, Object> myMap = new HashMap<>();
        try {
            log.info("Start: Deleting person.");
            Gateway gateway = FabricConnection.makeConnection();
            Network network = gateway.getNetwork("mychannel");
            Contract contract = network.getContract("basic");

            byteResponse = contract.submitTransaction("DeletePerson", key);
            log.info("End: Person have been deleted.");

            String result = new String(byteResponse, StandardCharsets.UTF_8);
            myMap.put("transactionId", result);

            if (byteResponse != null) {
                response = APIResponse.success(myMap);
            }
            gateway.close();
        }catch (Exception e) {
            log.error(e.toString());
            response = APIResponse.error(e.toString());
            return new ResponseEntity<>(response, HttpStatus.INTERNAL_SERVER_ERROR);
        }

        return new ResponseEntity<>(response, HttpStatus.OK);
    }
}
