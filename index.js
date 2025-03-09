const mongoose = require("mongoose")
const users = require("./user.json")
const User = new mongoose.Schema({
    first_name: String,
    last_name: String,
    email: String,
    gender: String,
    ip_address: String
})
const UserModel = mongoose.model("users", User)
const main = async () => {
    await mongoose.connect("mongodb://root:example@localhost:27017/")
    console.time("Inserting")
    for (const user of users) {
        const res = await UserModel.insertOne(user)
    }
    console.timeEnd("Inserting")
    // console.log('Done', res)
}
main()