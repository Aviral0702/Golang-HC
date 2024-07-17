import express from "express";
import cors from "cors";

const app = express();

app.use(express.json());
app.use(express.urlencoded({extended:true}));
app.use(cors());

app.get("/", (req, res) => {
    res.status(200).json("Welcome to the server");
});

app.get("/get",(req,res)=> {
    res.status(200).json({message : "New message from get request"});
});


app.post("/post",(req,res)=> {
    let myJson = req.body;
    res.status(200).json(myJson);
});

app.post("/postform",(req,res)=> {
    res.status(200).send(JSON.stringify(req.body));
});


app.listen(8000, () => {
    console.log("Server is running on port 8000");
});