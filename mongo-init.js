db.auth("root", "root");
db = db.getSiblingDB("linebot-go");

db.createUser({
  user: "user",
  pwd: "user_password",
  roles: [
    {
      role: "readWrite",
      db: "linebot-go",
    },
  ],
});
