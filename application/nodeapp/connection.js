const { Gateway, Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const { buildCAClient, registerAndEnrollUser, enrollAdmin } = require('./connection/javascript/CAUtil.js');
const { buildCCPOrg1, buildWallet } = require('./connection/javascript/AppUtil.js');
const path = require('path');
const consts = require('./utils/constants.js');

const walletPath = path.join(__dirname, 'wallet');

async function connection() {
    const ccp = buildCCPOrg1();
    const caClient = buildCAClient(FabricCAServices, ccp, 'ca.org1.example.com');
    const wallet = await buildWallet(Wallets, walletPath);

    await enrollAdmin(caClient, wallet, consts.mspOrg1);
    await registerAndEnrollUser(caClient, wallet, consts.mspOrg1, consts.org1UserId, 'org1.department1');
    const gateway = new Gateway();

    await gateway.connect(ccp, {
        wallet,
        identity: consts.org1UserId,
        discovery: { enabled: true, asLocalhost: true }
    });
    return gateway;

}

module.exports = connection;
