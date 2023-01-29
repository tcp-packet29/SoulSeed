<script>
    import axios from 'axios';

    let uname = "";
    let pword = "";
    

    function createUser() {
        if (uname == "" || pword == "" || zipcode == "" || confirm == "") {
            alert("Please fill out all fields");
            return;
        } else if (pword != confirm) {
            alert("Passwords do not match");
            return;
        }
        axios.post('http://localhost:8080/users', {
            "username": uname,
            "password": pword,
        })
        .then(function (response) {
            if (response.status == 201 || response.status == 200) {
                alert("User created successfully");
                window.location.href = "http://localhost:5173/login";
            }
            console.log(response);
            
        })
        .catch(function (error) {
            console.log(error);
        });
    }

</script>

<body class="bg-primary"></body>
<div class="navbar bg-base-100 bg-accent">
    <a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173">Auxili</a>
    <h1>--</h1>
    <h2><a class="btn btn-sm btn-ghost rounded-btn bg-secondary" href="http://localhost:5173/register">Register</a></h2>
    <h1>--</h1>
    <h2><a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173/login">Login</a></h2>
    
</div>
<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral">Log In</h2>
            
            <input type="text" placeholder="Username" class="input input-bordered input-accent w-full max-w-xs" bind:value={uname}/>
            <input type="password" placeholder="Password" class="input input-bordered input-accent w-full max-w-xs" bind:value={pword}/>
            <div class="card-actions">
                <button class="btn btn-outline btn-accent" on:click={createUser}>Login</button>
            </div>
        </div>
    </div>
</div>