<script>
    import { page } from '$app/stores';
    import axios from 'axios';
    import Navbar from '../../../lib/navbar.svelte';
    import { onMount } from 'svelte';
    import {browser} from "$app/environment";
    let items = []

    console.log($page.params.id)
    function replaceName(name, zip, ob) {
        if (browser) {
        document.getElementById("Name").innerHTML = name
        document.getElementById("Zipcode").innerHTML = zip
        items = ob
    }
}



    axios.get("http://localhost:8080/access/users/" + $page.params.id).then((response) => {
        if (response.status === 200) {
            console.log(response)
            let usr = JSON.parse(JSON.stringify(response.data))
            console.log(usr, usr.data.data.Username)
            replaceName(usr.data.data.Username, "Zipcode: " + +usr.data.data.Zipcode, usr.data.data.Items) //when elxed same since ocn catenation and adding stirngs when lexed abnd parsed and represetned in abrastc sybntax tree my code omcpieldnaSDOICAS4EDXTEND LSKP MACRO JWT AUTH MIDDLEARE JW TUAHT N TERCEPT RE
            if (usr.data.data.Items != null) {
                items = usr.data.data.Items
            } else {
                items = ["No Items"]
            }



            console.log(items)


        }

    }).catch((error) => {
        console.log(error)
        //if (error.response.status === 404) {
           // replaceName("404", "Zipcode: N/A", ["No Items"])
        //}


    })





</script>

<body class="bg-primary"></body>

<Navbar />
<div class="flex h-screen justify-center items-center">
<div class="card w-96 bg-secondary bg-base-500 text-neutral-content shadow-xl border-4 hover:border-8 transition-all border-accent">
    <div class="card-body text-center">
        <div class="divider">
            <h2  class="text-neutral text-xl uppercase text-center">USER</h2>
        </div>
        <br>
        <h2 id="Name" class="text-center text-neutral text-3xl uppercase font-bold">Golang</h2>
        <br>
        <br>
        <br>
        <p id="Zipcode" class="text-neutral">High Level, Static Typed, Compiled</p>
        <br>
<!--        <h2 class="text-neutral card-side"><b>Items</b></h2>-->
<!--        <ul class="menu bg-base-100 w-56 p-2 rounded-box bg-secondary">-->
<!--        &lt;!&ndash;{#each items as item}&ndash;&gt;-->
<!--        -->
<!--        &lt;!&ndash;        <li>&ndash;&gt;-->
<!--        &lt;!&ndash;            <div>&ndash;&gt;-->
<!--        &lt;!&ndash;            <b class="text-accent">&ndash;&gt;-->
<!--        &lt;!&ndash;                O&ndash;&gt;-->
<!--        &lt;!&ndash;            </b>&ndash;&gt;-->
<!--        &lt;!&ndash;            <div class="badge badge-accent gap-2">&ndash;&gt;-->
<!--        -->
<!--        &lt;!&ndash;                {item}&ndash;&gt;-->
<!--        &lt;!&ndash;            </div>&ndash;&gt;-->
<!--        &lt;!&ndash;            </div>&ndash;&gt;-->
<!--        &lt;!&ndash;        </li>&ndash;&gt;-->
<!--        -->
<!--        -->
<!--        &lt;!&ndash;{/each}&ndash;&gt;-->
<!--        </ul>-->
    </div>
</div>
</div>


