const express = require('express');

const { createPerson, updatePerson, getByNationalId, getByFirstName, getByLastName, getAll, deletePerson } = require('../impl/usersImpl.js');

const router = express.Router();

router.post('/create', createPerson);

router.post('/update', updatePerson);

router.get('/get-id', getByNationalId);

router.get('/get-firstname', getByFirstName);

router.get('/get-lastname', getByLastName);

router.get('/get-all', getAll);

router.post('/delete', deletePerson);

module.exports = router;