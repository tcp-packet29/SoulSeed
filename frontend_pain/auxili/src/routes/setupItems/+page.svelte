<script>
    import axios from 'axios';
    import {writable} from "svelte/store";
    import {browser} from "$app/environment";

    class Choices {
        constructor(mango, banana, apple, orange, persimmon, avocado, peach , apricot) { //SEPFICI INSTRACTURE, ANNOYING WHEN ADDINGT CVUSTOM CHOCIES
            this.mango = mango;
            this.banana = banana;
            this.apple = apple;
            this.orange = orange;
            this.persimmon = persimmon;
            this.avocado = avocado;
            this.peach = peach;
            this.apricot = apricot;
            this.choiceList = [this.mango, this.banana, this.apple, this.orange, this.persimmon, this.avocado, this.peach, this.apricot];

        }
    //conver to tlist and return for put requ
        convertToList() {
            let list = [];
            for (let i = 0; i < this.choiceList.length; i++) {
                if (this.choiceList[i][1]) {
                    list.push(this.choiceList[i][0]);
                }
            }
            //dont want toi make big elifs
            return list;

        }
    }

    let uname = "";
    let pword = "";
    let toke = ""; //xss vulnerability
    //const store = localStorage.content;

    let choices = new Choices(["Mango", false], ["Banana", false], ["Apple", false], ["Orange", false], ["Persimmon", false], ["Avocado", false], ["Peach", false], ["Apricot", false]);
    let test = choices.convertToList()

    function getToken() {
        if (browser) {
            return window.localStorage.getItem("token");
        }
    }

    function tests() {
        let testlist = choices.convertToList();
        console.log(testlist);

        axios.get('http://localhost:8080/access/users/token', {
            headers: {
                "Token": getToken()
            },
        })
            .then(function (response) {
                toke = JSON.parse(JSON.stringify(response.data)).data.Data;
                axios.put('http://localhost:8080/access/users/' + toke.Id, {
                    headers: {
                        "Token": getToken()
                    },
                    "username": toke.Username,
                    "email": toke.Email,
                    "password": toke.Password,
                    "zipcode": toke.Zipcode,
                    "organization_owned": toke.organization_owned,
                    "items": testlist

                })
                    .then(function (response) {
                        console.log(response)
                    })
                    .catch(function(erro) {
                        console.log(erro)
                    })
            })
            .catch(function (error) {
                console.log(error);
            });
    }

    function createUser() {
        if (uname == "" || pword == "" || zipcode == "" || confirm == "") {
            alert("Please fill out all fields");
            return;
        } else if (pword != confirm) {
            alert("Passwords do not match");
            return;
        }
        axios.post('http://localhost:8080/access/login', {
            "username": uname,
            "password": pword
        })
            .then(function (response) {
                if (response.status == 201 || response.status == 200) {
                    alert("Logged In successfully");
                    window.location.href = "http://localhost:5173/login";
                }
                console.log(response);
                let jso = JSON.parse(response.data)
                toke = writable(jso.token); //not immutable
                //localStorage.setItem("content", toke);


            })
            .catch(function (error) {
                console.log(error);
            });
    }

</script>
<body class="bg-primary"></body>
<button class="btn btn-accent" on:click={tests}>asdfasdf</button>
<div class="navbar bg-base-100 bg-accent">
    <a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173">Auxili</a>
    <h1>--</h1>
    <h2><a class="btn btn-sm btn-ghost rounded-btn bg-secondary" href="http://localhost:5173/register">Register</a>
    </h2>
    <h1>--</h1>
    <h2><a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173/login">Login</a></h2>

</div>
<br>
<div class="min-h-screen flex items-center justify-center">
    <div class="grid gap-20 grid-cols-3">
        <div class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
            <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
            <div class="card-body">
                <h2 class="card-title text-neutral">Golang</h2>
                <p class="text-neutral">High Level, Static Typed, Compiled</p>
                <div class="card-actions justify-end">
                    <input type="checkbox" class="checkbox checkbox-accent checkbox-lg" bind:checked={choices.choiceList[0][1]}/>
                </div>
            </div>
        </div>
        <div class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
            <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
            <div class="card-body">
                <h2 class="card-title text-neutral">Golang</h2>
                <p class="text-neutral">High Level, Static Typed, Compiled</p>
                <div class="card-actions justify-end">
                    <input type="checkbox" class="checkbox checkbox-accent checkbox-lg" bind:checked={choices.choiceList[1][1]}/>
                </div>
            </div>
        </div>
        <div class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
            <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
            <div class="card-body">
                <h2 class="card-title text-neutral">Golang</h2>
                <p class="text-neutral">High Level, Static Typed, Compiled</p>
                <div class="card-actions justify-end">
                    <input type="checkbox" class="checkbox checkbox-accent checkbox-lg" bind:checked={choices.choiceList[2][1]}/>
                </div>
            </div>
        </div>
        <div class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
            <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
            <div class="card-body">
                <h2 class="card-title text-neutral">Golang</h2>
                <p class="text-neutral">High Level, Static Typed, Compiled with garcol fvor saving and precise ocnteol over memory (not really precise very precise)</p>
                <div class="card-actions justify-end">
                    <input type="checkbox" class="checkbox checkbox-accent checkbox-lg" bind:checked={choices.choiceList[3][1]}/>
                </div>
            </div>
        </div>
        <div class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
            <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
            <div class="card-body">
                <h2 class="card-title text-neutral">Golang</h2>
                <p class="text-neutral">High Level, Static Typed, Compiled</p>
                <div class="card-actions justify-end">
                    <input type="checkbox" class="checkbox checkbox-accent checkbox-lg" bind:checked={choices.choiceList[4][1]}/>
                </div>
            </div>
        </div>
        <div class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
            <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
            <div class="card-body">
                <h2 class="card-title text-neutral">Golang</h2>
                <p class="text-neutral">High Level, Static Typed, Compiled</p>
                <div class="card-actions justify-end">
                    <input type="checkbox" class="checkbox checkbox-accent checkbox-lg" bind:checked={choices.choiceList[5][1]}/>
                </div>
            </div>
        </div>
        <div class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
            <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
            <div class="card-body">
                <h2 class="card-title text-neutral">Golang</h2>
                <p class="text-neutral">High Level, Static Typed, Compiled</p>
                <div class="card-actions justify-end">
                    <input type="checkbox" class="checkbox checkbox-accent checkbox-lg" bind:checked={choices.choiceList[6][1]}/>
                </div>
            </div>
        </div>
        <div class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
            <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
            <div class="card-body">
                <h2 class="card-title text-neutral">Golang</h2>
                <p class="text-neutral">High Level, Static Typed, Compiled</p>
                <div class="card-actions justify-end">
                    <input type="checkbox" class="checkbox checkbox-accent checkbox-lg" bind:checked={choices.choiceList[7][1]}/>
                </div>
            </div>
        </div>
    </div>
</div>
<br>
<br>

