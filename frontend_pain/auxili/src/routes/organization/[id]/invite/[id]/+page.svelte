<script>
    import axios from 'axios'
    import { page } from '$app/stores';
    import { getTok, getToken, prse } from '$lib/parse.svelte';
    import {browser} from "$app/environment";
    var exp = "";
    var emails = "";
    console.log($page.params.id)
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
                    var date = Date.now()
                    window.localStorage.setItem('cooldown', date.toString())

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

            <button class="btn btn-primary" on:click={aaa}>Generate</button>
        </div>
    </div>
</div>