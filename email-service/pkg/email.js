const nodemailer = require('nodemailer');
const mustache = require('mustache');
const fs = require('fs');

async function sendEmail(req, res) {
    const template = fs.readFileSync('./template/transaction_success.html', { encoding: 'utf-8' });
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
        to: req.email,
        subject: "Invoice Transaksi Arief Store",
        html: mustache.render(template, {
            playerId: req.playerId,
            product: req.product,
            price: price,
        })
    }

    let info = await transporter.sendMail(mailOptions);

    console.log("Message sent: %s", info.messageId);
}

module.exports = {
    sendEmail
}