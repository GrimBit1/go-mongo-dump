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
const InsertMany = async () => {
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
const Find = async () => {
    await mongoose.connect("mongodb://root:example@localhost:27017/")
    console.time("Finding")
    const users = await UserModel.find()
    console.log(users[0])
    console.timeEnd("Finding")
}

const FindByName = async () => {
    await mongoose.connect("mongodb://root:example@localhost:27017/")
    console.time("Finding")
    const users = await UserModel.findOne({
        first_name: 'Myrtia'
    })
    console.log(users)
    console.timeEnd("Finding")
}
const FindById = async () => {
    await mongoose.connect("mongodb://root:example@localhost:27017/")
    console.time("Finding")
    const users = await UserModel.findById("67cdb7acfb3bc8b95b8d7208")
    console.log(users)
    console.timeEnd("Finding")
}
const FindByRegex = async () => {
    await mongoose.connect("mongodb://root:example@localhost:27017/")
    console.time("Finding")
    const users = await UserModel.find({
        "first_name": {
            "$regex": "^A",
            "$options": "i",
        },
    })
    console.log(users.length)
    console.timeEnd("Finding")
}
FindByRegex()