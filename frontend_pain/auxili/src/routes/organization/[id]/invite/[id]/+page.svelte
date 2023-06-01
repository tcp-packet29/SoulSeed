<script>
    import axios from 'axios'
    import { page } from '$app/stores';
    import { getTok, getToken, prse } from '$lib/parse.svelte';
    import {browser} from "$app/environment";
    var exp = "";
    var emails = "";
    console.log($page.params.id)
    if (browser) {
        let script = document.createElement('script');
        script.src = "https://cdn.rawgit.com/davidshimjs/qrcodejs/gh-pages/qrcode.min.js"
        document.head.append(script);
    }


    if (browser) {
        let token = window.localStorage.getItem('token')
        if (token == null) {
            window.location.href = "/login"
        }
    }
    axios.get("http://localhost:8080/app/auth", {
        headers: {
            Token: getToken()
        }
    })
        .then(function(response) {
            console.log(response)
        })
        .catch(function (error) {
            console.log(error)
            // window.location.href = '/login'

        })
    function aaa() {
        if (browser) {
            var date = Date.now()
            var cooldown = window.localStorage.getItem('cooldown')
            if (cooldown != null) {
                if (date - parseInt(cooldown) < 10000) {
                    document.getElementById("token").innerHTML = "Please Wait 10 Seconds Before Generating Another Token"
                    document.getElementById("token").style.color = "red"
                    return
                }
            }
        }
        axios.post("http://localhost:8080/tokens/tokens/" + exp, {
            "OrganizationCode": $page.params.id,
            "emails": emails.split(","),
        },{
            headers: {
                Token: getToken()
            }
        })
            .then(function(response) {
                if (browser) {
                    document.getElementById("token").style.color = null
                    document.getElementById("token").innerHTML = response.data.data.data.Access;
                    document.getElementById("qr").innerHTML = "";
                    const qrcode = new QRCode("qr", {
                        text: "http://localhost:5173/joinOrganization/" + response.data.data.data.Access,
                        width: 128,
                        height: 128,
                        colorDark : "#000000",
                        colorLight : "#ffffff",
                    })

                    var date = Date.now()
                    window.localStorage.setItem('cooldown', date.toString())
                    axios.post("http://localhost:8080/email/join/test/" + emails + "/" + response.data.data.data.Access +"/" + exp, {
                        "organization_code": $page.params.id,
                        "emails": emails.split(","),
                    },{
                        headers: {
                            Token: getToken()
                        }
                    })
                        .then(function(response) {
                            console.log(response)
                        })
                        .catch(function(error) {
                            console.log(error)
                            if (error.statusCode == 400) {
                                alert("Invalid Email (keep less than 10)")
                            }
                        })

                }
                console.log(response)



            })
            .catch(function (error) {
                console.log(error)
            })
    }


</script>

<div class="flex h-screen justify-center items-center">
    <div class="card w-96 bg-secondary bg-base-500 text-neutral-content shadow-xl">

        <div class="card-body">
            <h2 class="card-title text-neutral ">Invite</h2>

            <input type="text" placeholder="Expiry" class="input input-bordered w-full" bind:value={exp} />
            <h1 class="text-neutral" id="token">Token To Be Generated!</h1>
            <input type="text" placeholder="Emails (separate by comma)" class="input input-bordered w-full" bind:value={emails} />
            <div id="qr"></div>
            <button class="btn btn-primary" on:click={aaa}>Generate</button>
        </div>
    </div>
</div>