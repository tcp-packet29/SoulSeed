<script>
    import axios from 'axios';
    import {browser} from "$app/environment";
    import { usrn, psord, em, zpc, emTok} from "../../stores.js";
    import {get} from "svelte/store";
    import Thisworks from "../../lib/thisworks.svelte";

    let uname = "";
    let pword = "";
    let confirm = "";
    let zipcode = "";
    let email= "";
    let code ="";

    let valid= false;

    function fillIn(one, two, btn) {
        document.getElementById("textone").innerHTML = one;
        document.getElementById("texttwo").innerHTML = two;
        //mycodehagthcomepidl
        document.getElementById("close").innerHTML = btn;
    } //decode then weas not assigned or issued by this server mthe provider

    //implemen jwt handlinglogin if cannot d

    //programin rabbitmq clone with tui
    //in addition implement profile api for processing request and sending response
    //i rember i should make my ownr fabitmq     clone and rep it thrugh a tui using golang tui library called tvuiew e vcayuse then u can represetn queures through trui using tview

    let tok = "";

    function leadToLogin() {
        if (valid) {
            if (browser) {
                window.localStorage.setItem("username", uname);
                window.localStorage.setItem("password", pword);
                window.localStorage.setItem("email", email);
                window.localStorage.setItem("zipcode", zipcode);
            }
            window.location.href = "http://localhost:5173/confirm"
        }
    }


    //checking zipcode on frontend, could do on bacvkend
    function validZipcode(a) {
        let regex = new RegExp("^[0-9]{5}(?:-[0-9]{4})?$");
        return regex.test(a)
    }






    function check() {
        if (uname.trim() === "" || pword.trim() === "" || zipcode.trim() === "" || confirm.trim() === "") {
            fillIn("Error", "Please fill in all fields", "Close");
            return
        } else if (pword !== confirm) {
            fillIn("Error", "Passwords do not match", "Close");
            return
        } else if (!validZipcode(zipcode)) {
            fillIn("Error", "Zipcode is not valid", "Close");
            return

        }
        axios.get('http://localhost:8080/users/' + uname)
            .then(function (response) {
                if (response.status == 200) {
                    fillIn("Error", "Username already exists", "Close");
                    console.log(response)
                }

            })
            .catch(function (error) {
                console.log(error);
                if (error.response.status == 404) {

                    valid = true

                    fillIn("Final Step", "Please enter the code sent to your email", "GOTO")

                }
            });

    }

</script>

<input type="checkbox" id="donemodal" class="modal-toggle" />
<div class="modal">
    <div class="modal-box relative bg-secondary">
        <h3 class="text-lg font-bold text-primary" id="textone">Error</h3>
        <p class="text-gray-600 text-primary" id="texttwo">Username already exists</p>

        <div class="modal-action">
            <label for="donemodal" class="btn btn-primary" id="close" on:click={leadToLogin}> </label>

        </div>
    </div>
</div>


<body class="bg-primary"></body>

<Thisworks/>

<div class="flex h-screen justify-center items-center">
    <div> <!-- ignore in html to oparse uib abtufe vriwedsr redbere oarse html by http requedst !-->
        <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
            <div class="card-body items-center text-center">
                <h2 class="card-title text-neutral">Register</h2>
                <p class="card-subtitle text-accent">Already have an account? <a class="link-hover" href="/login">Login Here.</a></p>
                <div class="form-control">
                    <label class="label">
                        <span class="label-text">Username</span>

                    </label>

                        <input type="text" placeholder="myCodeCompiled42" class="input input-bordered input-accent w-full max-w-xs text-white" bind:value={uname}/>

                    <label class="label">
                        <span class="label-text">Email</span>
                    </label>

                    <input type="text" placeholder="golang@golang.go" class="input input-bordered input-accent w-full max-w-xs text-white" bind:value={email}/>

                    <label class="label">
                        <span class="label-text">Password</span>
                    </label>

                        <input type="password" placeholder="clojureIsBestJVMLang" class="input input-bordered input-accent w-full max-w-xs text-white" bind:value={pword}/>


                    <label class="label">
                        <span class="label-text">Confirm Password</span>
                    </label>

                        <input type="password" placeholder="clojureIsBestJVMLang" class="input input-bordered input-accent w-full max-w-xs text-white" bind:value={confirm}/>

                    <label class="label">
                        <span class="label-text">Zipcode</span>
                    </label>

                        <input type="text" placeholder="11111" class="input input-bordered input-accent w-full max-w-xs text-white" bind:value={zipcode}/>

                </div>
                <div class="card-actions">
                    <label for="donemodal" class="btn btn-outline btn-accent" on:click={check}>Register</label>
                </div>
            </div>

        </div>
    </div>
</div>
