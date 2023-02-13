<script>
    import { page } from '$app/stores';
    import axios from 'axios';
    import { onMount } from 'svelte';



    console.log($page.params.id)
    function replaceName(name, zip) {
        document.getElementById("Name").innerHTML = name
        document.getElementById("Zipcode").innerHTML = zip
    }



    axios.get("http://localhost:8080/access/users/" + $page.params.id).then((response) => {
        if (response.status == 200) {
            let usr = JSON.parse(JSON.stringify(response.data))
            console.log(usr, usr.data.data.Username)
            replaceName(usr.data.data.Username, usr.data.data.Zipcode)


        }

    }).catch((error) => {
        console.log(error)

    })





</script>


<div class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
    <figure><img src="https://miro.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png" alt="Shoes" /></figure>
    <div class="card-body">
        <h2 id="Name" class="card-title text-neutral">Golang</h2>
        <p id="Zipcode" class="text-neutral">High Level, Static Typed, Compiled</p>
    </div>
</div>