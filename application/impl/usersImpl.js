const { Gateway } = require("fabric-network");
const gateway = require('../connection.js');
const consts = require('../utils/constants.js');

function prettyJSONString(inputString) {
    return JSON.stringify(JSON.parse(inputString), null, 2);
}

exports.deletePerson = async (req, res) => {

    console.log('Start: Calling DeletePerson');
    const gatewayConnection = gateway();
    const network = (await gatewayConnection).getNetwork(consts.channelName);
    const contract = (await network).getContract(consts.chaincodeName);
    console.log('\nRequest payload ==> ', JSON.stringify(req.body));

    try {
        let result = await contract.submitTransaction(
            'DeletePerson',
            req.body.key
        );

        (await gatewayConnection).disconnect();
        console.log('End: DeletePerson called!');
        return res.status(200).json({
            status: "ACCEPTED",
            timestamp: new Date().toISOString(),
            transactionId: new Buffer.from(result).toString(),
        });
    } catch (err) {
        console.log(
            'Error processing delete person request for asset ID %s: %s',
            req.body.key, err
        );

        return res.status(500).json({
            status: "INTERNAL_SERVER_ERROR",
            timestamp: new Date().toISOString(),
        });
    }
};

exports.getAll = async (req, res) => {
    console.log('Start: Calling GetAll');
    console.log('Request: ', req.body);

    const gatewayConnection = gateway();
    const network = (await gatewayConnection).getNetwork(consts.channelName);
    const contract = (await network).getContract(consts.chaincodeName);

    try {
        let result = await contract.submitTransaction(
            'GetAll'
        );

        (await gatewayConnection).disconnect();
        console.log('End: GetAll called!');
        return res.status(200).json({
            status: "ACCEPTED",
            timestamp: new Date().toISOString(),
            data: JSON.parse(result.toString())
        });
    } catch (err) {
        console.log(
            'Error processing get all people request for asset ID %s: %s',
            err
        );

        return res.status(500).json({
            status: "INTERNAL_SERVER_ERROR",
            timestamp: new Date().toISOString(),
            data: []
        });
    }
};

exports.getByLastName = async (req, res) => {
    console.log('Start: Calling GetByLastName');
    console.log('Request: ', req.body);

    const gatewayConnection = gateway();
    const network = (await gatewayConnection).getNetwork(consts.channelName);
    const contract = (await network).getContract(consts.chaincodeName);

    try {
        let result = await contract.submitTransaction(
            'GetByLastName',
            req.query.lastName,
            req.query.isMarried
        );

        (await gatewayConnection).disconnect();
        console.log('End: GetByLastName called!');
        return res.status(200).json({
            status: "ACCEPTED",
            timestamp: new Date().toISOString(),
            data: JSON.parse(result.toString())
        });
    } catch (err) {
        console.log(
            'Error processing get person by first name id request for asset ID %s: %s',
            err
        );

        return res.status(500).json({
            status: "INTERNAL_SERVER_ERROR",
            timestamp: new Date().toISOString(),
        });
    }
};


exports.getByFirstName = async (req, res) => {
    console.log('Start: Calling GetByFirstName');
    console.log('Request: ', req.body);

    const gatewayConnection = gateway();
    const network = (await gatewayConnection).getNetwork(consts.channelName);
    const contract = (await network).getContract(consts.chaincodeName);

    try {
        let result = await contract.submitTransaction(
            'GetByFirstName',
            req.query.firstName
        );

        (await gatewayConnection).disconnect();
        console.log('End: GetByFirstName called!');
        return res.status(200).json({
            status: "ACCEPTED",
            timestamp: new Date().toISOString(),
            data: JSON.parse(result.toString())
        });
    } catch (err) {
        console.log(
            'Error processing get person by first name id request for asset ID %s: %s',
            err
        );

        return res.status(500).json({
            status: "INTERNAL_SERVER_ERROR",
            timestamp: new Date().toISOString(),
        });
    }
};

exports.getByNationalId = async (req, res) => {
    console.log('Start: Calling getByNationalId');
    console.log('Request: ', req.body);

    const gatewayConnection = gateway();
    const network = (await gatewayConnection).getNetwork(consts.channelName);
    const contract = (await network).getContract(consts.chaincodeName);

    try {
        let result = await contract.submitTransaction(
            'GetByNationalId',
            req.query.nationalId
        );

        (await gatewayConnection).disconnect();
        console.log('End: getByNationalId called!');
        return res.status(200).json({
            status: "ACCEPTED",
            timestamp: new Date().toISOString(),
            data: JSON.parse(result.toString())
        });
    } catch (err) {
        console.log(
            'Error processing get person by national id request for asset ID %s: %s',
            err
        );

        return res.status(500).json({
            status: "INTERNAL_SERVER_ERROR",
            timestamp: new Date().toISOString(),
        });
    }
};

exports.updatePerson = async (req, res) => {
    console.log('Start: Calling update-product');
    console.log('Request: ', req.body);
    const nationalId = req.body.nationalId;

    const gatewayConnection = gateway();
    const network = (await gatewayConnection).getNetwork(consts.channelName);
    const contract = (await network).getContract(consts.chaincodeName);

    try {
        let result = await contract.submitTransaction(
            'UpdatePerson',
            req.body.nationalId,
            req.body.lastName
        );

        (await gatewayConnection).disconnect();
        console.log('End: update-product called!');
        return res.status(200).json({
            status: "ACCEPTED",
            timestamp: new Date().toISOString(),
            transactionId: new Buffer.from(result).toString(),
        });
    } catch (err) {
        console.log(
            'Error processing update person request for asset ID %s: %s',
            nationalId, err
        );

        return res.status(500).json({
            status: "INTERNAL_SERVER_ERROR",
            timestamp: new Date().toISOString(),
        });
    }
};

exports.createPerson = async (req, res) => {
    console.log('Start: Calling createPerson');
    const nationalId = req.body.nationalId;

    const gatewayConnection = gateway();
    const network = (await gatewayConnection).getNetwork(consts.channelName);
    const contract = (await network).getContract(consts.chaincodeName);

    console.log('\nRequest payload ==> ', JSON.stringify(req.body));

    try {
        let result = await contract.submitTransaction(
            'CreatePerson',
            JSON.stringify(req.body)
        );

        (await gatewayConnection).disconnect();
        console.log('End: createPerson called!');
        return res.status(200).json({
            status: "ACCEPTED",
            timestamp: new Date().toISOString(),
            transactionId: new Buffer.from(result).toString(),
        });
    } catch (err) {
        console.log(
            'Error processing create person request for asset ID %s: %s',
            nationalId, err
        );

        return res.status(500).json({
            status: "INTERNAL_SERVER_ERROR",
            timestamp: new Date().toISOString(),
        });
    }
};
