<script>
    import axios from 'axios';
    import {browser} from "$app/environment";

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
            axios.get("http://localhost:8080/email/confirm/" + email)
                .then(function (response) {
                    let jsooo = JSON.parse(JSON.stringify(response.data))
                    console.log(jsooo)
                    tok = jsooo.data.Token.toString();

                })
                .catch(function (error) {
                    console.log(error);
                });
        }
    }

    function create() {
        if tok === code {
            axios.post('http://localhost:8080/access/users', {
                "username": uname,
                "password": pword,
                "email": email,
                "zipcode": zipcode,
            })
                .then(function (response) {
                    if (response.status == 201 || response.status == 200) {
                        valid = true;
                        fillIn("Final Step!", "A registration code was emailed to you (check your spam folder)", "Continue");

                    } else {
                        fillIn("Error", "Username already exists", "Close");
                    }
                    console.log(response);

                })
                .catch(function (error) {
                    console.log(error);
                });
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
                    //not immutable
                    //idempotenmt
                    if (browser) {
                        window.localStorage.setItem("token", jso.Token.toString());
                        console.log(window.localStorage.getItem("token"))
                    }

                    console.log(window.localStorage.getItem("token"))

                })
                .catch(function (error) {
                    console.log(error);
                });
        }
    }
    //checking zipcode on frontend, could do on bacvkend
    function validZipcode(a) {
        let regex = new RegExp("^[0-9]{5}(?:-[0-9]{4})?$");
        return regex.test(a)
    }






    function check() {
        if (uname.trim() == "" || pword.trim() == "" || zipcode.trim() == "" || confirm.trim() == "") {
            fillIn("Error", "Please fill in all fields", "Close");
        } else if (pword != confirm) {
            fillIn("Error", "Passwords do not match", "Close");

        } else if (!validZipcode(zipcode)) {
            fillIn("Error", "Zipcode is not valid", "Close");

        }
    }

</script>

<input type="checkbox" id="donemodal" class="modal-toggle" />
<div class="modal">
    <div class="modal-box relative bg-neutral">
        <h3 class="text-lg font-bold text-primary" id="textone">Error</h3>
        <p class="text-gray-600 text-primary" id="texttwo">Username already exists</p>
        <div class="modal-action">
            <label for="donemodal" class="btn btn-primary" id="close" on:click={leadToLogin}>Close</label>
        </div>
    </div>
</div>


<body class="bg-primary"></body>

<div class="navbar bg-base-100 bg-accent">
    <a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173">Auxili</a>
    <h1>--</h1>
    <h2><a class="btn btn-sm btn-ghost rounded-btn bg-secondary" href="http://localhost:5173/register">Register</a></h2>
    <h1>--</h1>
    <h2><a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173/login">Login</a></h2>

</div>

<div class="flex h-screen justify-center items-center">
    <div>
        <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
            <div class="card-body items-center text-center">
                <h2 class="card-title text-neutral">Register</h2>
                <p class="card-subtitle text-primary">An app for vendors and customers; with emphasis on flexibility.</p>
                <div class="form-control">
                    <label class="label">
                        <span class="label-text">Username</span>

                    </label>

                        <input type="text" placeholder="myCodeCompiled42" class="input input-bordered input-accent w-full max-w-xs" bind:value={uname}/>

                    <label class="label">
                        <span class="label-text">Email</span>
                    </label>

                    <input type="text" placeholder="golang@golang.go" class="input input-bordered input-accent w-full max-w-xs" bind:value={email}/>

                    <label class="label">
                        <span class="label-text">Password</span>
                    </label>

                        <input type="password" placeholder="clojureIsBestJVMLang" class="input input-bordered input-accent w-full max-w-xs" bind:value={pword}/>


                    <label class="label">
                        <span class="label-text">Confirm Password</span>
                    </label>

                        <input type="password" placeholder="clojureIsBestJVMLang" class="input input-bordered input-accent w-full max-w-xs" bind:value={confirm}/>

                    <label class="label">
                        <span class="label-text">Zipcode</span>
                    </label>

                        <input type="text" placeholder="11111" class="input input-bordered input-accent w-full max-w-xs" bind:value={zipcode}/>

                </div>
                <div class="card-actions">
                    <label for="donemodal" class="btn btn-outline btn-accent" on:click={check}>Register</label>
                </div>
            </div>

            <div class="divider text-accent"> & </div>

            <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
                <div class="card-body items-center text-center">
                    <h2 class="card-title text-neutral">Confirmation</h2>
                    <p class="mb-5">Please enter your confirmation code that was sent to your email.</p>
                    <input type="password" placeholder="Password" class="input input-bordered input-accent w-full max-w-xs" bind:value={code}/>
                    <div class="card-actions">
                        <button class="btn btn-outline btn-accent" on:click={create}>Check</button>
                    </div>
                </div>
            </div>

        </div>
    </div>
</div>
