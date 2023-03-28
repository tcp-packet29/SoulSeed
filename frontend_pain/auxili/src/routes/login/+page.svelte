<script>
    import axios from 'axios';
    import { writable } from "svelte/store";
    import { browser } from "$app/environment";
    import setItem from "./login.js";


    let tok = ""

    let uname = "";
    let pword = "";
    let toke = ""; //xss vulnerability
   // const store = window.localStorage.content;

    //is there an env that determines brorwser?



    function forPass() {
        axios.get('localhost:8080/access/users/token', {
            headers: {
            "Token": tok
        }})
            .then(function (response) {
                console.log(response);
                let jso = JSON.parse(JSON.stringify(response.data))
                toke = jso.data.Data.Id
            })
            .catch(function (error) {
                console.log(error);
            });

        if (browser) {
            //ssr
            window.localStorage.setItem("token", toke.toString());
        }
    }
    function createUser() {
        if (uname == "" || pword == "" || confirm == "") {
            alert("Please fill out all fields");
            return;
        }
        axios.post('http://localhost:8080/access/login', {
            "username": uname,
            "password": pword
        })
            .then(function (response) {
                if (response.status == 201 || response.status == 200) {
                    alert("Logged In successfully");
                }
                console.log(response);
                let jso = JSON.parse(JSON.stringify(response.data))
                console.log(jso)
                console.log(jso.Token)
                //programminglkib to intercet http request amqp
                //not immutable
               //idempotenmt
                if (browser) {
                    window.localStorage.setItem("token", jso.Token.toString());
                }

                console.log(window.localStorage.getItem("token"))
                tok = window.localStorage.getItem("token")

            })
            .catch(function (error) {
                console.log(error);
            });
    }

</script>

<body class="bg-primary"></body>
<div class="navbar bg-base-100 bg-accent">
    <a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173">Auxili</a>
    <h1>--</h1>
    <h2><a class="btn btn-sm btn-ghost rounded-btn bg-secondary" href="http://localhost:5173/register">Register</a></h2>
    <h1>--</h1>
    <h2><a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173/login">Login</a></h2>

</div>
<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral">Log In</h2>

            <input type="text" placeholder="Username" class="input input-bordered input-accent w-full max-w-xs" bind:value={uname}/>
            <input type="password" placeholder="Password" class="input input-bordered input-accent w-full max-w-xs" bind:value={pword}/>
            <p class="text-neutral textarea-sm"><a on:click={forPass}>Forgot Password?</a></p>
            <div class="card-actions">
                <button class="btn btn-outline btn-accent" on:click={createUser}>Login</button>
            </div>
        </div>
    </div>
</div>

