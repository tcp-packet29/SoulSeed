<script>
    import { page } from '$app/stores';
    import axios from 'axios';
    import { onMount } from 'svelte';
    let items = []

    console.log($page.params.id)
    function replaceName(name, zip, ob) {
        document.getElementById("Name").innerHTML = name
        document.getElementById("Zipcode").innerHTML = zip
    }



    axios.get("http://localhost:8080/access/users/" + $page.params.id).then((response) => {
        if (response.status === 200) {
            let usr = JSON.parse(JSON.stringify(response.data))
            console.log(usr, usr.data.data.Username)
            replaceName(usr.data.data.Username, "Zipcode: " + +usr.data.data.Zipcode, usr.data.data.Items)
            items = usr.data.data.Items
            console.log(items)


        }

    }).catch((error) => {
        console.log(error)

    })





</script>

<body class="bg-primary"></body>

<div class="flex h-screen justify-center items-center">
<div class="card w-96 bg-secondary bg-base-500 text-neutral-content shadow-xl">
    <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
    <div class="card-body">
        <h2 id="Name" class="card-title text-neutral">Golang</h2>
        <p id="Zipcode" class="text-neutral">High Level, Static Typed, Compiled</p>
        <br>
        <h2 class="text-neutral card-side"><b>Items</b></h2>
        {#each items as item}
            <div class="badge badge-accent gap-2">
                {item}
            </div>
        {/each}
    </div>
</div>
</div>


