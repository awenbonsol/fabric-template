package com.fabrictemplate.template.controller;

import com.fabrictemplate.template.dto.DeleteKey;
import com.fabrictemplate.template.dto.Person;
import com.fabrictemplate.template.services.PersonService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@Slf4j
@CrossOrigin(origins = "*")
@RequestMapping("/api")
public class PersonController {

    @Autowired
    PersonService personService;

    @PostMapping(value = "/create")
    public ResponseEntity<Object> createPerson(@RequestBody Person person) {
        return personService.createPerson(person);
    }

    @PostMapping(value = "/update")
    public ResponseEntity<Object> updatePerson(@RequestBody Person person) {
        return personService.updatePerson(person.getNationalId(), person.getLastName());
    }

    @GetMapping(value = "/get-id")
    public ResponseEntity<Object> getPersonById(@RequestParam String nationalId) {
        return personService.getPersonById(nationalId);
    }

    @GetMapping(value = "/get-firstName")
    public ResponseEntity<Object> getPersonByFirstName(@RequestParam String firstName) {
        return personService.getPersonByFirstName(firstName);
    }

    @GetMapping(value = "/get-lastName")
    public ResponseEntity<Object> getPersonByLastName(@RequestParam String lastName, @RequestParam boolean isMarried) {
        return personService.getPersonByLastName(lastName, isMarried);
    }

    @GetMapping(value = "/get-all")
    public ResponseEntity<Object> getAll() {
        return personService.getAll();
    }

    @PostMapping(value = "/delete")
    public ResponseEntity<Object> deletePerson(@RequestBody DeleteKey deleteKey) {
        return personService.deletePerson(deleteKey.getKey());
    }

}
