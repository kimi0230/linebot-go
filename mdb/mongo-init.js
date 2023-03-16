print(
  "Start #################################################################"
);

// db.auth("root", "root");

db = db.getSiblingDB("linebot-go");
db.createUser({
  user: "user",
  pwd: "user_password",
  roles: [{ role: "readWrite", db: "linebot-go" }],
});
db.createCollection("users");

db = db.getSiblingDB("linebot-go");

db.createUser({
  user: "api_user",
  pwd: "api1234",
  roles: [{ role: "readWrite", db: "linebot-go" }],
});
db.createCollection("users");

print("END #################################################################");
