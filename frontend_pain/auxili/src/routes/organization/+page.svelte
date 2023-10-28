<script>
    import axios from 'axios';
    import { browser } from '$app/environment';
    import { getTok, getToken, prsse } from '$lib/parse.svelte';
    import Navbar from "../../lib/navbar.svelte";
    //http reference?
    //TODO:
    //1. send a request to check if user is logged in through jwt auth middleware to fertch dat aform the databs erolesnaocljerf and render orgs
    var ids = []
    var owned = []
    var dats = []
    var ins = []
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
            owned.splice(0,1);
            ids.splice(0, 1)
            axios.post('http://localhost:8080/app/organizations/idBulk', {
                "orgVal": owned,
            }, {
                headers: {
                    "Token": getToken(),
                }
            })
                .then(function (response) {
                    console.log(response)
                    let vals = prsse(response)
                    console.log(vals);
                    dats = vals.organizations;

                })
                .catch(function (error) {
                    if (browser) {
                        alert("not logged in");
                        console.log('why sir estapi')
                        console.log(error);
                        // window.location.href = "/login";
                    }

                })
            axios.post('http://localhost:8080/app/organizations/idBulk', {
                "orgVal": ids,
            }, {
                headers: {
                    "Token": getToken(),
                }
            })
                .then(function (response) {
                    console.log(response)
                    let vals = prsse(response)
                    console.log(vals);
                    ins = vals.organizations;

                })
                .catch(function (error) {
                    if (browser) {
                        alert("not logged in");
                        console.log('why sir estapi')
                        console.log(error);
                        // window.location.href = "/login";
                    }

                })
        })
        .catch(function (error) {

            if (browser) {
                alert("not logged in");
                console.log(error);
                // window.location.href = "/login";
            }})



</script>

<Navbar />
<body class="bg-primary"></body>
<div class="flex h-screen justify-center items-center">
<ul class="menu  bg-base-100 w-200 p-2 rounded-box bg-secondary">
    <li class="menu-title">
        <span class="hover:text-accent">Organizations In</span>
    </li>
    {#each ins as org}
        <li class="bg-secondary text-accent m-2"><a href="/organization/{org.Id}">{org.Name}</a></li>
    {/each}
    <li class="menu-title">
        <span class="hover:text-accent">Organizations Owned</span>
    </li>
    {#each dats as id}
        <li class="bg-secondary text-accent m-2"><a href="/organization/{id.Id}">{id.Name}</a></li>
    {/each}




    <li class="menu-title">
        <span class="hover:text-accent" >Util</span>
    </li>
    <li><a class="link-accent" href="/createOrg">Create</a></li>
</ul>
</div>