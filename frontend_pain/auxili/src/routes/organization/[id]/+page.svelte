<script>
    import axios from 'axios'
    import { page } from '$app/stores';
    import { getTok, getToken, prse } from '$lib/parse.svelte';
    import {browser} from "$app/environment";
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
                .then(function (response) {
                    const orgin = response.data.data.Data.OrganizationsIn.concat(response.data.data.Data.OrganizationOwned); //immutable

                    if (!orgin.includes(dat)) {
                        console.log(orgin)
                        console.log("edit keep in ast")
                    } else {
                        nam = response.data.data.Data.Name;
                        desc = response.data.data.Data.Description;
                        zipc = response.data.data.Data.Zipcode;
                    }


                })
                .catch(function (error) {

    
                    if (browser) {
                        alert("not logged in");
                        console.log(error);
                        // window.location.href = "/login";
                    }})
        })
        .catch(function (error) {
            console.log(error)
        })

</script>

<div class="flex h-screen justify-center items-center">
    <div class="card w-96 bg-secondary bg-base-500 text-neutral-content shadow-xl">
        <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
        <div class="card-body">
            <h2 id="Name" class="card-title text-neutral">{nam}</h2>
            <br>
            <p id="Zipcode" class="text-neutral">{desc}</p>
            <br>
            <h2 class="text-neutral card-side"><b>{zipc}</b></h2>
        </div>
    </div>
</div>