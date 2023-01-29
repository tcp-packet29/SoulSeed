<script>
    import axios from 'axios';

    let uname = "";
    let pword = "";
    let zipcode = "";
    

    function createUser() {
        if (uname == "" || pword == "" || zipcode == "") {
            alert("Please fill out all fields");
            return;
        }     
        axios.post('http://localhost:8080/users', {
            "username": uname,
            "password": pword,
            "zipcode": zipcode,
        })
        .then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            console.log(error);
        });
    }

</script>

<body class="bg-primary"></body>

<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral">Register</h2>
            <p class="card-subtitle text-primary">An app for vendors and customers; with emphasis on flexibility.</p>
            <input type="text" placeholder="Username" class="input input-bordered input-accent w-full max-w-xs" bind:value={uname}/>
            <input type="text" placeholder="Password" class="input input-bordered input-accent w-full max-w-xs" bind:value={pword}/>
            <input type="text" placeholder="Zipcode" class="input input-bordered input-accent w-full max-w-xs" bind:value={zipcode}/>
            <div class="card-actions">
                <button class="btn btn-outline btn-accent" on:click={createUser}>Register</button>
            </div>
        </div>
    </div>
</div>