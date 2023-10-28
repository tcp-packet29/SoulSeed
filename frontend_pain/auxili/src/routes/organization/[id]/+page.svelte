<script>
    import axios from 'axios'
    import { page } from '$app/stores';
    import { getTok, getToken, prse } from '$lib/parse.svelte';
    import {browser} from "$app/environment";
    import Navbar from "../../../lib/navbar.svelte";
    var nam = "";
    var desc = "";
    var zipc = "";
    axios.get("http://localhost:8080/app/organizations/" + $page.params.id, {
        headers: {
            Token: getToken()
        }
    })
        .then(function(response) {
            console.log(response)
            const dat = response.data.data.organization.Id
            console.log(dat)
            axios.get('http://localhost:8080/access/users/token', {
                headers: {
                    "Token": getToken(),
                }
            })
                .then(function (responsea) {
                    console.log(responsea)
                    const orgin = responsea.data.data.Data.OrganizationsIn.concat(responsea.data.data.Data.OrganizationOwned); //immutable

                    if (!orgin.includes(dat)) {
                        console.log(orgin)
                        console.log("edit keep in ast")
                        if (browser) {
                            window.location.href = "/organization"
                        }
                    } else {
                        nam = response.data.data.organization.Name;
                        desc = response.data.data.organization.Description;
                        zipc = response.data.data.organization.Zipcode;

                    }


                })
                .catch(function (error) {

    
                    if (browser) {
                        alert("Not logged in!");
                        console.log(error);
                        if (browser) {
                            window.location.href = "/organization"
                        }
                    }})
        })
        .catch(function (error) {
            console.log(error)
        })

</script>
<Navbar />

<div class="flex h-screen justify-center items-center">
    <div class="card w-96 bg-secondary bg-base-500 text-neutral-content shadow-xl">
        <div class="card-body">
            <h1 class="card-title text-neutral text-2xl">ORGANIZATION</h1>
            <h2 id="Name" class="card-title text-neutral">Name: {nam}</h2>
            <hr>
            <br>
            <p id="Zipcode" class="text-neutral">{desc}</p>
            <br>
            <h2 class="text-neutral card-side"><b>Zipcode: {zipc}</b></h2>
            <hr>
            <br>
            <button class="btn btn-accent mt-4"><a href="/mainPg/{$page.params.id}/">Check Trades in Organization</a></button>
            <button class="btn btn-accent mt-4"><a href="/organization/{$page.params.id}/invite/{$page.params.id}">Invite Someone!</a></button>
        </div>
    </div>
</div>