<script>
    import axios from 'axios';

    let uname = "";
    let pword = "";
    let confirm = "";
    let zipcode = "";
    

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
            "zipcode": zipcode,
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
            <h2 class="card-title text-neutral">Register</h2>
            <p class="card-subtitle text-primary">An app for vendors and customers; with emphasis on flexibility.</p>
            <div class="form-control">
                <label class="label">
                    <span class="label-text">Username</span>
                </label>
            
                    <input type="text" placeholder="myCodeCompiled42" class="input input-bordered input-accent w-full max-w-xs" bind:value={uname}/>
                
                <label class="label">
                    <span class="label-text">Password</span>
                </label>
                
                    <input type="password" placeholder="clojureIsBestJVMLang" class="input input-bordered input-accent w-full max-w-xs" bind:value={pword}/>
                
                <label class="label">
                    <span class="label-text">Confirm Password</span>
                </label>
                
                    <input type="password" placeholder="clojureIsBestJVMLang" class="input input-bordered input-accent w-full max-w-xs" bind:value={confirm}/>
                
                <label class="label">
                    <span class="label-text">Zipcode</span>
                </label>
                
                    <input type="text" placeholder="11111" class="input input-bordered input-accent w-full max-w-xs" bind:value={zipcode}/>
                
            </div>
            
            
            
           
            <div class="card-actions">
                <button class="btn btn-outline btn-accent" on:click={createUser}>Register</button>
            </div>
        </div>
    </div>
</div>