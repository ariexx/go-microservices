var PROTO_PATH = __dirname + "/../proto/rpc_send_email.proto";
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const packageDef = protoLoader.loadSync(PROTO_PATH, {});
const grpcObject = grpc.loadPackageDefinition(packageDef);
const emailPackage = grpcObject.emailPackage;
const { sendEmail } = require('../pkg/email');

const server = new grpc.Server();
server.addService(emailPackage.EmailService.service, {
    sendEmail: sendEmail
});

//get server address from env
const serverAddress = process.env.SERVER_ADDRESS || ""
console.log("server address: ", serverAddress)

server.bindAsync(serverAddress, grpc.ServerCredentials.createInsecure(), (error) => {
    if (error) {
        console.error("error from grpc server email service" + error);
        return;
    }
    server.start();
    console.log("gRPC server running at :50051");
});
