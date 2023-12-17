var PROTO_PATH = __dirname + "/../proto/rpc_send_email.proto";
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const packageDef = protoLoader.loadSync(PROTO_PATH, {});
const grpcObject = grpc.loadPackageDefinition(packageDef);
const emailPackage = grpcObject.pb;

const client = new emailPackage.EmailService("localhost:50051", grpc.credentials.createInsecure());

const email = {
    to: "ariefkerenss@gmail.com",
    playerId: "123456",
    productName: "PUBG",
    productPrice: 100000,
};

client.SendEmail(email, (error, response) => {
    if (error) {
        console.error(error);
        return;
    }
    console.log("Response from server: ", response);
});