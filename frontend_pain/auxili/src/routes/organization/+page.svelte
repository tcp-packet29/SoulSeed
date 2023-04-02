<script>
    import axios from 'axios';
    import { browser } from '$app/environment';
    import { getTok, getToken, prse } from '$lib/parse.svelte';
    //http reference?
    //TODO:
    //1. send a request to check if user is logged in through jwt auth middleware to fertch dat aform the databs erolesnaocljerf and render orgs
    var ids = []
    var owned = []
    axios.get('http://localhost:8080/app/auth', {
        headers: {//inverse kinematics reference
            "Token": getToken(),
        }
    })
        .then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            if (browser) {
                alert("not logged in");
                console.log(error);
                // window.location.href = "/login" //jwt reference
            }

        });

    axios.get('http://localhost:8080/access/users/token', {
        headers: {
            "Token": getToken(),
        }
    })
        .then(function (response) {
            console.log(response);
            ids = response.data.data.Data.OrganizationsIn;
            console.log(ids);
            owned = response.data.data.Data.OrganizationOwned;
        })
        .catch(function (error) {

            if (browser) {
                alert("not logged in");
                console.log(error);
                // window.location.href = "/login";
            }})



</script>
<body class="bg-primary"></body>
<div class="flex h-screen justify-center items-center">
<ul class="menu  bg-base-100 w-200 p-2 rounded-box bg-secondary">
    <li class="menu-title">
        <span class="hover:text-accent">Organizations In</span>
    </li>
    {#each ids as id}
        <li><a href="/organization/{id}">{id}</a></li>
    {/each}
    <li class="menu-title">
        <span class="hover:text-accent">Organizations Owned</span>
    </li>
    {#each owned as id}
        <li><a href="/organization/{id}">{id}</a></li>
    {/each}




    <li class="menu-title">
        <span class="hover:text-accent" >Util</span>
    </li>
    <li><a class>+ Create</a></li>
</ul>
</div>