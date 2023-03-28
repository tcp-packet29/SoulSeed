<script>
    import axios from 'axios';

    import { usrn, psord, em, zpc, emTok} from "../../stores.js";
    import {get} from "svelte/store";


    let val = "";
    let token = ""
    let valid= false;
    axios.post("http://localhost:8080/email/password/" + val, {

    })
        .then(function (response) {
            let jsooo = JSON.parse(JSON.stringify(response.data))
            console.log(jsooo)
            token = jsooo.data.Token.toString();



            //

            //lambdas anonymosu inolie fiunc s firs tclass functigons


            //stores not working????

        })
        .catch(function (error) {
            console.log(error);

        });

</script>

<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral">Confirmation</h2>
            <p class="mb-5">Please enter your email and checkin.</p>
            <input type="text" placeholder="Password" class="input input-bordered input-accent w-full max-w-xs" bind:value={val}/>
            <div class="card-actions">
                <button class="btn btn-outline btn-accent" on:click={Check}>Check</button>
            </div>
        </div>
    </div>
</div>
