<script>
    import axios from 'axios';
    import {browser} from "$app/environment";
    let val = '';
    function getToken() {
        if (browser) {
            return window.localStorage.getItem("token");

        }
    }

    axios.defaults.headers.put['Token'] = getToken();
    function getVal() {
        axios.get('http://localhost:8080/tokens/tokens/' + val)
            .then(response=> {
                let jsobj = JSON.parse(JSON.stringify(response.data));
                console.log(jsobj.data.data.OrganizationCode)
                console.log(jsobj)
                console.log(jsobj.data.data.access) //what they input
                axios.put('http://localhost:8080/app/organizations/' + jsobj.data.data.OrganizationCode+'/users', {
                    "id": "123456789876543212345678", //not just an int
                    "username": "test",
                    "email": "mycodehathcompiled@mycodehathcompiled.com",
                    "password": "test",

                })
                    .then(response => {
                        let jsobj = JSON.parse(JSON.stringify(response.data));
                        console.log(jsobj)
                        console.log("it worked")
                        alert(jsobj.data.data.organization.description)

                    })
                    .catch(error => {
                        console.log(error)
                    })

            })
            .catch(error => {
                console.log(error)
                alert("Invalid Code not va;od among us")
            })
    }
    /*axios.get('http://localhost:8080/tokens/tokens/')
        .then(response => {
           let jsobj = JSON.parse(JSON.stringify(response.data));
           console.log(jsobj.access)

        })
        .catch(error => {
            console.log(error)
        })]


        *
     */
</script>


<div class="hero min-h-screen bg-base-200 bg-secondary">
    <div class="hero-content flex-col lg:flex-row-reverse">
        <div class="text-center lg:text-right">
            <h1 class="mb-5 text-5xl font-bold text-neutral">Join A Group!</h1>
            <p class="text-neutral">Input the 10-Character Code</p>
        </div>
        <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl base-bg-100 bg-secondary">
            <div class="card-body">
                <div class="form-control">
                    <label class="label">
                        <span class="label-text text-neutral">Code</span>
                    </label>
                    <input type="text" placeholder="Code" class="input input-bordered text-accent" bind:value={val}>
                </div>
                <button class="btn btn-accent" on:click={getVal}>Join</button>
            </div>
        </div>
    </div>
</div>