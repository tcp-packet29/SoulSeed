<script>
    import axios from 'axios';
    import {browser} from "$app/environment";
    import Navbar from "../../lib/navbar.svelte";
    import { page } from '$app/stores';






    let val = '';
    let email = "";
    function getToken() {
        if (browser) {
            return window.localStorage.getItem("token");

        }
    }
    function getVal(val, another) {
        for (let i = 0; i < 10; i++) {
            if (val[i] == another) {
                return true
            }
        }
        return false
    }
    var data = [];
    axios.defaults.headers.put['Token'] = getToken();
    axios.get('http://localhost:8080/access/users/token', {
        headers: {
            'Token': getToken(),
        }
    })
    .then(response => {
            let jsobj = JSON.parse(JSON.stringify(response.data));
            console.log(jsobj)
            console.log(jsobj.data.Data.Username)
            console.log(jsobj.data.Data.Email)
        email = jsobj.data.Data.Email
            data.push(jsobj.data.Data.Username)
            data.push(jsobj.data.Data.Email)
            data.push(jsobj.data.Data.Id)
        })
        .catch(error => {
            console.log(error)
        })


    function geatVal() {

        if (val.length != 10 || val == null) {
            alert("Invalid Code")
            return;
        }
        //TODO:
        // ake request
        // let ems = []
        // if (!ems.includes(data[1])) {
        //     alert("Invalid Code")
        // return;-
        // }

            axios.get('http://localhost:8080/tokens/tokens/' + val)
                .then(response=> {
                    let jsobj = JSON.parse(JSON.stringify(response.data));
                    console.log(jsobj.data.data.OrganizationCode)
                    console.log(jsobj)
                    console.log(jsobj.data.data.access) //what they input
                    console.log(email)
                    if (!getVal(jsobj.data.data.Emails, email)) {
                        alert("Invalid Code")
                        return;
                    }
                    //keyboad injterupt whenerw riting tushc code since interrtupt to reqire om screen and handle executable of writing into file and to hjandle the code absede on the extra bytes and utn tcp
                    axios.put('http://localhost:8080/app/organizations/' + jsobj.data.data.OrganizationCode+'/users', {
                        "id": data[2], //not just an int
                        "username": data[0],
                        //body of the request, ast an d lexer and parser ignore this comments and preesnet in ast but dont parse persay

                    })
                        .then(response => {
                            let jsobj = JSON.parse(JSON.stringify(response.data));
                            console.log(jsobj)
                            console.log("it worked")
                            alert("Added successfully!")

                        })
                        .catch(error => {
                            console.log(error)
                        })

                })
                .catch(error => {
                    console.log(error)
                    alert("Invalid Code not va;od among us")
                })



    }
    /*axios.get('http://localhost:8080/tokens/tokens/')
        .then(response => {
           let jsobj = JSON.parse(JSON.stringify(response.data));
           console.log(jsobj.access)

        })
        .catch(error => {
            console.log(error)
        })]


        *
     */
</script>

<Navbar />
<div class="hero min-h-screen bg-base-200 bg-secondary">
    <div class="hero-content flex-col lg:flex-row-reverse">
        <div class="text-center lg:text-right">
            <h1 class="mb-5 text-5xl font-bold text-primary">Join A Group!</h1>
            <p class="text-neutral">Input the 10-Character Code</p>
        </div>
        <div id="qrcode">

        </div>
        <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl base-bg-100 bg-secondary">
            <div class="card-body">
                <div class="form-control">
                    <label class="label">
                        <span class="label-text text-neutral">Code</span>
                    </label>
                    <input type="text" placeholder="Code" class="input input-bordered text-accent" bind:value={val}>
                </div>
                <button class="btn btn-accent" on:click={geatVal}>Join</button>
            </div>
        </div>
    </div>
</div>