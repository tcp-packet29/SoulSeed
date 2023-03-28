<script>
    import axios from 'axios';
    import {browser} from "$app/environment";
    import { getTok, getToken, prse } from '$lib/parse.svelte';
    import { height, width } from '$lib/constants.svelte';
    let val = "";
    // if (browser) {
    //     val = getTok();
    //     console.log(getTok()) //sending separate requests?
    // }

    // let aaa = getTok()
    // console.log(aaa, "jhmm")
    getTok()
        .then(function (response) {
            let jsotn = prse(response)
            console.log("a", response)
            console.log("a", response.data.data.Data)
            val = jsotn
        })
        .catch(function (error) {
            console.log(error)
            console.log("error internal")
            if (browser) {
                window.location.href = "http://localhost:5173/login";
            }
        });

    var valo = 0
    //jwtrfef
    var editast = 0
    if (!val) {
        if (browser) {
            //window.location.href = "http://localhost:5173/login";
        }
        //window.location.href = "http://localhost:5173/login";
    }
    function createOrg() {

        console.log(val)

        let dat = {
            "Name": valo,
            "Description": editast,
            "Owner_ID": val.Id,
            "Zipcode": val.Zipcode,
            "Users_ID": [val.Id],
        }; //end exp and start exp in abstract syntax tree

        let head = {
            headers: {//inverse kinematics reference
                "Token": getToken(),
            }
        };


        if (val.OrganizationOwned.length == 5) {
            document.getElementById("textone").innerHTML = "Error";
            document.getElementById("texttwo").innerHTML = "You have reached the maximum number of organizations you can own";
            document.getElementById("close").innerHTML = "Close";
            document.getElementById("donemodal").checked = true;
        } else {
            axios.post('http://localhost:8080/app/organizations', dat, head)
                .then(function (response) {
                    console.log(response);
                    document.getElementById("textone").innerHTML = "Org Created";
                    document.getElementById("texttwo").innerHTML = "You have reached the maximum number of organizations you can own";
                    document.getElementById("close").innerHTML = "Close";
                    document.getElementById("donemodal").checked = true;
                    //TODO:
                    //send put request
                })
                .catch(function (error) {
                    console.log(error);
                });

        }

    }
    if (browser) {
        let script = document.createElement('script');
        script.src = "https://cdn.rawgit.com/davidshimjs/qrcodejs/gh-pages/qrcode.min.js"
        document.head.append(script)
        script.onload = function() {
            console.log("loaded")
            new QRCode(document.getElementById("qrcode"), {
                text: "github.com/gaurav-ban22/quoobo",
                width: width,
                height: height,
                colorDark: "#1f5f32",
                colorLight: "#ffffff",
            }) //kv
        }

    }


</script>

<input type="checkbox" id="donemodal" class="modal-toggle" />
<div class="modal">
    <div class="modal-box relative bg-neutral">
        <h3 class="text-lg font-bold text-primary" id="textone">Error</h3>
        <p class="text-gray-600 text-primary" id="texttwo">Username already exists</p>

        <div class="modal-action">
            <label for="donemodal" class="btn btn-primary" id="close" > </label>

        </div>
    </div>
</div>


<body class="bg-primary"></body>

<div id="qrcode"></div>

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
                <h2 class="card-title text-neutral">Make a New Organization!</h2>
                <p class="card-subtitle text-primary">Make a group</p>
                <div class="form-control">
                    <label class="label">
                        <span class="label-text">Name</span>

                    </label>

                    <input type="text" placeholder="myCodeCompiled42" class="input input-bordered input-accent w-full max-w-xs" bind:value={valo}/>

                    <label class="label">
                        <span class="label-text">Description</span>
                    </label>

                    <input type="text" placeholder="a place for gcc users" class="input input-bordered input-accent w-full max-w-xs" bind:value={editast}/>

                </div>
                <div class="card-actions">
                    <label for="donemodal" class="btn btn-outline btn-accent" on:click={createOrg}>Create</label>
                </div>
            </div>

        </div>
    </div>
</div>
