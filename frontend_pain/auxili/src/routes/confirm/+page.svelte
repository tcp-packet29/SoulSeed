<script>

    import { browser } from '$app/environment';
    import { usrn, psord, zpc, em, emTok} from "../../stores.js";
    import axios from "axios";

    let val = ""
    let token = emTok.toString()
    function create() {
        if (tok === code) {
            axios.post('http://localhost:8080/access/users', {
                "username": usrn.toString(),
                "password": psord.toString(),
                "email": em.toString(),
                "zipcode": zpc.toString(),
            })
                .then(function (response) {
                    console.log(response);
                    if (response.status == 201 || response.status == 200) {
                        valid = true;


                    } else {
                        fillIn("Error", "Username already exists", "Close");
                    }
                    console.log(response);

                })
                .catch(function (error) {
                    console.log(error);
                });
            axios.post('http://localhost:8080/access/login', {
                "Username": usrn.toString(),
                "Password": psord.toString()
            })
                .then(function (response) {
                    if (response.status == 201 || response.status == 200) {
                        alert("Logged In successfully");
                    }
                    console.log(response);
                    let jso = JSON.parse(JSON.stringify(response.data))
                    console.log(jso)
                    console.log(jso.Token)
                    //not immutable
                    //idempotenmt
                    if (browser) {
                        window.localStorage.setItem("token", jso.Token.toString());
                        console.log(window.localStorage.getItem("token"))
                    }

                    console.log(window.localStorage.getItem("token"))

                })
                .catch(function (error) {
                    console.log(error);
                });
        }
    }
    function Check() {
        if (val === token) {
            create()
        } else {
            console.log('incorrect')
        }
    }
//  could just look in local storage
    //shpuld proabably add usr here but okay



</script>

<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral">Confirmation</h2>
            <p class="mb-5">Please enter your confirmation code that was sent to your email.</p>
            <input type="password" placeholder="Password" class="input input-bordered input-accent w-full max-w-xs" bind:value={val}/>
            <div class="card-actions">
                <button class="btn btn-outline btn-accent" on:click={Check}>Check</button>
            </div>
        </div>
    </div>
</div>