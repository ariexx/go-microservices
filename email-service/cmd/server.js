var PROTO_PATH = __dirname + "/../proto/rpc_send_email.proto";
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const packageDef = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});
const grpcObject = grpc.loadPackageDefinition(packageDef);
const emailPackage = grpcObject.pb;
const { sendEmail } = require('../pkg/email');

let Server = new grpc.Server();
Server.addService(emailPackage.EmailService.service, {
    SendEmail: (call, callback) => {
        console.log("Request from client: ", call.request);
        sendEmail(call.request).then(r => {
            console.log("Response from sendEmail: ", r);
            //set message id to response
            callback(null, { status: "OK" });
        });
        callback(null, { status: "OK" });
    }
});

//get server address from env
// const serverAddress = process.env.SERVER_ADDRESS;
// console.log("server address: ", serverAddress)

Server.bindAsync("0.0.0.0:50051", grpc.ServerCredentials.createInsecure(), () => {
    console.log("Server running at 50051");
    Server.start();
});
