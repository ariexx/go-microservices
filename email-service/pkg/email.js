const nodemailer = require('nodemailer');
const mustache = require('mustache');
const fs = require('fs');

var TEMPLATE_PATH = __dirname + "/../template/transaction_success.html";

async function sendEmail(req) {
    const template = fs.readFileSync(TEMPLATE_PATH, { encoding: 'utf-8' });
    let transporter = nodemailer.createTransport({
        host: 'sandbox.smtp.mailtrap.io',
        port: 2525,
        secure: false,
        auth: {
            'user': '27de36f4ba6439',
            'pass': '1752eeafc75a03',
        }
    })

    //format rupiah
    let price = Intl.NumberFormat('id-ID', {
        style: 'currency',
        currency: 'IDR',
    }).format(req.price);

    let mailOptions = {
        from: "arief@microservices.com",
        to: "ariefkeren@gmail.com",
        subject: "Payment Pending - Arief Store",
        html: mustache.render(template, {
            player_id: req.player_id,
            product: req.product,
            price: price,
        })
    }

    let info = await transporter.sendMail(mailOptions);

    console.log("Message sent: %s", info.messageId);
    return info.messageId;
}

module.exports = {
    sendEmail
}