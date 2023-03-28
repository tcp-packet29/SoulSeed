<script>
    import axios from 'axios';
    import {browser} from "$app/environment";
    let val = '';
    function getToken() {
        if (browser) {
            return window.localStorage.getItem("token");

        }
    }
    var data = [];
    axios.defaults.headers.put['Token'] = getToken();
    axios.get('http://localhost:8080/access/users/token', {
        headers: {
            'Token': getToken(),
        }
    })
    .then(response => {
            let jsobj = JSON.parse(JSON.stringify(response.data));
            console.log(jsobj)
            console.log(jsobj.data.Data.Username)
            console.log(jsobj.data.Data.Email)
            data.push(jsobj.data.Data.Username)
            data.push(jsobj.data.Data.Email)
            data.push(jsobj.data.Data.Id)
        })
        .catch(error => {
            console.log(error)
        })


    function getVal() {

        if (val.length != 10 || val == null) {
            alert("Invalid Code")
            return;
        }
        //TODO:
        // ake request
        // let ems = []
        // if (!ems.includes(data[1])) {
        //     alert("Invalid Code")
        // return;-
        // }
        axios.get('http://localhost:8080/tokens/tokens/' + val)
            .then(response=> {
                let jsobj = JSON.parse(JSON.stringify(response.data));
                console.log(jsobj.data.data.OrganizationCode)
                console.log(jsobj)
                console.log(jsobj.data.data.access) //what they input
                axios.put('http://localhost:8080/app/organizations/' + jsobj.data.data.OrganizationCode+'/users', {
                    "id": data[2], //not just an int
                    "username": data[0],
                    //body of the request, ast an d lexer and parser ignore this comments and preesnet in ast but dont parse persay

                })
                    .then(response => {
                        let jsobj = JSON.parse(JSON.stringify(response.data));
                        console.log(jsobj)
                        console.log("it worked")
                        alert(jsobj.data.data.description)

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