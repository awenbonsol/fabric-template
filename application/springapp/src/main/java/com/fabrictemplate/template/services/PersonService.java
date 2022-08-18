package com.fabrictemplate.template.services;

import com.fabrictemplate.template.dto.DeleteKey;
import com.fabrictemplate.template.dto.Person;
import org.springframework.http.ResponseEntity;

public interface PersonService {

    public ResponseEntity<Object> createPerson(Person person);

    public ResponseEntity<Object> updatePerson(String nationalId, String lastName);

    public ResponseEntity<Object> getPersonById(String nationalId);

    public ResponseEntity<Object> getPersonByFirstName(String firstName);

    public ResponseEntity<Object> getPersonByLastName(String lastName, boolean isMarried);

    public ResponseEntity<Object> getAll();

    public ResponseEntity<Object> deletePerson(String key);
}
