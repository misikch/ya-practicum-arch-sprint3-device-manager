db.createUser({
    user: "mongouser",
    pwd: "mongopass",
    roles: [{
        role: "readWrite",
        db: "telemetry"
    }]
});

db = db.getSiblingDB('telemetry');

db.createCollection("devices");
db.createCollection("commands_log");