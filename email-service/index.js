const express = require("express");
const bodyParser = require('body-parser');
const { sendEmail } = require('./pkg/email');

const app = express();
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

const route = express.Router();
const port = process.env.PORT || 5000;

app.use('/v1', route);

route.post('/email', async (req, res) => {
    try {
        await sendEmail(new EmailRequest(req.body.email, req.body.player_id, req.body.product, req.body.price, req.body.expiredTime));
        console.log("email sent : ", req.body.email)
        res.status(200).send(new EmailResponse('success', 'Email sent'));
    } catch (err) {
        console.log("email error : ", err.message)
        res.status(500).send(new EmailResponse('error', err.message));
    }
});


app.listen(port, () => {
    console.log(`Server listening on port ${port}`);
});

//create dto for request
class EmailRequest {
    constructor(email, player_id, product, price, expiredTime) {
        this.email = email;
        this.player_id = player_id;
        this.product = product;
        this.price = price;
        this.expiredTime = expiredTime;
    }
}

//create dto for response
class EmailResponse {
    constructor(status, message) {
        this.status = status;
        this.message = message;
    }
}