const bodyParser = require('body-parser');
const routes = require('./routes/users.js');

const consts = require('./utils/constants.js');
const express = require('express');
const cors = require('cors');


const app = express();

async function main() {

    app.use(bodyParser.json());
    app.use(cors());
    app.use('/api', routes);

    app.get('/', (req, res) => res.send('Welcome to the Homepage!'));

    app.use((_req, res) =>
        res.status(404).json({
            status: "Invalid URL!",
            timestamp: new Date().toISOString(),
        })
    );

    console.log('Starting REST server');
    app.listen(consts.PORT, () => console.log(`Server is running on port: http://localhost:${consts.PORT}`));

}

main().catch(async (err) => {
    console.log('Unxepected error %s', err);
});







