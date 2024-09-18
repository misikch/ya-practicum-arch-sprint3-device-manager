db.createUser({
    user: "mongouser",
    pwd: "mongopass",
    roles: [{
        role: "readWrite",
        db: "device-manager"
    }]
});

db = db.getSiblingDB('device-manager');

db.createCollection("devices");
db.createCollection("commands_log");