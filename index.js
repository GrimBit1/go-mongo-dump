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
    const BATCH_SIZE = 100000;
    const batch = []
    for (let i = 0; i < users.length; i += BATCH_SIZE) {
        const batc = users.slice(i, i + BATCH_SIZE);
        batch.push(batc)

    }
    console.log(batch[batch.length - 1].length)
    console.time("Inserting")
    await Promise.all(batch.map(async (small) => {
        await UserModel.insertMany(small)
    }))
    console.timeEnd("Inserting")
    // console.log('Done', res)
}
main()